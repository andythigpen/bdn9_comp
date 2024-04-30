package device

import (
	"fmt"
	"time"

	"github.com/karalabe/hid"

	pb "github.com/andythigpen/bdn9_comp/v2/proto"
)

var NotOpenErr = fmt.Errorf("Device not open")
var DeviceNotFound = fmt.Errorf("No BDN9 device found")

type BDN9Device interface {
	Open() error
	IsOpen() bool
	Close() error
	SetRGBMode(mode pb.RGBMode) error
	SetMatrixHSV(h uint8, s uint8, v uint8) error
	ToggleMatrix() error
	SetIndicatorRGB(layer pb.Layer, index uint8, h uint8, s uint8, v uint8) error
	ToggleIndicator(layer pb.Layer, index uint8) error
	EnableIndicator(layer pb.Layer, index uint8) error
	DisableIndicator(layer pb.Layer, index uint8) error
	ActivateLayer(layer pb.Layer) error
	SetSpeed(speed uint8) error
	SetMuteStatus(muted bool) error
	EndCall() error
	Reset() error
	Echo(b []byte) error
	EnablePersist()
	DisablePersist()
}

type bdn9Device struct {
	device  *hid.Device
	persist byte
	handler EventHandler
}

const (
	VID = 0xcb10
	PID = 0x2133
)

func boolToByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

func NewDevice(handler EventHandler) BDN9Device {
	return &bdn9Device{
		device:  nil,
		persist: 0,
		handler: handler,
	}
}

func handleEvents(d *bdn9Device) {
	buf := make([]byte, 32)
	for {
		// blocks until bytes are read
		sz, err := d.device.Read(buf)
		if err != nil {
			fmt.Printf("Error reading: %s\n", err)
			d.Close()
			return
		}
		if sz == 0 {
			fmt.Println("No bytes read")
			time.Sleep(500 * time.Millisecond)
			continue
		}

		if d.handler != nil {
			for i := 0; i < sz; i++ {
				d.handler.HandleEvent(d, Event(buf[i]))
			}
		}
	}
}

func (d *bdn9Device) Open() (err error) {
	deviceInfos := hid.Enumerate(VID, PID)
	if len(deviceInfos) < 2 {
		return DeviceNotFound
	}
	d.device, err = deviceInfos[1].Open()
	if err != nil {
		return
	}

	if d.handler != nil {
		go handleEvents(d)
	}

	pkt := NewPacket(COMMAND_CONNECT)
	err = d.writeAll(pkt)
	return
}

func (d *bdn9Device) IsOpen() bool {
	return d.device != nil
}

func (d *bdn9Device) Close() error {
	pkt := NewPacket(COMMAND_DISCONNECT)
	_ = d.writeAll(pkt)
	err := d.device.Close()
	d.device = nil
	return err
}

func (d *bdn9Device) EnablePersist() {
	d.persist = 1
}

func (d *bdn9Device) DisablePersist() {
	d.persist = 0
}

func (d *bdn9Device) SetRGBMode(mode pb.RGBMode) error {
	if d.device == nil {
		return NotOpenErr
	}
	if err := mode.IsValid(); err != nil {
		return err
	}
	pkt := NewPacket(COMMAND_SET_MODE, d.persist, byte(mode))
	return d.writeAll(pkt)
}

func (d *bdn9Device) SetMatrixHSV(h uint8, s uint8, v uint8) error {
	if d.device == nil {
		return NotOpenErr
	}
	pkt := NewPacket(COMMAND_SET_MATRIX_HSV, d.persist, h, s, v)
	return d.writeAll(pkt)
}

func (d *bdn9Device) ToggleMatrix() error {
	if d.device == nil {
		return NotOpenErr
	}
	pkt := NewPacket(COMMAND_TOGGLE_MATRIX, d.persist)
	return d.writeAll(pkt)
}

func (d *bdn9Device) SetIndicatorRGB(layer pb.Layer, index uint8, r uint8, g uint8, b uint8) error {
	if d.device == nil {
		return NotOpenErr
	}
	if err := layer.IsValid(); err != nil {
		return err
	}
	pkt := NewPacket(COMMAND_SET_INDICATOR_RGB, byte(layer), index, r, g, b)
	return d.writeAll(pkt)
}

func (d *bdn9Device) ToggleIndicator(layer pb.Layer, index uint8) error {
	if d.device == nil {
		return NotOpenErr
	}
	if err := layer.IsValid(); err != nil {
		return err
	}
	pkt := NewPacket(COMMAND_TOGGLE_INDICATOR, byte(layer), index)
	return d.writeAll(pkt)
}

func (d *bdn9Device) EnableIndicator(layer pb.Layer, index uint8) error {
	if d.device == nil {
		return NotOpenErr
	}
	if err := layer.IsValid(); err != nil {
		return err
	}
	pkt := NewPacket(COMMAND_ENABLE_INDICATOR, byte(layer), index)
	return d.writeAll(pkt)
}

func (d *bdn9Device) DisableIndicator(layer pb.Layer, index uint8) error {
	if d.device == nil {
		return NotOpenErr
	}
	if err := layer.IsValid(); err != nil {
		return err
	}
	pkt := NewPacket(COMMAND_DISABLE_INDICATOR, byte(layer), index)
	return d.writeAll(pkt)
}

func (d *bdn9Device) ActivateLayer(layer pb.Layer) error {
	if d.device == nil {
		return NotOpenErr
	}
	if err := layer.IsValid(); err != nil {
		return err
	}
	pkt := NewPacket(COMMAND_ACTIVATE_LAYER, byte(layer))
	return d.writeAll(pkt)
}

func (d *bdn9Device) SetSpeed(speed uint8) error {
	if d.device == nil {
		return NotOpenErr
	}
	pkt := NewPacket(COMMAND_SET_SPEED, d.persist, speed)
	return d.writeAll(pkt)
}

func (d *bdn9Device) SetMuteStatus(muted bool) error {
	if d.device == nil {
		return NotOpenErr
	}
	pkt := NewPacket(COMMAND_SET_MUTE_STATUS, boolToByte(muted))
	return d.writeAll(pkt)
}

func (d *bdn9Device) EndCall() error {
	if d.device == nil {
		return NotOpenErr
	}
	pkt := NewPacket(COMMAND_END_CALL)
	return d.writeAll(pkt)
}

func (d *bdn9Device) Reset() error {
	if d.device == nil {
		return NotOpenErr
	}
	// magic bytes required to reset
	pkt := NewPacket(COMMAND_RESET, 0xDE, 0xAD, 0xF0, 0x00)
	return d.writeAll(pkt)
}

func (d *bdn9Device) Echo(b []byte) error {
	if d.device == nil {
		return NotOpenErr
	}
	pkt := NewPacket(COMMAND_ECHO, 0x1, 0x2, 0x3)
	return d.writeAll(pkt)
}

func (d *bdn9Device) writeAll(pkt *bdn9Packet) error {
	written := 0
	total := pkt.Len()
	buf := pkt.Buffer()
	for {
		n, err := d.device.Write(buf[written:])
		if err != nil {
			return err
		}
		written += n
		if written >= total {
			break
		}
	}
	return nil
}

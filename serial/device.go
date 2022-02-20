package serial

import (
	"fmt"
	"time"

	"go.bug.st/serial"
	"go.bug.st/serial/enumerator"
)

var NotOpenErr = fmt.Errorf("Device not open")
var DeviceNotFound = fmt.Errorf("No BDN9 device found")

type BDN9SerialDevice interface {
	Open(port string) error
	IsOpen() bool
	Close() error
	Name() string
	SetRGBMode(mode RGBMode) error
	SetMatrixHSV(h uint8, s uint8, v uint8) error
	ToggleMatrix() error
	SetIndicatorHSV(layer Layer, index uint8, h uint8, s uint8, v uint8) error
	ToggleIndicator(layer Layer, index uint8) error
	EnableIndicator(layer Layer, index uint8) error
	DisableIndicator(layer Layer, index uint8) error
	ActivateLayer(layer Layer) error
	SetSpeed(speed uint8) error
	SetMuteStatus(muted bool) error
	EndCall() error
	Reset() error
	Echo(b []byte) error
	EnablePersist()
	DisablePersist()
}

type bdn9SerialDevice struct {
	port    serial.Port
	persist byte
	name    string
	handler EventHandler
}

const (
	VID = "cb10"
	PID = "2133"
)

func boolToByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

func NewDevice(handler EventHandler) BDN9SerialDevice {
	return &bdn9SerialDevice{port: nil, persist: 0, name: "", handler: handler}
}

func FindPort() (string, error) {
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		return "", err
	}
	for _, port := range ports {
		if !port.IsUSB {
			continue
		}
		if port.VID == VID && port.PID == PID {
			return port.Name, nil
		}
	}
	return "", DeviceNotFound
}

func GetDevices() ([]string, error) {
	return serial.GetPortsList()
}

func handleEvents(d *bdn9SerialDevice) {
	buf := make([]byte, 16)
	for {
		// blocks until bytes are read
		sz, err := d.port.Read(buf)
		if err != nil {
			fmt.Printf("Error reading: %s\n", err)
			d.port.Close()
			d.port = nil
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
		// fmt.Printf("cmd: %v\n", buf)
	}
}

func (d *bdn9SerialDevice) Open(port string) (err error) {
	mode := &serial.Mode{
		BaudRate: 9600, // TODO: configurable
	}
	d.port, err = serial.Open(port, mode)
	if err == nil && d.handler != nil {
		go handleEvents(d)
	}
	d.name = port
	return
}

func (d *bdn9SerialDevice) IsOpen() bool {
	return d.port != nil
}

func (d *bdn9SerialDevice) Close() error {
	err := d.port.Close()
	d.port = nil
	return err
}

func (d *bdn9SerialDevice) Name() string {
	return d.name
}

func (d *bdn9SerialDevice) EnablePersist() {
	d.persist = 1
}

func (d *bdn9SerialDevice) DisablePersist() {
	d.persist = 0
}

func (d *bdn9SerialDevice) SetRGBMode(mode RGBMode) error {
	if d.port == nil {
		return NotOpenErr
	}
	if err := mode.IsValid(); err != nil {
		return err
	}
	pkt := NewPacket(COMMAND_SET_MODE, d.persist, byte(mode))
	return d.writeAll(pkt)
}

func (d *bdn9SerialDevice) SetMatrixHSV(h uint8, s uint8, v uint8) error {
	if d.port == nil {
		return NotOpenErr
	}
	pkt := NewPacket(COMMAND_SET_MATRIX_HSV, d.persist, h, s, v)
	return d.writeAll(pkt)
}

func (d *bdn9SerialDevice) ToggleMatrix() error {
	if d.port == nil {
		return NotOpenErr
	}
	pkt := NewPacket(COMMAND_TOGGLE_MATRIX, d.persist)
	return d.writeAll(pkt)
}

func (d *bdn9SerialDevice) SetIndicatorHSV(layer Layer, index uint8, h uint8, s uint8, v uint8) error {
	if d.port == nil {
		return NotOpenErr
	}
	if err := layer.IsValid(); err != nil {
		return err
	}
	pkt := NewPacket(COMMAND_SET_INDICATOR_HSV, byte(layer), index, h, s, v)
	return d.writeAll(pkt)
}

func (d *bdn9SerialDevice) ToggleIndicator(layer Layer, index uint8) error {
	if d.port == nil {
		return NotOpenErr
	}
	if err := layer.IsValid(); err != nil {
		return err
	}
	pkt := NewPacket(COMMAND_TOGGLE_INDICATOR, byte(layer), index)
	return d.writeAll(pkt)
}

func (d *bdn9SerialDevice) EnableIndicator(layer Layer, index uint8) error {
	if d.port == nil {
		return NotOpenErr
	}
	if err := layer.IsValid(); err != nil {
		return err
	}
	pkt := NewPacket(COMMAND_ENABLE_INDICATOR, byte(layer), index)
	return d.writeAll(pkt)
}

func (d *bdn9SerialDevice) DisableIndicator(layer Layer, index uint8) error {
	if d.port == nil {
		return NotOpenErr
	}
	if err := layer.IsValid(); err != nil {
		return err
	}
	pkt := NewPacket(COMMAND_DISABLE_INDICATOR, byte(layer), index)
	return d.writeAll(pkt)
}

func (d *bdn9SerialDevice) ActivateLayer(layer Layer) error {
	if d.port == nil {
		return NotOpenErr
	}
	if err := layer.IsValid(); err != nil {
		return err
	}
	pkt := NewPacket(COMMAND_ACTIVATE_LAYER, byte(layer))
	return d.writeAll(pkt)
}

func (d *bdn9SerialDevice) SetSpeed(speed uint8) error {
	if d.port == nil {
		return NotOpenErr
	}
	pkt := NewPacket(COMMAND_SET_SPEED, d.persist, speed)
	return d.writeAll(pkt)
}

func (d *bdn9SerialDevice) SetMuteStatus(muted bool) error {
	if d.port == nil {
		return NotOpenErr
	}
	pkt := NewPacket(COMMAND_SET_MUTE_STATUS, boolToByte(muted))
	return d.writeAll(pkt)
}

func (d *bdn9SerialDevice) EndCall() error {
	if d.port == nil {
		return NotOpenErr
	}
	pkt := NewPacket(COMMAND_END_CALL)
	return d.writeAll(pkt)
}

func (d *bdn9SerialDevice) Reset() error {
	if d.port == nil {
		return NotOpenErr
	}
	// magic bytes required to reset
	pkt := NewPacket(COMMAND_RESET, 0xDE, 0xAD, 0xF0, 0x00)
	return d.writeAll(pkt)
}

func (d *bdn9SerialDevice) Echo(b []byte) error {
	if d.port == nil {
		return NotOpenErr
	}
	pkt := NewPacket(COMMAND_ECHO, 0x1, 0x2, 0x3)
	return d.writeAll(pkt)
}

func (d *bdn9SerialDevice) writeAll(pkt *bdn9Packet) error {
	written := 0
	total := pkt.Len()
	buf := pkt.Buffer()
	for {
		n, err := d.port.Write(buf[written:])
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

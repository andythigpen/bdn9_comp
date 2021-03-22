package serial

import (
	"fmt"

	"go.bug.st/serial"
)

var NotOpenErr = fmt.Errorf("Device not open")

type BDN9SerialDevice interface {
	Open(port string) error
	IsOpen() bool
	Close() error
	SetRGBMode(mode RGBMode) error
	SetMatrixHSV(h uint8, s uint8, v uint8) error
	ToggleMatrix() error
	SetIndicatorHSV(index uint8, h uint8, s uint8, v uint8) error
	ToggleIndicator(index uint8) error
	EnableIndicator(index uint8) error
	DisableIndicator(index uint8) error
	ActivateLayer(layer Layer) error
	SetSpeed(speed uint8) error
	SetMuteStatus(muted bool) error
	EndCall() error
	Reset() error
	EnablePersist()
	DisablePersist()
}

type bdn9SerialDevice struct {
	port    serial.Port
	persist byte
}

func boolToByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

func NewDevice() BDN9SerialDevice {
	return &bdn9SerialDevice{port: nil, persist: 0}
}

func (d *bdn9SerialDevice) Open(port string) (err error) {
	mode := &serial.Mode{
		BaudRate: 9600, // TODO: configurable
	}
	d.port, err = serial.Open(port, mode)
	return
}

func (d *bdn9SerialDevice) IsOpen() bool {
	return d.port != nil
}

func (d *bdn9SerialDevice) Close() error {
	return d.port.Close()
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

func (d *bdn9SerialDevice) SetIndicatorHSV(index uint8, h uint8, s uint8, v uint8) error {
	if d.port == nil {
		return NotOpenErr
	}
	pkt := NewPacket(COMMAND_SET_INDICATOR_HSV, index, h, s, v)
	return d.writeAll(pkt)
}

func (d *bdn9SerialDevice) ToggleIndicator(index uint8) error {
	if d.port == nil {
		return NotOpenErr
	}
	pkt := NewPacket(COMMAND_TOGGLE_INDICATOR, index)
	return d.writeAll(pkt)
}

func (d *bdn9SerialDevice) EnableIndicator(index uint8) error {
	if d.port == nil {
		return NotOpenErr
	}
	pkt := NewPacket(COMMAND_ENABLE_INDICATOR, index)
	return d.writeAll(pkt)
}

func (d *bdn9SerialDevice) DisableIndicator(index uint8) error {
	if d.port == nil {
		return NotOpenErr
	}
	pkt := NewPacket(COMMAND_DISABLE_INDICATOR, index)
	return d.writeAll(pkt)
}

func (d *bdn9SerialDevice) ActivateLayer(layer Layer) error {
	if d.port == nil {
		return NotOpenErr
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

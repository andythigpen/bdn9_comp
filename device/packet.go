package device

type bdn9Packet struct {
	cmd Command
	buf []byte
}

func NewPacket(cmd Command, b ...byte) *bdn9Packet {
	return &bdn9Packet{
		cmd: cmd,
		buf: b,
	}
}

func (p *bdn9Packet) Len() int {
	// command + buf
	return len(p.buf) + 1
}

func (p *bdn9Packet) Add(b ...byte) {
	p.buf = append(p.buf, b...)
	// packets to the raw HID interface must be exactly 32 bytes
	if len(p.buf) > 31 {
		p.buf = p.buf[:31]
	}
}

func (p *bdn9Packet) Buffer() []byte {
	// packets to the raw HID interface must be exactly 32 bytes
	b := make([]byte, 32)
	b[0] = byte(p.cmd)
	for i := 0; i < len(p.buf); i++ {
		b[i+1] = p.buf[i]
	}
	return b
}

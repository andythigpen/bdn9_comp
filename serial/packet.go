package serial

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
	return len(p.buf)
}

func (p *bdn9Packet) Add(b ...byte) {
	p.buf = append(p.buf, b...)
}

func (p *bdn9Packet) Buffer() []byte {
	b := []byte{byte(p.cmd), byte(p.Len())}
	return append(b, p.buf...)
}

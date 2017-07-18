package netflow5

import (
	"io"
	"net"
)

type Decoder struct {
}

func NewDecoder() *Decoder {
	return &Decoder{}
}

func (d *Decoder) Read(r io.Reader) error {
	return nil
}

func Read(r io.Reader, routerAddr *net.UDPAddr) (*Packet, error) {
	p := new(Packet)
	return p, p.Unmarshal(r, routerAddr)
}

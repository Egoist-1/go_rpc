package codec

import (
	"bufio"
	"encoding/gob"
	"io"
)

var _ Codec = (*GobCode)(nil)

func NewGobCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)
	return &GobCode{
		conn: conn,

		buf: buf,
		dec: gob.NewDecoder(conn),
		enc: gob.NewEncoder(buf),
	}
}

type GobCode struct {
	conn io.ReadWriteCloser
	//防止阻塞而创建的带缓冲的writer
	buf *bufio.Writer
	dec *gob.Decoder
	enc *gob.Encoder
}

func (g *GobCode) Close() error {
	return g.conn.Close()
}

func (g *GobCode) ReadHeader(header *Header) error {
	return g.dec.Decode(header)
}

func (g *GobCode) ReadBody(body any) error {
	return g.dec.Decode(body)
}

func (g *GobCode) Write(header *Header, body any) (err error) {
	defer func() {
		_ = g.buf.Flush()
		if err != nil {
			_ = g.Close()
		}
	}()
	err = g.enc.Encode(header)
	if err != nil {
		return err
	}
	err = g.enc.Encode(body)
	return err
}

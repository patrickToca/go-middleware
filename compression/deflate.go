package compression

import (
	"compress/flate"
	"io"
)

type DeflateCompressor struct {
	Level int
}

func (c *DeflateCompressor) AcceptEncoding() string {
	return "deflate"
}

func (c *DeflateCompressor) ContentEncoding() string {
	return "deflate"
}

func (c *DeflateCompressor) NewReader(r io.Reader) (io.ReadCloser, error) {
	return flate.NewReader(r), nil
}

func (c *DeflateCompressor) NewWriter(w io.Writer) (io.WriteCloser, error) {
	return flate.NewWriter(w, c.Level)
}

func NewDeflateCompressor(level int) *DeflateCompressor {
	return &DeflateCompressor{Level: level}
}

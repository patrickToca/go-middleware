package compression

import (
	"compress/gzip"
	"io"
)

type GzipCompressor struct{}

func (c *GzipCompressor) AcceptEncoding() string {
	return "gzip"
}

func (c *GzipCompressor) ContentEncoding() string {
	return "gzip"
}

func (c *GzipCompressor) NewReader(r io.Reader) (io.ReadCloser, error) {
	return gzip.NewReader(r)
}

func (c *GzipCompressor) NewWriter(w io.Writer) (io.WriteCloser, error) {
	return gzip.NewWriter(w), nil
}

func NewGzipCompressor() *GzipCompressor {
	return &GzipCompressor{}
}

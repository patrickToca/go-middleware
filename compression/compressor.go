package compression

import (
	"io"
)

type Compressor interface {
	AcceptEncoding() string
	ContentEncoding() string
	NewReader(io.Reader) (io.ReadCloser, error)
	NewWriter(io.Writer) (io.WriteCloser, error)
}

type RequestDecompressor interface {
	Compressor
	Decompress() bool
}

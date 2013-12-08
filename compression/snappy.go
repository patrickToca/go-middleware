package compression

import (
	"bytes"
	"code.google.com/p/snappy-go/snappy"
	"io"
	"io/ioutil"
)

type SnappyCompressor struct{}

func (c *SnappyCompressor) AcceptEncoding() string {
	return "snappy"
}

func (c *SnappyCompressor) ContentEncoding() string {
	return "snappy"
}

func (c *SnappyCompressor) NewReader(r io.Reader) (io.ReadCloser, error) {
	src, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	dst, err := snappy.Decode(nil, src)
	if err != nil {
		return nil, err
	}

	return ioutil.NopCloser(bytes.NewReader(dst)), nil
}

func (c *SnappyCompressor) NewWriter(w io.Writer) (io.WriteCloser, error) {
	return &bufferedSnappyWriter{b: bytes.NewBuffer(make([]byte, 0, 128)), w: w}, nil
}

func (c *SnappyCompressor) Decompress() bool {
	return true
}

func NewSnappyCompressor() *SnappyCompressor {
	return &SnappyCompressor{}
}

type bufferedSnappyWriter struct {
	b *bytes.Buffer
	w io.Writer
}

func (w bufferedSnappyWriter) Write(p []byte) (n int, err error) {
	return w.b.Write(p)
}

func (w *bufferedSnappyWriter) Close() error {
	if w.b.Len() == 0 {
		return nil
	}

	_, err := w.w.Write(w.b.Bytes())
	w.b.Reset()
	return err
}

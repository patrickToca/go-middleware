package compression

import (
	"bytes"
	"io"
	"net/http"
	"strings"
)

// CompressionWriter acts as a buffered ResponseWriter that conditionally applies compression
type compressionWriter struct {
	http.ResponseWriter
	r *http.Request
	h *Handler
	b *bytes.Buffer
}

func newCompressionWriter(w http.ResponseWriter, r *http.Request, h *Handler) *compressionWriter {
	return &compressionWriter{ResponseWriter: w, r: r, b: bytes.NewBuffer(make([]byte, 0, 1024))}
}

func (w *compressionWriter) Write(b []byte) (n int, err error) {
	return w.b.Write(b)
}

func (w *compressionWriter) Close() (err error) {
	if w.b.Len() == 0 {
		return
	}

	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", http.DetectContentType(w.b.Bytes()))
	}

	accept := w.r.Header.Get("Accept-Encoding")

	write := func(writer io.Writer) {
		_, err = writer.Write(w.b.Bytes())
		w.b.Reset()
	}

	if w.h.Compress(w.Header().Get("Content-Type"), w.b.Len()) == false || accept == "" {
		write(w.ResponseWriter)
		return
	}

	var (
		accepted   []string = strings.Split(accept, ",")
		ok         bool
		compressor Compressor
		writer     io.WriteCloser
	)

	for _, acceptedEncoding := range accepted {
		acceptedEncoding = lowerTrim(acceptedEncoding)
		if compressor, ok = w.h.Compressors[acceptedEncoding]; !ok {
			continue
		}

		if writer, err = compressor.NewWriter(w.ResponseWriter); err != nil {
			return
		}
		defer writer.Close()

		w.Header().Add("Vary", "Accept-Encoding")
		w.Header().Set("Content-Encoding", compressor.ContentEncoding())

		write(writer)
		return
	}

	// No compressor found, write plain
	write(w.ResponseWriter)
	return
}

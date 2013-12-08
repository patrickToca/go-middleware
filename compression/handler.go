package compression

import (
	"compress/flate"
	"net/http"
	"strings"
)

// See: https://github.com/codegangsta/martini-contrib/blob/master/gzip/gzip.go

type Handler struct {
	Compressors   map[string]Compressor
	Decompressors map[string]RequestDecompressor
	Compress      func(string, int) bool
}

func (h *Handler) Process(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// Decompress the Request Body if required
	contentEncoding := lowerTrim(r.Header.Get("Content-Encoding"))
	if decompressor, ok := h.Decompressors[contentEncoding]; ok && decompressor.Decompress() {
		if reader, err := decompressor.NewReader(r.Body); err == nil {
			r.Body = reader
			defer reader.Close()
		}
	}

	wr := newCompressionWriter(w, r, h)
	defer wr.Close()

	next(wr, r)
}

func defaultCompress(contentType string, length int) bool {
	return length > 75 && !strings.HasPrefix(contentType, "image")
}

func NewHandler(compressors ...Compressor) *Handler {
	if compressors == nil || len(compressors) == 0 {
		compressors = []Compressor{
			NewGzipCompressor(),
			NewDeflateCompressor(flate.DefaultCompression),
			NewSnappyCompressor(),
		}
	}

	m := make(map[string]Compressor, 0)
	d := make(map[string]RequestDecompressor, 0)
	for _, c := range compressors {
		m[c.AcceptEncoding()] = c
		if cd, ok := c.(RequestDecompressor); ok {
			d[cd.ContentEncoding()] = cd
		}
	}

	return &Handler{Compressors: m, Compress: defaultCompress, Decompressors: d}
}

func lowerTrim(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

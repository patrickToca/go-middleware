package cors

import (
	"net/http"
)

// See: https://github.com/ServiceStack/ServiceStack/blob/6304fd1bb8e8f9fd604c5ee03933a1cc406819e1/src/ServiceStack/CorsFeature.cs
// See: https://github.com/ServiceStack/ServiceStack/blob/5d09d439cd1a13712411552e2b3ede5a71af2ee5/src/ServiceStack/EnableCorsAttribute.cs
// See: https://github.com/streadway/handy/blob/master/cors/cors.go

type CorsOptions struct {
	AllowOrigin      []string
	AllowCredentials bool
	ExposeHeaders    []string
	AllowMethods     []string
	AllowHeaders     []string
}

func NewCorsOptions() *CorsOptions {
	return &CorsOptions{
		AllowOrigin:      make([]string, 0),
		AllowCredentials: false,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Accept-Encoding", "Authorization", "Content-Type", "Origin"},
	}
}

type Handler struct {
	*CorsOptions
	EnableCors bool
}

func (h *Handler) Process(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if h.CorsOptions == nil {
		h.CorsOptions = NewCorsOptions()
	}

	if h.EnableCors {
		EnableCors(w, r, h.CorsOptions)
	}

	next(w, r)
}

func NewHandler(options *CorsOptions) *Handler {
	return &Handler{CorsOptions: options, EnableCors: true}
}

func NewDefaultHandler() *Handler {
	return NewHandler(NewCorsOptions())
}

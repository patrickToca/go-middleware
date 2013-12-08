package cors

import (
	"net/http"
	"strings"
)

func EnableCors(w http.ResponseWriter, r *http.Request, options *CorsOptions) {
	if len(options.AllowOrigin) == 0 {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	} else {
		origin := r.Header.Get("Origin")
		if len(origin) > 0 && containsString(options.AllowOrigin, origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
	}

	if len(w.Header().Get("Access-Control-Allow-Origin")) == 0 {
		return
	}

	if options.AllowCredentials {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	}

	if len(options.ExposeHeaders) > 0 {
		for _, header := range options.ExposeHeaders {
			if len(header) > 0 {
				w.Header().Add("Access-Control-Expose-Headers", header)
			}
		}
	}

	if len(options.AllowMethods) > 0 {
		for _, method := range options.AllowMethods {
			if len(method) > 0 {
				w.Header().Add("Access-Control-Allow-Methods", strings.ToUpper(method))
			}
		}
	}

	if len(options.AllowHeaders) > 0 {
		for _, header := range options.AllowHeaders {
			if len(header) > 0 {
				w.Header().Add("Access-Control-Allow-Headers", header)
			}
		}
	} else if len(r.Header.Get("Access-Control-Request-Headers")) > 0 {
		w.Header().Set("Access-Control-Allow-Headers", "")
	}
}

// Helpers

func containsString(arr []string, str string) bool {
	if arr == nil || len(arr) == 0 {
		return false
	}
	for _, s := range arr {
		if s == str {
			return true
		}
	}
	return false
}

package global_error_handling

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// See https://github.com/pilu/traffic for including source issue location **
// See https://github.com/astaxie/beego/blob/master/middleware/error.go

type GetErrorTextFunc func(recovered interface{}) string

type Handler struct {
	IsDebug      bool
	GetErrorText GetErrorTextFunc
}

func (h *Handler) Process(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		if m := recover(); m != nil {
			if h.GetErrorText == nil {
				h.GetErrorText = DefaultGetErrorText
			}

			err := h.GetErrorText(m)

			if !h.IsDebug {
				err = http.StatusText(http.StatusInternalServerError)
			}

			http.Error(w, err, http.StatusInternalServerError)
		}
	}()

	next(w, r)
}

func NewHandler(isDebug bool, getErrorText GetErrorTextFunc) *Handler {
	return &Handler{IsDebug: isDebug, GetErrorText: getErrorText}
}

func NewDefaultHandler() *Handler {
	return NewHandler(false, DefaultGetErrorText)
}

// Helpers

func DefaultGetErrorText(m interface{}) string {
	switch v := m.(type) {
	case string:
		return fmt.Sprintf("PANIC: %s\n%s", v, debug.Stack())
	case error:
		return fmt.Sprintf("PANIC: %s\n%s", v.Error(), debug.Stack())
	default:
		return fmt.Sprintf("PANIC: %v\n%s", v, debug.Stack())
	}
}

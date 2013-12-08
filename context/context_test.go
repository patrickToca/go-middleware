package context

import (
	"github.com/shelakel/go-middleware"
	"net/http"
	"testing"
)

func TestRequestLifecycle(t *testing.T) {
	r := new(http.Request)
	valid := true
	// Initialize the Request Context, Set a Key in middleware 1 and then ensure it can be read in middleware 2
	middleware.Compose(Initialize(),
		func(w http.ResponseWriter, r *http.Request) {
			Set(r, "test", "Test")
		},
		func(w http.ResponseWriter, r *http.Request) {
			var (
				val interface{}
				v   string
				ok  bool
			)

			if val, ok = GetOk(r, "test"); !ok {
				valid = false
				return
			}

			if v, ok = val.(string); !ok || v != "Test" {
				valid = false
			}
		})(nil, r)

	if !valid {
		t.Fatalf("Expected values to be passable between middleware")
		return
	}

	defer func() {
		recover() // Recovered from the expected panic caused by accessing a Request Context that isn't initialized
	}()

	// This should cause a panic because the Request Context isn't initialized (only the case in a middleware chain)
	Get(r, "test")

	// This should not be called
	t.Fatalf("Expected Get(r, \"test\") to cause a panic as Request Context isn't initialized")
}

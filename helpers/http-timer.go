package helpers

import (
	"fmt"
	"net/http"
	"time"
)

func TimedHandler(w http.ResponseWriter, r *http.Request, f func(w http.ResponseWriter, r *http.Request)) {
	now := time.Now()

	f(w, r)

	elapsed := time.Since(now)
	fmt.Printf("%s - %s - %s\n", r.Method, r.URL.Path, elapsed.String())
}

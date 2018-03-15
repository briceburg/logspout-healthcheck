package healthcheck

import (
	"net/http"
	"github.com/gliderlabs/logspout/router"
)

func init() {
	router.HttpHandlers.Register(HealthCheck, "healthcheck")
}

// HealthCheck returns a http.Handler for the health check
func HealthCheck() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthcheck", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("OK\n"))
	})
	return mux
}

package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/DanielHilton/go-amqp-consumer/services"
	"net/http"
)

func StartHttpServer(port int) {
	http.HandleFunc("/sample", func(w http.ResponseWriter, r *http.Request) {
		sample, _ := services.GetSample()
		b, err := json.Marshal(sample)
		if err != nil {
			fmt.Errorf("failed to marshal samples %w", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, string(b))
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	ExitOnFail(err, "Failed to start HTTP server")
}

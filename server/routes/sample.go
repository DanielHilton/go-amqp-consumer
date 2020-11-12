package routes

import (
	"encoding/json"
	"fmt"
	"github.com/DanielHilton/go-amqp-consumer/db"
	"github.com/DanielHilton/go-amqp-consumer/helpers"
	"net/http"
)

func getSample(w http.ResponseWriter, r *http.Request) {
	sample, _ := db.GetSample()
	b, err := json.Marshal(sample)
	if err != nil {
		fmt.Errorf("failed to marshal samples %w", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(b))
}

func GetSampleRoute(w http.ResponseWriter, r *http.Request) {
	helpers.TimedHandler(w, r, getSample)
}

package server

import (
	"fmt"
	"github.com/DanielHilton/go-amqp-consumer/helpers"
	"github.com/DanielHilton/go-amqp-consumer/services/server/routes"
	"net/http"
)

func StartHttpServer(port int) {
	http.HandleFunc("/sample", routes.GetSampleRoute)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	helpers.ExitOnFail(err, "Failed to start HTTP server")
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/senseyeio/roger"
)

func main() {
	// connect to RServe using Roger
	rClient, err := roger.NewRClient("127.0.0.1", 6311)
	if err != nil {
		fmt.Printf("Failed to connect to RServe: %s", err.Error())
		return
	}

	// define function to run when API is called
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// call generateCorrelationPlot R function, gathering the response
		returnVar, err := rClient.Eval("generateCorrelationPlot()")
		if err != nil {
			fmt.Fprintf(w, "Graph generation failed with error %s", err.Error())
			return
		}

		// convert response to a byte array
		imageBytes, ok := returnVar.([]byte)
		if !ok {
			fmt.Fprint(w, "Unexpected response from R")
			return
		}

		// return the binary image data along with suitable headers
		w.Header().Set("Content-Type", "image/png")
		w.Header().Set("Content-Length", strconv.Itoa(len(imageBytes)))
		if _, err := w.Write(imageBytes); err != nil {
			fmt.Fprint(w, "Failed to write image to response")
		}
	})

	// listen for HTTP calls
	log.Fatal(http.ListenAndServe(":8080", nil))
}

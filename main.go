package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if (len(os.Args) != 3) {
		os.Stderr.WriteString("Usage: " + os.Args[0] + 
				"<listen-addr> <health endpoint>\n")
		os.Exit(1)
	}

	listenAddr := os.Args[1]	// e.g. :8080
	healthPath := os.Args[2]	// e.g. /health_check

	healthy := func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		io.WriteString(w, "{}")
	}

	// one explicit endpoint
	http.HandleFunc(healthPath, healthy)

	// everything else gets an error
	http.Handle("/", http.NotFoundHandler())

	err := http.ListenAndServe(listenAddr, nil)
	if (err != nil) {
		fmt.Println(err)
		os.Exit(2)
	}
}

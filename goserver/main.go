package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello go server....!")
	})

	fmt.Println("running...")
	log.Fatal(http.ListenAndServe(":8081", mux))
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wzshiming/ctc"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Printf("%s%s%s: %s\n", ctc.ForegroundGreen, name, ctc.Reset, h)
		}
	}

	fmt.Printf("IP Address: %s\n", r.RemoteAddr)
	fmt.Fprint(w, "ok")
}

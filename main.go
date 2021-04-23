package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

type EchoDump struct{}

func (echoDump EchoDump) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println(string(dump))

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(dump)
}

func run() error {
	addr := os.Getenv("ECHODUMP_ADDR")
	if addr == "" {
		addr = ":10000"
	}
	log.Println("listening on " + addr)

	return http.ListenAndServe(addr, EchoDump{})
}

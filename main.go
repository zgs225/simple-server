package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	port = flag.Int("p", 8080, "Port of the server")
	root = flag.String("w", ".", "Root dir of the server")
)

func init() {
	flag.Parse()
}

func main() {
	a, err := filepath.Abs(*root)
	if err != nil {
		log.Fatal(err)
	}
	d := http.Dir(a)
	fs := http.FileServer(d)
	ls := &logServer{Next: fs, Logger: log.New(os.Stderr, "", log.LstdFlags)}
	log.Printf("Simple server on :%d...\n", *port)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", *port), ls))
}

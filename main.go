package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
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
	go func() {
		log.Println(http.ListenAndServe(fmt.Sprintf(":%d", *port), ls))
	}()
	time.Sleep(time.Second)
	url := fmt.Sprintf("http://localhost:%d", *port)
	openBrowser(url)
	// Keep your server running or perform other tasks
	select {}
}
func openBrowser(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	default:
		log.Println("Unsupported operating system:", runtime.GOOS)
		return
	}
	err := cmd.Start()
	if err != nil {
		log.Println("Failed to open browser:", err)
	}
}

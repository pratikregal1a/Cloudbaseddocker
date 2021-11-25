package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"image"
	"image/png"
)

func handler(w http.ResponseWriter, r *http.Request) {
	cmd := exec.CommandContext(r.Context(), "/bin/sh", "script.sh")
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		w.WriteHeader(500)
	}
	w.Write(out)
}

func main() {
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	{
      image1,err := os.Open("**/*.png")
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

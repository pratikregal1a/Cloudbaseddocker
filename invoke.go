package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
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
//  func getImageFromFilePath(bournvita.png string) (image.Image, error) {
//     f, err := os.Open(bournvita.png)
// //     if err != nil {
// //         return nil, err
// //     }
// // //     defer f.Close()
// // //     image, _, err := image.Decode(f)
// // //     return image, err
//  }
func pngHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    if !strings.HasSuffix(vars["item"], ".png") {
        http.Error(w, "Wrong File Extension", http.StatusNotFound)
        return
    }
    file, err = ioutil.ReadFile("img/" + vars["item"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    w.Header().Set("Content-type", "image/png")
    w.Write(file)
}
func main() {
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	
}


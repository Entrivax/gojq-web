package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"
)

//go:embed res
var res embed.FS

func main() {
	portEnv := os.Getenv("PORT")
	port, err := strconv.Atoi(portEnv)
	if err != nil {
		port = 8080
	}

	sRes, err := fs.Sub(res, "res")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", http.FileServer(http.FS(sRes)))

	fmt.Printf("Server started on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

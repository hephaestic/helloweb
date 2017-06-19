package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.PrintF("port:", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello, Web")
	})
	fmt.Printf("Hello, Web")
	http.ListenAndServe(":" + port, nil)
	fmt.Printf("listening and serving")
}



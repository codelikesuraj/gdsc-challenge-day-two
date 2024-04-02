package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT string = "3000"

func APIhandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"Hello GDSC"}`))
}

func HTMLhandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<strong>Hello GDSC</strong>"))
}

func main() {
	http.HandleFunc("/api/hello", APIhandler)
	http.HandleFunc("/hello", HTMLhandler)

	fmt.Printf("Server listening on %q\n", PORT)
	fmt.Println("Available routes:")
	fmt.Printf("\thttp://127.0.0.1:%s/hello\n", PORT)
	fmt.Printf("\thttp://127.0.0.1:%s/api/hello\n", PORT)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil))
}

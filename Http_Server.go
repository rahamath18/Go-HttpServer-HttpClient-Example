// Go - HTTP Server
// http://localhost:8090/
// http://localhost:8090/hello
// http://localhost:8090/headers

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	Name    string
	Address string
}

func root(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "{fine}\n")
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("hello invoked")
	var p Person

	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("Req body: ", p)

	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Person: %+v", p)

}

func headers(w http.ResponseWriter, req *http.Request) {
	fmt.Println("headers: ", req)
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/", headers)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}

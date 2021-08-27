// Go - HTTP Client

// http://localhost:8090/
// http://localhost:8090/hello
// http://localhost:8090/headers

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Person struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func main() {

	body := &Person{
		Name:    "Rahamath S",
		Address: "India",
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)
	req, _ := http.NewRequest("POST", "http://localhost:8090/hello", buf)

	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		log.Fatal(e)
	}

	defer res.Body.Close()

	fmt.Println("response Status:", res.Status)

	// Print the body to the stdout
	io.Copy(os.Stdout, res.Body)
}

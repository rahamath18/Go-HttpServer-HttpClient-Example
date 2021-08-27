## Go Programming Language - Examples - Simple HTTP Server & HTTP Client

The Go programming language is an open-source project to make programmers more productive.

Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.

https://golang.org/doc/

### Used technology stack in this example

1. go1.17.windows-amd64.msi - Download - https://golang.org/dl/
2. MS Visual Studio Code
3. Chrome Browser
4. POSTMAN

### Run GO HTTP Server

$ go run .\Http_Server.go

### Access the following POST URL from POSTMAN

POST http://localhost:8090/hello
Req JSON: 
{
    "name": "Rahamath S",
    "address": "India"
}

### Also, access the following GET URL from POSTMAN or Chrome browser

http://localhost:8090/

http://localhost:8090/headers

### Also, you can run the following http client go program

$ go run .\Http_Client.go

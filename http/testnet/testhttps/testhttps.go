// gohttps/2-https/server.go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w,
        "Hi, This is an example of https service in golang!")
}

func main() {
	http.HandleFunc("/", handler)
	//http.ListenAndServe(":8081", nil)
    http.ListenAndServeTLS(":8081", "myserver.com.crt","myserver.com.key", nil)
}
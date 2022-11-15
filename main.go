package main

// https://gist.github.com/denji/12b3a568f092ab951456

import (
	"fmt"
	"net/http"
	"log"
)

func HelloServer(w http.ResponseWriter, req *http.Request){
	//w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hello, World.\n")
}

func main(){
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	//http.ListenAndServe(":8080",nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


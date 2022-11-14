package main

import (
	"fmt"
	"net/http"
	//"log"
)

func HelloServer(w http.ResponseWriter, req *http.Request){
	//w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hello, World.\n")
}

func main(){
	http.HandleFunc("/hello", HelloServer)
	//err := http.ListenAndServerTLS(":443", "server.crt", "server.key", nil)
	http.ListenAndServe(":8080",nil)
	/*
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	*/
}


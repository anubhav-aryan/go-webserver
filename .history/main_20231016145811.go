package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", formHandler)
	http
}
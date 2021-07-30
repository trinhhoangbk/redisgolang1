package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)
func main(){
	 router := mux.NewRouter()
	const port string  = ":8000"
	 router.HandleFunc("/", func(respone http.ResponseWriter, request *http.Request) {
		 fmt.Fprintln(respone,"hello trinh hoang")
	 })
	log.Println("server listening on port",port)
	 log.Fatalln(http.ListenAndServe(port,router))
}
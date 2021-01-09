package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"welcome to the club")
	fmt.Println("endpoint hit: homepage")
}
func handleRequest()  {
	http.HandleFunc("/",homePage)
	log.Fatalln(http.ListenAndServe(":5555",nil))
}

func main()  {
	handleRequest()
}
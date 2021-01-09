package main

import (
	"log"
	"net/http"
)

/*func homePage(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"welcome to the club")
	fmt.Println("endpoint hit: homepage")
}
func handleRequest()  {
	http.HandleFunc("/",homePage)
	log.Fatalln(http.ListenAndServe(":5555",nil))
}
*/
func middlewareOne(next http.Handler) http.Handler  {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println("executing middlewareOne")
		next.ServeHTTP(writer,request)
		log.Println("executing middlewareone again")
	})
}

func middlewareTwo(next http.Handler) http.Handler  {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Executing middlewareTwo")
		if request.URL.Path=="/foo"{
			return
		}
		next.ServeHTTP(writer,request)
		log.Println("executing middlewaretwoagain")
	})
}

func final(w http.ResponseWriter, r *http.Request)  {
	log.Println("executing finalHandler")
	w.Write([]byte("ok"))
}
func main()  {
	//handleRequest()
	mux:= http.NewServeMux()
	finalHandler := http.HandlerFunc(final)
	mux.Handle("/",middlewareOne(middlewareTwo(finalHandler)))

	log.Println("listening on 3000....")
	err:= http.ListenAndServe(":3000",mux)
	log.Fatal(err)
}
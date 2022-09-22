package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/kjais1720/bazel-monorepo/projects/go_hello"
)

func homeRouteHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	log.Println("Recieved request")
	w.Write([]byte(go_hello.HelloWorld()))
}

func main() {
    r := mux.NewRouter()
    // IMPORTANT: you must specify an OPTIONS method matcher for the middleware to set CORS headers
    r.HandleFunc("/", homeRouteHandler).Methods(http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions)
    r.Use(mux.CORSMethodMiddleware(r))
		log.Println("Going to listen on port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
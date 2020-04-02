package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/tobiade/go-exercises/urlshort"
)

var mux = defaultMux()

func main() {
	filePath := flag.String("yaml-path", "", "path to yaml file containing url mappings")
	flag.Parse()

	yamlHandler := yamlHandler(filePath)
	jsonHandler := jsonHandler()
	dbHandler := dbHandler()
	http.HandleFunc("/yaml/", yamlHandler)
	http.HandleFunc("/json/", jsonHandler)
	http.HandleFunc("/db/", dbHandler)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", nil)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func yamlHandler(filePath *string) http.HandlerFunc {
	mappings, err := urlshort.GetMappingsFromYAML(*filePath)
	check(err)
	return urlshort.MapHandler(mappings, mux)
}

func jsonHandler() http.HandlerFunc {
	mappings, err := urlshort.GetMappingsFromJSON()
	check(err)
	return urlshort.MapHandler(mappings, mux)
}

func dbHandler() http.HandlerFunc {
	db, err := urlshort.GetDBConnection()
	check(err)
	return urlshort.DBHandler(db, mux)
}

package main

import (
	"fmt"
	groupieTracker "groupieTracker/handlers"
	"net/http"
	
)

func main() {

	http.HandleFunc("/aboutUs/", groupieTracker.AboutUsHandler)
	http.HandleFunc("/artist/", groupieTracker.GetArtist)
	http.HandleFunc("/", groupieTracker.GetHandler)

	http.HandleFunc("/styleArtists.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/styleArtists.css")
	})
	http.HandleFunc("/styleArtist.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/styleArtist.css")
	})
	http.HandleFunc("/error/error.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "error/error.css")
	})
	http.HandleFunc("/aboutUs.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/aboutUs.css")
	})

	port := "localhost:8080"
	fmt.Printf("Server is working on http://" + port + "\n")
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {

		}
	}()
	// nil means DefaultServeMux is used
	groupieTracker.OpenBrowser("http://localhost:8080")
	select {}
}



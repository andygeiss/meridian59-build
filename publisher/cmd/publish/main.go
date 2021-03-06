package main

import (
	"log"
	"net/http"
	"os"

	"github.com/andygeiss/meridian59-build/publisher/internal/accounts"
	"github.com/andygeiss/meridian59-build/publisher/internal/god"
	"github.com/xi2/httpgzip"
)

func main() {
	// check directories
	if _, err := os.Stat("client"); err != nil {
		log.Printf("client directory cannot be found! %v", err)
	}
	if _, err := os.Stat("download"); err != nil {
		log.Printf("download directory cannot be found! %v", err)
	}
	if _, err := os.Stat("static"); err != nil {
		log.Printf("static directory cannot be found! %v", err)
	}
	// setup the client patcher
	http.Handle("/client/", http.StripPrefix("/client/", httpgzip.NewHandler(http.FileServer(http.Dir("client")), nil)))
	// setup download page
	http.Handle("/download/", http.StripPrefix("/download/", httpgzip.NewHandler(http.FileServer(http.Dir("download")), nil)))
	// setup api
	http.HandleFunc("/api/accounts/create", accounts.Create())
	http.HandleFunc("/api/accounts/online", accounts.Online())
	http.HandleFunc("/api/god/log", god.Log())
	// setup the web frontend
	http.Handle("/", httpgzip.NewHandler(http.FileServer(http.Dir("static")), nil))
	// start listening
	log.Printf("Start listening ...")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}

package tserver

import (
	"log"
	"net/http"
)

// route like Servlet request mapping
func route() {
	http.HandleFunc("/student", GetHandler)
	http.HandleFunc("/student/list", ListHandler)
	http.HandleFunc("/student/saveOrUpdate", saveOrUpdateHandler)
}

func startUpServer() {
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("http tserver start unsuccessfully with error", err)
	}
}

// init for http server
func init() {
	route()
	startUpServer()
}

func Start() {
}

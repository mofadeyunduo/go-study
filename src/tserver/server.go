package tserver

import (
	"log"
	"net/http"
	"sync"
)

// route like Servlet request mapping
func route() {
	http.HandleFunc("/student", GetHandler)
	http.HandleFunc("/student/list", ListHandler)
	http.HandleFunc("/student/saveOrUpdate", saveOrUpdateHandler)
}

func httpHandler(fn func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	wg.Add(2)
	go valid(wg)
	go logHttp(wg)
	wg.Wait()

	return fn
}

func valid(wg sync.WaitGroup)  {
	defer wg.Done()

}

func logHttp(wg sync.WaitGroup)  {
	defer wg.Done()
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

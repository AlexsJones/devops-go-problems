package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SeedJobs/devops-go-problems/bit.ly/problem"
	"github.com/gorilla/mux"
)

func main() {
	store := problem.NewStorer(2 * time.Minute)
	r := mux.NewRouter()
	r.SkipClean(true)
	r.HandleFunc("/get/{link}", func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		l, err := store.GetLink(vars["link"])
		if err != nil {
			http.Error(w, fmt.Sprintf("Link does not exist %s", vars["link"]), http.StatusNotFound)
			return
		}
		if l == nil {
			http.Error(w, "Service unimplemented", http.StatusNotFound)
			return
		}
		switch {
		case l.Expired():
			http.Error(w, fmt.Sprintf("Link has expired %s", vars["link"]), http.StatusNotFound)
		default:
			http.Redirect(w, req, fmt.Sprintf("http://%s", l.GetURL()), 301)
		}
	}).Methods("GET")
	r.HandleFunc("/add/{link}", func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		l, err := store.AddLink(vars["link"])
		if err != nil {
			http.Error(w, fmt.Sprintf("Link already exists %s", vars["link"]), http.StatusNotFound)
			return
		}
		if l == nil {
			http.Error(w, "Service unimplemented", http.StatusNotFound)
			return
		}
		fmt.Fprintln(w, "New link created for", l.GetURL())
		fmt.Fprintln(w, "Link is", l.GetShortLink())
	}).Methods("POST")
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Print("Listening on port 8000")
	log.Fatal(srv.ListenAndServe())
}

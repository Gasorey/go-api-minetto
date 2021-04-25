package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gasorey/go-api-minetto/core/beer"
	"github.com/gorilla/mux"
)

func main() {
	db, err := sql.Open("sqlite3", "data.beer.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	service := beer.NewService(db)

	r := mux.NewRouter()

	n := negroni.new(
		negroni.NewLogger(),
	)
	r.Handle("/v1/beer", n.With(
		negroni.Wrap(hello(service)),
	)).Method("GET", "OPTIONS")
	http.Handle("/", r)

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         "4000",
		Handler:      http.DefaultServeMux,
		ErrorLog:     logger,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func hello(service beer.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		all, _ := service.GetAll()
		for _, i := range all {
			fmt.Println(i)
		}
	})
}

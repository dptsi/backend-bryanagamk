package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		// handle returned error here.
		w.WriteHeader(500)
		_, err := w.Write([]byte("empty or invalid id"))
		if err != nil {
			return
		}
	}
}

type Person struct {
	id        int
	name      string
	character string
}

var listPerson = []Person{
	{
		id:        1,
		name:      "Slack",
		character: "Harry Potter",
	},
	{
		id:        2,
		name:      "Anna",
		character: "John Weasley",
	},
	{
		id:        3,
		name:      "Bob",
		character: "Malfoy",
	},
}

func main() {
	r := chi.NewRouter()
	r.Method("GET", "/", Handler(customHandler))
	if err := http.ListenAndServe(":8080", r); err != nil {
		return
	}
}

func customHandler(w http.ResponseWriter, r *http.Request) error {
	idQuery := r.URL.Query().Get("id")
	if idQuery == "" {
		return errors.New(idQuery)
	}

	id, err := strconv.Atoi(idQuery)
	if err != nil || id > 3 || id < 0 {
		return errors.New(idQuery)
	}

	_, err = w.Write([]byte(listPerson[id].name))
	if err != nil {
		return err
	}
	return nil
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ffiat/nostr"
	"github.com/gorilla/mux"
)

func main() {

	repository := Repository{
		db: make(map[string]*nostr.Event),
	}

	for i := 0; i < 21; i++ {

		p := &nostr.Event{
			Content: fmt.Sprintf("alice%d", i),
			Kind:    nostr.KindTextNote,
			PubKey:  fmt.Sprintf("npub%d", i),
		}

		repository.Store(p)
	}

	r := mux.NewRouter()

	handler := Handler{
		repository: repository,
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	r.HandleFunc("/", handler.IndexHandler)
	r.HandleFunc("/contact", handler.SearchProfile).Methods("GET")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln("There's an error with the server", err)
	}
}

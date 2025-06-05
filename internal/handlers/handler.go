package handlers

import (
	"cutURL/internal/storage"
	"cutURL/internal/urlshortener"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"strconv"
)

func CreateURLHandler(storage *storage.URLStorage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if !urlshortener.URLCheck(string(body)) {
			http.Error(w, "Invalid URL!", http.StatusBadRequest)
			return
		}
		newURL := urlshortener.RandURL(8)
		JSONForDb := urlshortener.CreateJson(string(body), newURL)
		storage.SetData(JSONForDb)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(newURL))
	}
}

func GetURLHandler(storage *storage.URLStorage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data := storage.GetData(intId)
		if data == "" {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		oldURL, _, err := urlshortener.ParseJSON(data)
		if err != nil {
			http.Error(w, "Got error while parsing JSON from database", http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, oldURL, http.StatusTemporaryRedirect)
	}
}

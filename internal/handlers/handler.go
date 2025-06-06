package handlers

import (
	"cutURL/internal/storage"
	"cutURL/internal/urlshortener"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
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
		u := urlshortener.NewURLBuilder(8)
		newURL := u.CreateURL()
		storage.SetData(string(body), u.StringID)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(newURL))
	}
}

func GetURLHandler(storage *storage.URLStorage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		newURL := chi.URLParam(r, "id")
		oldURL := storage.GetData(newURL)
		if oldURL == "" {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		http.Redirect(w, r, oldURL, http.StatusTemporaryRedirect)
	}
}

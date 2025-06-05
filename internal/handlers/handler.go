package handlers

import (
	"cutURL/internal/storage"
	"cutURL/internal/urlShortener"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"strconv"
)

func CreateUrlHandler(storage *storage.UrlStorage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if !urlShortener.UrlCheck(string(body)) {
			http.Error(w, "Invalid url!", http.StatusBadRequest)
			return
		}
		newUrl := urlShortener.RandUrl(8)
		jsonForDb := urlShortener.CreateJson(string(body), newUrl)
		storage.SetData(jsonForDb)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(newUrl))
	}
}

func GetUrlHandler(storage *storage.UrlStorage) func(http.ResponseWriter, *http.Request) {
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
		oldUrl, _, err := urlShortener.ParseJson(data)
		if err != nil {
			http.Error(w, "Got error whiparsing json from database", http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, oldUrl, http.StatusTemporaryRedirect)
	}
}

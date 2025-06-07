package handlers

import (
	"cutURL/internal/storage"
	"cutURL/internal/urlshortener"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func CreateURLHandler(storage *storage.URLStorage) func(c *gin.Context) {
	return func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if !urlshortener.URLCheck(string(body)) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
			return
		}

		u := urlshortener.NewURLBuilder(8)
		newURL := u.CreateURL()
		storage.SetData(string(body), u.StringID)

		c.String(http.StatusCreated, newURL)
	}
}

func GetURLHandler(storage *storage.URLStorage) func(c *gin.Context) {
	return func(c *gin.Context) {
		newURL := c.Param("id")
		oldURL := storage.GetData(newURL)

		if oldURL == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			return
		}
		http.Redirect(c.Writer, c.Request, oldURL, http.StatusTemporaryRedirect)
	}
}

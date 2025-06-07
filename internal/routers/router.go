package routers

import (
	"cutURL/internal/handlers"
	"cutURL/internal/storage"
	"github.com/gin-gonic/gin"
)

//func RouterNew() *chi.Mux {
//	r := chi.NewRouter()
//	s := storage.NewStorage()
//
//	r.Post("/", handlers.CreateURLHandler(s))
//	r.Get("/{id}", handlers.GetURLHandler(s))
//
//	return r
//}

func RouterNew() *gin.Engine {
	r := gin.Default()
	s := storage.NewStorage()

	r.POST("/", handlers.CreateURLHandler(s))
	r.GET("/:id", handlers.GetURLHandler(s))

	return r
}

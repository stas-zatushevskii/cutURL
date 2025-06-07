package routers

import (
	"cutURL/internal/handlers"
	"cutURL/internal/storage"
	"github.com/gin-gonic/gin"
)

func RouterNew(BaseURL string) *gin.Engine {
	r := gin.Default()
	s := storage.NewStorage()

	r.POST("/", handlers.CreateURLHandler(s, BaseURL))
	r.GET("/:id", handlers.GetURLHandler(s))

	return r
}

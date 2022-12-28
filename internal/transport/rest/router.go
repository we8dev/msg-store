package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/pokrovsky-io/msg-store/internal/usecase"
	"net/http"
	"strconv"
)

// TODO: Здесь добавим и настроим роутер, а веб-сервер уже в gin

// TODO: Добавить инъекцию логгера в роутер
func NewRouter(router *gin.Engine, uc usecase.Order) {
	// Options
	//handler.Use(gin.Logger())
	//handler.Use(gin.Recovery())

	// Routers
	router.GET("/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		order, _ := uc.Get(id)

		c.JSON(http.StatusOK, order)
	})

	// TODO Вызвать метод Get у usecase и передать ему параметры

}

package controllers

import (
	"github.com/GagulProject/go-whisky/internal/http"
	"github.com/gin-gonic/gin"
)

type WhiskyController struct{}

func NewWhiskyController() *WhiskyController {
	return &WhiskyController{}
}

func (c *WhiskyController) Path() string {
	return "/ping"
}

func (c *WhiskyController) Pong(ctx *gin.Context) {
	ctx.String(200, "pong")
	//c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func (c *WhiskyController) Register(router http.Router) {
	router.GET("", c.Pong)
}

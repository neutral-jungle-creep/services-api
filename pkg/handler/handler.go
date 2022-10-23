package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"services-front/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{}
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/show")
	{
		api.GET("/messages", h.messages)
	}
	return router
}

type ExceptResp struct {
	Message string `json:"message"`
}

func NewExceptResp(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, message)
}

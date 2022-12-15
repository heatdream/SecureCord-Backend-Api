package internal

import (
	route "ShelterChatBackend/Api/internal/routes";

	gin "github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/auth/signup", route.RegisterRoute)
	router.GET("/auth/login", route.LoginRoute)

	return router

}

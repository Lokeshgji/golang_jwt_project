package routes

import (
	"github.com/gin-gonic/gin"

	controller "github.com/lokeshgji/golang_jwt_project/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())
}

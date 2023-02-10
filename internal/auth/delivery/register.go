package delivery

import (
	"fmt"
	"github.com/22Fariz22/gophermart/internal/auth"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc auth.UseCase) {
	h := NewHandler(uc)

	authEndpoints := router.Group("")
	{
		authEndpoints.POST("/register", h.SignUp)
		authEndpoints.POST("/login", h.SignIn)
	}
	fmt.Println("user-register-RegisterHTTPEndpoints: ")
}

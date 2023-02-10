package delivery

import (
	"fmt"
	"github.com/22Fariz22/gophermart/internal/auth"
	"github.com/22Fariz22/gophermart/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	useCase auth.UseCase
	l       logger.Interface
}

func NewHandler(useCase auth.UseCase, l logger.Interface) *Handler {
	fmt.Println("auth-NewHandler")
	return &Handler{
		useCase: useCase,
		l:       l,
	}
}

type signInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *Handler) SignUp(c *gin.Context) {
	fmt.Println("auth-handler")

	inp := new(signInput)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.useCase.SignUp(c.Request.Context(), inp.Login, inp.Password); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	fmt.Println("handler-signUp-login&password:", inp.Login, inp.Password)
	c.Status(http.StatusOK)
}

type signInResponse struct {
	Token string `json:"token"`
}

func (h *Handler) SignIn(c *gin.Context) {
	inp := new(signInput)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	token, err := h.useCase.SignIn(c.Request.Context(), inp.Login, inp.Password)
	if err != nil {
		if err == auth.ErrUserNotFound {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
	c.JSON(http.StatusOK, signInResponse{Token: token})
}

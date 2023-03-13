package handler

import (
	"go-template/handler/middleware"
	"go-template/model"
	"go-template/sdk/apires"
	"go-template/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{
	userService service.UserService
}
func NewUserHandler(userService *service.UserService, app *gin.Engine)(UserHandler){
	h := UserHandler{userService: *userService}
	h.route(app)
	return h
}

func (h *UserHandler) route(app *gin.Engine){
	user := app.Group("/user")
	user.POST("/login", h.Login)
	user.POST("/register", h.Register)
	user.GET("/debug", middleware.UserAuth(), h.Debug)
}



func (h *UserHandler) Register(c *gin.Context) {
	ctx := c.Request.Context()

	req := model.CreateUserRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		apires.FailOrError(c, http.StatusBadRequest, "Bad req", err)
		return
	}
	// trim extra whitespace
	req.FName = strings.Join(strings.Fields(req.FName), " ")
	req.LName = strings.Join(strings.Fields(req.LName), " ")

	err = req.Validate()
	if err != nil {
		apires.FailOrError(c, http.StatusBadRequest, "bad request", err)
		return
	}

	response, errw := h.userService.Register(ctx, req)
	if errw != nil {
		apires.FailOrError(c, errw.Code(), "Register failed", errw.Errors()...)
		return
	}
	apires.Success(c, http.StatusCreated, "Register success", response)
}

func (h *UserHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	req := model.LoginUserRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		apires.FailOrError(c, http.StatusBadRequest, "Login failed", err)
		return
	}

	err = req.Validate()
	if err != nil {
		apires.FailOrError(c, http.StatusBadRequest, "bad request", err)
		return
	}

	response, errw := h.userService.Login(ctx, req)
	if errw != nil {
		apires.FailOrError(c, errw.Code(), "Login failed", errw.Errors()...)
		return
	}
	apires.Success(c, http.StatusOK, "Login success", response)
}

// middleware userauth debug
func (h UserHandler) Debug(c *gin.Context){
	res, ok := c.Get("user")
	if !ok {
		c.Status(http.StatusUnauthorized)
		return
	}
	userClaims := res.(model.UserClaims)

	c.JSON(http.StatusOK, userClaims)
}
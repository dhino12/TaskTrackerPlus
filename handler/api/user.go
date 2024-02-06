package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserAPI interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetUserTaskCategory(c *gin.Context)
}

type userAPI struct {
	userService service.UserService
}

func NewUserAPI(userService service.UserService) *userAPI {
	return &userAPI{userService}
}

func (u *userAPI) Register(c *gin.Context) {
	var user model.UserRegister

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	if user.Email == "" || user.Password == "" || user.Fullname == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("register data is empty"))
		return
	}

	var recordUser = model.User{
		Fullname: user.Fullname,
		Email:    user.Email,
		Password: user.Password,
	}

	recordUser, err := u.userService.Register(&recordUser)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}

	c.JSON(http.StatusCreated, model.NewSuccessResponse("register success"))
}

func (u *userAPI) Login(c *gin.Context) {
	// TODO: answer here
	var userLogin model.UserLogin
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		fmt.Println("masuk iflogin")
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: "invalid decode json",
		})
		return
	}

	user := model.User{
		Email: userLogin.Email,
		Password: userLogin.Password,
	}
	token, err := u.userService.Login(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse {
			Error: "error internal server",
		})
		return
	}

	cookie := &http.Cookie{
		Name: "session_token",
		Value: *token,
		Expires: time.Now().Add(5 * time.Minute),
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookie)
	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"user_id": user.ID,
	})
	// TODO: answer here
}

func (u *userAPI) GetUserTaskCategory(c *gin.Context) {
	// TODO: answer here
	tasks, err := u.userService.GetUserTaskCategory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error: "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

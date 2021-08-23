package users

import (
	"github.com/gin-gonic/gin"
	"github.com/gokhantamkoc/auth-service/api"
	"github.com/gokhantamkoc/auth-service/util"
	"net/http"
)

type UserService struct {
	KCW *util.KeycloakClientWrapper
}

func (u UserService) CreateUser(c *gin.Context) {
	if err := u.KCW.LoginAdmin("sharpershape", "1", "KeycloakDemo"); err != nil {
		response := util.CreateErrorResponse(http.StatusInternalServerError, "Could not create User!")
		c.JSON(response.Code, response)
		return
	}
	user := api.User{}
	if err := c.BindJSON(&user); err == nil {
		u.KCW.CreateUser(user, "KeycloakDemo")
		response := util.CreateSuccessResponse(nil)
		c.JSON(response.Code, response)
		u.KCW.Logout("KeycloakDemo")
		return
	}
	response := util.CreateErrorResponse(http.StatusBadRequest, "Not a valid JSON!")
	c.JSON(response.Code, response)
}

func (u UserService) GetUsers(c *gin.Context) {
	if err := u.KCW.LoginAdmin("sharpershape", "1", "KeycloakDemo"); err != nil {
		response := util.CreateErrorResponse(http.StatusBadRequest, "Could not create User!")
		c.JSON(response.Code, response)
		return
	}
	users, err := u.KCW.GetUsers("KeycloakDemo", 0, 100)
	if err != nil {
		response := util.CreateErrorResponse(http.StatusInternalServerError, "Could not retrieve users!")
		c.JSON(response.Code, response)
		return
	}
	response := util.CreateSuccessResponse(users)
	c.JSON(response.Code, response)
	u.KCW.Logout("KeycloakDemo")
}

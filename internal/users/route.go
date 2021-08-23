package users

import (
	"github.com/gin-gonic/gin"
	"github.com/gokhantamkoc/auth-service/util"
)

func Route(gc *gin.Engine) {
	kcw := &util.KeycloakClientWrapper{}
	kcw.NewKeycloakClient("http://localhost:8080/")

	userService := UserService{KCW: kcw}

	grp := gc.Group("/users")
	grp.GET("", userService.GetUsers)
	grp.POST("", userService.CreateUser)
}
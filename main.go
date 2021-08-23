package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gokhantamkoc/auth-service/internal/users"
)

var (
	router = gin.Default()
)

func main() {
	users.Route(router)
	router.Run(":5000",)
}

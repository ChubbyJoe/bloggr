package routes

import (
	"github.com/ChubbyJoe/bloggr/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/signup", handlers.SignUpHandler)
	r.POST("/login", handlers.LoginHandler)

}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wharf/wharf/internal/controllers"
	"github.com/wharf/wharf/pkg/auth"
)

func ImageRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.Use(auth.AuthMiddleWare())
	incommingRoutes.GET("/api/images", controllers.GetImages())
}

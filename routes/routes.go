package routes

import(
	"bookmanagement/services"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		 api.GET("/users", services.GetUserBook)
		// api.POST("/users", controllers.CreateUser)
	}

	return r
}
package routes

import (
	"blogapi/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// CORS
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Health check
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Blog API Ready! 🚀"})
	})

	// Public routes — bina token ke
	public := r.Group("/api")
	{
		public.POST("/register", controllers.Register)
		public.POST("/login",    controllers.Login)
		public.GET("/posts",     controllers.GetAllPosts)
		public.GET("/posts/:id", controllers.GetPostByID)
	}

	// Protected routes — token zaroori
	protected := r.Group("/api")
	protected.Use(AuthMiddleware())
	{
		protected.POST("/posts",        controllers.CreatePost)
		protected.PUT("/posts/:id",     controllers.UpdatePost)
		protected.DELETE("/posts/:id",  controllers.DeletePost)
	}

	return r
}
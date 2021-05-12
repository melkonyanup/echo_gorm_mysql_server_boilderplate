package v1

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"myapp/internal/service"
	"myapp/pkg/middlewares"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouter(e *echo.Echo) {
	// http://localhost:8001/swagger/index.html
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	api := e.Group("/api/v1")
	{
		api.POST("/sign-up", h.signUp)
		api.POST("/sign-in", h.signIn)

		oauth := api.Group("/oauth")
		{
			// http://localhost:8001/api/v1/oauth
			oauth.GET("", h.oAuthStart)

			oauth.GET("/google", h.googleLogin)
			oauth.GET("/google/callback", h.googleLoginCallback)

			oauth.GET("/fb", h.fBLogin)
			oauth.GET("/fb/callback", h.fBLoginCallback)
		}

		authenticated := api.Group("")
		authenticated.Use(middlewares.Authz)
		{
			user := authenticated.Group("/user")
			{
				user.GET("/profile", h.getUserProfile)
			}

			post := authenticated.Group("/post")
			{
				post.GET("", h.getAllPosts)
				post.GET("/:id", h.getPost)
				post.POST("", h.createPost)
				post.PUT("/:id", h.updatePost)
				post.DELETE("/:id", h.deletePost)
			}
		}
	}
}

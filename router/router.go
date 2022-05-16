package router

import (
	"github.com/gin-gonic/gin"
	"lj-go-practice/handler"
	"lj-go-practice/router/middle"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middle.NoCache)
	g.Use(middle.Options)
	g.Use(middle.Secure)
	g.Use(mw...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "the incorrect api route")
	})

	check := g.Group("/api/check")
	{
		check.GET("/health", handler.HealthCheck)
		check.GET("/disk", handler.DiskCheck)
		check.GET("/cpu", handler.CPUCheck)
		check.GET("/ram", handler.RAMCheck)
	}

	account := g.Group("/api/account")
	{
		account.POST("/login", handler.LoginAccount)
		account.POST("/updatePass", handler.UpdateAccountPwd)
	}

	user := g.Group("/api/user/common")
	{
		user.GET("/queryAll", handler.QueryUsers)
	}

	userAuth := g.Group("/api/user/auth")
	userAuth.Use(middle.AuthMiddleware())
	{
		userAuth.POST("/add", handler.AddUser)
		userAuth.POST("/update", handler.UpdateUser)
	}

	return g
}

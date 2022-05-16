package middle

import (
	"github.com/LianJianTech/lj-go-common/errno"
	"github.com/LianJianTech/lj-go-common/log"
	"github.com/gin-gonic/gin"
	"lj-go-practice/handler"
	"lj-go-practice/pkg"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := pkg.ParseRequest(c); err != nil {
			log.Errorf(err, "pkg.ParseRequest error")
			handler.SendResponse(c, errno.AuthError, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}

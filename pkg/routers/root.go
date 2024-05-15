package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "my.domain/guestbook/apis/apps/v1"
	"my.domain/guestbook/pkg/builders"
	v1Routers "my.domain/guestbook/pkg/routers/v1"
)

var (
	rootUrl = fmt.Sprintf("/apis/%s/%s", v1.SchemeGroupVersion.Group, v1.SchemeGroupVersion.Version)
)

func Register(r *gin.Engine) {
	// 根
	r.GET(rootUrl, func(c *gin.Context) {
		c.JSON(200, builders.ApiResourceList())
	})

	v1Routers.RegisterRoute(r)
}

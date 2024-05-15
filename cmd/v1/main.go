package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"my.domain/guestbook/pkg/k8sconfig"
	"my.domain/guestbook/pkg/routers"
)

func main() {
	k8sconfig.K8sInitInformer() // 启动 Informer 监听
	r := gin.New()
	r.Use(func(c *gin.Context) {
		c.Next()
	})

	routers.Register(r)
	//  8443  没有为啥
	if err := r.RunTLS(":8443",
		"certs/aaserver.crt", "certs/aaserver.key"); err != nil {
		log.Fatalln(err)
	}
}

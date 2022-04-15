package routes

import (
	"github.com/gin-gonic/gin"
)

var (
	r = gin.Default()
)

const (
	port = ":8080"
)

func StartApp() {
	Url_maps()
	r.Run(port)
}

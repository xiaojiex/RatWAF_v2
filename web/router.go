package web

import (
	"RatWAF_v2/utils"

	"github.com/gin-gonic/gin"
)

func InitWeb() {
	r := gin.Default()
	r.LoadHTMLGlob("web/templates/*")
	r.GET("/index", utils.GetInfo)
	r.GET("/detail/:id", utils.GetDetail)
	r.Run(":9090")
}

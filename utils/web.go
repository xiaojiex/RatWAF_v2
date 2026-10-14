package utils

import (
	"RatWAF_v2/dao"
	"RatWAF_v2/global"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetInfo(ctx *gin.Context) {
	var wafs []dao.Waf
	global.DB.Select("id", "attack_type", "remote", "createtime", "attacker_ip").Order("id desc").Find(&wafs)
	// global.DB.Order("id desc").Find(&wafs)
	ctx.HTML(http.StatusOK, "index.html", gin.H{"log": wafs})
	// ctx.JSON(http.StatusOK, gin.H{"data": wafs})
}
func GetDetail(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"log": err})
		return
	}
	var wafs dao.Waf
	err = global.DB.Where("id=?", id).First(&wafs).Error
	if err != nil {
		fmt.Println("err", err)
		ctx.JSON(http.StatusNotFound, gin.H{"log": err})
		return
	}
	// ctx.JSON(http.StatusOK, gin.H{"log": wafs})
	ctx.HTML(http.StatusOK, "detail.html", gin.H{"log": wafs})
}

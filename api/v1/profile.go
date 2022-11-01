package v1

import (
	"duryun-blog/model"
	"duryun-blog/service"
	"duryun-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetProfile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	ps := service.ProfileService{}
	data, code := ps.GetProfile(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

func UpdateProfile(c *gin.Context) {
	var data model.Profile
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	ps := service.ProfileService{}
	code := ps.UpdateProfile(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

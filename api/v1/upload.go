package v1

import (
	"duryun-blog/service"
	"duryun-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UpLoad 上传图片接口
func UpLoad(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"upload error": err.Error(),
		})
		return
	}

	url, err := service.UpLoadFile(file, fileHeader)

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": errmsg.GetErrMsg(200),
		"url":     url,
	})
}

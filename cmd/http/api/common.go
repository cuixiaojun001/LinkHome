package api

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/cuixiaojun001/LinkHome/common/response"
	"github.com/cuixiaojun001/LinkHome/services/common"
	"github.com/gin-gonic/gin"
)

func AreaInfo(c *gin.Context) {
	if item, err := common.AreaInfo(); err != nil {
		c.JSON(http.StatusOK, response.InternalServerError(err))
		return
	} else {
		c.JSON(http.StatusOK, response.Success(item))
	}
}

func Upload(c *gin.Context) {
	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.InternalServerError(err))
		return
	}
	log.Println(file.Filename)

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.InternalServerError(err))
		return
	}
	defer src.Close()

	data, err := io.ReadAll(src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.InternalServerError(err))
		return
	}

	res := common.UploadFile(context.Background(), file.Filename, data)

	c.JSON(http.StatusOK, response.Success(res))
}

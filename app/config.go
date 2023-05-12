package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type mysql struct {
	Host     string `json:"host" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	DBName   string `json:"db_name" binding:"required"`
	Port     string `json:"port" binding:"required"`
}

func SaveConfig(c *gin.Context) {
	// 接受信息
	var form mysql
	if c.ShouldBind(&form) == nil {
		fmt.Println(form)
		c.JSON(http.StatusOK, struct {
			Msg  string
			Code int
		}{
			Code: 200,
			Msg:  "success",
		})
	}
	// 修改启动项目的配置文件
	// todo

}

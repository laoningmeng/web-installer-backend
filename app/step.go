package app

import (
	"github.com/gin-gonic/gin"
)

type StepListItem struct {
	Title string `json:"title"`
	Sort  int64  `json:"sort"`
}
type StepListResult struct {
	Data        []StepListItem `json:"data"`
	CurrentStep int64          `json:"current_step"`
}

var data = []StepListItem{
	{
		Title: "环境检测",
		Sort:  1,
	},
	{
		Title: "配置信息",
		Sort:  2,
	},
	{
		Title: "项目初始化",
		Sort:  3,
	},
	{
		Title: "初始化成功",
		Sort:  4,
	},
}

func GetStepList(c *gin.Context) {
	result := StepListResult{
		Data:        data,
		CurrentStep: 0,
	}
	c.JSON(200, result)
}

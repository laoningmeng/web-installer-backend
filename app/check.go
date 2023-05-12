package app

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-version"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
)

type IGetVersion interface {
	getVersion() string
}
type goVersion struct{}

func (g *goVersion) getVersion() string {
	v := runtime.Version()
	s := v[2:]
	s = strings.Trim(s, "\n")
	return s
}

type nodeVersion struct{}

func (g *nodeVersion) getVersion() string {
	v, _ := exec.Command("node", "-v").Output()
	s := string(v)

	return s[:len(s)-1]
}

var checkData = []checkItem{
	{
		Title:       "go",
		Status:      0,
		Desc:        "最低版本要求go1.13",
		Version:     "1.13",
		CurrVersion: "",
	},
	{
		Title:       "node",
		Status:      0,
		Desc:        "最低版本要求v1.13",
		Version:     "1.0.1",
		CurrVersion: "",
	},
}

type checkItem struct {
	Title       string `json:"title"`        // 检测的条目
	Status      int    `json:"status"`       // 0 待检测  1 符合条件 2 不符合条件
	Desc        string `json:"desc"`         // 提示信息
	Version     string `json:"version"`      // 最低版本要起
	CurrVersion string `json:"curr_version"` // 当前的版本
}

type CheckListResult struct {
	Data        []checkItem `json:"data"`
	CurrentStep int         `json:"current_step"`
}

func (c *checkItem) check() {
	// 检测命令是否存在
	isExist := cmdExists(c.Title)
	if isExist != nil {
		c.Desc = "命令" + c.Title + "不存在"
		c.Status = 2
		return
	}
	// 检测版本号
	switch c.Title {
	case "go":
		goVersion := new(goVersion)
		c.CurrVersion = goVersion.getVersion()
	case "node":
		nodeVersion := new(nodeVersion)
		c.CurrVersion = nodeVersion.getVersion()
	default:
		c.Desc = c.Title + "未查询到版本号"
		return
	}
	v := compareVersion(c.CurrVersion, c.Version)
	if v < 0 {
		c.Desc = c.Title + "版本不符合要求，最低版本 " + c.Version + " 当前版本 " + c.CurrVersion
		c.Status = 2
		return
	}
	c.Desc = "success"
	c.Status = 1
	return
}

func GetCheckList(c *gin.Context) {
	//
	if check := c.DefaultQuery("is_check", "0"); check != "0" {
		// 需要进行检测
		var res []checkItem
		for _, e := range checkData {
			e.check()
			res = append(res, e)
		}
		c.JSON(http.StatusOK, CheckListResult{
			Data:        res,
			CurrentStep: 0,
		})
		return
	}
	// 初始化界面展示的数据
	c.JSON(http.StatusOK, CheckListResult{
		Data:        checkData,
		CurrentStep: 0,
	})
}

func cmdExists(cmd string) error {
	_, err := exec.LookPath(cmd)
	return err
}

// 1 v1>v2  0 v1=v2 -1 v1<v2
func compareVersion(v1, v2 string) int {
	version1, _ := version.NewVersion(v1)
	version2, _ := version.NewVersion(v2)
	return version1.Compare(version2)
}

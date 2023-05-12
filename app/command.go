package app

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"os/exec"
	"time"
)

type Cmd struct {
	Cmd string
	Arg []string
}

type Shell struct {
	cmd     Cmd
	percent int
}

var shellData = []Shell{
	{
		cmd: Cmd{
			Cmd: "node",
			Arg: []string{"-h"},
		},
		percent: 10,
	},
	{
		cmd: Cmd{
			Cmd: "ls",
			Arg: []string{"-la"},
		},
		percent: 10,
	},
	{
		cmd: Cmd{
			Cmd: "ls",
			Arg: []string{"-la"},
		},
		percent: 10,
	},
}

func Exec(c *gin.Context) {
	conn, _, _, err := ws.UpgradeHTTP(c.Request, c.Writer)
	if err != nil {
		// handle error
	}

	go func() {
		defer conn.Close()
		for {
			msg, op, err := wsutil.ReadClientData(conn)
			if len(msg) > 0 {
				str := string(msg)
				fmt.Println("收到消息：", str, op)
				if str == "start" {
					for _, e := range shellData {
						fmt.Println("执行", e)
						command := exec.Command(e.cmd.Cmd, e.cmd.Arg...)
						pipe, err := command.StdoutPipe()
						if err != nil {
							fmt.Println("err:", err)
						}
						command.Start()
						scanner := bufio.NewScanner(pipe)
						for scanner.Scan() {
							wsutil.WriteServerMessage(conn, ws.OpText, []byte(">"+scanner.Text()))
							time.Sleep(200 * time.Millisecond)
						}

					}

				}
			}
			if err != nil {
				// handle error
				//fmt.Println(err)
			}

		}
	}()

}

package main

import (
	"fmt"
	"time"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "web_backend/internal/logic"

	"web_backend/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	time.Local = loc
	cmd.Main.Run(gctx.New())
}

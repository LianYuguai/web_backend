package main

import (
	"context"
	"fmt"
	"time"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/v2/frame/g"

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
	var (
		ctx = context.Background()
	)
	conn, _ := g.Redis().Conn(ctx)
	defer conn.Close(ctx)
	v, _ := g.Redis().Do(ctx, "GET", "token")
	fmt.Println(v.String())
	cmd.Main.Run(gctx.New())
}

package main

import (
	_ "mygf/internal/packed"

	_ "mygf/internal/logic"


	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"mygf/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}

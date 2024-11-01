package main

import (
	_ "github.com/ZYallers/fine-example/internal/logic"
	"github.com/ZYallers/fine-example/internal/router"

	"github.com/ZYallers/fine/database/fredis"
	"github.com/ZYallers/fine/frame/fapp"
	"github.com/ZYallers/fine/frame/fhandler"
	"github.com/ZYallers/fine/os/fsession"
	"github.com/ZYallers/fine/util/fvaild"
)

// @title fine-example
// @version 2.0.0
// @host fine-example.hxsapp.com
// @description fine example application
func main() {
	app := fapp.New(
		fsession.WithClient(fredis.ClientFunc("session")),
	)
	app.Run(
		fhandler.WithRecover(),
		fhandler.WithLogger(),
		fhandler.WithNoRoute(),
		fhandler.WithTracing(),
		// fhandler.WithRootPath(func(ctx *gin.Context) { ctx.String(http.StatusOK, "OK") }),
		fhandler.WithPing(),
		fhandler.WithHealth(),
		// fhandler.WithExpVar(),
		// fhandler.WithPrometheus(),
		fhandler.WithSwagger(),
		// fhandler.WithPProf(),
		// fhandler.WithStats("", nil),
		fhandler.WithRouter(router.Router),
		fvaild.WithBinding(fvaild.V10()),
	)
}

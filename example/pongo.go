package main

import (
	"gopkg.in/echo-contrib/pongor.v1"
	"gopkg.in/labstack/echo.v1"
)

func main() {
	serv := echo.New()
	r := pongor.GetRenderer()
	// r := pongor.GetRenderer(pongor.PongorOption{
	// 	Reload: true, // if you want to reload template every request, set Reload to true.
	// })
	serv.SetRenderer(r)
	serv.Static("/static", "./static")
	serv.Get("/", func(ctx *echo.Context) error {
		ctx.Render(200, "index.html", map[string]interface{}{
			"title": "你好，世界",
		})
		return nil
	})

	serv.Run("127.0.0.1:9000")
}

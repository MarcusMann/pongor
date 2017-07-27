func main() {
	serv := echo.New()
	r := pongor.GetRenderer()

	serv.Renderer = r

	serv.Static("/static", "./static")

	serv.GET("/", func() echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Render(200, "index.html", map[string]interface{}{
				"title": "你好，世界",
			})
			return nil
		}
	}())

	serv.Start(":3000")
}

func contato(c echo.Context) error {
	return c.Render(200, "index.html", nil)
}

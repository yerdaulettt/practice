package app

import (
	v1 "p7/internal/contoller/http/v1"
)

func Run() {
	// r := gin.Default()

	// r2 := r.Group("/admin")
	// {
	// 	r2.GET("/", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"admin": "page"}) })
	// 	r2.GET("/edit", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"admin": "edit"}) })
	// }

	// r.GET("/ok", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"status": "ok"}) })

	// log.Fatal(http.ListenAndServe(":8080", r))

	v1.StartServer()
}

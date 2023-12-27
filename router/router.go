package router

import (
	"api/api"
	"api/middleware"
	"github.com/gin-gonic/gin"
	"html/template"
	"os"
	"path/filepath"
)

func INitRouter() *gin.Engine  {
	router := gin.New()

	router.Use(gin.Recovery())

	router.Static("/static","./static")

	router.SetFuncMap(template.FuncMap{
		"safe": func(str string)template.HTML {
			return template.HTML(str)
		},
	})

	router.LoadHTMLGlob(filepath.Join(os.Getenv("GOPATH"),"/src/demo_panXinRen/public/*"))

	router.Use(middleware.Cors())
	router.Use(middleware.Logger())

	register(router)

	return router
}

func register(router *gin.Engine)  {
    router.GET("/",api.Index)
    router.POST("/dapp/transaction",api.Faucets)
}
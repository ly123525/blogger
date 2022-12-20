package main

import (
	"blogger/dal/db"

	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	dns := "root:@tcp(localhost:3306)/blogger?parseTime=true"
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}
	ginpprof.Wrapper(router)
	router.Static("/static/", "./static")
	router.LoadHTMLGlob("views/*")
}

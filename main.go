package main

import (
	"api/dao"
	"api/router"
)

func main() {
	dao.InitDB("main.db")
	r := router.SetupRouter()
	r.Run()
}

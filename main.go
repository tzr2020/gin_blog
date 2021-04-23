package main

import (
	"github.com/tzr2020/gin_blog/model"
	"github.com/tzr2020/gin_blog/routes"
	"github.com/tzr2020/gin_blog/utils"
)

func main() {
	model.InitDB()
	r := routes.InitRouter()
	r.Run(utils.HttpPort)
}

package main

import (
	"github.com/20gu00/go_bookManager/dao"
	"github.com/20gu00/go_bookManager/router"
)

func main() {
	dao.InitMysql()
	r := router.InitRouter()
	r.Run(":8000")
}

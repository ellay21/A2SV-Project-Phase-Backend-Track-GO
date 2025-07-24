package main

import (
	"task_manager/controllers"
	"task_manager/data"
	"task_manager/router"
)

func main() {
	db, _ := data.InitDB()
	ts := data.NewTaskService(db)
	us := data.NewUserService(db)
	tCont := controllers.NewTaskController(ts)
	uCont := controllers.NewUserController(us)
	r := router.SetupRouter(tCont, uCont)
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}

}

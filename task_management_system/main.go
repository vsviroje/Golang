package main

import (
	"log"

	"task_management_system/app"
)

func main() {

	taskManagementApp := app.NewApplication()

	err := taskManagementApp.Init("config/config.json")
	if err != nil {
		log.Fatalf("unable to init application :%v", err)
	}

	err = taskManagementApp.Start()
	if err != nil {
		log.Fatalf("unable to start application :%v", err)
	}
}

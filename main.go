package main

import (
	"io"
	"log"
	"os"
	_ "rateserver/routers"

	beego "github.com/beego/beego/v2/adapter"
)

func main() {
	fileName := "logs/blockcare_rate.log"
	var exist bool
	for !exist {
		err := os.MkdirAll("logs", os.ModePerm)
		if err == nil {
			exist = true
		}
	}
	// open log file
	logFile, logErr := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE, 0644)
	if logErr != nil {
		log.Panic(logErr)
	}
	defer logFile.Close()
	// redirect all the output to file
	wrt := io.MultiWriter(os.Stdout, logFile)

	// set log out put
	log.SetOutput(wrt)
	// optional: log date-time, filename, and line number
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	beego.BConfig.AppName = "Rateserver"
	beego.BConfig.Log.AccessLogs = true
	beego.Run()
}

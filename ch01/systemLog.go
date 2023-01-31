package main

import (
	"log"
	"log/syslog"
)

func main() {
	syslog, err := syslog.New(syslog.LOG_SYSLOG, "systemLog.go")
    if err!=nil{
        log.Println(err)
        return
    }else{
        log.SetOutput(syslog)
        log.Print("Everything is fine!")
    }
}

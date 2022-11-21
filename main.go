package main

import (
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var build = "develop"

func main() {
	numberCpu := runtime.NumCPU()
	log.Printf("Starting Service [%s] and number cpu [%d]", build, numberCpu)
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown
	log.Println("Stopping the service")
}

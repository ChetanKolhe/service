package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"go.uber.org/automaxprocs/maxprocs"
)

var build = "develop"

func main() {
	// numberCpu := runtime.NumCPU()
	// =========================================================================
	// GOMAXPROCS

	// Set the correct number of threads for the service
	// based on what is available either by the machine or quotas.
	if _, err := maxprocs.Set(); err != nil {
		fmt.Println("maxprocs: %w", err)
		os.Exit(1)
	}
	fmt.Println(build)
	maxCpu := runtime.GOMAXPROCS(0)
	log.Printf("Starting Service [%s] and number cpu [%d] chetan", build, maxCpu)
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	value := <-shutdown

	fmt.Println(value.String())

	log.Println("Stopping the service")
}

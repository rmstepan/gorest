// main.go

package main

import (
	cfg "./config"
	datab "./database"
	log "./logger"
	routing "./routing"
)

func init() {
	log.InfoLogger.Println("Initializing packages...")

	cfg.Setup()
	datab.Setup()
	routing.Setup()
}

func main() {
	log.InfoLogger.Println("Starting main app!")
	routing.Run()

}

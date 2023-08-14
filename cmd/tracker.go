package main

import (
	"flag"
	"os"

	"github.com/buraksenn/expense-tracker/pkg/logger"
)

var (
	logPath = flag.String("log_path", "expense_tracker.log", "Log file path")
)

func main() {
	flag.Parse()

	var file *os.File
	if *logPath == "" {
		var err error
		file, err = os.OpenFile(*logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic("Failed to open log file, err: " + err.Error())
		}
		defer func() {
			if err := file.Close(); err != nil {
				panic("Failed to close log file, err: " + err.Error())
			}
		}()
	}
	logger.Init(file)
}

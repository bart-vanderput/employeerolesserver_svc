package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"golang.org/x/sys/windows/svc/debug"
)

var exeFolder string

// if setup returns an error, the service doesn't start
func setup(wl debug.Log, svcName string) (server, error) {
	var s server

	s.winlog = wl

	// Note: any logging here goes to Windows App Log
	// I suggest you setup local logging

	//Find current path
	exeFolder := filepath.Dir(os.Args[0])
	configFilePath := exeFolder + "\\config.txt"

	// Read env variables
	myEnv, err := godotenv.Read(configFilePath)
	if err != nil {
		s.winlog.Error(1, fmt.Sprintf("Error loading config file: %v", err))
		log.Fatalf("Error loading config file: %v", err)
	}

	// Create logger object
	ml := myLogger{logfile: myEnv["log_filepath"]}
	if err = ml.Write("Application started"); err != nil {
		s.winlog.Error(1, fmt.Sprintf("Error writing to logfile: %v", err))
		log.Fatalf("Error writing to logfile: %v", err)
	}
	s.myLog = &ml

	// Check if the csv file is readable, if not, kill
	fp := myEnv["csv_filepath"]
	if err := testCSV(fp); err != nil {
		s.winlog.Error(1, fmt.Sprintf("Error reading csv file (%s): %v", fp, err))
		ml.Write(fp)
		ml.Write(fmt.Sprintf("Error reading CSV file: %v", err))
		os.Exit(2)
	}

	s.csvFilePath = myEnv["csv_filepath"]
	s.port = ":" + myEnv["http_port"]
	s.rootPath = exeFolder

	ml.Write(fmt.Sprintf("Rootpath: %s", exeFolder))

	return s, nil
}

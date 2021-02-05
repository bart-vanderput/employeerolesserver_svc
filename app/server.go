package app

import (
	"golang.org/x/sys/windows/svc/debug"
)

type server struct {
	winlog      debug.Log
	csvFilePath string
	rootPath    string
	myLog       *myLogger
	port        string
}

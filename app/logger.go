package app

import (
	"log"
	"os"
	"sync"
)

type myLogger struct {
	logfile string
	mu      sync.Mutex
}

func (l *myLogger) Write(line string) error {

	l.mu.Lock()
	defer l.mu.Unlock()
	f, err := os.OpenFile(l.logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	logger.Println(line)
	return nil
}

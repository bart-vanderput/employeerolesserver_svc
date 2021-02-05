package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func myEmployeeRolesServer(s server) {

	rp := roleProcessor{s.csvFilePath, nil}

	templateFilePath := s.rootPath + "\\templates\\templateEmployeeRoles.html"
	staticPath := s.rootPath + "\\static"

	s.myLog.Write(fmt.Sprintf("Staticpath: %s", staticPath))
	s.myLog.Write(fmt.Sprintf("Templatefilepath: %s", templateFilePath))

	// Create pagehandler object
	handlePages := pageHandler{rp: &rp, logger: s.myLog, templateFile: templateFilePath}

	// Start the router
	router := mux.NewRouter()
	router.HandleFunc("/", handlePages.Home)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(staticPath))))

	s.myLog.Write(fmt.Sprintf("Starting HTTP server on port: %v", s.port))
	s.winlog.Info(1, fmt.Sprintf("Starting HTTP server on port: %v", s.port))

	srv := &http.Server{
		Handler:      router,
		Addr:         s.port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()

}

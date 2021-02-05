package app

import (
	"fmt"
	"html/template"
	"net/http"
)

type pageHandler struct {
	rp           *roleProcessor
	logger       *myLogger
	templateFile string
}

type templateData struct {
	Managers        []string
	SelectedManager string
	Employees       []employee
}

func (h *pageHandler) Home(w http.ResponseWriter, r *http.Request) {

	// Output result to browser
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")

	td := templateData{}
	sm, err := h.rp.getSortedManagers()
	if err != nil {
		h.logger.Write(fmt.Sprintf("Error getting data from csv: %v", err))
		return
	}
	td.Managers = sm

	// Check for posted data = manager selection
	if r.Method == "POST" {

		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			h.logger.Write(fmt.Sprintf("ParseForm() err: %v", err))
			return
		}

		manager := r.FormValue("manager")
		if manager == "" {
			fmt.Fprint(w, "Selecteer een manager !!")
			return
		}
		td.SelectedManager = manager
		h.logger.Write(fmt.Sprintf("Manager selected: %v", manager))
		efm, err := h.rp.getEmployeesForManager(manager)
		if err != nil {
			h.logger.Write(fmt.Sprintf("Error getting data from csv: %v", err))
			return
		}
		td.Employees = efm
	}

	// Build template

	t, err := template.ParseFiles(h.templateFile)
	if err != nil {
		h.logger.Write(fmt.Sprintf("Error with accessing employee template: %v", err))
		return
	}

	err = t.Execute(w, td)
	if err != nil {
		h.logger.Write(fmt.Sprintf("Error with executing employee template: %v", err))
		return
	}
}

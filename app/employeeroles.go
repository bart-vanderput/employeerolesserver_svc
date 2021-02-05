package app

import (
	"sort"
)

// Employee is combination of attributes for a single employee
type employee struct {
	Name       string
	Manager    string
	Department string
	Timelines  string
	Roles      []roleInfo
}

//
type roleInfo struct {
	RoleName string
	Status   string
}

type roleProcessor struct {
	csvFileName  string
	allEmployees []employee
}

func (f *roleProcessor) getEmployeesForManager(manager string) ([]employee, error) {
	err := getEmployees(f)
	if err != nil {
		return nil, err
	}

	empList := []employee{}
	for _, emp := range f.allEmployees {
		if emp.Manager == manager {
			empList = append(empList, emp)
		}
	}
	return empList, nil
}

func (f *roleProcessor) getSortedManagers() ([]string, error) {
	// Get all distinct managers from the employeeList and sort the list
	err := getEmployees(f)
	if err != nil {
		return nil, err
	}

	managerList := []string{}
	keys := make(map[string]bool)

	for _, e := range f.allEmployees {
		if _, value := keys[e.Manager]; !value {
			keys[e.Manager] = true
			managerList = append(managerList, e.Manager)
		}
	}
	sort.Strings(managerList)
	return managerList, nil
}

func getEmployees(f *roleProcessor) error {
	if f.allEmployees != nil {
		return nil
	}
	records, header, err := readCSV(f.csvFileName, ';')
	if err != nil {
		return err
	}

	var empList []employee
	prevEmpName := ""
	var emp employee = employee{}

	for _, record := range records {

		if prevEmpName != record[0] && emp.Name != "" {

			// Sort Roles when new employee is found
			sort.Slice(emp.Roles, func(i, j int) bool {
				return emp.Roles[i].RoleName < emp.Roles[j].RoleName // sort return
			})
			// Add to result
			empList = append(empList, emp)
			emp = employee{}
		}
		emp.Name = record[find(header, "medewerker")]
		emp.Manager = record[find(header, "manager")]
		emp.Department = record[find(header, "afdeling")]
		emp.Roles = append(emp.Roles, roleInfo{RoleName: record[find(header, "rol")]})
		prevEmpName = emp.Name
	}

	// Sort result by name
	sort.Slice(empList, func(i, j int) bool {
		return empList[i].Name < empList[j].Name // sort return
	})

	f.allEmployees = empList
	return nil
}

func find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return 0
}

package employee

type EmployeeFormatter struct {
	ID   int
	Name string
	NIP  string
	Role string
	No   int
}

func FormatEmployee(employee Employee) EmployeeFormatter {
	employeeFormatter := EmployeeFormatter{}

	employeeFormatter.ID = employee.ID
	employeeFormatter.Name = employee.Name
	employeeFormatter.NIP = employee.Nip
	employeeFormatter.Role = employee.Role

	return employeeFormatter
}

func FormatEmployees(employees []Employee) []EmployeeFormatter {
	employeesFormatter := []EmployeeFormatter{}

	for index, employee := range employees {
		employeeFormatter := FormatEmployee(employee)
		employeeFormatter.No = index + 1

		employeesFormatter = append(employeesFormatter, employeeFormatter)
	}

	return employeesFormatter
}

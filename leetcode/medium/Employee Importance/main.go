package medium

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func recurTotalImportance(employeeTable map[int]*Employee, id int) int {
	subordinatesImportanceSum := 0

	for _, subEmpId := range employeeTable[id].Subordinates {
		subordinatesImportanceSum += recurTotalImportance(employeeTable, subEmpId)
	}

	return employeeTable[id].Importance + subordinatesImportanceSum
}

func getImportance(employees []*Employee, id int) int {
	employeeTable := make(map[int]*Employee)
	for _, employee := range employees {
		employeeTable[employee.Id] = employee
	}

	return recurTotalImportance(employeeTable, id)
}

package time_needed_to_inform_all_employees

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

type Employee struct {
	inform_time  int
	subordinates []*Employee
}

func NewEmployee(inform_time int) *Employee {
	return &Employee{inform_time, []*Employee{}}
}

func (manager *Employee) addSubordinate(employee *Employee) {
	manager.subordinates = append(manager.subordinates, employee)
}

func (manager *Employee) informSubordinates() int {
	maxTimeForSubordinates := 0
	for _, employee := range manager.subordinates {
		maxTimeForSubordinates = max(maxTimeForSubordinates, employee.informSubordinates())
	}
	return manager.inform_time + maxTimeForSubordinates
}

func NumOfMinutes(_ int, headId int, manager []int, informTime []int) int {
	employeeRecord := make(map[int]*Employee)

	for employeeId, informTime := range informTime {
		employeeRecord[employeeId] = NewEmployee(informTime)
	}

	for employeeId, managerId := range manager {
		if manager, ok := employeeRecord[managerId]; ok {
			manager.addSubordinate(employeeRecord[employeeId])
		}
	}

	return employeeRecord[headId].informSubordinates()
}

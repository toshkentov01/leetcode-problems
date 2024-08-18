package main

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	emplMap := make(map[int]*Employee)

	for i := range employees {
		emplMap[employees[i].Id] = employees[i]
	}

	importanceSum := dfs(emplMap, id)

	return importanceSum
}

func dfs(emplMap map[int]*Employee, id int) int {
	var importanceSum int

	empl := emplMap[id]
	importanceSum += empl.Importance

	for i := range empl.Subordinates {
		importanceSum += dfs(emplMap, empl.Subordinates[i])
	}

	return importanceSum
}

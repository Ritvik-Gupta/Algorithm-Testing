package medium

import "sync"

type flag = struct{}

type Course struct {
	outLinks []int
}

func NewCourse() *Course {
	return &Course{outLinks: make([]int, 0)}
}

type CourseRecord struct {
	courses map[int]*Course
}

func (record *CourseRecord) addPrereq(fromId, toId int) {
	if _, ok := record.courses[toId]; !ok {
		record.courses[toId] = NewCourse()
	}
	if _, ok := record.courses[fromId]; !ok {
		record.courses[fromId] = NewCourse()
	}
	fromCourse := record.courses[fromId]
	fromCourse.outLinks = append(fromCourse.outLinks, toId)
}

func (record *CourseRecord) checkIfPrereq(fromId, toId int) bool {
	return record.dfsCheckIfPrereq(fromId, toId, make(map[int]flag))
}

func (record *CourseRecord) dfsCheckIfPrereq(fromId, toId int, visitedCourses map[int]flag) bool {
	if fromId == toId {
		return true
	}
	if _, ok := visitedCourses[fromId]; ok {
		return false
	}
	visitedCourses[fromId] = flag{}

	if fromCourse, ok := record.courses[fromId]; ok {
		for _, linkToId := range fromCourse.outLinks {
			if record.dfsCheckIfPrereq(linkToId, toId, visitedCourses) {
				return true
			}
		}
	}
	return false
}

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	courses := CourseRecord{courses: make(map[int]*Course, numCourses)}
	for _, prereq := range prerequisites {
		courses.addPrereq(prereq[0], prereq[1])
	}

	var waitForResults sync.WaitGroup
	waitForResults.Add(len(queries))
	result := make([]bool, len(queries))

	for idx, query := range queries {
		fromId, toId := query[0], query[1]
		i := idx
		go func() {
			result[i] = courses.checkIfPrereq(fromId, toId)
			waitForResults.Done()
		}()
	}

	waitForResults.Wait()
	return result
}

package main

import (
	"github.com/m_yokomichi/design-pattern/iterator/iterator"
)

func main() {
	var teacher Teacher

	teacher.createStudentList(Student{name : "Tom", sex : 1})
	teacher.createStudentList(Student{name : "Jelly", sex : 2})
	teacher.createStudentList(Student{name : "Bob", sex : 1})

	teacher.callStudents()

}
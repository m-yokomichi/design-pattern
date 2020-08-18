package main

import (
	iterator "1_iterator/iterator"
	//"fmt"
)

func main() {
	var teacher iterator.Teacher
	teacher.Init()

	var student iterator.Student

	student.SetStudent("Jelly", 2)
	teacher.CreateStudentList(student)
	student.SetStudent("Tom", 1)
	teacher.CreateStudentList(student)
	student.SetStudent("Bob", 1)
	teacher.CreateStudentList(student)

	teacher.CallStudents()

}

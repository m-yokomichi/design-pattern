package main

import (
	iterator "1_iterator/iterator"
)

func main() {
	var teacher iterator.Teacher


	teacher.CreateStudentList(iterator.Student{})
	//teacher.CreateStudentList(iterator.Student{name : "Jelly", sex : 2})
	//teacher.CreateStudentList(iterator.Student{name : "Bob", sex : 1})

	teacher.CallStudents()

}
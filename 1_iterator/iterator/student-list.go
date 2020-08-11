package iterator

type StudentList struct {
	students []*Student
	index int
}

func (studentList *StudentList) StudentList() {
	studentList.students = append(studentList.students, &Student{})
} 

func (studentList *StudentList) Add(student Student) {
	studentList.students = append(studentList.students, &student)
}

func (studentList *StudentList) HasNext() bool {
	return len(studentList.students) > studentList.index 
}

func (studentList *StudentList) Next() Student {
	index := studentList.index
	studentList.index = index + 1

	return *studentList.students[index]
}
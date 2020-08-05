package iterator

type StudentList struct {
	students []*Student
}

func (studentList *StudentList) StudentList() {
	studentList.students = append(studentList.students, &Student{})
} 

func (studentList *StudentList) Add(student *Student) {
	studentList.students = append(studentList.students, student)
}
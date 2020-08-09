package iterator

import (
	"fmt"
)


type Teacher struct {
	studentList *StudentList
}


// 学校から与えられた情報を名簿に書き込む
func (t *Teacher) createStudentList(s Student) {
	t.studentList.Add(&s)
}

// 名簿順に生徒の名前を全て呼ぶ
func (t *Teacher) callStudents() {
	for _, student := range t.studentList.students {
		fmt.Println(student.GetName())
	}
}
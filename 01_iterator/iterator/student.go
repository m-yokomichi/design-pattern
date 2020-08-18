package iterator

type Student struct {
	name string
	sex int
}

func (s *Student) SetStudent(name string, sex int) {
	s.name = name
	s.sex = sex
}

func (s *Student) GetName() string {
	return s.name
}

func (s *Student) GetSex() int {
	return s.sex
}
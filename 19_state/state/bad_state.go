package state

type BadState struct{}

func (s *BadState) ShowTension() string {
	return "too Bad"
}

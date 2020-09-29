package state

type GoodState struct {}

func (s *GoodState) ShowTension() string {
	return "so Good"
}
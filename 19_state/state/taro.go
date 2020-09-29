package state

type Taro struct {
	name string
	State
}

func (t *Taro) ChangeState(state State) {
	t.State = state
}

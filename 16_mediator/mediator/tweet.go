package mediator

type Tweet struct{}

func (t *Tweet) NewTweet() bool {
	// 新しいtweetがあるかの色々の処理
	return true
}

func CreateTweet() *Tweet {
	return &Tweet{}
}

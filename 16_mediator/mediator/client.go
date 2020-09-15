package mediator

type Client interface {
	SetMediator(Mediator)
	CheckTweet() bool
}

package mediator

type User struct {
	mediator Mediator
}

func (u *User) SetMediator(mediator Mediator) {
	u.mediator = mediator
}

func (u *User) CheckTweet() bool {
	return u.mediator.NewTweet()
}

func CreateUser() Client {
	var user Client
	var tweet Mediator
	tweet = &Tweet{}
	user = &User{
		mediator: tweet,
	}

	return user

}

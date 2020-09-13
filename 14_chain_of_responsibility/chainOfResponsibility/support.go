package chainOfResponsibility

type Support interface {
	SetNext(Support) Support
	Resolve(Level)
}

type Level int
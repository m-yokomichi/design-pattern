package command

type Beaker struct {
	salt  int
	water int
}

func NewBeaker(salt, water int) *Beaker {
	return &Beaker{
		salt:  salt,
		water: water,
	}
}

package command

import (
	"fmt"
)

type Beaker struct {
	salt   int
	water  int
	melted bool
}

func NewBeaker(salt, water int) *Beaker {
	beaker := &Beaker{
		salt:  salt,
		water: water,
	}

	beaker.Mix()

	return beaker
}

func (b *Beaker) AddSalt(salt int) {
	b.salt += salt
}

func (b *Beaker) AddWater(water int) {
	b.water += water
}

func (b *Beaker) Mix() {
	var salinity float64
	salinity = float64(b.salt) / float64(b.water+b.salt) * 100

	b.melted = salinity >= 26
}

func (b *Beaker) IsMelted() bool {
	return b.melted
}

func (b *Beaker) Note() {
	var salinity float64
	salinity = float64(b.salt) / float64(b.water+b.salt) * 100
	fmt.Println("水", b.water, "g")
	fmt.Println("食塩", b.salt, "g")
	fmt.Println("濃度", salinity, "%")
}

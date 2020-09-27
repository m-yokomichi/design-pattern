package observer

import (
	"fmt"
)
type Degit struct {

}

func NewDegit() Observer {
	var observer Observer
	observer = &Degit{}

	return observer
}

func (d *Degit) Update(generator *Generator) {
	fmt.Println(generator.GetNumber())
}
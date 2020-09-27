package observer

import (
	"fmt"
)
type Degit struct {

}

func (d *Degit) Update(generator *Generator) {
	fmt.Println(generator.GetNumber())
}
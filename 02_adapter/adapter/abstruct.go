package adapter

import (
	"fmt"
)

// Hello はセットされているものに対して挨拶する
type Hello string

// CallHello こんにちわする
func (hello *Hello) CallHello() {
	fmt.Println("Hello ", *hello)
}

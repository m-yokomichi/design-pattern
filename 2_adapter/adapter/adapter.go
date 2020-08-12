package adapter

import "fmt"

// NewHello はHelloとHelloInterfaceのAdapter
type NewHello struct {
	hello Hello
}

// CallHello は挨拶する
func (newHello *NewHello) CallHello() {
	newHello.hello.CallHello()
}

// CallToHelloString は挨拶する先だけ表示する
func (newHello *NewHello) CallToHelloString() {
	fmt.Println(newHello.hello)
}

// SetString 挨拶する相手をセットする
func (newHello *NewHello) SetString(s string) {
	newHello.SetString(s)
}

package adapter

import "fmt"

// NewHello はHelloとHelloInterfaceのAdapter
type AdapterHello struct {
	*Hello
}

func NewHello() *AdapterHello {
	var hello Hello
	adapterHello := &AdapterHello{
		&hello,
	}

	return adapterHello
}

// CallToHelloString は挨拶する先だけ表示する
func (hello *AdapterHello) CallToHelloString() {
	fmt.Println(*hello.Hello)
}

// SetString 挨拶する相手をセットする
func (hello *AdapterHello) SetString(s Hello) {
	hello.Hello = &s
}

// SetString 挨拶する相手をセットする
func (hello *AdapterHello) CallHello() {
	hello.Hello.CallHello()
}

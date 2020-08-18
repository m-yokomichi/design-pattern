package adapter

// HelloInterface Helloに実装必須な機能を指定する
type HelloInterface interface {
	SetString(Hello)
	CallToHelloString()
}

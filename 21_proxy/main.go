package main

import (
	"./proxy"
)

func main() {
	var proxy1 proxy.ProxyInterface
	proxy1 = &proxy.Proxy{
		Principal: &proxy.Principal{
			Name : "principal",
		},
		Name : "proxy",
	}

	proxy1.Method1()
	proxy1.Method2()
}
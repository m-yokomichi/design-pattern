package singleton

import(
	"fmt"
)

type SingletonInterface interface{}
type Singleton struct {}

var singletonStruct SingletonInterface

func GetSingletonStruct() *SingletonInterface {
	if singletonStruct != nil {
		fmt.Println("既存struct返却")
		return &singletonStruct
	}

	fmt.Println("新規struct返却")
	singletonStruct = Singleton{}
	return &singletonStruct
}
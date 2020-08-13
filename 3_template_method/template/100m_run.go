package template

import (
	"fmt"
	"time"
)

type Run100 struct {
	*Template
}

func NewRun100() *Run100 {
	run100 := &Run100{
		Template: &Template{},
	}
	run100.tmp = run100
	return run100
}

func (r *Run100) StartCall() {
	fmt.Println("位置について...")
	time.Sleep(1 * time.Second)
	fmt.Println("ヨーイ...")
	time.Sleep(2 * time.Second)
}

func (r *Run100) Run() {
	fmt.Println("ドンッ！！")
	time.Sleep(10 * time.Second)
}

func (r *Run100) GoalCall() {
	fmt.Println("ゴール！")
	time.Sleep(time.Second)
}

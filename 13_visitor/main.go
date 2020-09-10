package main

import (
	"./visitor"
)

func main() {
	var t visitor.Teacher
	t = &visitor.RookieTeacher{}

	var home visitor.Home
	home = &visitor.SatoHome{}
	home.Accept(t)
}

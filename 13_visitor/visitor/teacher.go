package visitor

import "fmt"

type Teacher interface {
	Visit(interface{})
}

type RookieTeacher struct {
	name string
}

func (t *RookieTeacher) Visit(e interface{}) {
	switch e.(type) {
	case *SuzukiHome:
		t.suzukiVisit()
	case *SatoHome:
		t.satoVisit()
	}
	fmt.Printf("%T", e)
}

func (_ *RookieTeacher) suzukiVisit() {
	fmt.Println("鈴木さんこんにちは")
}

func (_ *RookieTeacher) satoVisit() {
	fmt.Println("佐藤さんこんにちは")
}

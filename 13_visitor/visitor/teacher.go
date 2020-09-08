package visitor

type Teacher interface {
	Visit(int)
	Visit(string)
}

type RookieTeacher struct {
	name string
}

func (t *RookieTeacher) Visit(i int) {

}

// 同じメソッドはエラー(method redeclared)
func (t *RookieTeacher) Visit(i string) {

}
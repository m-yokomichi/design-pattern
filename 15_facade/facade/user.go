package facade

type User struct {

}

func (u *User) GetUser() string {
	db := &DB{}
	return db.Select("select * from users")
}
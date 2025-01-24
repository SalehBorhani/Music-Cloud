package entity

type User struct {
	ID       uint
	Email    string
	Name     string
	UserName string
	Password string
	PlayList []Music
}

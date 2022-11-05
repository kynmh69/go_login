package model

type User struct {
	Id        string
	UserId    string
	Password  string
	CreatedAt string
	UpdatedAt string
}

func SignUp(userId, password string) (*User, error)  {
	user := User{}
}
package model

type User struct {
	Login    string
	Password string
}

func (s *User) Get() string {
	return s.Password
}

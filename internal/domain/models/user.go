package models

type User struct {
	ID        int64
	Email     string
	PassHash  []byte
	FirstName string
	Name      string
	Phone     string
	Sex       string
}

package entity

const (
	Reader Role = "reader"
	Writer Role = "writer"
)

type Role string

type User struct {
	ID   int
	Role Role
}

// Entity может содержать не только данные,
// но и функции или объекты с методами, которые реализуют логику бизнеса

// func NewUser
// func ChangeUserRole

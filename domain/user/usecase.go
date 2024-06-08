package domain

import (
	"bookinfo/storage"
)

//import "bookinfo/storage/db"

type Repository interface {
	// TODO implement
	// указывать как зависимость вместо repo db.BookDB
}

// выкинуть
//type UseCase struct {
//	repo db.BookDB
//}

func NewUseCase(repo storage.BookDB) *UseCase {
	return &UseCase{repo: repo}
}

package domain

import "bookinfo/storage/db"

type UseCase struct {
	repo db.BookDB
}

func NewUseCase(repo db.BookDB) *UseCase {
	return &UseCase{repo: repo}
}

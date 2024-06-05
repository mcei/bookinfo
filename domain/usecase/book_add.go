package domain

import (
	"bookinfo/entity"
)

type AddUseCase interface {
	Get(id int) (*entity.Book, error)
}

func (uc *UseCase) Add(b entity.Book) (*entity.Book, error) {
	uc.repo.AddBook(nil, b)
	return nil, nil
}

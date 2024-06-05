package domain

import (
	"bookinfo/entity"
)

type GetUseCase interface {
	Get(id int) (*entity.Book, error)
}

func (uc *UseCase) Get(id int) (*entity.Book, error) {
	b := uc.repo.GetBook(nil, entity.Book{}, id)
	return &b, nil
}

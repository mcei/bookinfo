package book

import domain "bookinfo/domain/user"

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

type GetUseCase interface {
	Get(id int) (*Book, error)
}

func (uc *domain.UseCase) Get(id int) (*Book, error) {
	b := uc.repo.GetBook(nil, Book{}, id)
	// TODO handle errors
	return &b, nil
}

type AddUseCase interface {
	Add(id int) (*Book, error)
}

func (uc *domain.UseCase) Add(b Book) (*Book, error) {
	uc.repo.AddBook(nil, b)
	// TODO handle errors
	return nil, nil
}

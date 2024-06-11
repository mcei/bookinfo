package book

import "context"

// как передавать контекст из мейн?

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

type GetUseCase interface {
	Get(id int) (*Book, error)
}

func (uc *UseCase) Get(id int) (*Book, error) {
	b, err := uc.repo.GetBook(context.Background(), nil, Book{}, id)
	return &b, err
}

type AddUseCase interface {
	Add(id int) (*Book, error)
}

func (uc *UseCase) Add(b Book) (*Book, error) {
	_, err := uc.repo.AddBook(context.Background(), nil, b)
	return &b, err
}

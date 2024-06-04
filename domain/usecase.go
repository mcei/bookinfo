package domain

import (
	"bookinfo/entity"
	"bookinfo/storage/db"
)

type UseCase interface {
	// USER HAS ACCESS
	Get(id int) (*entity.Book, error)
	Add(book entity.Book) (*entity.Book, error)
	// Update(id int) (*entity.Book, error)
	// Remove(id int) (*entity.Book, error)

	// USER HAS NO ACCESS

	// Get(id int) (*entity.Book, error)
	// Add(id int) (*entity.Book, error)
	// Update(id int) (*entity.Book, error)
	// Remove(id int) (*entity.Book, error)
}

func NewUseCase(repo db.BookDB) UseCase {
	return &usecase{repo: repo}
}

type usecase struct {
	repo db.BookDB
}

func (uc *usecase) Get(id int) (*entity.Book, error) {
	b := uc.repo.GetBook(nil, entity.Book{}, id)
	//if err != nil {
	//	if errors.Is(err, repository.ErrBookNotFound) {
	//		return nil, ErrNotFound
	//	}
	//	return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	//}
	return &b, nil
}

func (uc *usecase) Add(b entity.Book) (*entity.Book, error) {
	uc.repo.AddBook(nil, b)
	return nil, nil
}

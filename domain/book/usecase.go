package book

import "bookinfo/storage"

type UseCase struct {
	repo storage.BookDB
}

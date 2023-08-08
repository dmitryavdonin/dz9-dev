package service

import (
	"store/internal/model"
	"store/internal/repository"
)

type StoreBookService struct {
	repo repository.StoreBook
}

func NewStoreBookService(repo repository.StoreBook) *StoreBookService {
	return &StoreBookService{repo: repo}
}

func (s *StoreBookService) Create(input model.StoreBook) (int, error) {
	return s.repo.Create(input)
}

func (s *StoreBookService) GetById(book_id int) (model.StoreBook, error) {
	return s.repo.GetById(book_id)
}

func (s *StoreBookService) GetAll(limit int, offset int) ([]model.StoreBook, error) {
	return s.repo.GetAll(limit, offset)
}

func (s *StoreBookService) Delete(book_id int) error {
	return s.repo.Delete(book_id)
}

func (s *StoreBookService) Update(book_id int, input model.StoreBook) error {
	return s.repo.Update(book_id, input)
}

func (s *StoreBookService) AlreadyExists(book_id int) bool {
	return s.repo.AlreadyExists(book_id)
}

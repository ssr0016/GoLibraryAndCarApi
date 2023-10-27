package libraryimpl

import (
	"context"
	"main/pkg/infra/storage/postgres"
	book "main/pkg/tsm/library"
	"time"

	"go.uber.org/zap"
)

type service struct {
	store  *store
	db     postgres.DB
	logger *zap.Logger
}

func NewService(db postgres.DB) book.Service {
	return &service{
		store:  newStore(db),
		db:     db,
		logger: zap.L().Named("book.service"),
	}
}

func (s *service) Search(ctx context.Context, query *book.SearchBookQuery) (*book.SearchBookResult, error) {
	result, err := s.store.search(ctx, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *service) Create(ctx context.Context, cmd *book.CreateBookCommand) error {
	return s.db.InTransaction(ctx, func(ctx context.Context) error {
		result, err := s.store.bookTaken(ctx, 0, cmd.Title)
		if err != nil {
			return err
		}

		if len(result) > 0 {
			return book.ErrBookNameExisting
		}

		err = s.store.create(ctx, &book.Book{
			Title:       cmd.Title,
			AuthorID:    cmd.AuthorID,
			CategoryID:  cmd.CategoryID,
			PublishedAt: time.Now().UTC().Format(time.RFC3339Nano),
		})
		if err != nil {
			return err
		}

		return nil
	})
}

func (s *service) Get(ctx context.Context, id int64) (*book.BookDTO, error) {
	result, err := s.store.get(ctx, id)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, book.ErrBookNotFound
	}

	return result, nil
}

func (s *service) Update(ctx context.Context, cmd *book.UpdateBookCommand) error {
	return s.db.InTransaction(ctx, func(ctx context.Context) error {
		result, err := s.store.bookTaken(ctx, cmd.ID, cmd.Title)
		if err != nil {
			return err
		}

		if len(result) > 1 || (len(result) == 1 && result[0].ID != cmd.ID) {
			return book.ErrBookNotFound
		}

		err = s.store.update(ctx, &book.Book{})
		if err != nil {
			return err
		}

		return nil
	})
}

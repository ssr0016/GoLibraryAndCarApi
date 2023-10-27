package libraryimpl

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"main/pkg/infra/storage/postgres"
	book "main/pkg/tsm/library"
	"main/pkg/util/pointer"
	"strings"

	"go.uber.org/zap"
)

type store struct {
	db     postgres.DB
	logger *zap.Logger
}

func newStore(db postgres.DB) *store {
	return &store{
		db:     db,
		logger: zap.L().Named("book.store"),
	}
}

func (s *store) create(ctx context.Context, entity *book.Book) error {
	return s.db.WithTransaction(ctx, func(tx postgres.Tx) error {
		rawSQL := `
			INSERT INTO book(
				title,
				author_id,
				category_id,
				published_at
			)
			VALUEs(
				:title,
				:author_id,
				:category_id,
				:published_at
			)
		`

		_, err := tx.NamedExec(ctx, rawSQL, entity)
		if err != nil {
			return err
		}

		return nil
	})

}

func (s *store) search(ctx context.Context, query *book.SearchBookQuery) (*book.SearchBookResult, error) {
	var (
		result = &book.SearchBookResult{
			Books: make([]*book.Book, 0),
		}
		sql             bytes.Buffer
		whereConditions = make([]string, 0)
		whereParams     = make([]interface{}, 0)
	)
	sql.WriteString(`
		SELECT
			id,
			title,
            author_id,
            category_id,
            published_at
		FROM book
	`)

	if len(query.Title) > 0 {
		whereConditions = append(whereConditions, "name ILKIE ?")
		whereParams = append(whereParams, pointer.LikeString(query.Title))
	}

	if len(whereConditions) > 0 {
		sql.WriteString(" WHERE ")
		sql.WriteString(strings.Join(whereConditions, " AND "))
	}

	sql.WriteString(`ORDER BY published_at DESC `)

	err := s.db.Select(ctx, &result.Books, sql.String(), whereParams...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *store) get(ctx context.Context, id int64) (*book.BookDTO, error) {
	var result book.BookDTO

	rawSQL := `
		SELECT
			id,
			title,
			author_id,
			category_id,
			published_at
		FROM book
		WHERE
			id = ?
	`

	err := s.db.Get(ctx, &result, rawSQL, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &result, nil

}

func (s *store) update(ctx context.Context, entity *book.Book) error {
	return s.db.WithTransaction(ctx, func(tx postgres.Tx) error {
		rawSQL := `
			UPDATE assigment
			SET
				title = :title,
                author_id = :author_id,
                category_id = :category_id,
                published_at = :published_at
			WHERE
				id = :id
		`

		_, err := tx.NamedExec(ctx, rawSQL, entity)
		if err != nil {
			return err
		}

		return nil
	})
}

func (s *store) bookTaken(ctx context.Context, id int64, title string) ([]*book.Book, error) {
	var result []*book.Book

	rawSQL := `
		SELECT 
			id,
			title,
			author_id,
			category_id,
			published_at 
		FROM book
		WHERE 
			id = ? OR
			title = ?
	`

	err := s.db.Select(ctx, &result, rawSQL, id, title)
	if err != nil {
		return nil, err
	}

	return result, nil
}

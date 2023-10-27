package library

import "context"

type Service interface {
	Create(ctx context.Context, cmd *CreateBookCommand) error
	Search(ctx context.Context, query *SearchBookQuery) (*SearchBookResult, error)
	Get(ctx context.Context, id int64) (*BookDTO, error)
	Update(ctx context.Context, cmd *UpdateBookCommand) error
}

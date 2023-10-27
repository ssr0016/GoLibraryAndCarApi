package carbrand

import "context"

type Service interface {
	Search(ctx context.Context, query *SearchCarBrandQuery) (*SearchCarBrandResult, error)
	Create(ctx context.Context, cmd *CreateCarBrandCommand) error
}

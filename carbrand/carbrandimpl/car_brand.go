package carbrandimpl

import (
	"context"
	"main/pkg/api/request"
	"main/pkg/infra/storage/postgres"
	"main/pkg/tsm/carbrand"
	"time"

	"go.uber.org/zap"
)

type service struct {
	store  *store
	db     postgres.DB
	logger *zap.Logger
}

func NewService(db postgres.DB) carbrand.Service {
	return &service{
		store:  newStore(db),
		db:     db,
		logger: zap.L().Named("assignment.service"),
	}
}

func (s *service) Search(ctx context.Context, query *carbrand.SearchCarBrandQuery) (*carbrand.SearchCarBrandResult, error) {
	result, err := s.store.search(ctx, query)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) Create(ctx context.Context, cmd *carbrand.CreateCarBrandCommand) error {
	return s.db.InTransaction(ctx, func(ctx context.Context) error {
		signedInUser, _ := request.UserFrom(ctx)

		result, err := s.store.carBrandTaken(ctx, cmd.Name)
		if err != nil {
			return err
		}

		if len(result) > 0 {
			return carbrand.ErrCarBrandNameExisting
		}

		err = s.store.created(ctx, &carbrand.CarBrand{
			Name:     cmd.Name,
			CreateBy: signedInUser.Username,
			CreateAt: time.Now().UTC().Format(time.RFC3339Nano),
		})
		if err != nil {
			return err
		}

		return nil
	})
}

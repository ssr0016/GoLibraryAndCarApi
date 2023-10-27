package carbrandimpl

import (
	"bytes"
	"context"
	"main/pkg/infra/storage/postgres"
	"main/pkg/tsm/carbrand"
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
		logger: zap.L().Named("carbrand.store"),
	}
}

func (s *store) search(ctx context.Context, query *carbrand.SearchCarBrandQuery) (*carbrand.SearchCarBrandResult, error) {
	var (
		result = &carbrand.SearchCarBrandResult{
			CarBrands: make([]*carbrand.CarBrand, 0),
		}
		sql             bytes.Buffer
		whereConditions = make([]string, 0)
		whereParams     = make([]interface{}, 0)
	)

	sql.WriteString(`
		SELECT
			id,
			name,
            create_by,
            create_at,
			update_by,
            update_at
		FROM car_brand
	`)

	if len(query.Name) > 0 {
		whereConditions = append(whereConditions, "name ILKIE ?")
		whereParams = append(whereParams, pointer.LikeString(query.Name))
	}

	if len(whereConditions) > 0 {
		sql.WriteString(" WHERE ")
		sql.WriteString(strings.Join(whereConditions, " AND "))
	}

	sql.WriteString(` ORDER BY create_at DESC`)

	err := s.db.Select(ctx, &result.CarBrands, sql.String(), whereParams...)
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (s *store) created(ctx context.Context, entity *carbrand.CarBrand) error {
	return s.db.WithTransaction(ctx, func(tx postgres.Tx) error {
		rawSQL := `
			INSERT INTO car_brand(name,create_at,create_by)
			VALUES (:name,:create_at,:create_by)
		`

		_, err := tx.NamedExec(ctx, rawSQL, entity)
		if err != nil {
			return err
		}

		return nil
	})
}

func (s *store) carBrandTaken(ctx context.Context, name string) ([]*carbrand.CarBrand, error) {
	var result []*carbrand.CarBrand

	rawSQL := `
		SELECT
            id,
			name
		FROM car_brand
		WHERE 
			name = ?
	`

	err := s.db.Select(ctx, &result, rawSQL, name)
	if err != nil {
		return nil, err
	}

	return result, nil
}

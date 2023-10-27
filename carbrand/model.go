package carbrand

import (
	"main/pkg/api/errors"
	"time"
)

var (
	ErrNameEmpty            = errors.New("carbrand.name-empty", "Name is empty")
	ErrCarBrandNameExisting = errors.New("carbrand.name-existing", "Name already exists")
)

type CarBrand struct {
	ID       int64   `db:"id" json:"id"`
	Name     string  `db:"name" json:"name"`
	CreateBy string  `db:"create_by" json:"createBy"`
	UpdateBy *string `db:"update_by" json:"updateBy"`
	CreateAt string  `db:"create_at" json:"createAt"`
	UpdateAt *string `db:"update_at" json:"updateAt"`
}

type SearchCarBrandQuery struct {
	Name     string     `schema:"name"`
	DateFrom *time.Time `schema:"date_from"`
	DateTo   *time.Time `schema:"date_to"`
}

type SearchCarBrandResult struct {
	CarBrands []*CarBrand `json:"result"`
}

type CreateCarBrandCommand struct {
	Name string `json:"name"`
}

func (cmd *CreateCarBrandCommand) Validate() error {
	if len(cmd.Name) == 0 {
		return ErrNameEmpty
	}

	return nil
}

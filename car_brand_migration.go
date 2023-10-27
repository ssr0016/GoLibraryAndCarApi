package migrations

import (
	. "main/pkg/migrator"
)

func addCarBrandMigrations(mg *Migrator) {
	carBrand := Table{
		Name: "car_brand",
		Columns: []*Column{
			{Name: "id", Type: DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "name", Type: DB_NVarchar, Length: 190, Nullable: false},
			{Name: "create_by", Type: DB_NVarchar, Length: 190, Nullable: false},
			{Name: "update_by", Type: DB_NVarchar, Length: 190, Nullable: true},
			{Name: "create_at", Type: DB_DateTime, Nullable: false, Default: DB_NowTimeZoneUTC},
			{Name: "update_at", Type: DB_DateTime, Nullable: true, Default: DB_NowTimeZoneUTC},
		},
	}

	mg.AddMigration("create car_brand table", NewAddTableMigration(carBrand))
}

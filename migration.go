package migrations

import (
	. "main/pkg/migrator"
)

type Migrations struct {
}

func New() *Migrations {
	return &Migrations{}
}

func (m *Migrations) AddMigration(mg *Migrator) {
	addMigrationLogMigrations(mg)
	addSchedulerAssigneeMigrations(mg)
	addSchedulerMigrations(mg)
	addAssignmentMigrations(mg)
	addAssignmentLogMigrations(mg)
	addCarBrandMigrations(mg)
	//library1
	addBookMigration(mg)
	addCategoryMigration(mg)
	addAuthorMigration(mg)

}

func addMigrationLogMigrations(mg *Migrator) {
	migrationLogV1 := Table{
		Name: "migration_log",
		Columns: []*Column{
			{Name: "id", Type: DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "migration_id", Type: DB_NVarchar, Length: 255},
			{Name: "sql", Type: DB_Text},
			{Name: "success", Type: DB_Bool},
			{Name: "error", Type: DB_Text},
			{Name: "timestamp", Type: DB_DateTime},
		},
	}

	mg.AddMigration("create migration_log table", NewAddTableMigration(migrationLogV1))
}

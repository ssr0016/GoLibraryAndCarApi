package migrations

import (
	. "main/pkg/migrator"
)

func addAuthorMigration(mg *Migrator) {
	author := Table{
		Name: "author",
		Columns: []*Column{
			{Name: "id", Type: DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "name", Type: DB_NVarchar, Length: 190, Nullable: false},
		},
	}

	mg.AddMigration("create author table", NewAddTableMigration(author))

}

func addCategoryMigration(mg *Migrator) {
	category := Table{
		Name: "category",
		Columns: []*Column{
			{Name: "id", Type: DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "name", Type: DB_NVarchar, Length: 190, Nullable: false},
		},
	}
	mg.AddMigration("create category table", NewAddTableMigration(category))
}

func addBookMigration(mg *Migrator) {
	book := Table{
		Name: "book",
		Columns: []*Column{
			{Name: "id", Type: DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "title", Type: DB_NVarchar, Length: 190, Nullable: false},
			{Name: "author_id", Type: DB_BigInt, Nullable: false},
			{Name: "category_id", Type: DB_BigInt, Nullable: false},
			{Name: "published_at", Type: DB_DateTime, Nullable: false, Default: DB_NowTimeZoneUTC},
		},
		Indices: []*Index{
			{Cols: []string{"author_id", "category_id"}, Type: UniqueIndex},
		},
	}

	mg.AddMigration("create book table", NewAddTableMigration(book))
	mg.AddMigration("create index book.author_id.category_id", NewAddIndexMigration(book, book.Indices[0]))
}

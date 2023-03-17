package models

func CreateTables() {
	UserTable.Generate()
	ArticleTable.Generate()
}

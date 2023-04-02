package models

func CreateTables() {
	UserTable.Generate()
	ImgTable.Generate()
}

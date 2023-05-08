package models

var (
	UserTable = User{}
)

func CreateTables() {
	UserTable.Generate()
	ImgTable.Generate()
}

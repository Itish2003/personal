package initializers

import (
	"itish.github.io/model"
)

func SyncDB() {
	DB.AutoMigrate(&model.User{})
	CONTENTDB.AutoMigrate(&model.Blog{})
}

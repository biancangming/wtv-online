package model

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"os"
)

var (
	dbs *gorm.DB //SQLLITE
)

func InitModel() {
	cwd, _ := os.Getwd()
	dbs, _ = gorm.Open(sqlite.Open(cwd+"/gorm.db"), &gorm.Config{})
	err := dbs.AutoMigrate(&Share{})
	fmt.Println(err)
}

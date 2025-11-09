package initial

import (
	"go-web-template/internal/global"
	"go-web-template/internal/model/database"
	"go-web-template/internal/utils"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDBSqlite() {
	execPath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	execPath, err = filepath.EvalSymlinks(execPath)
	if err != nil {
		panic(err)
	}

	execDir := filepath.Dir(execPath)
	dbPath := filepath.Join(execDir, "data.db")

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logrus.Fatalln("Failed to connect database", err)
	}
	db.AutoMigrate(
		&database.User{},
	)
	var usersCount int64 = 0
	err = db.Model(&database.User{}).Count(&usersCount).Error
	if err != nil {
		logrus.Fatalln("Failed to count users table", err)
	}
	if usersCount == 0 {
		hash, err := utils.HashPassword("admin")
		if err != nil {
			logrus.Fatalln("Failed to generate password", err)
		}
		user := database.User{
			Name:    "admin",
			PWDHash: hash,
			IsAdmin: true,
		}
		err = db.Create(&user).Error
		if err != nil {
			logrus.Fatalln("Failed to create user", err)
		}
	}
	global.DB = db
}

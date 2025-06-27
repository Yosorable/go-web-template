package global

import (
	"go-web-template/internal/config"

	"gorm.io/gorm"
)

var (
	CONFIG config.Config
)

var (
	DB             *gorm.DB
	SKIP_AUTH_PATH = []string{"/login"}
)

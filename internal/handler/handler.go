package handler

import (
	"github.com/CrysLef/gopportunities/config"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *config.Logger
)

func Init() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}

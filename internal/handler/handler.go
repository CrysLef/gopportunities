package handler

import (
	"fmt"
	"reflect"

	"github.com/CrysLef/gopportunities/config"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *config.Logger
)

func typeOf(v interface{}) string {
	return fmt.Sprint(reflect.TypeOf(v))
}

func Init() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}

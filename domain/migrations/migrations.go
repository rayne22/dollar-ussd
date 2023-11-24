package migrations

import (
	"dollar-ussd/domain/models/screens"
	"gorm.io/gorm"
)

func Migrations(db *gorm.DB) {
	_ = db.Debug().AutoMigrate(&screens.Screen{})
	_ = db.Debug().AutoMigrate(&screens.Option{})

	//_ = db.Debug().AutoMigrate(&controllers.Session{})
}

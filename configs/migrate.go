package configs

import (
	"github.com/OctavianoRyan25/be-agriculture/modules/admin"
	"github.com/OctavianoRyan25/be-agriculture/modules/user"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&user.User{}, &admin.Admin{}); err != nil {
		return err
	}
	return nil
}

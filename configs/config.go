package configs

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	//Konfigurasi koneksi ke database MySQL
	// dbUser := os.Getenv("DB_USER")
	// dbPass := os.Getenv("DB_PASS")
	// dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")
	// dbName := os.Getenv("DB_NAME")

	dbUser := "root"
	dbPass := "@Demin123#"
	dbHost := "localhost"
	dbPort := "3306"
	dbName := "test_capstone"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	//dsn := "root:@tcp(localhost:3306)/minpro?charset=utf8mb4&parseTime=True&loc=Local"

	// Membuka koneksi ke database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

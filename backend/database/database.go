package database

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"airbd/database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func get_envs(db *models.Database) {
	var err error

	godotenv.Load()

	db.Port, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err)
	}
	db.Host = os.Getenv("DB_HOST")
	db.Pass = os.Getenv("DB_PASS")
	db.Name = os.Getenv("DB_NAME")
	db.User = os.Getenv("DB_USER")
}

func Newdatabase() models.Database {
	var db models.Database
	get_envs(&db)

	var dbtemp *gorm.DB
	fmt.Println(os.Getenv("DATABASE_URL"))
	dbtemp, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	db.Db = dbtemp
	if err != nil {
		panic(err)
	}

	db.Db.AutoMigrate(&models.User{})
	db.Db.AutoMigrate(&models.Article{})

	fmt.Print("Sucessfully connected !\n")
	return db
}

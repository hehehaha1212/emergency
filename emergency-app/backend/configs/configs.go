package configs

import(
"fmt"
"log"
"os"
"gorm.io/driver/postgres"
"gorm.io/gorm"
)

var DB *gorm.DB

func connectDB(){

  dsn:= fmt.sprintf(
	"host=%s user=%s password=%s name=%s port=%s sslmode=disable"
	os.Getenv=("DB_HOST"), 
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB.NAME"),
	os.Getenv("DB_PORT"),
	)
var err error
DB, err =gorm.open(postgres.open(dsn),&gorm.config{})
if err !=nil {
	log.fatal("failed to connect to database:",err)
}
fmt.println("connected to database!")
}

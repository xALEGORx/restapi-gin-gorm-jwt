package core

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Api struct {
	Db  *gorm.DB
	App *gin.Engine
	Log *logrus.Logger
}

// Initializes gorm, gin, logrus and run http server
func Run() {
	var err error
	api := Api{}

	// Init logrus
	api.Log = &logrus.Logger{
		Out:   os.Stdout,
		Level: logrus.DebugLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%]: %time% - %msg%\n",
		},
	}

	// Init .env config
	if err := godotenv.Load(); err != nil {
		api.Log.Fatal("failed load env config")
	}

	// Init gorm connection
	api.Db, err = gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SERVER"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed connection database: %s", err.Error())
	}
	api.Migrate()

	// Init gin engine
	gin.SetMode(gin.ReleaseMode)
	api.App = gin.Default()
	api.Routes()

	// Run http server
	api.App.Run(fmt.Sprintf(
		"%s:%s",
		os.Getenv("IP"),
		os.Getenv("PORT"),
	))
}

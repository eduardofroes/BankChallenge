package configs

import (
	"bankchallenge/commons"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
)

type Config struct {
	// Database struct is responsible to get all database configuration in application.yml.
	Database struct {
		Host      string `yaml:"host"`
		Port      string `yaml:"post"`
		DbName    string `yaml:"dbname"`
		User      string `yaml:"user"`
		Password  string `yaml:"password"`
		DriveName string `yaml:"drivename"`
	}
}

var (
	config *Config
)

func GetDatabase() *sql.DB {
	return databaseBuilder()
}

// DatabaseBuilder function is responsible to get the database conection.
func databaseBuilder() *sql.DB {

	absPath, _ := filepath.Abs("application.yml")

	if config == nil {
		config = loadDatabaseConfig(absPath)
	}

	databaseConfig := config.Database

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", databaseConfig.Host, databaseConfig.Port, databaseConfig.User, databaseConfig.Password, databaseConfig.DbName)

	database, err := sql.Open(databaseConfig.DriveName, dataSourceName)

	commons.CheckError(err, "Error to opening the database connection.")

	return database
}

// LoadDatabaseConfig function is responsible to get database configuration from environment variables.
func loadDatabaseConfig(applicationPath string) *Config {

	config := Config{}

	config.Database.Host = os.Getenv("DATABASE_HOST")
	config.Database.Port = os.Getenv("DATABASE_PORT")
	config.Database.User = os.Getenv("DATABASE_USER")
	config.Database.Password = os.Getenv("DATABASE_PASS")
	config.Database.DbName = os.Getenv("DATABASE_NAME")
	config.Database.DriveName = os.Getenv("DATABASE_DRIVE")

	return &config
}

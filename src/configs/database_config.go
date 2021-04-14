package configs

import (
	"bankchallenge/commons"
	"database/sql"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"sigs.k8s.io/yaml"

	_ "github.com/lib/pq"
)

type Config struct {
	// Database struct is responsible to get all database configuration in application.yml.
	Database struct {
		Host      string `yaml:"host"`
		Port      int32  `yaml:"post"`
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

	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", databaseConfig.Host, databaseConfig.Port, databaseConfig.User, databaseConfig.Password, databaseConfig.DbName)

	database, err := sql.Open(databaseConfig.DriveName, dataSourceName)

	commons.CheckError(err, "Error to opening the database connection.")

	return database
}

// LoadDatabaseConfig function is responsible to get database configuration from application file.
func loadDatabaseConfig(applicationPath string) *Config {

	buffer, err := ioutil.ReadFile(applicationPath)

	commons.CheckError(err, "Error to open the application.yml file.")

	config := &Config{}

	err = yaml.Unmarshal(buffer, config)

	commons.CheckError(err, "Error to parsing the application.yml.")

	return config
}

package config

import (
	"github.com/spf13/viper"
	"log"
)

// Initilize this variable to access the env values
var EnvConfigs *envConfigs

// Call this to load the env variables
func InitEnvConfigs() *envConfigs {
	EnvConfigs = loadEnvVariables()
	return EnvConfigs
}

type envConfigs struct {
	DBPort              int    `mapstructure:"DB_PORT"`
	DBHostName          string `mapstructure:"DB_HOST_NAME"`
	DBUserName          string `mapstructure:"DB_USER_NAME"`
	DBDatabaseName      string `mapstructure:"DB_DATABASE_NAME"`
	DBDatabasePassword  string `mapstructure:"DB_DATABASE_PASSWORD"`
	DBMaxOpenConnection int    `mapstructure:"DB_MAX_OPEN_CONNECTION"`
	DBMaxIdleConnection int    `mapstructure:"DB_MAX_IDLE_CONNECTION"`
	DBReadTimeout       string `mapstructure:"DB_READ_TIMEOUT"`
	DBDialTimeout       string `mapstructure:"DB_DIAL_TIMEOUT"`
}

// Call to load the variables from env
func loadEnvVariables() (config *envConfigs) {
	// Tell viper the path/location of your env file. If it is root just add "."
	viper.AddConfigPath(".")

	// Tell viper the name of your file
	viper.SetConfigName("mathOperation")

	// Tell viper the type of your file
	viper.SetConfigType("env")

	// Viper reads all the variables from env file and log error if any found
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	// Viper unmarshalls the loaded env variables into the struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return
}

package util

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Neo4jConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Uri      string `mapstructure:"uri"`
}

type Config struct {
	GinMode     string      `mapstructure:"gin_mode"`
	Neo4jConfig Neo4jConfig `mapstructure:"neo_4j"`
}

// DO NOT expose variables from config file (etc. DO NOT push to GitHub)
// config file is only for development use
// for app deployed on Koyeb, it uses environment variables set under Settings
func LoadConfig() *Config {
	ginMode, hasGinMode := os.LookupEnv("GIN_MODE")
	if hasGinMode {
		gin.SetMode(ginMode)
	} else {
		ginMode = gin.Mode()
	}

	fmt.Printf("Started with Gin Mode: %s\n", ginMode)

	// Tell viper location of config file {mode}.config.yaml on local machine
	configFileName := fmt.Sprintf("%s.config", ginMode)
	configFileExt := "yaml"
	configFilePath := "./config"
	viper.AddConfigPath(configFilePath)
	viper.SetConfigName(configFileName)
	viper.SetConfigType(configFileExt)
	viper.AutomaticEnv()

	// Load environment variables from config file
	err := viper.ReadInConfig()
	fmt.Printf("Finding and reading %s file...\n", viper.ConfigFileUsed())
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found, should be loaded using runtime environment with viper.AutomaticEnv()
			fmt.Println("Config file not found. Loaded runtime environment variables")
		} else {
			// Config file was found but another error was produced
			log.Fatalf("Error loading %s file\n", viper.ConfigFileUsed())

		}
	}

	// fmt.Println("Config file found. Loading environment variables...")

	// var config Config

	// if err := viper.Unmarshal(&config); err != nil {
	// 	log.Fatalf("Failed to decode config into struct, %v", err)
	// }

	// for _, key := range viper.AllKeys() {
	// 	// set in os env
	// 	err := os.Setenv(key, viper.GetString(key))
	// 	if err == nil {
	// 		fmt.Printf("Loaded env key %s\n", key)
	// 	} else {
	// 		log.Fatalf("Error loading env key %s\n", key)
	// 	}
	// }

	var config *Config

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Failed to decode config into struct, %v", err)
	}
	fmt.Println("Loaded environment variables")

	// Watch changes from environment variables powered by viper
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()
	fmt.Println("Start watching changes from environment variables...")

	return config
}

var GlobalConfig *Config = LoadConfig()

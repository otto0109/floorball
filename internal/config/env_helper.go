package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var (
	ErrNoEnvFile          = errors.New("no .env file in the provided directory")
	ErrUnknownEnvironment = errors.New("the provided environment is unknown (valid environments are production, development, test)")
)

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func loadEnv(env string, folder string) error {
	validEnv := map[string]bool{"production": true, "development": true, "test": true}
	if !validEnv[env] {
		return ErrUnknownEnvironment
	}

	envFileLocal := fmt.Sprintf("%s.env.%s.local", folder, env)
	if err := loadConfigFile(envFileLocal); err != nil {
		return err
	}

	localEnvFile := fmt.Sprintf("%s.env.local", folder)
	if err := loadConfigFile(localEnvFile); err != nil {
		return err
	}

	envFile := fmt.Sprintf("%s.env.%s", folder, env)
	if err := loadConfigFile(envFile); err != nil {
		return err
	}

	defaultEnvFile := fmt.Sprintf("%s.env", folder)
	if err := loadConfigFile(defaultEnvFile); err != nil {
		return err
	}
	if !fileExists(defaultEnvFile) {
		return ErrNoEnvFile
	}
	return nil
}

func loadConfigFile(fileName string) error {
	err := godotenv.Load(fileName)
	if err != nil && fileExists(fileName) {
		return err
	}
	return nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

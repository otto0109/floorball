package config

import "fmt"

type DatabaseConfig struct {
	BaseURL  string
	Port     int
	Database string
	Config   string
	UserName string
	Password string
}

func LoadDatabaseConfig(environment string, folder string) (*DatabaseConfig, error) {
	if err := loadEnv(environment, folder); err != nil {
		return nil, err
	}

	return &DatabaseConfig{
		BaseURL:  getEnv("SQLDATABASE_ENDPOINT", "http://localhost"),
		Port:     getEnvAsInt("SQLDATABASE_PORT", 1433),
		Database: getEnv("SQLDATABASE_NAME", "Floorball"),
		Config:   getEnv("SQLDATABASE_CONFIG", "parseTime=true"),
		UserName: getEnv("SQLDATABASE_ENDPOINT_USER", "root"),
		Password: getEnv("SQLDATABASE_ENDPOINT_PASSWORD", ""),
	}, nil
}

func (config *DatabaseConfig) ConnectionString() string {
	return fmt.Sprintf("server=%s;user id =%s;password=%s;port=%d;database=%s;%s",
		config.BaseURL,
		config.UserName,
		config.Password,
		config.Port,
		config.Database,
		config.Config,
	)
}

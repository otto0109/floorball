package config

type VicciConfig struct {
	BaseURL  string
	UserName string
	Password string
}

func LoadVicciBaseURL(environment string, folder string) (*VicciConfig, error) {
	if err := loadEnv(environment, folder); err != nil {
		return nil, err
	}

	return &VicciConfig{
		BaseURL:  getEnv("USERSERVICE_VICCI_BASE_URL", "http://localhost"),
		UserName: getEnv("auth_user", "root"),
		Password: getEnv("password", ""),
	}, nil
}

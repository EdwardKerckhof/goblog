package configs

import "github.com/spf13/viper"

type Config struct {
	ApiPort           int    `mapstructure:"API_PORT"`
	ApiVersion        string `mapstructure:"API_VERSION"`
	ApiOriginsAllowed string `mapstructure:"ORIGINS_ALLOWED"`
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBHost            string `mapstructure:"DB_HOST"`
	DBPort            string `mapstructure:"DB_PORT"`
	DBUser            string `mapstructure:"DB_USER"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	DBName            string `mapstructure:"DB_NAME"`
}

// LoadConfig reads in an env file and loads into a config struct
func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

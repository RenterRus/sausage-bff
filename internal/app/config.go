package app

import (
	"fmt"

	validator "github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type Server struct {
	Host string `validate:"required"`
	Port int    `validate:"required"`
}

type Config struct {
	Server Server `validate:"required"`
}

func ReadConfig(path string, fileName string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(fileName)
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("ReadConfig: %w", err)
	}

	b, err := yaml.Marshal(viper.AllSettings())
	if err != nil {
		return nil, fmt.Errorf("ReadConfig (Marshal): %w", err)
	}

	res := &Config{}
	err = yaml.Unmarshal(b, res)
	if err != nil {
		return nil, fmt.Errorf("ReadConfig (Unmarshal): %w", err)
	}

	if err := valid(res); err != nil {
		return nil, fmt.Errorf("ReadConfig (Validate): %w", err)
	}

	return res, nil
}

func valid(conf *Config) error {
	v := validator.New(validator.WithRequiredStructEnabled())
	if err := v.Struct(conf); err != nil {
		return fmt.Errorf("validation faild: %w", err)
	}

	return nil
}

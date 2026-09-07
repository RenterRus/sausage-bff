package app

import (
	"fmt"
)

type App struct {
	conf *Config
}

func NewApp(configPath string) (*App, error) {
	lastSlash := 0
	for i, v := range configPath {
		if v == '/' {
			lastSlash = i
		}
	}

	conf, err := ReadConfig(configPath[:lastSlash], configPath[lastSlash+1:])
	if err != nil {
		return nil, fmt.Errorf("ReadConfig: %w", err)
	}

	return &App{
		conf: conf,
	}, nil
}

func (a *App) Run() error {

	return nil
}

func (a *App) Close() {
}

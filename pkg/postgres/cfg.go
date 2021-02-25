package postgres

import (
	"fmt"

	"github.com/spf13/viper"
)

type Cfg struct {
	ConnString string
}

func NewCfg(filename string, configType string) (Cfg, error) {
	viper.SetConfigName(filename)
	viper.SetConfigType(configType)
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return Cfg{}, fmt.Errorf("%w error in read", err)
	}

	cfg := Cfg{
		viper.GetString("db.connstring"),
	}

	return cfg, nil
}

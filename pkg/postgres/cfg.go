package postgres

import (
	"log"

	"github.com/spf13/viper"
)

type Cfg struct {
	ConnString string
}

func NewCfg(filename string) Cfg {
	viper.SetConfigName(filename)
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("can't read", err)
	}

	cfg := Cfg{
		viper.GetString("db.connstring"),
	}
	return cfg
}

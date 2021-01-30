package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Dest     string
	Emoticon map[string]string
}

func Load() (*Config, error) {
	v := viper.New()
	v.SetConfigName("emote") // name of config file (without extension!!!)
	v.AddConfigPath(".")     // path to look for the config file in

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = v.Unmarshal(c)
	return c, err
}

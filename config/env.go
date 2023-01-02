/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * created at : 02/01/23, 09:46
 * Copyright (c) 2022
 */

package config

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

type Environment interface {
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	Init(string)
}

type viperEnvironment struct{}

func (v *viperEnvironment) Init(prefix string) {
	viper.SetEnvPrefix(`go-clean`)
	viper.AutomaticEnv()

	osEnv := os.Getenv("OS_ENV")

	env := "env"
	if osEnv != "" {
		env = osEnv
	}

	if prefix != "" {
		env = prefix + "." + env
	}

	replacer := strings.NewReplacer(`.`, `_`)
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType(`yaml`)
	viper.SetConfigFile(env + `.yaml`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
}

func (v *viperEnvironment) GetString(key string) string {
	return viper.GetString(key)
}

func (v *viperEnvironment) GetStringMap(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}

func NewEnvironment() Environment {
	v := &viperEnvironment{}
	v.Init("")
	return v
}

/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 17/10/22, 10:17
 * Copyright (c) 2022
 */

package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type MysqlConfig struct {
	User     string `json:"users"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

func NewMysqlConnector(param map[string]interface{}) (*sql.DB, error) {
	paramsByte, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	conf := new(MysqlConfig)
	if err := json.Unmarshal(paramsByte, &conf); err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

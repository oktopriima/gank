/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * created at : 02/01/23, 09:46
 * Copyright (c) 2022
 */

package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"io/fs"
	"log"
	"os"
	"time"
)

type MysqlConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	LogPath  string `json:"log_path"`
	Debug    string `json:"debug"`
}

func NewMysqlConnector(param map[string]interface{}) (*gorm.DB, error) {
	paramsByte, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	conf := new(MysqlConfig)
	if err := json.Unmarshal(paramsByte, &conf); err != nil {
		return nil, err
	}

	var dbLogFile *os.File
	dbLogFile, err = os.OpenFile(fmt.Sprintf("%s/database.log", conf.LogPath), os.O_CREATE|os.O_RDWR|os.O_APPEND, fs.ModeAppend)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		_ = os.Mkdir(conf.LogPath, os.ModePerm)
		dbLogFile, err = os.Create(fmt.Sprintf("%s/database.log", conf.LogPath))
	}

	dbLogger := logger.New(
		log.New(io.MultiWriter(os.Stdout, dbLogFile), "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	gormConfig := &gorm.Config{
		// enhance performance config
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		Logger:                 dbLogger,
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
	)

	sqlMaster, err := apmsql.Open("mysql", dsn)
	if err != nil {
		panic("error during construct APM, please check the config")
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlMaster.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlMaster.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlMaster.SetConnMaxLifetime(time.Hour)

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlMaster,
	}), gormConfig)
	if err != nil {
		panic("couldn't connect to database, caused: " + err.Error())
	}

	return db, nil
}

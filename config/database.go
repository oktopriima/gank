/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * created at : 02/01/23, 09:46
 * Copyright (c) 2022
 */

package config

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type databaseInstance struct {
	mysql *gorm.DB
	pool  *redis.Pool
	mDb   *mongo.Database
}

type DatabaseInstance interface {
	MySql() *gorm.DB
	Redis() *redis.Pool
	MongoDb() *mongo.Database
}

func NewDatabaseInstance(env Environment) DatabaseInstance {
	ins := new(databaseInstance)

	mysql, err := NewMysqlConnector(env.GetStringMap("mysql"))
	if err != nil {
		panic(fmt.Sprintf("failed connect to database. error : %s", err.Error()))
	}

	ins.mysql = mysql

	return ins
}

func (d *databaseInstance) MySql() *gorm.DB {
	return d.mysql
}

func (d *databaseInstance) Redis() *redis.Pool {
	return d.pool
}

func (d *databaseInstance) MongoDb() *mongo.Database {
	return d.mDb
}

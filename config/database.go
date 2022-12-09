/*
 * Name : Okto Prima Jaya
 * GitHub : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 17/10/22, 10:07
 * Copyright (c) 2022
 */

package config

import (
	"database/sql"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

type databaseInstance struct {
	mysql *sql.DB
	pool  *redis.Pool
	mDb   *mongo.Database
}

type DatabaseInstance interface {
	MySql() *sql.DB
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

func (d *databaseInstance) MySql() *sql.DB {
	return d.mysql
}

func (d *databaseInstance) Redis() *redis.Pool {
	return d.pool
}

func (d *databaseInstance) MongoDb() *mongo.Database {
	return d.mDb
}

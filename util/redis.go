package util

import (
	"strconv"

	"github.com/gomodule/redigo/redis"
)

type RedisUtil struct {
	client redis.Conn
}

// 全局变量, 外部使用utils.RedisClient来访问
var RedisClient RedisUtil

// 初始化redis
func InitRedisUtil(address string, port int, pwd string) (*RedisUtil, error) {
	//连接redis
	client, err := redis.Dial("tcp", address+":"+strconv.Itoa(port))
	if err != nil {
		panic("failed to redis:" + err.Error())
	}
	//验证redis redis的配置文件redis.conf中一定要设置quirepass=password, 不然连不上
	_, err = client.Do("auth", pwd)
	if err != nil {
		panic("failed to auth redis:" + err.Error())
	}
	//初始化全局redis结构体
	RedisClient = RedisUtil{client: client}
	return &RedisClient, nil
}

// 设置数据到redis中（string）
func (rs *RedisUtil) SetStr(key string, value string) error {
	_, err := rs.client.Do("Set", key, value)
	return err
}

// 设置数据到redis中（string）
func (rs *RedisUtil) SetStrNotExist(key string, value string, expireSecond int) bool {
	val, err := rs.client.Do("SET", key, value, "EX", expireSecond, "NX")
	if err != nil || val == nil {
		return false
	}
	return true
}

// 设置数据到redis中（string）
func (rs *RedisUtil) SetStrWithExpire(key string, value string, expireSecond int) error {
	_, err := rs.client.Do("Set", key, value, "ex", expireSecond)
	return err
}

// 获取redis中数据（string）
func (rs *RedisUtil) GetStr(key string) (string, error) {
	val, err := rs.client.Do("Get", key)
	if err != nil {
		return "", err
	}
	return string(val.([]byte)), nil
}

// 设置数据到redis中（hash）
func (rs *RedisUtil) HSet(key string, field string, value string) error {
	_, err := rs.client.Do("HSet", key, field, value)
	return err
}

// 设置数据到redis中（hash）
func (rs *RedisUtil) HGet(key string, field string) (string, error) {
	val, err := rs.client.Do("HGet", key, field)
	if err != nil {
		return "", err
	}
	return string(val.([]byte)), nil
}

// 删除
func (rs *RedisUtil) DelByKey(key string) error {
	_, err := rs.client.Do("DEL", key)
	return err
}

// 设置key过期时间
func (rs *RedisUtil) SetExpire(key string, expireSecond int) error {
	_, err := rs.client.Do("EXPIRE", key, expireSecond)
	return err
}

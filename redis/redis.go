// redis.go
package redis

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

var String = redis.String

type RedisRes struct {
	Cont interface{}
}

type RedisInstance struct {
	Flag         string
	Dsn          string
	MaxIdle      int
	ConnTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Password     string
}

var redisPool = make(map[string]*redis.Pool)
var RedisPool = &redisPool

func New(flag string) (r redis.Conn) {

	r = (*RedisPool)[flag].Get()
	return
}

func regRedisInstance(redisInstance RedisInstance) (err error) {

	flag := redisInstance.Flag
	dsn := redisInstance.Dsn
	maxIdle := redisInstance.MaxIdle
	idleTimeout := 240 * time.Second

	connTimeout := redisInstance.ConnTimeout * time.Millisecond
	readTimeout := redisInstance.ReadTimeout * time.Millisecond
	writeTimeout := redisInstance.WriteTimeout * time.Millisecond

	redisPool[flag] = &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: idleTimeout,
		Dial: func() (c redis.Conn, err error) {
			c, err = redis.DialTimeout("tcp", dsn, connTimeout, readTimeout, writeTimeout)
			if err != nil {
				return
			}
			if _, err = c.Do("AUTH", redisInstance.Password); err != nil {
				return
			}
			return
		},
		TestOnBorrow: pingRedis,
	}
	instanceFlag := "RegRedisInstance"
	fmt.Printf("%-20s: %-10s [ %s ]\n", instanceFlag, flag, dsn)
	return
}

func regRedisInstanceMulti(redisInstances []RedisInstance) (err error) {
	for _, redisInstance := range redisInstances {
		err = regRedisInstance(redisInstance)
		if err != nil {
			return
		}
	}
	return
}

func Init(redisInstances []RedisInstance) {
	regRedisInstanceMulti(redisInstances)

}

func pingRedis(c redis.Conn, t time.Time) (err error) {
	_, err = c.Do("ping")
	if err != nil {
		return err
	}
	return
}

func (this *RedisInstance) Init(flag string,
	dsn string,
	maxIdle int,
	connTimeout int,
	readTimeout int,
	writeTimeout int,
	password string) {

	this.Flag = flag
	this.Dsn = dsn
	this.MaxIdle = maxIdle
	this.ConnTimeout = time.Duration(connTimeout) * time.Millisecond
	this.ReadTimeout = time.Duration(readTimeout) * time.Millisecond
	this.WriteTimeout = time.Duration(writeTimeout) * time.Millisecond
	this.Password = password

}

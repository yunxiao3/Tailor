package client

import (
	"github.com/go-redis/redis"
	"log"
)

type RedisCli struct {
	Addr string
	conn *redis.Client
	Mode string
}

func (this *RedisCli) Conn() {

	this.conn = redis.NewClient(&redis.Options{
		Addr: this.Addr,
		Password: "",
		DB: 0,
	})

}

func (this *RedisCli) GetV(key string) string {

	val, err := this.conn.Get(key).Result()
	if err != nil {
		log.Fatal(err)
	}
	return val

}

func (this *RedisCli) SetV(key string, value string) error {

	err := this.conn.Set(key, value, 0).Err()

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (this *RedisCli) Close() {
	this.conn.Close()
}


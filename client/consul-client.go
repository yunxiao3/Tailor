package client

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/ngaut/log"
)

type ConsulCli struct {
	Addr string
	conn *api.Client
	Mode string
}

func (this *ConsulCli) Conn() {

	var err error

	this.conn, err = api.NewClient(api.DefaultConfig())

	if err != nil {
		log.Fatal(err)
	}

}

func (this *ConsulCli) GetV(key string) string{

	kv := this.conn.KV()
	if p, _, err := kv.Get(key, nil); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(p.Key, "=>", string(p.Value))
	}
	return ""
}

func (this *ConsulCli) SetV(key string, value string) error {

	kv := this.conn.KV()
	pair := &api.KVPair{Key:key, Value: []byte(value)}
	if _, err := kv.Put(pair, nil); err != nil {
		log.Fatal(err)
	}

	return nil
}

func (this *ConsulCli) Close() {

}
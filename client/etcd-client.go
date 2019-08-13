package client

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/client"
	"github.com/ngaut/log"
	"strings"
	"time"
)

type EtcdCli struct {
	Addr string
	conn client.Client
	Mode string
}

func (this *EtcdCli) Conn() {

	var err error
	cfg := client.Config{
		Endpoints:               strings.Split(this.Addr, ","),
		Transport:               client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}

	this.conn, err = client.New(cfg)

	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

}

func (this *EtcdCli) GetV(key string) string{

	kapi := client.NewKeysAPI(this.conn)
	resp, err := kapi.Get(context.Background(), key, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Get is done. Metadata is %q\n", resp)
		fmt.Println("%q key has %q value\n", resp.Node.Key, resp.Node.Value)
	}

	return ""
}

func (this *EtcdCli) SetV(key string, value string) error {

	kapi := client.NewKeysAPI(this.conn)

	resp, err := kapi.Set(context.Background(), key, value, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Set is done. Metadata is %q\n", resp)
	
	return nil
}

func (this *EtcdCli) Close() {
}
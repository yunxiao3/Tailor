package client

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"strings"
	"time"
)

type ZkCli struct {
	Addr string
	conn *zk.Conn
	Mode string
}

func (this *ZkCli) Conn() {

	var hosts= strings.Split(this.Addr, ",")
	var err error
	this.conn, _, err = zk.Connect(hosts, time.Second*5)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func (this *ZkCli) GetV(key string) string{

	var path = "/zk_test/" + key
	v, _, err := this.conn.Get(path)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	fmt.Printf("value of path[%s]=[%s].\n", path, v)
	return string(v)
}

func (this *ZkCli) SetV(key string, value string) error {

	var path = "/zk_test/" + key
	var data = []byte(value)
	var flags int32 = 0
	var acls = zk.WorldACL(zk.PermAll)

	p, err := this.conn.Create(path, data, flags, acls)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("created:", p)
	return nil
}

func (this *ZkCli) Close() {
	this.conn.Close()
}
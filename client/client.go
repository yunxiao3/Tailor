package client

type Client interface {
	Conn()
	GetV(key string) string
	SetV(key string, value string) error
	Close()
}

func Factory(name string, addr string, mode string) Client {
	switch name {
	case "zk":
		return &ZkCli{Addr: addr, Mode:mode}
	case "etcd":
		return &EtcdCli{Addr: addr, Mode:mode}
	case "consul":
		return &ConsulCli{Addr: addr, Mode:mode}
	case "redis":
		return &RedisCli{Addr: addr, Mode:mode}
	case "mongo":
		return &MongoCli{Addr: addr, Mode:mode}
	default:
		panic("No such client")
	}
}

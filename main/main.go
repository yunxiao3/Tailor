package main

import (
	"github.com/Tailor/client"
	"github.com/gin-gonic/gin"
	"github.com/jessevdk/go-flags"
	"log"
	"os"
)

type setKeyData struct {
	Value string `json:"value"`
}

type Opts struct {
	Addr  string `long:"addr" env:"ADDR" default:"127.0.0.1:2181" description:"addr"`
	ServerType  string `long:"server_type" env:"SERVER_TYPE" default:"127.0.0.1:2181" description:"server type"`
	Mode  string `long:"mode" env:"MODE" default:"0" description:"consistency mode"`
}

func getKeyView(addr string, server_type string, mode string) func(*gin.Context)  {

	view := func(c *gin.Context) {

		cli := client.Factory(server_type,  addr, mode)
		cli.Conn()
		key := c.Param("key")
		result := cli.GetV(key)

		c.JSON(200, gin.H{
			"value": result,
		})
		defer cli.Close()

	}

	return view
}

func setKeyView(addr string, server_type string, mode string) func(*gin.Context) {

	view := func(c *gin.Context) {
		key := c.Param("key")
		var data setKeyData
		err := c.BindJSON(&data)
		if err != nil {
			panic(err)
		}

		cli := client.Factory(server_type,  addr, mode)
		cli.Conn()
		err = cli.SetV(key, data.Value)

		if err == nil {
			c.JSON(200, gin.H{
				"value": data.Value,
			})
		}

		defer cli.Close()

	}

	return view
}



func main()  {

	var opts Opts
	p := flags.NewParser(&opts, flags.Default)
	if _, err := p.ParseArgs(os.Args[1:]); err != nil {
		log.Panicln(err)
	}
	router := gin.Default()
	router.POST("/keys/:key/", setKeyView(opts.Addr, opts.ServerType, opts.Mode))
	router.GET("/keys/:key/", getKeyView(opts.Addr, opts.ServerType, opts.Mode))

	router.Run(":9988")

}

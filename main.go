package main

import (
	"flag"
	"fmt"
	. "ginweb02/Controller"
	"ginweb02/Server"
	"ginweb02/internal/config"
	"ginweb02/internal/handler"
	"ginweb02/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "./etc/open-api.yaml", "the config file")

func main() {
	gozero()
}

func gozero() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func ginweb01() {
	Server.
		Init().
		Route(NewUserController()).
		GroupRouter("v1", NewUserController()).
		GroupRouter("v2", NewUserController()).
		GroupRouter("v3", NewUserController()).
		Listen()
}

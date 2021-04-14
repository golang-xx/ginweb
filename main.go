package main

import (
	. "ginweb02/Controller"
	"ginweb02/Server"
)

func main() {
	Server.
		Init().
		Route(NewUserController()).
		GroupRouter("v1", NewUserController()).
		Listen()
}

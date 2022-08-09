package main

import (
	"github.com/mikhailbuslaev/eshop/server"
)

func main() {
	server := server.New()
	server.Run()
}

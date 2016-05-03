package main

import (
	"github.com/MDFS/MDFS/server"
)

func main() {
	var s server.StorageNode
	server.Start(&s)
}

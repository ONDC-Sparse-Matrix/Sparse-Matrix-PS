package utils

import (
	"cache-server/types"
)

func InitServer() {
	server := types.ServerInfo{
		Host: "192.168.220.208",
		Port: "4000",
	}
	types.Server = server
}

package main

import (
	api "github.com/scalent-sushil/user-management-go/cmd"
	"github.com/scalent-sushil/user-management-go/database"
)

func main() {
	database.Connect()
	api.Run()
}

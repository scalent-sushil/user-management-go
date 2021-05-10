package main

import (
	"github.com/scalent-sushil/user-management-go/database"
	api "github.com/scalent-sushil/user-management-go/pkg"
)

func main() {
	database.Connect()
	api.Run()
}

package cmd

import (
	"fmt"
	"log"
	"net/http"
	"github.com/scalent-sushil/user-management-go/cmd/auto"
	"github.com/scalent-sushil/user-management-go/cmd/config"

	"github.com/scalent-sushil/user-management-go/cmd/router"
)

func Run() {
	config.Load()
	// auto.Load()
	fmt.Printf("\n\tListening [::]:%d\n", config.PORT)
	listen(config.PORT)
}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

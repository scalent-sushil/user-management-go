package pkg

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	// "github.com/scalent-sushil/user-management-go/cmd/auto"

	"github.com/scalent-sushil/user-management-go/cmd/router"
)

var (
	PORT      = 0
	SECRETKEY []byte
	DBURL     = ""
	// DBDRIVER = ""
)

func Run() {
	var err error
	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		PORT = 8080
	}
	DBURL = fmt.Sprintf(os.Getenv("DB_URL"))
	SECRETKEY = []byte(os.Getenv("API_SECRET"))
	// auto.Load()
	fmt.Printf("\n\tListening [::]:%d\n", PORT)
	listen(PORT)
}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

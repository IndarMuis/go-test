package main

import (
	"fmt"
	"github.com/IndarMuis/pt-xyz.git/src/common/config"
	"github.com/IndarMuis/pt-xyz.git/src/common/exception"
	"github.com/IndarMuis/pt-xyz.git/src/common/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	//init database
	config.NewDatabase()

	router := httprouter.New()

	// run server
	serverAddress := fmt.Sprintf("localhost:%s", utils.ViperEnv("PORT"))
	server := http.Server{
		Handler: router,
		Addr:    serverAddress,
	}
	err := server.ListenAndServe()
	exception.PanicIfError(err)
}

package main

import (
	"api/conf"
	"api/db"
	"api/router"
	"api/service_chain"
	"api/utils"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main()  {

	defer db.Db.Close()
	defer service_chain.Client.Close()

	log := utils.Log()

	gin.SetMode(conf.Conf.Server.Model)

	router := router.INitRouter()

	srv := &http.Server{
		Addr: conf.Conf.Server.Address,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed{
			log.Fatal("listen: %s \n",err)
		}
		log.Fatal("listen: %s \n",conf.Conf.Server.Address)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit,os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:",err)
	}
	log.Println("Server exiting")
}

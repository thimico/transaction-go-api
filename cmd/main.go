package main

import (
	"flag"
	"log"
	"net/http"
	"projeto-pismo-api-go/pkg/api/controller"
	"projeto-pismo-api-go/pkg/api/provider/mongo/dao"
	"projeto-pismo-api-go/pkg/api/router"
	"projeto-pismo-api-go/pkg/api/service"
	"projeto-pismo-api-go/pkg/utl/config"
	"projeto-pismo-api-go/pkg/utl/mg"
)

// @title Pismo Challenge API
// @version 1.0
// @description Pismo Challenge API
// @termsOfService http://swagger.io/terms/
// @contact.name Thiago Menezes
// @contact.email thg.mnzs@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {

	cfgPath := flag.String("p", "./cmd/conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	if err != nil {
		log.Fatal("error get config file")
	}

	client, database, err := mg.New(cfg.DB.PSN, cfg.DB.DB)
	if err != nil {
		log.Fatal("error get database")
	}

	halthDao := dao.NewMongoAccount(client, database)
	accountDAO := dao.NewMongoAccount(client, database)
	transactionDAO := dao.NewMongoTransaction(client, database)

	accountService := service.NewAccount(accountDAO)
	transactionService := service.NewTransaction(transactionDAO)

	accountHandler := controller.NewAccountController(accountService)
	transactionHandler := controller.NewTransactionController(transactionService)
	healthHandler := controller.NewHealthHandler(halthDao)

	routers := router.NewRARouter(accountHandler, transactionHandler, healthHandler)

	log.Println("Servidor esta rodando na porta " + cfg.Server.Port)
	log.Fatal(http.ListenAndServe(cfg.Server.Port, routers))
}

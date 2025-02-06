package main

import (
	_ "GoProjects/docs"
	"GoProjects/pkg/controller"
	"GoProjects/pkg/db"
	"GoProjects/pkg/repository"
	"GoProjects/pkg/router"
	"GoProjects/pkg/service"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"log/slog"
	"net/http"
)

func main() {

	//	@title			SWAGGER REST API BACKEND 0.1
	//	@version		0.1
	//  @BasePath       /
	//	@description	THIS SIMPLE TEST
	//	@termsOfService	http://swagger.io/terms/
	//	@host			localhost:8080
	connString := "postgres://postgres:postgres@localhost:5432/postgres"
	database, err := db.New(connString)
	if err != nil {
		slog.Error(err.Error())
	}
	repo := repository.New(database)
	s := service.New(repo)
	qc := controller.NewQ(repo)
	c := controller.New(s)
	r := router.New("localhost:8080", c, qc)

	r.Router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler).Methods(http.MethodGet)

	err = r.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

package main

import (
	"github.com/dveleztx/golang-echo-gorm-graphql/handler"
	"github.com/dveleztx/golang-echo-gorm-graphql/datastore"
	"github.com/dveleztx/golang-echo-gorm-graphql/graphql"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	db, err := datastore.NewDB()
	logFatal(err)

	db.LogMode(true)
	defer db.Close()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handler.Welcome())

	// GraphQL Endpoint Here
	h, err := graphql.NewHandler(db)
	logFatal(err)
	e.POST("/graphql", echo.WrapHandler(h))

	err = e.Start(":3000")
	logFatal(err)

}

func logFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

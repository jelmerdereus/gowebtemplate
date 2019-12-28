package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jelmerdereus/goweb3/cli"
	"github.com/jelmerdereus/goweb3/controllers"
	"github.com/jelmerdereus/goweb3/datastore"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//RunApp is responsible for running the web application
func RunApp(address string, orm *datastore.DBORM) error {
	// root URL
	rootCtr := controllers.GetRoot

	// user API
	users, err := controllers.NewUserHandle(orm)
	if err != nil {
		return err
	}

	// routing
	r := gin.Default()
	r.GET("/", rootCtr)
	r.GET("/users", users.GetAll)
	r.GET("/users/:alias", users.GetByAlias)
	r.POST("/users", users.Create)
	r.PUT("/users/:id", users.Update)
	r.DELETE("/users/:id", users.Delete)

	// run the application
	return http.ListenAndServe(address, r)
}

func main() {
	// use arguments or environment variables; exit with reason and help text
	arg, err := cli.GetParams()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Fprintf(os.Stderr, "\nUsage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	// connect to postgres
	postgresConnectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		arg["address"], arg["pgport"], arg["pguser"], arg["pgdatabase"], arg["pgpass"])

	orm, err := datastore.NewORM("postgres", postgresConnectionString)
	if err != nil {
		panic(err)
	}

	// run the web application
	log.Fatal(RunApp(":8080", orm))
}

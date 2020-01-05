package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jelmerdereus/gowebtemplate/cli"
	"github.com/jelmerdereus/gowebtemplate/controllers"
	"github.com/jelmerdereus/gowebtemplate/datastore"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//RunApp is responsible for running the web application
func RunApp(address string, orm *datastore.DBORM) error {
	// repos
	eventRepo, err := datastore.NewEventRepo(orm)
	if err != nil {
		log.Fatal(err)
	}

	userRepo, err := datastore.NewUserRepo(orm)
	if err != nil {
		log.Fatal(err)
	}

	// apis
	rootCtr := controllers.GetRoot

	events := controllers.NewEventAPI(eventRepo)
	if err != nil {
		return err
	}

	users := controllers.NewUserAPI(userRepo)
	if err != nil {
		return err
	}

	// routing
	r := gin.Default()
	r.GET("/", rootCtr)

	userRoute := r.Group("/users")
	{
		userRoute.GET("/", users.GetAll)
		userRoute.GET("/:alias", users.GetByAlias)
		userRoute.POST("/", users.Create)
		userRoute.PUT("/:id", users.Update)
		userRoute.DELETE("/:id", users.Delete)
	}

	eventRoute := r.Group("/events")
	{
		eventRoute.GET("/", events.GetAll)
		eventRoute.POST("/", events.Create)
		eventRoute.GET("/:id", events.GetByID)
		eventRoute.PUT("/:id", events.Update)
		eventRoute.DELETE("/:id", events.Delete)
	}

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

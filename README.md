# gowebtemplate
base structure for simple REST API implementations in Go

\*work in progress\*

## About
this package is meant to be used as a base structure for Go projects involving REST APIs

Goals:
* Offer a practical structure for a basic web application
* Use a modular and practical setup
* Support for any database supported by GORM

Dependencies:
* Gin framework
* Gorm database ORM

## Features
The following features are currently supported:
* REST API
* Users and operations

## Roadmap
Important items:
* events (new user, deleted user, logged in)
* roles
* authentication and authorization

## Requirements
you need a database that is supported by gorm see also [the Gorm documentation](https://gorm/io) . you can insert the ORM into the `RunApp` function like below:

```go
// connect to postgres
postgresConnectionString := //...

orm, err := datastore.NewORM("postgres", postgresConnectionString)
if err != nil {
    panic(err)
}

// run the web application
log.Fatal(RunApp(":8080", orm))
```

you need to have the following environment variables or cli flags to connect to postgres:
* pghost
* pgdatabase
* pguser
* pgpass


## Design
There is an interface for REST API handlers that is extended by subinterfaces to add additional methods.
The `GetAll` method does not have to mean 'get all'. It can for instance have a limit and handle filter criteria.

There is also a datastore.DBLayer interface for abstracting database operations

```go
//RESTAPI is an interface for implementation by REST controllers
type RESTAPI interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

//Handle contains the orm layer and is embedded by the implementations of RESTAPI
type Handle struct {
	DB datastore.DBLayer
}
```

The User and Event APIs have their own implementation of these types like for the User object

```go
//UserAPI is the interface for the user API
type UserAPI interface {
	RESTAPI

	//additional methods
	GetByAlias(c *gin.Context)
}

//UserHandle is a handle for the user API
type UserHandle struct {
	Handle
}

//NewUserAPI takes an ORM and returns a UserAPI
func NewUserAPI(orm *datastore.DBORM) (UserAPI, error) {
	if orm == nil {
		return nil, errors.New("No ORM provided")
	}
	orm.AutoMigrate(&models.User{})
	return &UserHandle{Handle: Handle{DB: orm}}, nil
}
```


## Getting started


in the main function, you can see that the environment variables are parsed and that an ORM instance was provided to the RunApp function.


```go
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
	postgresConnectionString := //...

	orm, err := datastore.NewORM("postgres", postgresConnectionString)
	if err != nil {
		panic(err)
	}

	// run the web application
    log.Fatal(RunApp(":8080", orm))
}
```

the `GetParams` function can be adapted to your needs with regards to environment variables and cli flags.


the `RunApp` function below specifies all the routes and starts the application.

```go
//RunApp is responsible for running the web application
func RunApp(address string, orm *datastore.DBORM) error {
	// root handler
	rootCtr := controllers.GetRoot

	// user api handler
	users, err := controllers.NewUserHandle(orm)
	if err != nil {
		return err
	}

	// event handler
	events, err := controllers.NewEventHandle(orm)
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
	}

	// run the application
	return http.ListenAndServe(address, r)
}
```

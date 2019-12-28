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

## Getting started
in the main function, you can see that the environment variables are parsed and that an ORM instance was provided to the RunApp function that specifies all the current routes. the `GetParams` function can be adapted to your needs.


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

the `RunApp` function specifies all the routes and starts the application

```go
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
```

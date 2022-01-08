
# Store Management App

A simple RESTful API for store management application built using Go language

Installation & Run
# Download this project
go get github.com/sangramkonde/store-management-app

Before running API server, you should set the database config with yours or set the your database config with my values on config.go

func Connect(){
	d, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/storedb?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	db = d
}

# Build and Run
cd store-management-app
cd cmd
go build
./store-management-app

# API Endpoint : http://127.0.0.1:9010
Structure
├── app
│   ├── cmd
|   |   ├── main.go
│   ├── pkg          // Our API core handlers
│   │   ├──config
|   |   |    |──app.go    // Common response functions
│   │   ├──controllers
|   |   |    |──store-controller.go  // APIs for Project model
│   │   └──models 
|   |   |    |──tasks.go
|   |   |──routes // APIs for Task model
│   |   |     |──store-routes.go
│   |   └── utils
|   |   |     |──utils.go

# API
/stores
* GET : Get all stores
* POST : Create a new store
/stores/:storeId
* GET : Get a project
* PUT : Update a project
* DELETE : Delete a project


# TODO

 Authentication with user for securing the APIs.
 Database Config
 Write the tests for all APIs.
 Organize the code with packages
 Make docs with GoDoc
 Building a deployment process
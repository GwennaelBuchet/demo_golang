# demo_golang
Just a simple and tiny project of Golang for a short demo.

It provides :
 - web service with CORS management
 - Unit test
 - interface, struct and virtual inheritance

## Setup
Install required libraries for our tiny web server
```commandline
go get github.com/rs/cors
go get -u github.com/gorilla/mux
```

## tests
Only few UT are provided here.
To launch them
```commandline
go test
```

## build and run
As fo every Golang projects, just use ```go run``` tool to build and run the project.
 
Server is setup to listen on 8090 port

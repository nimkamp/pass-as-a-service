# Pass-as-a-service API

This project contains several api endpoints displaying user and group information from files that have been parsed.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

Any Unix based system


### Run Unit tests
```
cd pass-as-a-service/etc/
go test
```

## Running Project 
You need to pass in two paramters when executing the binary for pass-as-a-service. The first parameter is the passwords file and the second parameter is the groups file.
```
go get -u github.com/go-chi/chi
git clone https://github.com/nimkamp/pass-as-a-service
cd pass-as-a-service
go build
./pass-as-a-service (/etc/passwd) (/etc/group)
```

## Built With

* [Chi](https://github.com/go-chi/chi) - Router for building HTTP services



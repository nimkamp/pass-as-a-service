# pass-as-a-service

This project contains several api endpoints displaying user and group information from files that have been parsed.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

Any Unix based system
Root access
Router for HTTP services



## Running the tests
cd pass-as-a-service/etc/
go test



## Deployment
git clone https://github.com/nimkamp/pass-as-a-service
cd pass-as-a-service
go build
./pass-as-a-service (/etc/passwd) (/etc/group)


## Built With

* [Chi](https://github.com/go-chi/chi) - Router for building HTTP services

## Authors

* **Nick Imkamp**


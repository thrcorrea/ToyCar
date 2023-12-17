# Toy Car Simulator


### Description

This application is a simulation of a toy car moving on a tabletop. It consists of a golang server application and a react app for showing the car and the tabletop graphically.

### Getting Started

To run both applications is necessary to run the docker-compose command, which will build and run both
server and web app

```
  docker-compose up
```


To run the server unit tests, its necessary to:

1. first to install the golang runtime

2. Install the used testing tools:
```
  go get https://github.com/vektra/mockery
  go get https://github.com/gotestyourself/gotestsum
```
3. Run the unit tests:
```golang
  // by makefile
  make mocks
  make unit-test
```



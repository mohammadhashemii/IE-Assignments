# RESTful API in GoLang
This project was the fifth project of the Internet Engineering course at Shahid Beheshti university.

# Requirements

The necessary package for running the server of this project is [Echo](https://echo.labstack.com). So you need to install it:

```shell
$ cd <PROJECT IN $GOPATH>
$ go get -u github.com/labstack/echo/
```

# How to Run

First, you must run the server:

```shell
$ cd PATH/TO/PROJECT
$ go build src/main.go
$ go run src/main.go
```
Now, your sever is listening on [http:/localhost:1323](http:/localhost:1323).

# Test and Debug
To test the server, I provided a script. You can run it through your terminal:
```shell
$ sh test.sh
```
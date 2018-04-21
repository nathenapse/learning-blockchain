# learning-blockchain
My first blockchain in Golang

- [INSTALLATION](#installation)
- [RUN](#run)
- [API](#api)
- [TODO](#todo)

## REQUIREMENTS
- Golang version > 1.9.1
- [Dep](https://github.com/golang/dep) which is a go dependency manager

## INSTALLATION
- Get the project `go get github.com/nathenapse/learning-blockchain`
- Go to directory `cd $GOPATH/src/github.com/nathenapse/learning-blockchain/`
- Set Env variable `cp .env.example .env`
- Change server port in `.env` which the default is ``PORT=8080``

## RUN
- Go to project `cd $GOPATH/src/github.com/nathenapse/learning-blockchain/`
- Run the project `go run cmd/main.go`

## API

### Get Blockchain
- `localhost:PORT`

### Create a new Block
- POST to `localhost:PORT`
- payload 
```JSON
{
    "Money": 150
}
```

## TODO
- [ ] Create Proof of work
- [ ] Improve data structure
- [ ] Add tests




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
- Get the project `go get -u github.com/nathenapse/learning-blockchain`
- Go to directory `cd $GOPATH/src/github.com/nathenapse/learning-blockchain/`
- Set Env variable `cp .env.example .env`
- If `.env` exists it will use it if not it will use the default values found in `.env.example`
- Install dependencies `dep ensure`
- It will save the database in your home directory in folder `.blockchain`
- Change server port in `.env` which the default is ``PORT=8080``

## BUILD
- Go to project `cd $GOPATH/src/github.com/nathenapse/learning-blockchain/`
- Run the project `make install`
- This will not use `.env`. It will use the default values if you don't set the variables in `~/.bash_rc`

## RUN
- Go to project `cd $GOPATH/src/github.com/nathenapse/learning-blockchain/`
- Run the project `make run`
- Or if you have built the project use `learning-blockchain`



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
- [x] Add Persistance using [BoltDB](https://github.com/coreos/bbolt)
- [x] create a Makefile
- [ ] Create Proof of work
- [ ] Improve data structure
- [ ] Add tests
# go-rest-api

## Swagger
![Screenshot](https://github.com/gulizay91/go-rest-api/blob/main/etc/ss-go-rest-api.png?raw=true)
### download swagger
```sh
go install github.com/swaggo/swag/cmd/swag@latest
```

### generate swagger doc
```sh
swag init -g ./cmd/main.go
```

## Folder Structure
```
├── Readme.md
├── cmd
│   ├── main.go                 // entry point
│   └── services
│       ├── services.go         // run all services
│       ├── config.service.go   // init config
│       └── router.service.go   // init router
├── config
│   └── config.go               // all config models
├── go.mod
├── go.sum
├── docs
│   └── docs.go
│   └── swagger.yaml            // swagger files


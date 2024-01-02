# go-rest-api
Sample Go RESTful Api


## Folder Structure
```
├── README.md
├── app
│   ├── cmd
│   │   ├── main.go                     // entry point
│   │   └── services
│   │       ├── services.go             // run all services
│   │       ├── config.service.go       // init config
│   │       └── router.service.go       // init router
│   ├── config
│   │   └── config.go                   // all config models
│   ├── go.mod
│   ├── go.sum
│   ├── docs
│   │   └── docs.go
│   │   └── swagger.yaml                // swagger files
│   ├── routers
│   │   └── router.go                   // all endpoints
│   ├── pkg
│   │   ├── handlers                    // all handlers
│   │   ├── models                      // all dtos
│   │   ├── repository                  // all repositories
│   │   │   └── entities                // all db entities
│       └── service                     // all services
├── k8s-manifests
│   └── base
│       ├── .env                        // define base environment variables
│       ├── deployment.yaml             // kubernetes deployment base
│       ├── kustomization.yaml          // kustomiza kubernetes service base
│       └── service.yaml                // kubernetes service base on cluster
│   └── dev
│       ├── .env                        // define base environment variables
│       ├── patch-deployment.yaml       // kubernetes patch deployment for specific environment
│       ├── kustomization.yaml          // kustomiza kubernetes service for specific environment
│       └── hpa.yaml                    // kubernetes horizontal pod autoscaler
│   └── prod
│       ├── .env                        // define base environment variables
│       ├── patch-deployment.yaml       // kubernetes patch deployment for specific environment
│       ├── kustomization.yaml          // kustomiza kubernetes service for specific environment
│       └── hpa.yaml                    // kubernetes horizontal pod autoscaler

```

## Swagger
![Screenshot](https://github.com/gulizay91/go-rest-api/blob/main/etc/ss-go-rest-api.png?raw=true)
### download swagger
```sh
go install github.com/swaggo/swag/cmd/swag@latest
```

## Generate Swagger Doc
```sh
swag init -g ./app/cmd/main.go
```

## Docker
### build docker image
```sh
# /go-rest-api>
docker build -t go-rest-api ./app
```

### push image to docker hub
```sh
docker login --username=guliz91 # you will prompted for your password
docker tag go-rest-api guliz91/go-rest-api:latest # tag docker image
docker push guliz91/go-rest-api:latest # push docker image to docker hub
```

## K8S - Deployment
### prepare minikube
```sh
minikube start
minikube dashboard
minikube addons enable metrics-server # for hpa.yaml
minikube tunnel
```
### push image to docker hub
```sh
cd k8s-manifests
cd base
# /go-rest-api/k8s-manifests/base>
kubectl apply -f namespace.yaml # create namespaces
cd ../
cd dev
# /go-rest-api/k8s-manifests/dev>
kubectl apply -k .
```

## Minikube Dashboard 
![Screenshot](https://github.com/gulizay91/go-rest-api/blob/main/etc/ss-minikube-go-rest-api.png?raw=true)



### remove
```sh
# /go-rest-api/k8s-manifests/dev>
kubectl delete -k .
```
```sh
minikube delete
```

 

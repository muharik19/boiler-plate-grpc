> Proposed grpc gateway based service .This project is still in development stage, any critic and suggestion all very based including but not limited to project name, function naming, folder structure etc. please refer to CONTRIBUTING.md.

## Prerequisite

- Install [go](https://golang.org/doc/install) in local machine for convenience development experience (auto complete, code sugestion, etc)
- Install golang plugin to your editor choice (ie. VSCode, Atom, Goland, Intellij IDE)
- [Docker](https://docs.docker.com/install/) and [docker-compose](https://docs.docker.com/compose/)
- Protoc and proto-gen-go plugin (described in the next session)

## Install make

### You first need Makefile

```sh
# install windows
C:\> choco install make
```

```sh
# install linux
$ sudo apt-get install make
```

```sh
# install MacOS with Homebrew
$ brew install make
```

### Verify Version Makefile

```sh
# verify that Makefile is installed correctly by checking the version
$ make --version
```

## Protoc and Proto-gen-go

### Install protoc and protoc-gen-go in Mac

- Install protoc-gen-go use homebrew `brew install protoc-gen-go` it will also install protoc
- The compiler plugin, protoc-gen-go, will be installed in `$GOPATH/bin` unless `$GOBIN` is set. It must be in your $PATH for the protocol compiler, protoc, to find it
- Install grpc-gateway  `go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway`
- Intall protoc-gen-swagger `go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger`

### Install protoc and protoc-gen-go in Ubuntu

- Install protoc-gen-go `go get -u github.com/golang/protobuf/protoc-gen-go`. It will also install protoc
- Install grpc-gateway  `go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway`
- Intall protoc-gen-swagger `go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger`

### Install protoc and protoc-gen-go in Windows

- Install protoc-gen-go `go get -u github.com/golang/protobuf/protoc-gen-go`. It will also install protoc
- Install grpc-gateway  `go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway`
- Intall protoc-gen-swagger `go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger`

## How to run

Despite, it is possible to run this project in local machine Please follow this steps:
- Changes file .env to your config `.cp .env.example to .env`.
- Go to directory proto.
- Generate protobuf file `protoc -I ./api/grpc/api/proto --go_out ./api/grpc/api/pb --go_opt paths=source_relative --go-grpc_out ./api/grpc/api/pb --go-grpc_opt paths=source_relative --grpc-gateway_out ./api/grpc/api/pb --grpc-gateway_opt paths=source_relative ./api/grpc/api/proto/v1/role/role.proto --go-grpc_opt=require_unimplemented_servers=false --swagger_out ./api/grpc/api/proto/swagger`.
- Generate protobuf file make `make proto`.
- Run apps gp to root project `go run main.go`.
- Run apps gp to root project make `make server`.
- Run apps gp to change directory deployments use docker `docker compose up -d`.

# Project-Structure

    boiler-plate-grpc/
    ├── README.md               # Project description
    ├── LICENSE                 # Project license
    ├── Makefile                # Build scripts and common commands
    ├── go.mod                  # Go module
    ├── go.sum                  # Dependency list
    ├── api/                    # API definitions (e.g., gRPC protocol)
    │   ├── graphql/            # GraphQL API definitions
    │   ├── grpc/               # gRPC API definitions
    │   │   ├── api/            # gRPC API implementation
    │   │   │   ├── client.go   # gRPC client implementation
    │   │   │   ├── info/       # Directory for gRPC info
    │   │   │   ├── pb/         # Directory for protocol buffers
    │   │   |   |   └── v1/     # Directory for version protocol buffers
    │   │   │   ├── proto/      # Directory for definitions proto
    │   │   |   |   └── v1/     # Directory for version definitions proto
    │   │   │   └── server.go   # gRPC server implementation
    │   ├── rest/               # REST API definitions
    │   │   ├── api/            # REST API implementation
    │   │   │   ├── routes/     # Routes for REST API
    │   │   │   └── server.go   # Server setup for REST API
    │   ├── worker/             # Worker API definitions
    │   │   ├── consumer/       # Consumer for messages
    │   │   └── scheduler/      # Scheduler for tasks
    ├── build/                  # Build-related files
    │   ├── package/           # Package-related files
    │   |   └── Dockerfile     # Dockerfile for building the application
    ├── cmd/                    # Contains main applications
    │   ├── graphql/            # GraphQL application
    │   │   └── graphql.go      # Entry point Cobra Command for GraphQL app
    │   ├── grpc/               # gRPC application
    │   │   └── grpc.go         # Entry point for gRPC app
    │   ├── rest/               # REST application
    │   │   └── rest.go         # Entry point Cobra Command for REST app
    │   ├── worker/             # Worker application
    │   │   ├── consumer/       # Consumer for messages
    │   │   │   └── consumer.go # Entry point Cobra Command for consumer
    │   │   └── scheduler/      # Scheduler for tasks
    │   │      └── scheduler.go # Entry point Cobra Command for scheduler
    │   ├── root.go             # Entry point for the root Cobra Command
    ├── configs/                # Configuration files
    │   ├── configs.go          # Configuration Global
    ├── deployments/            # Deployment-related files
    │   ├── docker-compose.yml  # Docker Compose configuration file
    ├── docs/                   # Project documentation
    │   ├── api/                # API documentation
    │   ├── architecture/       # Architecture documentation
    │   ├── images/             # Images related to documentation
    ├── internal/               # Code that can only be used by this project
    │   ├── application/        # Application logic
    │   │   ├── adapters/       # Adapters for external services
    │   │   │   └── info/       # Info related adapters
    │   │   │       └── info.go # Info adapter
    │   │   ├── controllers/    # Controllers for handling requests
    │   │   │   └── info/       # Info related controllers
    │   │   │       └── controllers.go # Info controller
    │   │   │       └── info.go # Info controller
    │   │   ├── presenters/     # Presenters for formatting responses
    │   │   │   ├── request/    # Request related presenters
    │   │   │   │   └── FindInfoDetailByIDRequest.go # Find info detail request
    │   │   │   └── response/   # Response related presenters
    │   │   │       └── InfoDetailResponse.go # Info detail response
    │   │   ├── repositories/   # Repositories for data access
    │   │   │   └── info/       # Info related repositories
    │   │   │       ├── info.go # Info repository
    │   │   │       ├── info_mock.go # Mock for info repository
    │   │   │       └── interface.go # Info repository interface
    │   │   ├── usecases/       # Use cases for application logic
    │   │   │   └── info/       # Info related use cases
    │   │   │       ├── add_info.go  # Use case for adding info
    │   │   │       ├── find_info.go # Use case for finding info
    │   │   │       ├── find_info_test.go # Tests for finding info
    │   │   │       └── interface.go # Info use case interface
    │   ├── domain/             # Domain models
    │   │   ├── entities/       # Domain entities
    │   │   |   ├── response.go # Response standardization entities
    │   │   │   └── info/       # Info related entities group
    │   │   │       └── info.go # Info entity definition
    │   ├── pkg/                # Internal packages
    │   │   ├── database/       # Database related code
    │   │   │   ├── nosql/      # NoSQL database connection code
    │   │   │   └── sql/        # SQL database connection code
    │   │   ├── logger/         # Logging utilities
    │   │   └── utils/          # Utility functions
    ├── pkg/                    # Code that can be used by other projects
    │   ├── pkg1/               # Package global for used all application
    ├── scripts/                # Scripts for development and deployment
    │   ├── setup.sh            # Setup script
    ├── test/                   # Contains test files
    │   ├── testdata/           # Test data
    │   ├── integration_test.go # Integration test
    └── migrations/             # Database migration scripts
        └── 20210101_init.sql   # Initial migration script

## Documentation API Postman

[API](https://documenter.getpostman.com/view/42999233/2sAYk7SPpJ)
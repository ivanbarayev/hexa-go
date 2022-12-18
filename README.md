# go-microservice-boilerplate

    .
    ├── bin                                     # Binaries output directory
    ├── cmd                                     # Root of main
    │   ├── grpc                                # GRPC main
    │   └── http                                # HTTP main
    ├── config                                  # Configuration parser, structure and .yaml file
    │   ├── config.go                           # Config file reader and parser 
    │   ├── config.yaml                         # Config .yaml file
    │   └── models.go                           # Config model structure
    ├── docs                                    # Swagger documentation
    │   ├── docs.go                             # Document parser
    │   ├── swagger.json                        # Auto generated swagger JSON file 
    │   └── swagger.yaml                        # Auto generated swagger YAML file 
    ├── internal                                # Hexagonal - DDD Layer
    │   └── auth                                # Auth Domain
    │       ├── application                     # Application Layer
    │       │   └── service                     # Usecase Layer
    │       │   |   └── service.go              # Usecase file of current domain
    │       │   └── jobs                        # Special Jobs Ticker or Crons
    │       │       └── jobs.go                 # Special jobs runner by go Ticker with GoRoutines (like as Cron Jobs) (time specific) 
    │       ├── domain                          # Domain Layer
    │       │   └── entities                    # Models/Structures
    │       │       └── entities.go             # Models/Structures file of current domain
    │       │   └── errors                      # Custom Errors
    │       │       └── errors.go               # Custom Errors file of current domain
    │       │   └── ports                       # Ports / Interfaces
    │       │       └── http_handler.go         # HTTP Handlers Interfaces file of current domain
    │       │       └── postgresql_repo.go      # Postgresql Interfaces file of current domain
    │       │       └── service.go              # Usecase Interface file of current domain
    │       │   └── utils                       # Specific Utilities
    │       │       └── utils.go                # Specific Utilities file of current domain
    │       └── handler                         # Handler Layer 
    │       │   └── grpc                        # GRPC Handlers
    │       │       └── grpc                    # GRPC Handler file of current domain 
    │       │   └── http                        # HTTP Handlers
    │       │       └── handlers.go             # HTTP Handler file of current domain  
    │       │       └── routes.go               # HTTP routes file of current domain  
    │       ├── infrastructure                  # Infrastructure Layer
    │       │   └── adapters                    # 3rd party integration adapters 
    │       │       └── postgresql_repo         # PostgresqlRepository file of current domain
    │       │   └── repository                  # Repositories 
    │       │       └── postgresql_repo         # PostgresqlRepository file of current domain
    ├── pkg                                     # Custom packages for general common usage
    ├── proto                                   # GRPC proto directory
    │   └── auth                                # Auth Domain
    │       └── auth.pb.go                      # Auto generated domain specific proto functions
    │       └── auth.proto                      # Domain specific generated domain specific proto functions
    │       └── auth_grpc.pb.go                 # Auto generated domain specific proto functions
    └── ...

```shell
# Makefile commands 
make run #for run current application 

make swag # for generate swagger documentations

make genproto # for generate grpc files *.pb.go and *_grpc.pb.go

make build # compile binary for all -> Mac, Linux and Windows

make buildl # compile only for Linux OS

make buildm # compile only for Mac OS

make buildw # compile only for Windows OS

```
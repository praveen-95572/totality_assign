# User Related Grpc Calls
## Implmented the ddd model

Project structure is divided into 2 directory. One is cmd and another is internal.

1. cmd (directory)
    - Inside the server directory there is the main file. 
    - it is the starting point of the service.
    - All the instances are created here.

2. internal (directory)
    - Inside the internal directory there are multiple directory
    1.  domain
        -   Here the structure is declare and function related to them are declared.
    2. grpc
        - In this directory only the proto file in written
        - generated file is saved in infra/grpcserver/.gen directory
    3  infra
        - Here grpcserver directory is created.
        - In this grpcserver dir function related to grpc call is defined.
    4 service
        - Here function that declare in domain is defined.
        - Function related to user domain is defined inside the user directory.

## Features
    created the docker file so that docker image can be crreated easily.


## Installation

Install the dependencies and start the server.
```sh
go mod tidy
go run cmd/server/main.go
```


# Clean Architecture with Microservice sample template
 
## Introduction

This repo contains 2 services: **User** and **Role** which their structures adapt **Clean Architecture**.

![Clean Architecture Overview](abcxyz.jpg)
![Comparision gRPC vs REST](abcxyz.jpg)

## Folder hierarchy
Main workflow from folder **services**

```
services
├── role
│   ├── adapters
│   │   ├── api
│   │   │   ├── handler fuctions
│   │   ├── gRPC
│   │   │   ├── handler fuctions
│   ├── cmd
│   ├── constants
│   │   ├── envVar
│   │   ├── notis
│   │   ├── redisKey
│   ├── entities
│   ├── infrastructures
│   │   ├── db connection
│   │   ├── repositories
│   │   ├── routes
│   ├── interfaces
│   ├── mocks
│   ├── usecases
│   │   ├── businessLogics
│   │   ├── tests
├── user
│   ├── adapters
│   │   ├── api
│   │   │   ├── handler fuctions
│   │   ├── gRPC
│   │   │   ├── handler fuctions
│   ├── cmd
│   ├── constants
│   │   ├── envVar
│   │   ├── notis
│   │   ├── redisKey
│   ├── dtos
│   │   ├── request
│   ├── entities
│   ├── external_services
│   │   ├── role
│   ├── infrastructures
│   │   ├── db connection
│   │   ├── repositories
│   │   ├── routes
│   ├── interfaces
│   ├── mocks
│   ├── usecases
│   │   ├── businessLogics
│   │   ├── tests
```


- Services communicating with each others through Http/2 - gRPC which is significantly faster than the traditional protocol Http/1.1 - REST
- Protocol buffers supporting for communications between services generated at **protocols** folder

```
protocols
├── roleService
│   ├── pb
│   │   ├── role_grpc.pb.go
│   │   ├── role.pb.go
│   ├── role.proto
├── userService
│   ├── user.proto
```

Some other common services which support through variety of services

```
common_dtos
├── request
├── response
constants
external_services (3rd services)
helper
├── api_response
htmlTemplate
middlewares
```

## General about the structure
As Microservice based, each service in the template handles independent requirements and features with these main stacks:
- GIN (for HTTP request/response)
- gRPC (internal communication between services)
- Core database: PostgreSQL (with package database/sql)
- Cache: Redis (type of NoSQL)
- Container: Docker
- Message broker: RabbitMQ 
- Authentication: JWT
- Testing: unit tests with mock

### User Service
- REST API: user management with basic CRUD task. This service divides to 3 kind of connection with different level of authorizations. It also consists of a higher service with supporting locking account, reactivate account with authentication through mail confirmation and locking with specific fail attempts.

### Role Service
- REST API: role management with basic CRUD task. All of this accessions require almost highest level of authorization - required by "Admin" role.
- gRPC: provides a get-all-role service. In this template, user service will call this stuff to verify if an account's role provided by upper "Staff" role is valid or exist as a protection for faking roles.

### Other commons
Those are shared services with support multiple services with public data.

#### Common dtos
- This folder consists of 2 divided folder: request and response.
- With response, this folder contains common models wich usually transfered: 
+ apiResponse: generate response to client.
+ dataStorage: the general struct which used to save data to Redis cache.
- With request, ...

#### Constants
- envVar: consists of shared environment variables for services to fetch and use.
- grpcGateway: contains all available gRPC ports from all services which helps to get access.
- mailConst: includes html directory to generate mails with specific purposes.
- notis: contains all of common messages which support for logging information, generate notifications to users from errors to warnings and messages. There are also subjects for specific mails' purposes.
- queues: all of queues which are used to produce when running application to wait for consumers.

#### External Services/ 3rd-party services
Currently consists of message broker (RabbitMQ) and cache (Redis)

#### Helper
- Common functions such as validating empty input or get time and so on.
- It also has a special helper - generating api response to use across services.

#### Html Template
Contains html files which supporting for generating mails.

#### Middlewares
Mostly for authorization.


## Guidelines

### Download/clone template
Open your cmd/terminal/console and type this below command to clone this template:

```shell
git clone https://github.com/Phuchtq/go-microservice-clean-architecture
cd go-microservice-clean-architecture
```

### Start project
- Take the database script from one of services (as in this template this 2 service share same database) and execute the query in **pgAdmin 4** desktop - PostgreSQL. If you haven't downloaded, click [here](https://www.postgresql.org/).
- Move to your directory and type this command to execute:

```shell
go run main.go
```

- Command next 2 lines to start 2 available service:
```shell
go run main.go role-service
go run main.go user-service
```

- If you want to start it independently:
```shell
go run main.go <command-of-service>
```

- If you start by docker:
```shell
docker-compose up --build
```

Note: you should change all of environment variables to your own ones



























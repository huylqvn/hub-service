# hub-service

## Preface

## Install
Perform the following steps:
1. Download and install [Golang](https://golang.org/).
2. Clone this repository.

## Starting Server
There are 2 methods for starting server.

### Without Web Server
1. Starting this web application by the following command.
    ```bash
    go run main.go
    ```
1. When startup is complete, the console shows the following message:
    ```
    http server started on [::]:8080
    ```
1. Access [http://localhost:8080](http://localhost:8080) in your browser.
1. Login with the following username and password.
    - username : ``test``
    - password : ``test``

### With Web Server
#### Starting Application Server
1. Starting this web application by the following command.
    ```bash
    go run main.go
    ```
1. When startup is complete, the console shows the following message:
    ```
    http server started on [::]:8080
    ```
1. Access [http://localhost:8080/api/health](http://localhost:8080/api/health) in your browser and confirm that this application has started.
    ```
    healthy
    ```

## Using Swagger
In this sample, Swagger is enabled only when executed this application on the development environment.
Swagger isn't enabled on the another environments in default.

### Accessing to Swagger
1. Start this application according to the 'Starting Application Server' section.
2. Access [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) in your browser.

### Updating the existing Swagger document
1. Update some comments of some controllers.
2. Download Swag library. (Only first time)
    ```bash
    go install github.com/swaggo/swag/cmd/swag@latest
    ```
3. Update ``docs/docs.go``.
    ```bash
    swag init
    ```

## Build executable file
Build this source code by the following command.
```bash
go build main.go
```

## Project Map
The following figure is the map of this sample project.

```
- hub-service
  + config                  … Define configurations of this system.
  + logger                  … Provide loggers.
  + middleware              … Define custom middleware.
  + migration               … Provide database migration service for development.
  + router                  … Define routing.
  + controller              … Define controllers.
  + model                   … Define models.
  + repository              … Provide a service of database access.
  + service                 … Provide a service of book management.
  + session                 … Provide session management.
  + test                    … for unit test
  - main.go                 … Entry Point.
```

## Services
This sample provides User management for Hub and Team.
Regarding the detail of the API specification, please refer to the 'Using Swagger' section.

`Although the test requirements are more, the hub and team are similar to the user so to not waste time I didn't do it, hope you understand :d`

## Tests
Create the unit tests only for the packages such as controller, service, model/dto and util. The test cases is included the regular cases and irregular cases. Please refer to the source code in each packages for more detail.

The command for testing is the following:
```bash
go test ./... -v
```

## Libraries
This sample uses the following libraries.

| Library Name               | Version |
| :------------------------- | :-----: |
| echo                       | 4.11.4  |
| gorm                       | 1.25.9  |
| go-playground/validator.v9 | 9.31.0  |
| zap                        | 1.26.0  |


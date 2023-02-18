go rest api template using GIN, GORM(mysql) & ozzo-validation.
\*Still work in progress...

## Packages

#### repository

`repository` package is a place where to store repositories.

#### service

`service` package is a place where to store business logic of the app. <br />
each service returns a custom error (sdk/apires/errwrapper)

#### handler

`handler` package is responsible for connection of business logic layer and transport layer.

#### entity

`entity` package is a place where to store app entities.

#### model

`model` package is a place where to store request/response structs.<br />
i don't use gin/binding, instead i use ozzo-validation

#### database

`database` package is responsible for databases purposes like connection, migration, etc.

# go-oauth2-server
Scaffolding OAuth 2.0 Server

## Setup .env Configuratation

Please set some configurations such as database connection and others according to your environment.
You can copy from `.env.example` with the following command:

```
cp .env.example .env
```

## Run Database Migration
```
go run main.go db:migrate
```

## Create a New User
```
go run main.go create:user
```

## Create a New Client (OAuth Application)
```
go run main.go create:client
```

## Run OAuth 2.0 Server
```
go run main.go
```
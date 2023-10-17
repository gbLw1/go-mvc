# ToDo App

This is a ToDo [MVC](https://en.wikipedia.org/wiki/Model%E2%80%93view%E2%80%93controller) app with [Go](https://go.dev/)

## Dependencies

- [Air](https://github.com/cosmtrek/air) - Hot reloading
- [GORM](https://gorm.io/) - ORM
- [Fiber](https://gofiber.io/) - Web framework
- [GoDotEnv](https://github.com/joho/godotenv) - Env variables
- [Fiber template](https://github.com/gofiber/template) - HTML Template engine

## Database

- [PostgreSQL](https://www.postgresql.org/)

## Setting up the project

Follow the steps below to set up the project.

### Environment variables

create a new `.env` file in the root directory and add the following:

```bash
PORT=8080
DB_URL="host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable"
```

`PORT` is the port where the API will run (default: 8080). You can change it if you want.

`DB_URL` is the connection string for the database. Change it according to your database credentials.

## Running with Air `(optional)`

Air is a tool for running Go applications in the background, refreshing whenever it is modified.

ps: You can change the `.air.toml` file with your configs if you want.

With `air` installed, run the following command:

```sh
air
```

## Running the project

Run the following command

```sh
go run ./cmd/main.go
```

<!--
## You can add REACT to your project (bruh...)

`npm init -y`

`npm i react react-dom`

`npm i -D @types/react @types/react-dom`

`npm i -D esbuild`

`npm i -D react-router-dom`
-->

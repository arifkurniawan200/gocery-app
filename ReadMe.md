
# Migration Database

This is a service that can be used to perform database migrations.
In this case, two distinct databases are used.


## Authors

- [@arifkurniawan200](https://github.com/arifkurniawan200)


## Environment Variables

**Install Goose:**
```bash
 go install github.com/pressly/goose/v3/cmd/goose@latest
```

**Run Postgres:**
```bash
 sudo systemctl start postgresql
```

**Check Postgres:**
```bash
 sudo systemctl status postgresql
```

**Change App.yaml Configuration:**
```bash
 To run this project, you will need to add the following environment variables to your app.yaml file in folder config

`db` : database


Change the configuration according to your database 
```



## Run Locally

change app.yaml.example in folder config to app.yaml

install dependencies

```bash
  go mod tidy
```

running database migration

```bash
  go run main.go db:migrate up
```


reset database (delete database and existing data)

```bash
  go run main.go db:migrate reset
```

running api server

```bash
  go run main.go api
```




## API Reference

#### login as admin

```http
  GET /login
  
  param :
  username: admin
  password: admin
```

#### register user

```http
  GET /register
  
  param :
  username: {{username}}
  password: {{password}}
```

### login user
```http
  GET /login
  
  param :
  username: {{username}}
  password: {{password}}
```

### Authorization

To Access endpoint always using bearer Authorization

```
Bearer {{token from login}}
```


#### example operation

postman file already attached in repo


## Tech Stack

**Database:** PostgresSQL

**Framework:** Echo golang

**Migration:** Goose
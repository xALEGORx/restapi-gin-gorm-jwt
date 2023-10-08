# ❤️ Go RestApi (Gin + Gorm + JWT)
A very simple example of a RESTful API built on Gin+ Gorm+ JWT with an error handler

This is the basic structure for your project, it needs to be improved for each project.

## How to launch?
```bash
# Download this git repo
git clone https://github.com/xALEGORx/restapi-gin-gorm-jwt

# Copy .env.example to .env
cp .env.example .env

# Start the server
go run main.go
```

> Don't forget to specify the data from the MySQL database and generate a new JWT

## Structure
```
│   main.go           // main file
├───core
│   ├───api.go        // core rest api
│   ├───database.go   // migrations for db
│   └───routes.go     // routes list
├───handlers
│   ├───auth.go       // handler for login & register
│   ├───ping.go       // handler for ping/pong
│   └───user.go       // handler for getting information about an authorized user
├───middlewares
│   ├───authorized.go // check jwt token
│   ├───cors.go       // enable cors for all domains
│   ├───error.go      // error handler
│   └───handler.go    // core handler
├───models
│   └───user.go       // ORM table: user
└───types
    ├───error.go      // Types of registered errors
    └───response.go   // Structure for the response
```

## Examples of API requests
```
All parameters of POST requests are sent as JSON
```
`GET` /api/v1/ping
`ANSWER`: 
```
{
    "success": true,
    "data": {
        "message": "pong"
    }
}
```

`POST` /api/v1/login (email, password)
`ANSWER`:
```
{
    "success": true,
    "data": {
        "token": "eyJhbGciOiJIUzI...Y3iVcL0",
        "user": {
            "email": "admin@admin.com",
            "id": 1,
            "name": "ADMIN"
        }
    }
}
```

## Error handler
`When calling an error from the /types/error list.go is triggered /middlewares/error.go`
> For example, if the login or password is incorrect during authorization
```go
if !user.CheckPassword(body.Password) {
	c.Error(types.WRONG_PASSWORD)
	return
}
```
In the `/types/error.go` file is stored variable
```go
var (
	WRONG_PASSWORD   = &ApiError{203, "wrong login or password"}
	...
)
```
The answer then will be:
```
{
    "success": false,
    "error": "user does not exist"
}
```

## List of libraries
* [GIN](https://github.com/gin-gonic/gin)
* [GORM](https://gorm.io/docs/)

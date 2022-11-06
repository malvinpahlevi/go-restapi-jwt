
# Go RESTAPI JWT

Before, especially thanks for @seefnasrul for the tutorial on [Medium](https://seefnasrul.medium.com/create-your-first-go-rest-api-with-jwt-authentication-in-gin-framework-dbe5bda72817) 

## Pre-requisite


- Using [Air](https://github.com/cosmtrek/air) for Live Reload Go 




## How to usage

- Protected Endpoint
```bash
  GET http://localhost:8080/api/v1/admin/user
  Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2Njc3NDE4OTgsInVzZXJfaWQiOjI0fQ.-U_un1uzTt

  #response
  {
  "data": {
    "ID": 24,
    "CreatedAt": "2022-11-06T18:50:51.992+07:00",
    "UpdatedAt": "2022-11-06T18:50:51.992+07:00",
    "DeletedAt": null,
    "username": "admin",
    "password": ""
  },
  "message": "success"
}
```

- Register User

```bash
  POST http://localhost:8080/api/v1/register
  {
    "username": "brody",
    "password": "123456"
  }
  
  #response  
  {
    "message": "registration success !"
  }

```

- Login User

```bash
  POST http://localhost:8080/api/v1/login
  {
    "username": "brody",
    "password": "123456"
  }

  #response  
  {
    "message" : "success"
    "token" : "kjbskjh12938124kj123..."
  }
```



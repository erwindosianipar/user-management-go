# usermanagement

## Register

| Method | URL |
| -- | -- |
| POST | /user/register |

Request
```json
{
	"name" : "Erwindo Sianipar",
	"email" : "erwindosianipar@gmail.com",
	"password" : "12345678"
}
```
Response
```json
{
    "message": "success: register user",
    "response": {
        "id": 5,
        "created_at": "2020-02-23T18:31:13.598121+07:00",
        "updated_at": "2020-02-23T18:31:13.598121+07:00",
        "deleted_at": null,
        "email": "erwindosianipar@gmail.com",
        "password": "$2a$10$y97WYwJMjuX.jx7UWQb20ueZSm655rfEdf1TFT1LtSC7J.NGmXHMS",
        "name": "Erwindo Sianipar",
        "age": 18,
        "address": ""
    },
    "status": true
}
```

## Login

| Method | URL |
| -- | -- |
| POST | /user/login |

Request
```json
{
	"email" : "erwindosianipar@gmail.com",
	"password" : "12345678"
}
```

Response
```json
{
    "message": "success: user already logged in",
    "response": {
        "id": 5,
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImVyd2luZG9zaWFuaXBhckBnbWFpbC5jb20iLCJpZCI6NX0.1RKNYDQhhaBvUy_C3eKpp5j1NayX2zyJBWWzocKv0GA"
    },
    "status": true
}
```

## Get All Data User

| Method | URL |
| -- | -- |
| GET | /user/all |

Response
```json
{
    "message": "success: get all data user",
    "response": [
        {
            "id": 1,
            "created_at": "2020-02-23T13:24:43.229021+07:00",
            "updated_at": "2020-02-23T13:24:43.229021+07:00",
            "deleted_at": null,
            "email": "john@gmail.com",
            "password": "$2a$10$d7IcARpTZELjjWNqXinUiuR699rU6/K0azJDuClsdVr6bI4sH.bYu",
            "name": "John Doe",
            "age": 18,
            "address": ""
        },
        {
            "id": 2,
            "created_at": "2020-02-23T13:24:50.976155+07:00",
            "updated_at": "2020-02-23T18:41:20.650551+07:00",
            "deleted_at": null,
            "email": "jane@gmail.com",
            "password": "$2a$10$iY.UwHY2Qu5yy6v14P4me.bW6LGs4H3iRmgRbw0xVlWbtxBpZi7yy",
            "name": "Jane Doe",
            "age": 20,
            "address": ""
        },
        {
            "id": 5,
            "created_at": "2020-02-23T18:31:13.598121+07:00",
            "updated_at": "2020-02-23T18:39:48.809345+07:00",
            "deleted_at": null,
            "email": "erwindosianipar@gmail.com",
            "password": "$2a$10$y97WYwJMjuX.jx7UWQb20ueZSm655rfEdf1TFT1LtSC7J.NGmXHMS",
            "name": "Erwindo Sianipar",
            "age": 20,
            "address": ""
        }
    ],
    "status": true
}
```

## Get Data User By ID

| Method | URL |
| -- | -- |
| GET | /user/{id} |

Response
```json
{
    "message": "success: get data user by id",
    "response": {
        "id": 5,
        "created_at": "2020-02-23T18:31:13.598121+07:00",
        "updated_at": "2020-02-23T18:39:48.809345+07:00",
        "deleted_at": null,
        "email": "erwindosianipar@gmail.com",
        "password": "$2a$10$y97WYwJMjuX.jx7UWQb20ueZSm655rfEdf1TFT1LtSC7J.NGmXHMS",
        "name": "Erwindo Sianipar",
        "age": 20,
        "address": ""
    },
    "status": true
}
```

## Update User

| Method | URL |
| -- | -- |
| PUT | /user/{id} |

Request
```json
{
	"name" : "Erwindo Sianipar",
	"email" : "erwindosianipar@gmail.com",
	"age" : 20
}
```

Response
```json
{
    "message": "success: update data user",
    "response": {
        "id": 5,
        "created_at": "2020-02-23T18:31:13.598121+07:00",
        "updated_at": "2020-02-23T18:39:48.809345+07:00",
        "deleted_at": null,
        "email": "erwindosianipar@gmail.com",
        "password": "$2a$10$y97WYwJMjuX.jx7UWQb20ueZSm655rfEdf1TFT1LtSC7J.NGmXHMS",
        "name": "Erwindo Sianipar",
        "age": 20,
        "address": ""
    },
    "status": true
}
```

## Delete User

| Method | URL |
| -- | -- |
| DELETE | /user/{id} |

Response
```json
{
    "message": "success: delete data user",
    "response": true,
    "status": true
}
```
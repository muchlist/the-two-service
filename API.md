# The-Two-(Tower)Service API Specs

Documentation API for The-Two-(Tower)Service Auth app and Fetch app

## Authentication

- HTTP Authorization, Bearer Schema

## Auth App
`localhost:8080`

### Register

`POST /register`

*Request*
> Body Params : JSON

```json
{
  "phone": "081288888888",
  "name": "Muchlis",
  "role": "admin"
}
```

*Response*

> HTTP 201 : Example Response

```json
{
  "message": "success register phone number",
  "data": {
    "phone": "08128888888",
    "password": "9HxE"
  },
  "error": null
}
```

### Login

`POST /login`

*Request*
> Body Params : JSON

```json
{
  "phone": "081320243880",
  "password": "6fUJ"
}
```

*Response*

> HTTP 200 : Example Response

```json
{
  "data": {
    "access_token": "XXXXeXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJmcmVzaCI6dHJ1ZSwiaWF0IjoxNjU0MjcwMjExLCJqdGkiOiIxM2I2MmRmZi0zYTkwLTRkMDQtOWZjZS04OGViYjk3YWNjZjEiLCJ0eXBlIjoiYWNjZXNzIiwic3ViIjoiKzgxMjMxODQ3NDEiLCJuYmYiOjE2NTQyNzAyMTEsImV4cCI6MTY1NDI3MTExMSwibmFtZSI6Im1vcmFsYSIsInBob25lIjoiKzgxMjMxODQ3NDEiLCJyb2xlIjoiYWRtaW4iLCJ0aW1lc3RhbXAiOiIyMDIyLTA2LTAzIDE1OjI5OjUxIn0.sY4Ugfxc_sgJpskv_TALq4rhJOPeQPdvEMCVwO-XXXX",
    "name": "Muchlis",
    "phone": "081320243880",
    "role": "admin",
    "timestamp": "2022-06-03 15:29:51"
  },
  "error": null
}
```

> HTTP 400 : Example Response

```json
{
  "data": null,
  "error": "phone or password not valid"
}
```

### Profile

`GET /profile`

*Request*
> Header Params

```
Authorization : Bearer {JWT_Token}
```

*Response*

> HTTP 200 : Example Response

```json
{
  "data": {
    "exp": 1654277168,
    "iat": 1654277168,
    "nbf": 1654277168,
    "fresh": true,
    "jti": "a6bb8482-3a8a-442e-8834-9b5b0eb734aa",
    "name": "Muchlis",
    "phone": "081320243880",
    "role": "admin",
    "sub": "+81231741",
    "type": "access",
    "timestamp": "2022-06-03 15:29:51"
  },
  "error": null
}
```

> HTTP 401 : Example Response

```json
{
  "data": null,
  "error": "Invalid Token!",
}
```

## Fetch App 
`localhost:8081`

### Profile
`GET /profile`

*Request*
> Header Params

```
Authorization : Bearer {JWT_Token}
```

*Response*

> HTTP 200 : Example Response

```json
{
  "data": {
    "name": "Muchlis",
    "phone": "081320243880",
    "role": "admin",
    "timestamp": "2022-06-03 15:29:51",
    "type": "Access",
    "exp": 9999999999,
    "fresh": true
  },
  "error": null
}
```

> HTTP 401 : Example Response

```json
{
  "data": null,
  "error": "Invalid Token!",
}
```

### Fetch Fish Data
`POST /fish`

*Request*
> Header Params

```
Authorization : Bearer {JWT_Token}
```

*Response*

> HTTP 200 : Example Response

```json
{
  "data": [
    {
      "uuid": "10799adf-5284-456a-b6be-fb43e0113251",
      "commodity": "Haruan",
      "province": "KALIMANTAN SELATAN",
      "city": "BANJARMASIN",
      "size": 120,
      "price": 200000,
      "price_usd": 12.12312,
      "time_parsing": "2022/05/16 23:35:55",
      "timestamp": null
    }
  ],
  "error": null
}
```

> HTTP 401 : Example Response

```json
{
  "data": null,
  "error": "Invalid Token!",
}
```

> HTTP 403 : Example Response
```json
{
    "data": null,
    "error": "unauthorized, role [farmer] needed"
}
```

### Fetch Fish Data Aggregated
`GET /fish-aggregate`

*Request*
> Header Params

```
Authorization : Bearer {JWT_Token}
```

*Response*

> HTTP 200 : Example Response

```json
{
  "data": [
    {
        "year": "2022",
        "week": "20",
        "province": "BANTEN",
        "data_count": 127,
        "size": {
            "maximal": 1000,
            "minimal": 0,
            "median": 60,
            "average": 66.91338582677166
        },
        "price": {
            "maximal": 123566788,
            "minimal": 0,
            "median": 11111,
            "average": 2826829.5118110236
        },
        "price_usd": {
            "maximal": 8561.11051820282,
            "minimal": 0,
            "median": 0.769806357415,
            "average": 195.8519781766238
        }
    }
  ],
  "error": null
}
```

> HTTP 401 : Example Response

```json
{
  "data": null,
  "error": "Invalid Token!",
}
```

> HTTP 403 : Example Response
```json
{
    "data": null,
    "error": "unauthorized, role [admin] needed"
}
```
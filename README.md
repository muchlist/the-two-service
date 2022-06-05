# the-two-service

Two restfull-api service build with go and python

## Demo API (online)
"""Not ready yet"""


## Installation

Clone this project

```bash
git clone https://github.com/muchlist/the-two-service.git
```

Go to the project directory

```bash
cd the-two-service
```

Start the server with docker compose

```bash
docker-compose up -d
```


## Documentation

[Markdown API DOC](https://github.com/muchlist/the-two-service/blob/main/API.md)

You can also use swagger spec and paste to swagger editor ([Swagger Editor](https://editor.swagger.io/))

Swagger file : 
- [Auth swagger](https://github.com/muchlist/the-two-service/blob/main/auth/static/swagger.json)
- [Fetch swagger](https://github.com/muchlist/the-two-service/blob/main/fetch/swaggerui/swagger.json)


## C4 Document

### Context
![context](https://github.com/muchlist/the-two-service/blob/main/static/context.png)

### Container
![container](https://github.com/muchlist/the-two-service/blob/main/static/container.png)

### Component
![component](https://github.com/muchlist/the-two-service/blob/main/static/component.png)


## Authors

- [Muchlis - @muchlist](https://github.com/muchlist)


## Goals

- [x]  Servers bisa dinyalakan di port berbeda
- [x]  Semua endpoint berfungsi dengan semestinya (3 endpoint auth, 3 endpoint fetch)
- [x]  Dokumentasi endpoint dengan format OpenAPI (API.md)
- [x]  Dokumentasi system diagram-nya dalam format C4 Model (Context, container, component)
- [x]  Pergunakan satu repo git untuk semua apps (mono repo)
- [x]  Dockerfile untuk masing-masing app
- [x]  Petunjuk penggunaan dan instalasi di README.md yang memudahkan

## Additional Goals

- [ ]  Deployed ke Host/Penyedia Layanan (semacam surge, heroku, vercel, firebase, glitch,
host anda pribadi)
- [x]  Docker Compose
- [ ]  Unit Testing

## Warning
- All credential file exposed to make it easier in terms of testing. In Prod we need to gitignore or just change name of .env file.
- To keep it simple, many things are not implemented in this system. let say : better log management with request id, Metric, Profiling.
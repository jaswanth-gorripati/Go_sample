## TODO apis

### Run the api

```bash
go run main.go
```

### Generate the swagger documents

```bash
go install github.com/swaggo/swag/cmd/swag@latest

swag init -g main.go -o docs
```

### Swagger ui url

http://localhost:8081/swagger/index.html

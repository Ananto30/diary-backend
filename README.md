# Golpo
**WORK IN PROGRESS**

Social network backend made with fiber (WIP)


## Run
```bash
## start postgres
docker-compose up -d

## install dependencies
dep ensure

## migrate database
go run main.go migrate

## start server
go run main.go

```

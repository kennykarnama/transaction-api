# transaction-api

Service to handle order CRUD

## How To Run

### env files

please look up into .env file and also config.go to include your own configuration

### docker

windows & mac: 18.0.3+

```shell
docker-compose up
```

Linux: 20.04+

```shell
docker-compose up --add-host=host.docker.internal:host-gateway
```

### migration

for migration `UP` it will be handled automically, for `down` operation, you need to install
go-migrate (https://github.com/golang-migrate/migrate)

example command

```shell
migrate.exe -source file://database/migrations/mysql -database "mysql://kenny:kenny@tcp(localhost:3306)/user_api" down
```

## TO DO

- [ ] unit test 
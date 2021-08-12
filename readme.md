# transaction-api

# how to run

## db migrations

```shell
migrate.exe -source file://database/migrations/mysql -database "mysql://kenny:kenny@tcp(localhost:3306)/transaction_api" up
migrate.exe -source file://database/migrations/mysql -database "mysql://kenny:kenny@tcp(localhost:3306)/transaction_api" down
```
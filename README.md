### INFORMATION
mini project 3 is located at branch `deploy/mini-project-3`

### How to use

please generate mock file using [mockery](https://vektra.github.io/mockery) if you do not have one installed on you own machine
then run this command `go generate -v ./...`

to create a table you can use gorm feature for auto migration
see [migration](https://gorm.io/docs/migration.html)

please generate mock file using [mockery](https://vektra.github.io/mockery) if you do not have one installed on you own machine
then run this command `go generate -v ./...`

```go
// add this code to main.go after assignment to 'store' variable

store.AutoMigrate(&entity.Role{}, &entity.RegisterApproval{}, &entity.Account{}, &entity.Customer{})
```
or you can create schema based on the entities
or used the schema used from mini project 1.

### Environment variables
create `.env` file at the root of the project
the listing.1 show all environment values used at this repository

listing.1
```env
JWT_SECRET_KEY=<you_jwt_secret>
MYSQL_USERNAME=<you_sql_username>
MYSQL_PASSWORD=<you_sql_password>
MYSQL_HOST=<localhost>
MYSQL_PORT=<3306>
MYSQL_DB=<you_db_name>
```
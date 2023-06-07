### How to use

to create a table you can use gorm feature for auto migration
see [migration](https://gorm.io/docs/migration.html)

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
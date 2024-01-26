go mod dependancies for testing 

github.com/DATA-DOG/go-sqlmock
github.com/testcontainers/testcontainers-go

driver to allow access to db

github.com/jackc/pgx/v5

in order get the above driver to work add the following to main and test imports
_ "github.com/jackc/pgx/v5/stdlib"

normal example:
select name from projects where id = $1
expect query in mock expects regular expression, where $ is a recognised symbol, ergo i need to escape with the \ prefx as shown below:
select name from projects where id = \$1


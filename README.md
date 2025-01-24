# Golang REST API skeleton
This is a small template of a service written in Go, demonstrating:
1. My modest vision of a convenient architecture of applications in Go
2. Interaction with the database and the use of migrations
3. Generation of API from the Swagger specification (this eliminates the discrepancy between documentation and code)
4. Using mocks when writing unit tests
5. Injection of dependencies using Wire

## How to run
1. Copy the `config.yml.dist` file to `config.yml`
2. Run Docker containers: `make run_containers`
3. Apply migrations: `make migrate`
4. Run the service: `make run`
5. Open the Swagger UI: [http://localhost:44444/v1/docs](http://localhost:44444/v1/docs)
6. Use Swagger UI to interact with the API
7. To stop the service, press `Ctrl+C` in the terminal

## Dependency injection
Dependency injection is implemented using [Wire](https://github.com/google/wire).
To generate the DI code, run the following command:
```bash
make wire
```

## REST API
Code generation is used here from the Swagger 2.0 specification.
It is placed in the [api/rest/swagger.yml](api/rest/swagger.yml) file.
Swagger 2 is used, not OpenAPI 3, because at the moment there are no API code generation
libraries for Go that would fully cover the entire OpenAPI 3 specification. So I 
used [go-swagger](https://github.com/go-swagger/go-swagger) here because of it.

To generate the code from specification, run the following command:
```bash
make generate_server
```
## Database migrations
Working with the database structure (creating tables, indexes, etc.) is done **strictly through migrations**.
1. Create a file with the migration, for example: `goose -dir ./internal/storage/postgresql/migrator create create_table_article sql`.
`create_table_article` is an arbitrary migration name. The `sql` argument at the end of the command is needed to generate an SQL migration, not a Go file
2. In the `./internal/storage/postgresql/migrator/XXX_create_table_article.sql` file, write an SQL query for migration
3. Build the application
4. Run the application with the `migrate` flag: `./myapp migrate`

When developing locally, migrations can be used like this: `make migrate`.
Docker container with PostgreSQL must be running before it (see [docker-compose.yml](docker.compose.yml)).
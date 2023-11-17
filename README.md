# pro-profile-ai-api

## Load environment variables

1. Create `.env` file and populate it.
2. Execute command to load variables into shell

```bash
export $(grep -v '^#' .env | xargs)
```
## Migrations

1. Install golang-migrate CLI by following the instructions in the [official documentation](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate).

2. Create a new migration file using the following command:

  ```bash
    make migrate/create name=<migration_name>
  ```

Replace `<migration_name>` with a descriptive name for your migration.

3. Write your migration code in the newly created file.

4. Run the migration using the following command:

  ```bash
    make migrate/up db_con=<db_connection_string> up
  ```
Replace `db_connection_string`, example: `"postgres://user:password@localhost:5432/dbname?sslmode=disable"`
You can add optional `n=1` argument to apply only certain migrations.

5. To rollback a migration, use the following command:

  ```bash
    make migrate/down db_con=<db_connection_string>
  ```
Same arguments as above.

6. To apply up to certain migration:
```bash
    make migrate/go db_con=<db_connection_string> v=<up_to_version>
```
`up_to_version` replace with number you wish to apply.
`db_con` is optional as it is prefiled with env variables

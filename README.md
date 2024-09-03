# go-pgdump-cli

CLI tool using the library

## Usage of go-pgdump-cli

- `-u` string, username for PostgreSQL

- `-p` string, password for PostgreSQL

- `-h` string, hostname for PostgreSQL

- `-d` string, database name for PostgreSQL

- `-P` string, port number for PostgreSQL (default 5432)

- `-csv` bool, dump to CSV

- `-o` string, path to output directory

- `-px` string, prefix of tablen names for dump

- `-s` string, schema filter for dump

- `-sx` string, suffix of tablen names for dump

- `-tables` string, comma-separated list of table names to dump to CSV

### Usage for a database dump with default port

```bash
./go-pgdump-cli -u user -p example -h localhost -d test -o test -sx example -px test -s myschema
```

### Usage for a CSV dump with custom port

```bash
./go-pgdump-cli -u user -p example -h localhost -d test -P 5433 -o test -csv -tables employees,departments
```

## License

This project is licensed under MIT - see the LICENSE file for details.

# go-pgdump-cli

CLI tool using the library

## Usage of go-pgdump-cli:

-  `-csv` bool, dump to CSV

- `-o` string, path to output directory

- `-px` string, prefix of tablen names for dump

- `-s` string, schema filter for dump

- `-sx` string, suffix of tablen names for dump

- `-tables` string, comma-separated list of table names to dump to CSV

### Usage for a database dump

```bash
./go-pgdump-cli -o test -sx example -px test -s myschema
```

### Usage for a CSV dump

```bash
./go-pgdump-cli -o test -csv -tables employees,departments
```

## License

This project is licensed under MIT - see the LICENSE file for details.

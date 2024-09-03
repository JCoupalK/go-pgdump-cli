module go-pgdump-test

go 1.23.0

replace github.com/JCoupalK/go-pgdump => ../go-pgdump

require github.com/JCoupalK/go-pgdump v0.0.0

require (
	github.com/lib/pq v1.10.9 // indirect
	golang.org/x/sync v0.8.0 // indirect
)

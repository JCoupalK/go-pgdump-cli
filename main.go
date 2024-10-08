package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/JCoupalK/go-pgdump"
)

var (
	username  = flag.String("u", "", "username for PostgreSQL")
	password  = flag.String("p", "", "password for PostgreSQL")
	hostname  = flag.String("h", "", "hostname for PostgreSQL")
	db        = flag.String("d", "", "database name for PostgreSQL")
	port      = flag.Int("P", 5432, "port number for PostgreSQL")
	dumpCSV   = flag.Bool("csv", false, "dump to CSV")
	outputDIR = flag.String("o", "", "path to output directory")
	suffix    = flag.String("sx", "", "suffix of table names for dump")
	prefix    = flag.String("px", "", "prefix of table names for dump")
	schema    = flag.String("s", "", "schema filter for dump")
)

func BackupPostgreSQL(username, password, hostname, dbname, outputDir string, port int) {
	// PostgreSQL connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		hostname, port, username, password, dbname)

	// Create a new dumper instance with connection string and number of threads
	dumper := pgdump.NewDumper(psqlInfo, 50)

	// Check if CSV dump is requested
	if *dumpCSV {
		// Dump all tables to CSV files in the output directory
		err := dumper.DumpDBToCSV(*outputDIR, "go-pgdump.log", &pgdump.TableOptions{
			TableSuffix: *suffix,
			TablePrefix: *prefix,
			Schema:      *schema,
		})
		if err != nil {
			log.Fatal("Error dumping to CSV:", err)
		}
		fmt.Println("CSV files successfully saved in:", *outputDIR)
	} else {
		// Regular SQL dump
		currentTime := time.Now()
		dumpFilename := filepath.Join(
			outputDir,
			fmt.Sprintf("%s-%s.sql", dbname, currentTime.Format("20060102T150405")),
		)

		if err := dumper.DumpDatabase(dumpFilename, &pgdump.TableOptions{
			TableSuffix: *suffix,
			TablePrefix: *prefix,
			Schema:      *schema,
		}); err != nil {
			log.Fatal("Error dumping database:", err)
		}

		fmt.Println("Dump successfully saved to:", dumpFilename)
	}
}

func main() {
	flag.Parse()

	if *outputDIR == "" {
		log.Fatal("Output directory (-o) must be specified")
	}

	BackupPostgreSQL(*username, *password, *hostname, *db, *outputDIR, *port)
}

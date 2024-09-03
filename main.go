package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strings"
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
	csvTables = flag.String("tables", "", "comma-separated list of table names to dump to CSV")
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
		tableList := strings.Split(*csvTables, ",")
		csvFiles, err := dumper.DumpToCSV(outputDir, tableList...)
		if err != nil {
			log.Fatal("Error dumping to CSV:", err)
		}
		fmt.Println("CSV files successfully saved in:", csvFiles)
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
	outputDir := *outputDIR

	BackupPostgreSQL(*username, *password, *hostname, *db, outputDir, *port)
}

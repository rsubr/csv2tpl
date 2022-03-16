// TODO: header documentation
// How to use this application

package main

import (
	"encoding/csv"
	"flag"
	"html/template"
	"io"
	"log"
	"os"
	"strings"
)

func init() {
	log.SetFlags(0) // hide timestamp in log messages
}

func main() {

	var header []string                  // csv header
	var params = make(map[string]string) // go map of csv record
	var noempty bool = false             // error if csv line has empty records
	var verbose bool = false

	flag.BoolVar(&verbose, "verbose", false, "Work on log level info")
	flag.BoolVar(&noempty, "noempty", false, "Stop if any CSV columns are empty")
	flag.Parse()

	if len(flag.Args()) != 2 {
		log.Fatal("Usage: csv2tpl [-verbose] [-noempty] template-file csv-file")
	}

	tplFname := flag.Args()[0]
	csvFname := flag.Args()[1]

	tpl, err := template.ParseFiles(tplFname)
	if err != nil {
		log.Fatal("ERROR: parsing template: ", err)
	}

	csvFh, err := os.Open(csvFname)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFh.Close()

	csvReader := csv.NewReader(csvFh)

	header, err = csvReader.Read()
	if err != nil {
		log.Fatal("ERROR: reading CSV header: ", err)
	}

	for i := range header {
		header[i] = strings.TrimSpace((header[i]))
		if verbose {
			log.Println("Parsed CSV header column:", header[i])
		}
	}

	for line := 1; ; line++ {
		// Stop if end of file reached
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		// Abort on any other errors
		if err != nil {
			log.Fatal("ERROR: reading CSV line ", line, ": ", err)
		}

		// Map CSV line into parameters
		for i := range record {
			params[header[i]] = strings.TrimSpace(record[i])
			if noempty && len(params[header[i]]) == 0 {
				log.Fatal("ERROR: empty column ", header[i], " in CSV line ", i)
			}

			if verbose {
				log.Println("Parsed CSV line", line, "column:", header[i], " data:", params[header[i]])
			}
		}

		err = tpl.Execute(os.Stdout, params)
		if err != nil {
			log.Fatal("ERROR: executing template in CSV line ", line, ": ", err)
		}
	}

	return
}

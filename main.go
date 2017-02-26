package main

import (
	"flag"
	"fmt"
)

var xmlPath = flag.String("xml", "", "path to xml")
var schemaDir = flag.String("xsdDir", "", "path to directory with XSD files to check against")

func main() {
	flag.Parse()

	if len(*xmlPath) == 0 {
		panic("XML path not provided. `--xml=path/to/xml` Exiting.")
	}
	if len(*schemaDir) == 0 {
		panic("Schema Directory not provided. `--xsdDir=path/to/dir/with/xsds` Exiting.")
	}

	err := GetAnalyzer(*xmlPath, *schemaDir).Run()
	if err != nil {
		fmt.Errorf("Error occurred. %s", "")
	} else {
		fmt.Println("Success!")
	}
}

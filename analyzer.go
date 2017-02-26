package main

import (
	"fmt"
	"github.com/aeikenberry/xml-analyzer/lib"
)

type Analyzer struct {
	xmlPath    string
	schemaPath string
}

func GetAnalyzer(xml string, schema string) *Analyzer {
	return &Analyzer{xmlPath: xml, schemaPath: schema}
}

func (a *Analyzer) Run() error {
	xmlData, err := lib.OpenFile(a.xmlPath)
	if err != nil {
		return err
	}

	xmlDocument, err := lib.ParseXMLFile(xmlData)
	if err != nil {
		fmt.Print("Unable to Parse XML File. Exiting.")
		return err
	}

	schemas, err := lib.GetAllSchemas(a.schemaPath)
	if err != nil {
		fmt.Printf("Unable to get schemas for dir: %s", a.schemaPath)
		return err
	}

	for _, schema := range schemas {
		match := lib.MatchesSchema(xmlDocument, schema)
		var matchDisplay string
		if match {
			matchDisplay = "YES"
		} else {
			matchDisplay = "NO"
		}
		fmt.Printf("%s: %s \n", schema.Name, matchDisplay)
	}

	return nil
}

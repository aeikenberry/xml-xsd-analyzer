package lib

import (
	"fmt"
	"github.com/jbussdieker/golibxml"
	"github.com/krolaw/xsd"
	"io/ioutil"
	"os"
	"unsafe"
)

type SchemaFile struct {
	schema *xsd.Schema
	Name   string
}

func MatchesSchema(xml *golibxml.Document, schemaFile *SchemaFile) bool {
	if err := schemaFile.schema.Validate(xsd.DocPtr(unsafe.Pointer(xml.Ptr))); err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func OpenFile(path string) (string, error) {
	xml, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(xml), nil
}

func ParseXMLFile(xmlData string) (*golibxml.Document, error) {
	doc := golibxml.ParseDoc(xmlData)
	if doc == nil {
		fmt.Println("Error parsing document")
		return nil, nil
	}
	defer doc.Free()
	return doc, nil
}

func ParseXSDFile(xsdData string, name string) (*SchemaFile, error) {
	schema, err := xsd.ParseSchema([]byte(xsdData))
	if err != nil {
		return nil, err
	}

	return &SchemaFile{schema, name}, nil
}

func GetDirFileInfo(dirPath string) []os.FileInfo {
	files, _ := ioutil.ReadDir(dirPath)
	return files
}

func GetAllSchemas(dirPath string) ([]*SchemaFile, error) {
	files := GetDirFileInfo(dirPath)
	var schemas []*SchemaFile
	for _, f := range files {
		name := f.Name()
		openFile, err := OpenFile(dirPath + name)
		if err != nil {
			return nil, err
		}

		xsd, err := ParseXSDFile(openFile, name)
		if err != nil {
			return nil, err
		}

		schemas = append(schemas, xsd)
	}
	return schemas, nil
}

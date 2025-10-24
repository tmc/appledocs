// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewXMLParser

// ExampleXMLParser_AbortParsing demonstrates using AbortParsing on a XMLParser instance.
// Stops the parser object.
func ExampleXMLParser_AbortParsing() {
	obj := foundation.NewXMLParser()
	obj.AbortParsing()
	// Output:
	}

// ExampleXMLParser_Parse demonstrates using Parse on a XMLParser instance.
// Starts the event-driven parsing operation.
func ExampleXMLParser_Parse() {
	obj := foundation.NewXMLParser()
	_ = obj.Parse()
	// Output:
	}




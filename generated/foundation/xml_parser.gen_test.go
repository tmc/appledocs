// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewXMLParser

// ExampleNewXMLParserWithContentsOfURL demonstrates how to create a XMLParser instance using NewXMLParserWithContentsOfURL.
// Initializes a parser with the XML content referenced by the given URL.
func ExampleNewXMLParserWithContentsOfURL() {
	_ = foundation.NewXMLParserWithContentsOfURL(
		foundation.URL{}, // url URL
	)
	// Output:
}


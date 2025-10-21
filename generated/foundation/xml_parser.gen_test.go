// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewXMLParser

// ExampleNewXMLParserWithData demonstrates how to create a XMLParser instance using NewXMLParserWithData.
// Initializes a parser with the XML contents encapsulated in a given data object.
func ExampleNewXMLParserWithData() {
	_ = foundation.NewXMLParserWithData(
		foundation.NSData{}, // data NSData
	)
	// Output:
}
// ExampleNewXMLParserWithStream demonstrates how to create a XMLParser instance using NewXMLParserWithStream.
// Initializes a parser with the XML contents from the specified stream and parses it.
func ExampleNewXMLParserWithStream() {
	_ = foundation.NewXMLParserWithStream(
		foundation.NSInputStream{}, // stream NSInputStream
	)
	// Output:
}


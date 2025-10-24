// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit_test

import (
	"github.com/tmc/appledocs/generated/pdfkit"
)

// Suppress unused import errors
var _ = pdfkit.NewPDFActionNamed

// ExampleNewPDFActionNamedWithName demonstrates how to create a PDFActionNamed instance using NewPDFActionNamedWithName.
// Initializes the   object with the specified named action.
func ExampleNewPDFActionNamedWithName() {
	_ = pdfkit.NewPDFActionNamedWithName(
		pdfkit.PDFActionNamedName{}, // name PDFActionNamedName
	)
	// Output:
}

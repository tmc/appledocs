// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit_test

import (
	"github.com/tmc/appledocs/generated/pdfkit"
)

// Suppress unused import errors
var _ = pdfkit.NewPDFOutline

// ExampleNewPDFOutline demonstrates how to create a PDFOutline instance.
// Initializes a   object.
func ExampleNewPDFOutline() {
	_ = pdfkit.NewPDFOutline()
	// Output:
}
// ExamplePDFOutline_RemoveFromParent demonstrates using RemoveFromParent on a PDFOutline instance.
// Removes the outline object from its parent (does nothing if outline object is the root outline object).
func ExamplePDFOutline_RemoveFromParent() {
	obj := pdfkit.NewPDFOutline()
	obj.RemoveFromParent()
	// Output:
	}


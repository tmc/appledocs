// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit_test

import (
	"github.com/tmc/appledocs/generated/pdfkit"
)

// Suppress unused import errors
var _ = pdfkit.NewPDFSelection

// ExampleNewPDFSelectionWithDocument demonstrates how to create a PDFSelection instance using NewPDFSelectionWithDocument.
// Returns an empty   object.
func ExampleNewPDFSelectionWithDocument() {
	_ = pdfkit.NewPDFSelectionWithDocument(
		pdfkit.PDFDocument{}, // document PDFDocument
	)
	// Output:
}

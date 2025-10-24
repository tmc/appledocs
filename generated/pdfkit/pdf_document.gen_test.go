// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit_test

import (
	"github.com/tmc/appledocs/generated/pdfkit"
)

// Suppress unused import errors
var _ = pdfkit.NewPDFDocument

// ExampleNewPDFDocument demonstrates how to create a PDFDocument instance.
// Initializes a   object.
func ExampleNewPDFDocument() {
	_ = pdfkit.NewPDFDocument()
	// Output:
}
// ExamplePDFDocument_CancelFindString demonstrates using CancelFindString on a PDFDocument instance.
// Cancels a search initiated with  .
func ExamplePDFDocument_CancelFindString() {
	obj := pdfkit.NewPDFDocument()
	obj.CancelFindString()
	// Output:
	}

// ExamplePDFDocument_DataRepresentation demonstrates using DataRepresentation on a PDFDocument instance.
// Returns a representation of the document as an   object.
func ExamplePDFDocument_DataRepresentation() {
	obj := pdfkit.NewPDFDocument()
	_ = obj.DataRepresentation()
	// Output:
	}


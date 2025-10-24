// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit_test

import (
	"github.com/tmc/appledocs/generated/pdfkit"
)

// Suppress unused import errors
var _ = pdfkit.NewPDFSelection

// ExamplePDFSelection_ExtendSelectionForLineBoundaries demonstrates using ExtendSelectionForLineBoundaries on a PDFSelection instance.
func ExamplePDFSelection_ExtendSelectionForLineBoundaries() {
	obj := pdfkit.NewPDFSelection()
	obj.ExtendSelectionForLineBoundaries()
	// Output:
	}

// ExamplePDFSelection_SelectionsByLine demonstrates using SelectionsByLine on a PDFSelection instance.
// Returns an array of selections, one for each line of text covered by the receiver.
func ExamplePDFSelection_SelectionsByLine() {
	obj := pdfkit.NewPDFSelection()
	_ = obj.SelectionsByLine()
	// Output:
	}


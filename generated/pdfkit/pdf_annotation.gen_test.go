// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit_test

import (
	"github.com/tmc/appledocs/generated/pdfkit"
)

// Suppress unused import errors
var _ = pdfkit.NewPDFAnnotation

// ExampleNewPDFAnnotationWithBounds demonstrates how to create a PDFAnnotation instance using NewPDFAnnotationWithBounds.
// Creates a PDF annotation object.
func ExampleNewPDFAnnotationWithBounds() {
	_ = pdfkit.NewPDFAnnotationWithBounds(
		pdfkit.Rect /* not a class type */{}, // bounds Rect /* not a class type */
	)
	// Output:
}

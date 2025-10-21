// Code generated from Apple documentation for PDFKit. DO NOT EDIT.
// This file provides fallback type aliases for types that are referenced
// in the generated code but not explicitly defined.

package pdfkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/coregraphics"
)

// Point represents a 2D point (x, y).
// This type is referenced by PDFPage and PDFDocument but not fully documented.
// Using unsafe.Pointer as fallback.
type Point unsafe.Pointer

// Rect represents a rectangle (origin, size).
// This type is referenced by PDFPage and PDFDocument but not fully documented.
// Using unsafe.Pointer as fallback.
type Rect unsafe.Pointer

// Range represents a character range in a string.
// This type is referenced by PDFSelection but not fully documented.
// Using unsafe.Pointer as fallback.
type Range unsafe.Pointer

// CGContextRef is a reference to a CoreGraphics graphics context.
type CGContextRef = coregraphics.CGContextRef

// CGPDFPageRef is a reference to a PDF page object.
type CGPDFPageRef unsafe.Pointer

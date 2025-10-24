// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (

	"github.com/tmc/appledocs/generated/appkit"
)

// PPDFPageOverlayViewProvider is the PDFPageOverlayViewProvider protocol interface.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.pdfkit/documentation/PDFKit/PDFPageOverlayViewProvider
type PPDFPageOverlayViewProvider interface {
	// Required methods
	PdfViewOverlayViewForPage(view IPDFView, page IPDFPage) appkit.View/* debug [protocol_interface/required_method]: PdfViewOverlayViewForPage */
	// Optional methods
	PdfViewWillDisplayOverlayViewForPage(pdfView IPDFView, overlayView appkit.View, page IPDFPage)
	HasPdfViewWillDisplayOverlayViewForPage() bool
	PdfViewWillEndDisplayingOverlayViewForPage(pdfView IPDFView, overlayView appkit.View, page IPDFPage)
	HasPdfViewWillEndDisplayingOverlayViewForPage() bool
}

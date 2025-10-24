// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

// PWebDocumentRepresentation is the WebDocumentRepresentation protocol interface.
//
// This protocol is adopted by document representation classes that handle specific MIME types. You can implement your own document view classes and document representation classes to render data for specific MIME types, and register those classes using the     method.
//
// Availability:
//   - macOS 10.3+ (Deprecated in 10.14)
//
// See: doc://com.apple.webkit/documentation/WebKit/WebDocumentRepresentation
type PWebDocumentRepresentation interface {
	// Required methods
	CanProvideDocumentSource() bool
	DocumentSource() foundation.String
	FinishedLoadingWithDataSource(dataSource IWebDataSource)
	ReceivedDataWithDataSource(data objc.IObject /* cross-framework: NSData */, dataSource IWebDataSource)
	ReceivedErrorWithDataSource(error_ objc.IObject /* cross-framework: Error */, dataSource IWebDataSource)
	SetDataSource(dataSource IWebDataSource)
	Title() foundation.String
}

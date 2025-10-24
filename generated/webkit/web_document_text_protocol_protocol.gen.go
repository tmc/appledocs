// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import "github.com/tmc/appledocs/generated/foundation"

// PWebDocumentText is the WebDocumentText protocol interface.
//
//	is an optional protocol for document view objects that display text. This protocol defines methods for accessing document content as strings, and methods for text selection. Classes that adopt this protocol should also adopt   and inherit from  .
//
// Availability:
//   - macOS 10.3+ (Deprecated in 10.14)
//
// See: doc://com.apple.webkit/documentation/WebKit/WebDocumentText
type PWebDocumentText interface {
	// Required methods
	AttributedString() foundation.AttributedString
	DeselectAll()
	SelectAll()
	SelectedAttributedString() foundation.AttributedString
	SelectedString() foundation.String
	String() foundation.String
	SupportsTextEncoding() bool
}

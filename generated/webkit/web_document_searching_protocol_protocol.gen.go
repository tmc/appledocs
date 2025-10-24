// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import "github.com/tmc/appledocs/generated/objc"

// PWebDocumentSearching is the WebDocumentSearching protocol interface.
//
//	is an optional protocol for document view objects that support searching. Classes that adopt this protocol should also adopt   and inherit from  .
//
// Availability:
//   - macOS 10.3+ (Deprecated in 10.14)
//
// See: doc://com.apple.webkit/documentation/WebKit/WebDocumentSearching
type PWebDocumentSearching interface {
	// Required methods
	SearchForDirectionCaseSensitiveWrap(string_ objc.IObject /* cross-framework: NSString */, forward bool, caseFlag bool, wrapFlag bool) bool
}

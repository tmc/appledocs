// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PWebDocumentText is the WebDocumentText protocol interface.
//
//  is an optional protocol for document view objects that display text. This protocol defines methods for accessing document content as strings, and methods for text selection. Classes that adopt this protocol should also adopt   and inherit from  .
//
// Availability:
//   - macOS 10.3+ (Deprecated in 10.14)
//
// See: doc://com.apple.webkit/documentation/WebKit/WebDocumentText
type PWebDocumentText interface {
	// Required methods
	AttributedString() foundation.AttributedString/* debug [protocol_interface/required_method]: AttributedString */
	DeselectAll()/* debug [protocol_interface/required_method]: DeselectAll */
	SelectAll()/* debug [protocol_interface/required_method]: SelectAll */
	SelectedAttributedString() foundation.AttributedString/* debug [protocol_interface/required_method]: SelectedAttributedString */
	SelectedString() foundation.String/* debug [protocol_interface/required_method]: SelectedString */
	String() foundation.String/* debug [protocol_interface/required_method]: String */
	SupportsTextEncoding() bool/* debug [protocol_interface/required_method]: SupportsTextEncoding */
}

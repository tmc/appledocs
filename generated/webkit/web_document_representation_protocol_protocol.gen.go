// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"
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
	CanProvideDocumentSource() bool/* debug [protocol_interface/required_method]: CanProvideDocumentSource */
	DocumentSource() foundation.String/* debug [protocol_interface/required_method]: DocumentSource */
	FinishedLoadingWithDataSource(dataSource IWebDataSource)/* debug [protocol_interface/required_method]: FinishedLoadingWithDataSource */
	ReceivedDataWithDataSource(data objc.IObject /* cross-framework: NSData */, dataSource IWebDataSource)/* debug [protocol_interface/required_method]: ReceivedDataWithDataSource */
	ReceivedErrorWithDataSource(error_ objc.IObject /* cross-framework: Error */, dataSource IWebDataSource)/* debug [protocol_interface/required_method]: ReceivedErrorWithDataSource */
	SetDataSource(dataSource IWebDataSource)/* debug [protocol_interface/required_method]: SetDataSource */
	Title() foundation.String/* debug [protocol_interface/required_method]: Title */
}

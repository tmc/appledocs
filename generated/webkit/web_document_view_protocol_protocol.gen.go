// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (

	"github.com/tmc/appledocs/generated/appkit"
)

// PWebDocumentView is the WebDocumentView protocol interface.
//
// This protocol is adopted by the document view of a  . You can extend WebKit to support additional MIME types by implementing your own document view and document representation classes to render data for specific MIME types. You register those classes using the WebFrame   method. Classes that adopt this protocol are expected to be subclasses of  .
//
// Availability:
//   - macOS 10.3+ (Deprecated in 10.14)
//
// See: doc://com.apple.webkit/documentation/WebKit/WebDocumentView
type PWebDocumentView interface {
	// Required methods
	DataSourceUpdated(dataSource IWebDataSource)/* debug [protocol_interface/required_method]: DataSourceUpdated */
	Layout()/* debug [protocol_interface/required_method]: Layout */
	SetDataSource(dataSource IWebDataSource)/* debug [protocol_interface/required_method]: SetDataSource */
	SetNeedsLayout(flag bool)/* debug [protocol_interface/required_method]: SetNeedsLayout */
	ViewDidMoveToHostWindow()/* debug [protocol_interface/required_method]: ViewDidMoveToHostWindow */
	ViewWillMoveToHostWindow(hostWindow appkit.Window)/* debug [protocol_interface/required_method]: ViewWillMoveToHostWindow */
}

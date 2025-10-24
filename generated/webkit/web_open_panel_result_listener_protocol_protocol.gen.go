// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PWebOpenPanelResultListener is the WebOpenPanelResultListener protocol interface.
//
//  user interface delegates that implement the webView:runOpenPanelForFileButtonWithResultListener: method use the methods defined in this protocol to communicate with the listener object. The methods allow the delegate to send a cancel message, or set the selected file name.
//
// Availability:
//   - macOS 10.3+ (Deprecated in 10.14)
//
// See: doc://com.apple.webkit/documentation/WebKit/WebOpenPanelResultListener
type PWebOpenPanelResultListener interface {
	// Required methods
	Cancel()/* debug [protocol_interface/required_method]: Cancel */
	ChooseFilename(fileName objc.IObject /* cross-framework: NSString */)/* debug [protocol_interface/required_method]: ChooseFilename */
	ChooseFilenames(fileNames objc.IObject /* cross-framework: NSArray */)/* debug [protocol_interface/required_method]: ChooseFilenames */
}

// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

// PScriptMessageHandler is the WKScriptMessageHandler protocol interface.
//
// An interface for receiving messages from JavaScript code running in a webpage.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - visionOS +
//
// See: doc://com.apple.webkit/documentation/WebKit/WKScriptMessageHandler
type PScriptMessageHandler interface {
	// Required methods
	UserContentControllerDidReceiveScriptMessage(userContentController IWKUserContentController, message IWKScriptMessage)/* debug [protocol_interface/required_method]: UserContentControllerDidReceiveScriptMessage */
}

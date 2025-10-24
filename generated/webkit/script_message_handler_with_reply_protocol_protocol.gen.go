// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/func(objc"
)

// PScriptMessageHandlerWithReply is the WKScriptMessageHandlerWithReply protocol interface.
//
// An interface for responding to messages from JavaScript code running in a webpage.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - visionOS +
//
// See: doc://com.apple.webkit/documentation/WebKit/WKScriptMessageHandlerWithReply
type PScriptMessageHandlerWithReply interface {
	// Required methods
	UserContentControllerDidReceiveScriptMessageReplyHandler(userContentController IWKUserContentController, message IWKScriptMessage, replyHandler func(objc.ID, unsafe.Pointer))/* debug [protocol_interface/required_method]: UserContentControllerDidReceiveScriptMessageReplyHandler */
}

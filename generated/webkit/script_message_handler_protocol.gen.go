// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import "github.com/ebitengine/purego/objc"

// ScriptMessageHandlerProtocol is the WKScriptMessageHandler protocol.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - visionOS +
//
// Use this protocol when registering custom classes that conform to WKScriptMessageHandler.
var ScriptMessageHandlerProtocol *objc.Protocol

func init() {
	ScriptMessageHandlerProtocol = objc.GetProtocol("WKScriptMessageHandler")
}

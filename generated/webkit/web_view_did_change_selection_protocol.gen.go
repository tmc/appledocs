// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import "github.com/ebitengine/purego/objc"

// webViewDidChangeSelectionProtocol is the webViewDidChangeSelection: protocol.
//
// Availability:
//   - macOS 10.3+ (Deprecated in 10.14)
//
// Use this protocol when registering custom classes that conform to webViewDidChangeSelection:.
var webViewDidChangeSelectionProtocol *objc.Protocol

func init() {
	webViewDidChangeSelectionProtocol = objc.GetProtocol("webViewDidChangeSelection:")
}

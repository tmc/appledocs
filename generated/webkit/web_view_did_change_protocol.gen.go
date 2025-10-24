// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import "github.com/ebitengine/purego/objc"

// webViewDidChangeProtocol is the webViewDidChange: protocol.
//
// Availability:
//   - macOS 10.3+ (Deprecated in 10.14)
//
// Use this protocol when registering custom classes that conform to webViewDidChange:.
var webViewDidChangeProtocol *objc.Protocol

func init() {
	webViewDidChangeProtocol = objc.GetProtocol("webViewDidChange:")
}


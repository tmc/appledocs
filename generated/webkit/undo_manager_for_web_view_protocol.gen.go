// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import "github.com/ebitengine/purego/objc"

// undoManagerForWebViewProtocol is the undoManagerForWebView: protocol.
//
// Availability:
//   - macOS 10.3+ (Deprecated in 10.14)
//
// Use this protocol when registering custom classes that conform to undoManagerForWebView:.
var undoManagerForWebViewProtocol *objc.Protocol

func init() {
	undoManagerForWebViewProtocol = objc.GetProtocol("undoManagerForWebView:")
}


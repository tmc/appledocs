// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import "github.com/ebitengine/purego/objc"

// webViewProtocol is the webView: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to webView:.
var webViewProtocol *objc.Protocol

func init() {
	webViewProtocol = objc.GetProtocol("webView:")
}

// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import "github.com/ebitengine/purego/objc"

// webViewDidCloseProtocol is the webViewDidClose: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to webViewDidClose:.
var webViewDidCloseProtocol *objc.Protocol

func init() {
	webViewDidCloseProtocol = objc.GetProtocol("webViewDidClose:")
}


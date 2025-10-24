// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import "github.com/ebitengine/purego/objc"

// didReceiveDataProtocol is the didReceiveData: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.13+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to didReceiveData:.
var didReceiveDataProtocol *objc.Protocol

func init() {
	didReceiveDataProtocol = objc.GetProtocol("didReceiveData:")
}

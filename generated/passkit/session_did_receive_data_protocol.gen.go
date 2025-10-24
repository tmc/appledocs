// Code generated from Apple documentation for PassKit. DO NOT EDIT.

package passkit

import "github.com/ebitengine/purego/objc"

// sessionDidReceiveDataProtocol is the sessionDidReceiveData: protocol.
//
// Availability:
//   - Mac Catalyst 15.4+
//   - iOS 15.4+
//   - iPadOS 15.4+
//   - macOS 10.10+
//   - visionOS 1.0+
//   - watchOS 8.5+
//
// Use this protocol when registering custom classes that conform to sessionDidReceiveData:.
var sessionDidReceiveDataProtocol *objc.Protocol

func init() {
	sessionDidReceiveDataProtocol = objc.GetProtocol("sessionDidReceiveData:")
}


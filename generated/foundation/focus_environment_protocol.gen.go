// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import "github.com/ebitengine/purego/objc"

// FocusEnvironmentProtocol is the UIFocusEnvironment protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to UIFocusEnvironment.
var FocusEnvironmentProtocol *objc.Protocol

func init() {
	FocusEnvironmentProtocol = objc.GetProtocol("UIFocusEnvironment")
}

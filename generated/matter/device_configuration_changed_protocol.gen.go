// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import "github.com/ebitengine/purego/objc"

// deviceConfigurationChangedProtocol is the deviceConfigurationChanged: protocol.
//
// Availability:
//   - Mac Catalyst 17.6+
//   - iOS 17.6+
//   - iPadOS 17.6+
//   - macOS 14.6+
//   - tvOS 17.6+
//   - visionOS 1.0+
//   - watchOS 10.6+
//
// Use this protocol when registering custom classes that conform to deviceConfigurationChanged:.
var deviceConfigurationChangedProtocol *objc.Protocol

func init() {
	deviceConfigurationChangedProtocol = objc.GetProtocol("deviceConfigurationChanged:")
}


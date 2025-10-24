// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import "github.com/ebitengine/purego/objc"

// objectForKeyedSubscriptProtocol is the objectForKeyedSubscript: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 3.0+
//
// Use this protocol when registering custom classes that conform to objectForKeyedSubscript:.
var objectForKeyedSubscriptProtocol *objc.Protocol

func init() {
	objectForKeyedSubscriptProtocol = objc.GetProtocol("objectForKeyedSubscript:")
}


// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import "github.com/ebitengine/purego/objc"

// sessionWasSuspendedProtocol is the sessionWasSuspended: protocol.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - watchOS 7.3+
//
// Use this protocol when registering custom classes that conform to sessionWasSuspended:.
var sessionWasSuspendedProtocol *objc.Protocol

func init() {
	sessionWasSuspendedProtocol = objc.GetProtocol("sessionWasSuspended:")
}

// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import "github.com/ebitengine/purego/objc"

// sessionDidStartRunningProtocol is the sessionDidStartRunning: protocol.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - watchOS 9.0+
//
// Use this protocol when registering custom classes that conform to sessionDidStartRunning:.
var sessionDidStartRunningProtocol *objc.Protocol

func init() {
	sessionDidStartRunningProtocol = objc.GetProtocol("sessionDidStartRunning:")
}


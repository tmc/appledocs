// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import "github.com/ebitengine/purego/objc"

// sessionProtocol is the session: protocol.
//
// Availability:
//   - iOS 26.0+
//   - iPadOS 26.0+
//
// Use this protocol when registering custom classes that conform to session:.
var sessionProtocol *objc.Protocol

func init() {
	sessionProtocol = objc.GetProtocol("session:")
}

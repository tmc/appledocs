// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import "github.com/ebitengine/purego/objc"

// locationManagerDidChangeAuthorizationProtocol is the locationManagerDidChangeAuthorization: protocol.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+
//
// Use this protocol when registering custom classes that conform to locationManagerDidChangeAuthorization:.
var locationManagerDidChangeAuthorizationProtocol *objc.Protocol

func init() {
	locationManagerDidChangeAuthorizationProtocol = objc.GetProtocol("locationManagerDidChangeAuthorization:")
}



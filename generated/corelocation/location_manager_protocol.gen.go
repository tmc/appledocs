// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import "github.com/ebitengine/purego/objc"

// locationManagerProtocol is the locationManager: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.9+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// Use this protocol when registering custom classes that conform to locationManager:.
var locationManagerProtocol *objc.Protocol

func init() {
	locationManagerProtocol = objc.GetProtocol("locationManager:")
}



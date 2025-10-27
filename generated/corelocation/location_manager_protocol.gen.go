// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import "github.com/ebitengine/purego/objc"

// locationManagerProtocol is the locationManager: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 14.0)
//   - iOS 4.2+ (Deprecated in 14.0)
//   - iPadOS 4.2+ (Deprecated in 14.0)
//   - macOS 10.7+ (Deprecated in 11.0)
//   - tvOS 9.0+ (Deprecated in 14.0)
//   - watchOS 1.0+ (Deprecated in 7.0)
//
// Use this protocol when registering custom classes that conform to locationManager:.
var locationManagerProtocol *objc.Protocol

func init() {
	locationManagerProtocol = objc.GetProtocol("locationManager:")
}


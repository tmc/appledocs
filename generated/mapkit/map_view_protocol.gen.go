// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import "github.com/ebitengine/purego/objc"

// mapViewProtocol is the mapView: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 3.0+
//   - iPadOS 3.0+
//   - macOS 10.9+
//   - tvOS 9.2+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to mapView:.
var mapViewProtocol *objc.Protocol

func init() {
	mapViewProtocol = objc.GetProtocol("mapView:")
}

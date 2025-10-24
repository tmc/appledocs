// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

// PMKOverlay is the MKOverlay protocol interface.
//
// An interface for associating content with a specific map region.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.9+
//   - tvOS 9.2+
//   - visionOS 1.0+
//   - watchOS 1.0+
//
// See: doc://com.apple.mapkit/documentation/MapKit/MKOverlay
type PMKOverlay interface {
	// Optional methods
	CanReplaceMapContent() bool
	HasCanReplaceMapContent() bool
	IntersectsMapRect(mapRect objc.IObject /* cross-framework: MKMapRect */) bool
	HasIntersectsMapRect() bool
}

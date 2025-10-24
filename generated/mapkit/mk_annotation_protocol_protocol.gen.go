// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

// PMKAnnotation is the MKAnnotation protocol interface.
//
// An interface for associating your content with a specific map location.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - tvOS +
//   - visionOS +
//   - watchOS +
//
// See: doc://com.apple.mapkit/documentation/MapKit/MKAnnotation
type PMKAnnotation interface {
	// Required methods
	SetCoordinate(newCoordinate LocationCoordinate2D /* not a class type */)/* debug [protocol_interface/required_method]: SetCoordinate */
}

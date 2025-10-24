// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKMapCamera

// ExampleNewMKMapCameraLookingAtCenterCoordinateFromDistancePitchHeading demonstrates how to create a MKMapCamera instance using NewMKMapCameraLookingAtCenterCoordinateFromDistancePitchHeading.
// Returns a new camera object using the specified distance, pitch, and heading information.
func ExampleNewMKMapCameraLookingAtCenterCoordinateFromDistancePitchHeading() {
	_ = mapkit.NewMKMapCameraLookingAtCenterCoordinateFromDistancePitchHeading(
		mapkit.LocationCoordinate2D /* not a class type */{}, // centerCoordinate LocationCoordinate2D /* not a class type */
		mapkit.LocationDistance /* not a class type */{}, // distance LocationDistance /* not a class type */
		0.0, // pitch float64
		mapkit.LocationDirection /* not a class type */{}, // heading LocationDirection /* not a class type */
	)
	// Output:
}
// ExampleNewMKMapCameraLookingAtCenterCoordinateFromEyeCoordinateEyeAltitude demonstrates how to create a MKMapCamera instance using NewMKMapCameraLookingAtCenterCoordinateFromEyeCoordinateEyeAltitude.
// Returns a new camera object using the specified viewing angle information.
func ExampleNewMKMapCameraLookingAtCenterCoordinateFromEyeCoordinateEyeAltitude() {
	_ = mapkit.NewMKMapCameraLookingAtCenterCoordinateFromEyeCoordinateEyeAltitude(
		mapkit.LocationCoordinate2D /* not a class type */{}, // centerCoordinate LocationCoordinate2D /* not a class type */
		mapkit.LocationCoordinate2D /* not a class type */{}, // eyeCoordinate LocationCoordinate2D /* not a class type */
		mapkit.LocationDistance /* not a class type */{}, // eyeAltitude LocationDistance /* not a class type */
	)
	// Output:
}

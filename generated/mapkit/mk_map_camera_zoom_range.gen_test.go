// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKMapCameraZoomRange

// ExampleNewMKMapCameraZoomRangeWithMaxCenterCoordinateDistance demonstrates how to create a MKMapCameraZoomRange instance using NewMKMapCameraZoomRangeWithMaxCenterCoordinateDistance.
// Create a camera zoom range by specifying the maximum distance from your map view’s center coordinate, measured in meters.
func ExampleNewMKMapCameraZoomRangeWithMaxCenterCoordinateDistance() {
	_ = mapkit.NewMKMapCameraZoomRangeWithMaxCenterCoordinateDistance(
		mapkit.LocationDistance /* not a class type */{}, // maxDistance LocationDistance /* not a class type */
	)
	// Output:
}
// ExampleNewMKMapCameraZoomRangeWithMinCenterCoordinateDistance demonstrates how to create a MKMapCameraZoomRange instance using NewMKMapCameraZoomRangeWithMinCenterCoordinateDistance.
// Create a camera zoom range by specifying the minimum distance from your map view’s center coordinate, measured in meters.
func ExampleNewMKMapCameraZoomRangeWithMinCenterCoordinateDistance() {
	_ = mapkit.NewMKMapCameraZoomRangeWithMinCenterCoordinateDistance(
		mapkit.LocationDistance /* not a class type */{}, // minDistance LocationDistance /* not a class type */
	)
	// Output:
}
// ExampleNewMKMapCameraZoomRangeWithMinCenterCoordinateDistanceMaxCenterCoordinateDistance demonstrates how to create a MKMapCameraZoomRange instance using NewMKMapCameraZoomRangeWithMinCenterCoordinateDistanceMaxCenterCoordinateDistance.
// Create a camera zoom range by specifying a minimum and maximum distance from your map view’s center coordinates, measured in meters.
func ExampleNewMKMapCameraZoomRangeWithMinCenterCoordinateDistanceMaxCenterCoordinateDistance() {
	_ = mapkit.NewMKMapCameraZoomRangeWithMinCenterCoordinateDistanceMaxCenterCoordinateDistance(
		mapkit.LocationDistance /* not a class type */{}, // minDistance LocationDistance /* not a class type */
		mapkit.LocationDistance /* not a class type */{}, // maxDistance LocationDistance /* not a class type */
	)
	// Output:
}

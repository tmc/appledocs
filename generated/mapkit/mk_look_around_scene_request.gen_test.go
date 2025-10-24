// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKLookAroundSceneRequest

// ExampleNewMKLookAroundSceneRequestWithCoordinate demonstrates how to create a MKLookAroundSceneRequest instance using NewMKLookAroundSceneRequestWithCoordinate.
// Creates a LookAround scene at the specified coordinates.
func ExampleNewMKLookAroundSceneRequestWithCoordinate() {
	_ = mapkit.NewMKLookAroundSceneRequestWithCoordinate(
		mapkit.LocationCoordinate2D /* not a class type */{}, // coordinate LocationCoordinate2D /* not a class type */
	)
	// Output:
}
// ExampleMKLookAroundSceneRequest_Cancel demonstrates using Cancel on a MKLookAroundSceneRequest instance.
// Cancels the pending scene request.
func ExampleMKLookAroundSceneRequest_Cancel() {
	obj := mapkit.NewMKLookAroundSceneRequest()
	obj.Cancel()
	// Output:
	}


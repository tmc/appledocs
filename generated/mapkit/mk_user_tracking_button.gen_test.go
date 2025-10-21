// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit_test

import (
	"github.com/tmc/appledocs/generated/mapkit"
)

// Suppress unused import errors
var _ = mapkit.NewMKUserTrackingButton

// ExampleNewMKUserTrackingButtonWithMapView demonstrates how to create a MKUserTrackingButton instance using NewMKUserTrackingButtonWithMapView.
// Initializes the button with the map view that it should control.
func ExampleNewMKUserTrackingButtonWithMapView() {
	_ = mapkit.NewMKUserTrackingButtonWithMapView(
		mapkit.MKMapView{}, // mapView MKMapView
	)
	// Output:
}

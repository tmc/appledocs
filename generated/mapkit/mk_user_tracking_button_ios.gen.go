//go:build darwin && ios

// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for MKUserTrackingButton


// iOS-only properties

// The map view associated with the button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserTrackingButton/mapView
func (m_ MKUserTrackingButton) MapView() IMKMapView {
	rv := objc.Send[MKMapView](m_.ID, objc.Sel("mapView"))
	return rv
}
func (m_ MKUserTrackingButton) SetMapView(value IMKMapView) {
	m_.ID.Send(objc.RegisterName("setMapView:"), value)
}





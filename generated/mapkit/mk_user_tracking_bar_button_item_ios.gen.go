//go:build darwin && ios

// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for MKUserTrackingBarButtonItem


// iOS-only properties

// The map view associated with this bar button item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserTrackingBarButtonItem/mapView
func (m_ MKUserTrackingBarButtonItem) MapView() IMKMapView {
	rv := objc.Send[MKMapView](m_.ID, objc.Sel("mapView"))
	return rv
}
func (m_ MKUserTrackingBarButtonItem) SetMapView(value IMKMapView) {
	m_.ID.Send(objc.RegisterName("setMapView:"), value)
}





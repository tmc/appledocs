//go:build darwin && ios

// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for MKScaleView


// iOS-only properties

// The alignment of the distance information in the scale view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKScaleView/legendAlignment
func (m_ MKScaleView) LegendAlignment() MKScaleViewAlignment {
	rv := objc.Send[MKScaleViewAlignment](m_.ID, objc.Sel("legendAlignment"))
	return rv
}
func (m_ MKScaleView) SetLegendAlignment(value MKScaleViewAlignment) {
	m_.ID.Send(objc.RegisterName("setLegendAlignment:"), value)
}

// The map view that provides the scale information to the scale view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKScaleView/mapView
func (m_ MKScaleView) MapView() IMKMapView {
	rv := objc.Send[MKMapView](m_.ID, objc.Sel("mapView"))
	return rv
}
func (m_ MKScaleView) SetMapView(value IMKMapView) {
	m_.ID.Send(objc.RegisterName("setMapView:"), value)
}

// The visibility of the scale view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKScaleView/scaleVisibility
func (m_ MKScaleView) ScaleVisibility() MKFeatureVisibility {
	rv := objc.Send[MKFeatureVisibility](m_.ID, objc.Sel("scaleVisibility"))
	return rv
}
func (m_ MKScaleView) SetScaleVisibility(value MKFeatureVisibility) {
	m_.ID.Send(objc.RegisterName("setScaleVisibility:"), value)
}





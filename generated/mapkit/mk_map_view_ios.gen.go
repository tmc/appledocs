//go:build darwin && ios

// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for MKMapView


// iOS-only properties

// The property that describes which selectable features the map responds to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapView/selectableMapFeatures
func (m_ MKMapView) SelectableMapFeatures() MKMapFeatureOptions {
	rv := objc.Send[MKMapFeatureOptions](m_.ID, objc.Sel("selectableMapFeatures"))
	return rv
}
func (m_ MKMapView) SetSelectableMapFeatures(value MKMapFeatureOptions) {
	m_.ID.Send(objc.RegisterName("setSelectableMapFeatures:"), value)
}






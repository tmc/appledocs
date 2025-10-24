//go:build darwin && ios

// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for MKReverseGeocoder


// iOS-only properties

// The coordinate whose placemark data you want to retrieve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKReverseGeocoder/coordinate
func (m_ MKReverseGeocoder) Coordinate() LocationCoordinate2D /* not a class type */ {
	rv := objc.Send[LocationCoordinate2D](m_.ID, objc.Sel("coordinate"))
	return rv
}

// The reverse geocoder’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKReverseGeocoder/delegate
func (m_ MKReverseGeocoder) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("delegate"))
	return rv
}
func (m_ MKReverseGeocoder) SetDelegate(value unsafe.Pointer) {
	m_.ID.Send(objc.RegisterName("setDelegate:"), value)
}

// The result of the reverse-geocoding operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKReverseGeocoder/placemark
func (m_ MKReverseGeocoder) Placemark() IMKPlacemark {
	rv := objc.Send[MKPlacemark](m_.ID, objc.Sel("placemark"))
	return rv
}

// A Boolean value indicating whether the receiver is in the middle of reverse-geocoding its coordinate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKReverseGeocoder/querying
func (m_ MKReverseGeocoder) Querying() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("querying"))
	return rv
}





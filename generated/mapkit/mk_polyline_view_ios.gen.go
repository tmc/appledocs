//go:build darwin && ios

// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for MKPolylineView


// iOS-only properties

// The polyline overlay object that contains the information used to draw the overlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolylineView/polyline
func (m_ MKPolylineView) Polyline() IMKPolyline {
	rv := objc.Send[MKPolyline](m_.ID, objc.Sel("polyline"))
	return rv
}





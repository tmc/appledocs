//go:build darwin && ios

// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for MKPolygonView


// iOS-only properties

// The polygon overlay object that contains the information used to draw the overlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolygonView/polygon
func (m_ MKPolygonView) Polygon() IMKPolygon {
	rv := objc.Send[MKPolygon](m_.ID, objc.Sel("polygon"))
	return rv
}





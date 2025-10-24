//go:build darwin && ios

// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for MKCircleView


// iOS-only properties

// The circle overlay object that contains the information used to draw the overlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircleView/circle
func (m_ MKCircleView) Circle() IMKCircle {
	rv := objc.Send[MKCircle](m_.ID, objc.Sel("circle"))
	return rv
}





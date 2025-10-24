//go:build darwin && ios

// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for AltitudeData


// iOS-only properties

// The recorded pressure, in kilopascals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAltitudeData/pressure
func (a_ AltitudeData) Pressure() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](a_.ID, objc.Sel("pressure"))
	return rv
}

// The change in altitude (in meters) since the first reported event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAltitudeData/relativeAltitude
func (a_ AltitudeData) RelativeAltitude() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](a_.ID, objc.Sel("relativeAltitude"))
	return rv
}






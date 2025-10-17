// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Beacon] class.
var beaconClass = _BeaconClass{objc.GetClass("CLBeacon")}

type _BeaconClass struct {
	class objc.Class
}

// Information about an observed iBeacon device and its relative distance to a person’s device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeacon

type Beacon struct {
	objectivec.Object
}

// BeaconFrom constructs a [Beacon] from an unsafe.Pointer.
//
// Information about an observed iBeacon device and its relative distance to a person’s device.
func BeaconFrom(ptr unsafe.Pointer) Beacon {
	return Beacon{objectivec.Object{objc.ID(ptr)}}
}




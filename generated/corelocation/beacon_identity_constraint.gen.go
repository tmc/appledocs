// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BeaconIdentityConstraint] class.
var beaconIdentityConstraintClass = _BeaconIdentityConstraintClass{objc.GetClass("CLBeaconIdentityConstraint")}

type _BeaconIdentityConstraintClass struct {
	class objc.Class
}

// Identity characteristics that can match one or more beacons. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeaconIdentityConstraint

type BeaconIdentityConstraint struct {
	BeaconIdentityCondition
}

// BeaconIdentityConstraintFrom constructs a [BeaconIdentityConstraint] from an unsafe.Pointer.
//
// Identity characteristics that can match one or more beacons.
func BeaconIdentityConstraintFrom(ptr unsafe.Pointer) BeaconIdentityConstraint {
	return BeaconIdentityConstraint{
		BeaconIdentityCondition: BeaconIdentityConditionFrom(ptr),
	}
}




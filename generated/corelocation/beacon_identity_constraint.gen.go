// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BeaconIdentityConstraint] class.
var (
	beaconIdentityConstraintClass     _BeaconIdentityConstraintClass
	beaconIdentityConstraintClassOnce sync.Once
)

func getBeaconIdentityConstraintClass() _BeaconIdentityConstraintClass {
	beaconIdentityConstraintClassOnce.Do(func() {
		beaconIdentityConstraintClass = _BeaconIdentityConstraintClass{objc.GetClass("CLBeaconIdentityConstraint")}
	})
	return beaconIdentityConstraintClass
}

type _BeaconIdentityConstraintClass struct {
	class objc.Class
}

// An interface definition for the [BeaconIdentityConstraint] class.
type IBeaconIdentityConstraint interface {
	IBeaconIdentityCondition
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

// Alloc allocates a new instance without initialization.
func (bc _BeaconIdentityConstraintClass) Alloc() BeaconIdentityConstraint {
	rv := objc.Send[BeaconIdentityConstraint](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BeaconIdentityConstraintClass) New() BeaconIdentityConstraint {
	rv := objc.Send[BeaconIdentityConstraint](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BeaconIdentityConstraint) Init() BeaconIdentityConstraint {
	rv := objc.Send[BeaconIdentityConstraint](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BeaconIdentityConstraint) Autorelease() BeaconIdentityConstraint {
	rv := objc.Send[BeaconIdentityConstraint](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBeaconIdentityConstraint creates a new BeaconIdentityConstraint instance.
func NewBeaconIdentityConstraint() BeaconIdentityConstraint {
	return getBeaconIdentityConstraintClass().New()
}





// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [BeaconIdentityConstraint] class.
var (
	BeaconIdentityConstraintClass     _BeaconIdentityConstraintClass
	BeaconIdentityConstraintClassOnce sync.Once
)

func getBeaconIdentityConstraintClass() _BeaconIdentityConstraintClass {
	BeaconIdentityConstraintClassOnce.Do(func() {
		BeaconIdentityConstraintClass = _BeaconIdentityConstraintClass{objc.GetClass("CLBeaconIdentityConstraint")}
	})
	return BeaconIdentityConstraintClass
}

type _BeaconIdentityConstraintClass struct {
	class objc.Class
}

// An interface definition for the [BeaconIdentityConstraint] class.
type IBeaconIdentityConstraint interface {
	IBeaconIdentityCondition
	Major() foundation.Number
	SetMajor(value foundation.INumber)
	Minor() foundation.Number
	SetMinor(value foundation.INumber)
	Uuid() foundation.UUID
	SetUuid(value foundation.IUUID)
}

// Identity characteristics that can match one or more beacons.
//
// A constraint specifies beacon identity characteristics. Use constraints to check for matching beacons by comparing the beacon’s identity characteristics ( , , and ) to those in the constraint. Constraints always specify a UUID value, but the major and minor values are optional. A beacon satisfies the constraint if all three identity characteristics of the beacon match the same characteristic of the constraint. Major and minor characteristics are wildcards if they have no value. A major or minor wildcard value matches any value in the beacon’s corresponding characteristic.
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


// The major value that the observed beacon transmitted.
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clbeacon/major
func (b_ BeaconIdentityConstraint) Major() foundation.Number {
	rv := objc.Send[foundation.Number](b_.ID, objc.Sel("major"))
	return rv
}


// SetMajor sets the value of the major property.
// The major value that the observed beacon transmitted.

//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clbeacon/major
func (b_ BeaconIdentityConstraint) SetMajor(value foundation.INumber) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setMajor:"), value)
}

// The minor value that the observed beacon transmitted.
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clbeacon/minor
func (b_ BeaconIdentityConstraint) Minor() foundation.Number {
	rv := objc.Send[foundation.Number](b_.ID, objc.Sel("minor"))
	return rv
}


// SetMinor sets the value of the minor property.
// The minor value that the observed beacon transmitted.

//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clbeacon/minor
func (b_ BeaconIdentityConstraint) SetMinor(value foundation.INumber) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setMinor:"), value)
}

// The UUID that the observed beacon transmitted.
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clbeacon/uuid
func (b_ BeaconIdentityConstraint) Uuid() foundation.UUID {
	rv := objc.Send[foundation.UUID](b_.ID, objc.Sel("uuid"))
	return rv
}


// SetUuid sets the value of the uuid property.
// The UUID that the observed beacon transmitted.

//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clbeacon/uuid
func (b_ BeaconIdentityConstraint) SetUuid(value foundation.IUUID) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setUuid:"), value)
}




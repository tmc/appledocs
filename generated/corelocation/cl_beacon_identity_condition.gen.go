// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BeaconIdentityCondition] class.
var (
	BeaconIdentityConditionClass     _BeaconIdentityConditionClass
	BeaconIdentityConditionClassOnce sync.Once
)

func getBeaconIdentityConditionClass() _BeaconIdentityConditionClass {
	BeaconIdentityConditionClassOnce.Do(func() {
		BeaconIdentityConditionClass = _BeaconIdentityConditionClass{objc.GetClass("CLBeaconIdentityCondition")}
	})
	return BeaconIdentityConditionClass
}

type _BeaconIdentityConditionClass struct {
	class objc.Class
}

// An interface definition for the [BeaconIdentityCondition] class.
type IBeaconIdentityCondition interface {
	ICondition
}

// A condition that describes the identity characteristics of a beacon.
//
// Core Location defines a beacon identity by UUID, and major and minor values. You need to specify the UUID. If you only specify a UUID, the framework treats the major and minor values as wildcards and any beacons with the same UUID satisfy the condition. Similarly, if you specify only a UUID and a major value, the framework treats the minor value as a wildcard and any beacons with the same UUID and major value satisfy the condition.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeaconIdentityCondition
type BeaconIdentityCondition struct {
	Condition
}

// BeaconIdentityConditionFrom constructs a [BeaconIdentityCondition] from an unsafe.Pointer.
//
// A condition that describes the identity characteristics of a beacon.
func BeaconIdentityConditionFrom(ptr unsafe.Pointer) BeaconIdentityCondition {
	return BeaconIdentityCondition{
		Condition: ConditionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BeaconIdentityConditionClass) Alloc() BeaconIdentityCondition {
	rv := objc.Send[BeaconIdentityCondition](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BeaconIdentityConditionClass) New() BeaconIdentityCondition {
	rv := objc.Send[BeaconIdentityCondition](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BeaconIdentityCondition) Init() BeaconIdentityCondition {
	rv := objc.Send[BeaconIdentityCondition](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BeaconIdentityCondition) Autorelease() BeaconIdentityCondition {
	rv := objc.Send[BeaconIdentityCondition](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBeaconIdentityCondition creates a new BeaconIdentityCondition instance.
func NewBeaconIdentityCondition() BeaconIdentityCondition {
	return getBeaconIdentityConditionClass().New()
}


// Creates a new beacon identity condition with the identifier you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeaconIdentityCondition/initWithUUID:
func NewBeaconIdentityConditionWithUUID(uuid unsafe.Pointer) BeaconIdentityCondition {
	instance := getBeaconIdentityConditionClass().Alloc()
	rv := objc.Send[BeaconIdentityCondition](instance.ID, objc.Sel("initWithUUID:"), uuid)
	rv.Autorelease()
	return rv
}

// Creates a new beacon identity condition with the identifier and major value you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeaconIdentityCondition/initWithUUID:major:
func NewBeaconIdentityConditionWithUUIDMajor(uuid unsafe.Pointer, major unsafe.Pointer) BeaconIdentityCondition {
	instance := getBeaconIdentityConditionClass().Alloc()
	rv := objc.Send[BeaconIdentityCondition](instance.ID, objc.Sel("initWithUUID:major:"), uuid, major)
	rv.Autorelease()
	return rv
}

// Creates a new beacon identity condition with the identifier, and major and minor values you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeaconIdentityCondition/initWithUUID:major:minor:
func NewBeaconIdentityConditionWithUUIDMajorMinor(uuid unsafe.Pointer, major unsafe.Pointer, minor unsafe.Pointer) BeaconIdentityCondition {
	instance := getBeaconIdentityConditionClass().Alloc()
	rv := objc.Send[BeaconIdentityCondition](instance.ID, objc.Sel("initWithUUID:major:minor:"), uuid, major, minor)
	rv.Autorelease()
	return rv
}


// A universally unique identifier that represent the beacon’s identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeaconIdentityCondition/UUID
func (b_ BeaconIdentityCondition) UUID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("UUID"))
	return rv
}

// The most significant value associated with the beacon.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeaconIdentityCondition/major
func (b_ BeaconIdentityCondition) Major() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("major"))
	return rv
}

// The least significant value associated with the beacon.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBeaconIdentityCondition/minor
func (b_ BeaconIdentityCondition) Minor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("minor"))
	return rv
}



// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKWheelchairUseObject] class.
var (
	HKWheelchairUseObjectClass     _HKWheelchairUseObjectClass
	HKWheelchairUseObjectClassOnce sync.Once
)

func getHKWheelchairUseObjectClass() _HKWheelchairUseObjectClass {
	HKWheelchairUseObjectClassOnce.Do(func() {
		HKWheelchairUseObjectClass = _HKWheelchairUseObjectClass{objc.GetClass("HKWheelchairUseObject")}
	})
	return HKWheelchairUseObjectClass
}

type _HKWheelchairUseObjectClass struct {
	class objc.Class
}

// An interface definition for the [HKWheelchairUseObject] class.
type IHKWheelchairUseObject interface {
	objectivec.IObject
}

// This class acts as a wrapper for the wheelchair use enumeration.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKWheelchairUseObject
type HKWheelchairUseObject struct {
	objectivec.Object
}

// HKWheelchairUseObjectFrom constructs a [HKWheelchairUseObject] from an unsafe.Pointer.
//
// This class acts as a wrapper for the wheelchair use enumeration.
func HKWheelchairUseObjectFrom(ptr unsafe.Pointer) HKWheelchairUseObject {
	return HKWheelchairUseObject{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKWheelchairUseObjectClass) Alloc() HKWheelchairUseObject {
	rv := objc.Send[HKWheelchairUseObject](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKWheelchairUseObjectClass) New() HKWheelchairUseObject {
	rv := objc.Send[HKWheelchairUseObject](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKWheelchairUseObject) Init() HKWheelchairUseObject {
	rv := objc.Send[HKWheelchairUseObject](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKWheelchairUseObject) Autorelease() HKWheelchairUseObject {
	rv := objc.Send[HKWheelchairUseObject](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKWheelchairUseObject creates a new HKWheelchairUseObject instance.
func NewHKWheelchairUseObject() HKWheelchairUseObject {
	return getHKWheelchairUseObjectClass().New()
}


// A value indicating the user’s wheelchair use.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkwheelchairuseobject/wheelchairuse
func (h_ HKWheelchairUseObject) WheelchairUse() HKWheelchairUse {
	rv := objc.Send[HKWheelchairUse](h_.ID, objc.Sel("wheelchairUse"))
	return rv
}


// SetWheelchairUse sets the value of the wheelchairUse property.
// A value indicating the user’s wheelchair use.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkwheelchairuseobject/wheelchairuse
func (h_ HKWheelchairUseObject) SetWheelchairUse(value IHKWheelchairUse) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setWheelchairUse:"), value)
}




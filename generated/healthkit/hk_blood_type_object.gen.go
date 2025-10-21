// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [HKBloodTypeObject] class.
var (
	HKBloodTypeObjectClass     _HKBloodTypeObjectClass
	HKBloodTypeObjectClassOnce sync.Once
)

func getHKBloodTypeObjectClass() _HKBloodTypeObjectClass {
	HKBloodTypeObjectClassOnce.Do(func() {
		HKBloodTypeObjectClass = _HKBloodTypeObjectClass{objc.GetClass("HKBloodTypeObject")}
	})
	return HKBloodTypeObjectClass
}

type _HKBloodTypeObjectClass struct {
	class objc.Class
}

// An interface definition for the [HKBloodTypeObject] class.
type IHKBloodTypeObject interface {
	objectivec.IObject
}

// This class acts as a wrapper for the enumeration.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBloodTypeObject
type HKBloodTypeObject struct {
	objectivec.Object
}

// HKBloodTypeObjectFrom constructs a [HKBloodTypeObject] from an unsafe.Pointer.
//
// This class acts as a wrapper for the enumeration.
func HKBloodTypeObjectFrom(ptr unsafe.Pointer) HKBloodTypeObject {
	return HKBloodTypeObject{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKBloodTypeObjectClass) Alloc() HKBloodTypeObject {
	rv := objc.Send[HKBloodTypeObject](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKBloodTypeObjectClass) New() HKBloodTypeObject {
	rv := objc.Send[HKBloodTypeObject](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKBloodTypeObject) Init() HKBloodTypeObject {
	rv := objc.Send[HKBloodTypeObject](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKBloodTypeObject) Autorelease() HKBloodTypeObject {
	rv := objc.Send[HKBloodTypeObject](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKBloodTypeObject creates a new HKBloodTypeObject instance.
func NewHKBloodTypeObject() HKBloodTypeObject {
	return getHKBloodTypeObjectClass().New()
}





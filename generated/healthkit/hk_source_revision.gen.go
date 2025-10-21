// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKSourceRevision] class.
var (
	HKSourceRevisionClass     _HKSourceRevisionClass
	HKSourceRevisionClassOnce sync.Once
)

func getHKSourceRevisionClass() _HKSourceRevisionClass {
	HKSourceRevisionClassOnce.Do(func() {
		HKSourceRevisionClass = _HKSourceRevisionClass{objc.GetClass("HKSourceRevision")}
	})
	return HKSourceRevisionClass
}

type _HKSourceRevisionClass struct {
	class objc.Class
}

// An interface definition for the [HKSourceRevision] class.
type IHKSourceRevision interface {
	objectivec.IObject
}

// An object indicating the source of a HealthKit sample.
//
// The class acts as a wrapper for the class, adding information about the source’s version, operating system, and product type. Source revision objects are immutable: you set the source revision’s properties when you create the object, and they cannot change. When an instance is created, its property is set to . When the object is saved to the HealthKit store, HealthKit assigns a new source revision to the object’s property. The source revision can be accessed only on objects retrieved from the HealthKit store.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKSourceRevision
type HKSourceRevision struct {
	objectivec.Object
}

// HKSourceRevisionFrom constructs a [HKSourceRevision] from an unsafe.Pointer.
//
// An object indicating the source of a HealthKit sample.
func HKSourceRevisionFrom(ptr unsafe.Pointer) HKSourceRevision {
	return HKSourceRevision{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKSourceRevisionClass) Alloc() HKSourceRevision {
	rv := objc.Send[HKSourceRevision](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKSourceRevisionClass) New() HKSourceRevision {
	rv := objc.Send[HKSourceRevision](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKSourceRevision) Init() HKSourceRevision {
	rv := objc.Send[HKSourceRevision](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKSourceRevision) Autorelease() HKSourceRevision {
	rv := objc.Send[HKSourceRevision](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKSourceRevision creates a new HKSourceRevision instance.
func NewHKSourceRevision() HKSourceRevision {
	return getHKSourceRevisionClass().New()
}





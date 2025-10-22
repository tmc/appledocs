// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKBiologicalSexObject] class.
var (
	HKBiologicalSexObjectClass     _HKBiologicalSexObjectClass
	HKBiologicalSexObjectClassOnce sync.Once
)

func getHKBiologicalSexObjectClass() _HKBiologicalSexObjectClass {
	HKBiologicalSexObjectClassOnce.Do(func() {
		HKBiologicalSexObjectClass = _HKBiologicalSexObjectClass{objc.GetClass("HKBiologicalSexObject")}
	})
	return HKBiologicalSexObjectClass
}

type _HKBiologicalSexObjectClass struct {
	class objc.Class
}

// An interface definition for the [HKBiologicalSexObject] class.
type IHKBiologicalSexObject interface {
	objectivec.IObject
	BiologicalSex() HKBiologicalSex
}

// This class acts as a wrapper for the enumeration.


// This class acts as a wrapper for the enumeration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBiologicalSexObject

type HKBiologicalSexObject struct {
	objectivec.Object
}

// HKBiologicalSexObjectFrom constructs a [HKBiologicalSexObject] from an unsafe.Pointer.
//
// This class acts as a wrapper for the enumeration.
func HKBiologicalSexObjectFrom(ptr unsafe.Pointer) HKBiologicalSexObject {
	return HKBiologicalSexObject{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKBiologicalSexObjectClass) Alloc() HKBiologicalSexObject {
	rv := objc.Send[HKBiologicalSexObject](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKBiologicalSexObjectClass) New() HKBiologicalSexObject {
	rv := objc.Send[HKBiologicalSexObject](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKBiologicalSexObject) Init() HKBiologicalSexObject {
	rv := objc.Send[HKBiologicalSexObject](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKBiologicalSexObject) Autorelease() HKBiologicalSexObject {
	rv := objc.Send[HKBiologicalSexObject](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKBiologicalSexObject creates a new HKBiologicalSexObject instance.
func NewHKBiologicalSexObject() HKBiologicalSexObject {
	return getHKBiologicalSexObjectClass().New()
}



// The biological sex.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKBiologicalSexObject/biologicalSex

func (h_ HKBiologicalSexObject) BiologicalSex() HKBiologicalSex {
	rv := objc.Send[HKBiologicalSex](h_.ID, objc.Sel("biologicalSex"))
	return rv
}




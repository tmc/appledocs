// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKLensSpecification] class.
var (
	HKLensSpecificationClass     _HKLensSpecificationClass
	HKLensSpecificationClassOnce sync.Once
)

func getHKLensSpecificationClass() _HKLensSpecificationClass {
	HKLensSpecificationClassOnce.Do(func() {
		HKLensSpecificationClass = _HKLensSpecificationClass{objc.GetClass("HKLensSpecification")}
	})
	return HKLensSpecificationClass
}

type _HKLensSpecificationClass struct {
	class objc.Class
}

// An interface definition for the [HKLensSpecification] class.
type IHKLensSpecification interface {
	objectivec.IObject
}

// An abstract superclass for lens specifications.
//
// Don’t instantiate this class directly. Instead, use one of its concrete subclasses: or .
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKLensSpecification
type HKLensSpecification struct {
	objectivec.Object
}

// HKLensSpecificationFrom constructs a [HKLensSpecification] from an unsafe.Pointer.
//
// An abstract superclass for lens specifications.
func HKLensSpecificationFrom(ptr unsafe.Pointer) HKLensSpecification {
	return HKLensSpecification{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKLensSpecificationClass) Alloc() HKLensSpecification {
	rv := objc.Send[HKLensSpecification](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKLensSpecificationClass) New() HKLensSpecification {
	rv := objc.Send[HKLensSpecification](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKLensSpecification) Init() HKLensSpecification {
	rv := objc.Send[HKLensSpecification](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKLensSpecification) Autorelease() HKLensSpecification {
	rv := objc.Send[HKLensSpecification](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKLensSpecification creates a new HKLensSpecification instance.
func NewHKLensSpecification() HKLensSpecification {
	return getHKLensSpecificationClass().New()
}





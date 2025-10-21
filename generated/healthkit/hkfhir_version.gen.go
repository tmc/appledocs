// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKFHIRVersion] class.
var (
	HKFHIRVersionClass     _HKFHIRVersionClass
	HKFHIRVersionClassOnce sync.Once
)

func getHKFHIRVersionClass() _HKFHIRVersionClass {
	HKFHIRVersionClassOnce.Do(func() {
		HKFHIRVersionClass = _HKFHIRVersionClass{objc.GetClass("HKFHIRVersion")}
	})
	return HKFHIRVersionClass
}

type _HKFHIRVersionClass struct {
	class objc.Class
}

// An interface definition for the [HKFHIRVersion] class.
type IHKFHIRVersion interface {
	objectivec.IObject
}

// The FHIR version.
//
// Use an instance to represent the version of the Fast Healthcare Interoperability Resources (FHIR) standard used to create a sample.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKFHIRVersion
type HKFHIRVersion struct {
	objectivec.Object
}

// HKFHIRVersionFrom constructs a [HKFHIRVersion] from an unsafe.Pointer.
//
// The FHIR version.
func HKFHIRVersionFrom(ptr unsafe.Pointer) HKFHIRVersion {
	return HKFHIRVersion{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKFHIRVersionClass) Alloc() HKFHIRVersion {
	rv := objc.Send[HKFHIRVersion](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKFHIRVersionClass) New() HKFHIRVersion {
	rv := objc.Send[HKFHIRVersion](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKFHIRVersion) Init() HKFHIRVersion {
	rv := objc.Send[HKFHIRVersion](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKFHIRVersion) Autorelease() HKFHIRVersion {
	rv := objc.Send[HKFHIRVersion](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKFHIRVersion creates a new HKFHIRVersion instance.
func NewHKFHIRVersion() HKFHIRVersion {
	return getHKFHIRVersionClass().New()
}





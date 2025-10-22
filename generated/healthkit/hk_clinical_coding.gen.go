// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKClinicalCoding] class.
var (
	HKClinicalCodingClass     _HKClinicalCodingClass
	HKClinicalCodingClassOnce sync.Once
)

func getHKClinicalCodingClass() _HKClinicalCodingClass {
	HKClinicalCodingClassOnce.Do(func() {
		HKClinicalCodingClass = _HKClinicalCodingClass{objc.GetClass("HKClinicalCoding")}
	})
	return HKClinicalCodingClass
}

type _HKClinicalCodingClass struct {
	class objc.Class
}

// An interface definition for the [HKClinicalCoding] class.
type IHKClinicalCoding interface {
	objectivec.IObject
	Code() string
	System() string
	Version() string
}

// A clinical coding that represents a medical concept using a standardized coding system.
//
// A clinical coding pairs a , an optional , and a which identify a medical concept. This model is closely related to the .
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKClinicalCoding
type HKClinicalCoding struct {
	objectivec.Object
}

// HKClinicalCodingFrom constructs a [HKClinicalCoding] from an unsafe.Pointer.
//
// A clinical coding that represents a medical concept using a standardized coding system.
func HKClinicalCodingFrom(ptr unsafe.Pointer) HKClinicalCoding {
	return HKClinicalCoding{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKClinicalCodingClass) Alloc() HKClinicalCoding {
	rv := objc.Send[HKClinicalCoding](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKClinicalCodingClass) New() HKClinicalCoding {
	rv := objc.Send[HKClinicalCoding](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKClinicalCoding) Init() HKClinicalCoding {
	rv := objc.Send[HKClinicalCoding](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKClinicalCoding) Autorelease() HKClinicalCoding {
	rv := objc.Send[HKClinicalCoding](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKClinicalCoding creates a new HKClinicalCoding instance.
func NewHKClinicalCoding() HKClinicalCoding {
	return getHKClinicalCodingClass().New()
}




// Creates a clinical coding with the specified system, version, and code.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKClinicalCoding/init(system:version:code:)
func NewHKClinicalCodingWithSystemVersionCode(system string, version string, code string) HKClinicalCoding {
	instance := getHKClinicalCodingClass().Alloc()
	rv := objc.Send[HKClinicalCoding](instance.ID, objc.Sel("initWithSystem:version:code:"), objc.String(system), objc.String(version), objc.String(code))
	rv.Autorelease()
	return rv
}


// The clinical code that represents a medical concept inside the coding system.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKClinicalCoding/code
func (h_ HKClinicalCoding) Code() string {
	rv := objc.Send[string](h_.ID, objc.Sel("code"))
	return rv
}

// The string that identifies the coding system that defines this clinical code.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKClinicalCoding/system
func (h_ HKClinicalCoding) System() string {
	rv := objc.Send[string](h_.ID, objc.Sel("system"))
	return rv
}

// The version of the coding system.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKClinicalCoding/version
func (h_ HKClinicalCoding) Version() string {
	rv := objc.Send[string](h_.ID, objc.Sel("version"))
	return rv
}



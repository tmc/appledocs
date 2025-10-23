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
	// properties:
	Code() string /* primitive/slice/pointer. */
	SetCode(value string /* primitive/slice/pointer. */)
	System() string /* primitive/slice/pointer. */
	SetSystem(value string /* primitive/slice/pointer. */)
	Version() string /* primitive/slice/pointer. */
	SetVersion(value string /* primitive/slice/pointer. */)
	// methods:
}

// A clinical coding that represents a medical concept using a standardized coding system.
//
// A clinical coding pairs a , an optional , and a which identify a medical concept. This model is closely related to the .


// A clinical coding that represents a medical concept using a standardized coding system.
//
// [Full Topic]
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



// The clinical code that represents a medical concept inside the coding system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkclinicalcoding/code
func (h_ HKClinicalCoding) Code() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](h_.ID, objc.Sel("code"))
	return rv
}


// The clinical code that represents a medical concept inside the coding system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkclinicalcoding/code
func (h_ HKClinicalCoding) SetCode(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setCode:"), objc.String(value))
}


// The string that identifies the coding system that defines this clinical code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkclinicalcoding/system
func (h_ HKClinicalCoding) System() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](h_.ID, objc.Sel("system"))
	return rv
}


// The string that identifies the coding system that defines this clinical code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkclinicalcoding/system
func (h_ HKClinicalCoding) SetSystem(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSystem:"), objc.String(value))
}


// The version of the coding system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkclinicalcoding/version
func (h_ HKClinicalCoding) Version() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](h_.ID, objc.Sel("version"))
	return rv
}


// The version of the coding system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkclinicalcoding/version
func (h_ HKClinicalCoding) SetVersion(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setVersion:"), objc.String(value))
}




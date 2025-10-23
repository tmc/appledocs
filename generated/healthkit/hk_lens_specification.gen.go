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
	AddPower() HKQuantity
	SetAddPower(value IHKQuantity)
	Axis() HKQuantity
	SetAxis(value IHKQuantity)
	Cylinder() HKQuantity
	SetCylinder(value IHKQuantity)
	Sphere() HKQuantity
	SetSphere(value IHKQuantity)
}

// An abstract superclass for lens specifications.
//
// Don’t instantiate this class directly. Instead, use one of its concrete subclasses: or .


// An abstract superclass for lens specifications.
//
// [Full Topic]
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



// The correction for nearsightedness.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hklensspecification/addpower
func (h_ HKLensSpecification) AddPower() HKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("addPower"))
	return rv
}


// The correction for nearsightedness.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hklensspecification/addpower
func (h_ HKLensSpecification) SetAddPower(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAddPower:"), value)
}


// Part of the correction for astigmatism that measures the orientation fo the correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hklensspecification/axis
func (h_ HKLensSpecification) Axis() HKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("axis"))
	return rv
}


// Part of the correction for astigmatism that measures the orientation fo the correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hklensspecification/axis
func (h_ HKLensSpecification) SetAxis(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAxis:"), value)
}


// Part of the correction for astigmatism that measures the strength of the correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hklensspecification/cylinder
func (h_ HKLensSpecification) Cylinder() HKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("cylinder"))
	return rv
}


// Part of the correction for astigmatism that measures the strength of the correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hklensspecification/cylinder
func (h_ HKLensSpecification) SetCylinder(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setCylinder:"), value)
}


// The correction for farsightedness.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hklensspecification/sphere
func (h_ HKLensSpecification) Sphere() HKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("sphere"))
	return rv
}


// The correction for farsightedness.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hklensspecification/sphere
func (h_ HKLensSpecification) SetSphere(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSphere:"), value)
}




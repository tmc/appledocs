// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKContactsLensSpecification] class.
var (
	HKContactsLensSpecificationClass     _HKContactsLensSpecificationClass
	HKContactsLensSpecificationClassOnce sync.Once
)

func getHKContactsLensSpecificationClass() _HKContactsLensSpecificationClass {
	HKContactsLensSpecificationClassOnce.Do(func() {
		HKContactsLensSpecificationClass = _HKContactsLensSpecificationClass{objc.GetClass("HKContactsLensSpecification")}
	})
	return HKContactsLensSpecificationClass
}

type _HKContactsLensSpecificationClass struct {
	class objc.Class
}

// An interface definition for the [HKContactsLensSpecification] class.
type IHKContactsLensSpecification interface {
	IHKLensSpecification
	BaseCurve() HKQuantity
	SetBaseCurve(value IHKQuantity)
	Diameter() HKQuantity
	SetDiameter(value IHKQuantity)
}

// An object that contains the contacts prescription data for one eye.
//
// To create a sample that stores a contacts prescription, start by defining a specification for each eye. Each lens specification object requires a parameter. This measures the lens’s strength for correcting either nearsightedness or farsightedness (measured in units). Next, create values for any of the prescription’s optional parameters. For example, if the prescription corrects for astigmatism, create the and values. The value uses units, while the uses . To add a multifocal correction for reading, create an value using units. To add fitting information for the contact lens, create and values. Both of these values use millimeters. Then you can create the lens specification. After you create your lens specifications, you can create an sample. Then save the sample to the HealthKit store. Finally, add an image or PDF of the prescription to the sample as an attachment.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKContactsLensSpecification
type HKContactsLensSpecification struct {
	HKLensSpecification
}

// HKContactsLensSpecificationFrom constructs a [HKContactsLensSpecification] from an unsafe.Pointer.
//
// An object that contains the contacts prescription data for one eye.
func HKContactsLensSpecificationFrom(ptr unsafe.Pointer) HKContactsLensSpecification {
	return HKContactsLensSpecification{
		HKLensSpecification: HKLensSpecificationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKContactsLensSpecificationClass) Alloc() HKContactsLensSpecification {
	rv := objc.Send[HKContactsLensSpecification](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKContactsLensSpecificationClass) New() HKContactsLensSpecification {
	rv := objc.Send[HKContactsLensSpecification](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKContactsLensSpecification) Init() HKContactsLensSpecification {
	rv := objc.Send[HKContactsLensSpecification](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKContactsLensSpecification) Autorelease() HKContactsLensSpecification {
	rv := objc.Send[HKContactsLensSpecification](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKContactsLensSpecification creates a new HKContactsLensSpecification instance.
func NewHKContactsLensSpecification() HKContactsLensSpecification {
	return getHKContactsLensSpecificationClass().New()
}


// Part of the contact’s fit, it measures the curve of the back side of the contact, measured in mm.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcontactslensspecification/basecurve
func (h_ HKContactsLensSpecification) BaseCurve() HKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("baseCurve"))
	return rv
}


// SetBaseCurve sets the value of the baseCurve property.
// Part of the contact’s fit, it measures the curve of the back side of the contact, measured in mm.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcontactslensspecification/basecurve
func (h_ HKContactsLensSpecification) SetBaseCurve(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setBaseCurve:"), value)
}

// Part of the contact’s fit, it measures the diameter of the lens, measured in mm.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcontactslensspecification/diameter
func (h_ HKContactsLensSpecification) Diameter() HKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("diameter"))
	return rv
}


// SetDiameter sets the value of the diameter property.
// Part of the contact’s fit, it measures the diameter of the lens, measured in mm.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcontactslensspecification/diameter
func (h_ HKContactsLensSpecification) SetDiameter(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setDiameter:"), value)
}




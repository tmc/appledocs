// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKContactsPrescription] class.
var (
	HKContactsPrescriptionClass     _HKContactsPrescriptionClass
	HKContactsPrescriptionClassOnce sync.Once
)

func getHKContactsPrescriptionClass() _HKContactsPrescriptionClass {
	HKContactsPrescriptionClassOnce.Do(func() {
		HKContactsPrescriptionClass = _HKContactsPrescriptionClass{objc.GetClass("HKContactsPrescription")}
	})
	return HKContactsPrescriptionClass
}

type _HKContactsPrescriptionClass struct {
	class objc.Class
}

// An interface definition for the [HKContactsPrescription] class.
type IHKContactsPrescription interface {
	IHKVisionPrescription
	// properties:
	Brand() string /* primitive/slice/pointer. */
	SetBrand(value string /* primitive/slice/pointer. */)
	LeftEye() IHKContactsLensSpecification
	SetLeftEye(value IHKContactsLensSpecification)
	RightEye() IHKContactsLensSpecification
	SetRightEye(value IHKContactsLensSpecification)
	// methods:
}

// A sample that store a prescription for contacts.
//
// To create a sample that stores a contacts prescription, start by defining a specification for each eye. Each lens specification object requires a parameter. This measures the lens’s strength for correcting either nearsightedness or farsightedness (measured in units). Next, create values for any of the prescription’s optional parameters. For example, if the prescription corrects for astigmatism, create the and values. The value uses units, while the uses . To add a multifocal correction for reading, create an value using units. To add fitting information for the contact lens, create and values. Both of these values use millimeters. Then you can create the lens specification. After you create your lens specifications, you can create an sample. Then save the sample to the HealthKit store. Finally, add an image or PDF of the prescription to the sample as an attachment.


// A sample that store a prescription for contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKContactsPrescription
type HKContactsPrescription struct {
	HKVisionPrescription
}

// HKContactsPrescriptionFrom constructs a [HKContactsPrescription] from an unsafe.Pointer.
//
// A sample that store a prescription for contacts.
func HKContactsPrescriptionFrom(ptr unsafe.Pointer) HKContactsPrescription {
	return HKContactsPrescription{
		HKVisionPrescription: HKVisionPrescriptionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKContactsPrescriptionClass) Alloc() HKContactsPrescription {
	rv := objc.Send[HKContactsPrescription](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKContactsPrescriptionClass) New() HKContactsPrescription {
	rv := objc.Send[HKContactsPrescription](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKContactsPrescription) Init() HKContactsPrescription {
	rv := objc.Send[HKContactsPrescription](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKContactsPrescription) Autorelease() HKContactsPrescription {
	rv := objc.Send[HKContactsPrescription](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKContactsPrescription creates a new HKContactsPrescription instance.
func NewHKContactsPrescription() HKContactsPrescription {
	return getHKContactsPrescriptionClass().New()
}



// The name of the prescribed brand, based on the contact lens fitting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcontactsprescription/brand
func (h_ HKContactsPrescription) Brand() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](h_.ID, objc.Sel("brand"))
	return rv
}


// The name of the prescribed brand, based on the contact lens fitting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcontactsprescription/brand
func (h_ HKContactsPrescription) SetBrand(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setBrand:"), objc.String(value))
}


// The lens specification for the left eye.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcontactsprescription/lefteye
func (h_ HKContactsPrescription) LeftEye() IHKContactsLensSpecification {
	rv := objc.Send[HKContactsLensSpecification](h_.ID, objc.Sel("leftEye"))
	return rv
}


// The lens specification for the left eye.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcontactsprescription/lefteye
func (h_ HKContactsPrescription) SetLeftEye(value IHKContactsLensSpecification) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setLeftEye:"), value)
}


// The lens specification for the right eye.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcontactsprescription/righteye
func (h_ HKContactsPrescription) RightEye() IHKContactsLensSpecification {
	rv := objc.Send[HKContactsLensSpecification](h_.ID, objc.Sel("rightEye"))
	return rv
}


// The lens specification for the right eye.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkcontactsprescription/righteye
func (h_ HKContactsPrescription) SetRightEye(value IHKContactsLensSpecification) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setRightEye:"), value)
}




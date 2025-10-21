// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKVisionPrescription] class.
var (
	HKVisionPrescriptionClass     _HKVisionPrescriptionClass
	HKVisionPrescriptionClassOnce sync.Once
)

func getHKVisionPrescriptionClass() _HKVisionPrescriptionClass {
	HKVisionPrescriptionClassOnce.Do(func() {
		HKVisionPrescriptionClass = _HKVisionPrescriptionClass{objc.GetClass("HKVisionPrescription")}
	})
	return HKVisionPrescriptionClass
}

type _HKVisionPrescriptionClass struct {
	class objc.Class
}

// An interface definition for the [HKVisionPrescription] class.
type IHKVisionPrescription interface {
	IHKSample
}

// A sample that stores a vision prescription.
//
// Use this class to create an image-only prescription. Here, you attach the prescription as an image or PDF to a simple sample. The sample contains only basic information about the prescription, such as the issue and expiration dates. To see the prescription data, people must view the attached image or PDF. To create an image-only prescription, start by creating an sample object. Next, save the sample to the HealthKit store. Then, you can attach the image or PDF to the sample. For more information about adding images or pdfs as attachments, see . To create a vision prescription sample that contains the full data for the prescription, use or instead.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVisionPrescription
type HKVisionPrescription struct {
	HKSample
}

// HKVisionPrescriptionFrom constructs a [HKVisionPrescription] from an unsafe.Pointer.
//
// A sample that stores a vision prescription.
func HKVisionPrescriptionFrom(ptr unsafe.Pointer) HKVisionPrescription {
	return HKVisionPrescription{
		HKSample: HKSampleFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKVisionPrescriptionClass) Alloc() HKVisionPrescription {
	rv := objc.Send[HKVisionPrescription](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKVisionPrescriptionClass) New() HKVisionPrescription {
	rv := objc.Send[HKVisionPrescription](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKVisionPrescription) Init() HKVisionPrescription {
	rv := objc.Send[HKVisionPrescription](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKVisionPrescription) Autorelease() HKVisionPrescription {
	rv := objc.Send[HKVisionPrescription](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKVisionPrescription creates a new HKVisionPrescription instance.
func NewHKVisionPrescription() HKVisionPrescription {
	return getHKVisionPrescriptionClass().New()
}


// The date when the prescription expires.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkvisionprescription/expirationdate
func (h_ HKVisionPrescription) ExpirationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("expirationDate"))
	return rv
}


// SetExpirationDate sets the value of the expirationDate property.
// The date when the prescription expires.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkvisionprescription/expirationdate
func (h_ HKVisionPrescription) SetExpirationDate(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setExpirationDate:"), value)
}

// The date when the doctor issued the prescription.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkvisionprescription/dateissued
func (h_ HKVisionPrescription) DateIssued() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("dateIssued"))
	return rv
}


// SetDateIssued sets the value of the dateIssued property.
// The date when the doctor issued the prescription.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkvisionprescription/dateissued
func (h_ HKVisionPrescription) SetDateIssued(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setDateIssued:"), value)
}

// The type of vision prescription.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkvisionprescription/prescriptiontype
func (h_ HKVisionPrescription) PrescriptionType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("prescriptionType"))
	return rv
}


// SetPrescriptionType sets the value of the prescriptionType property.
// The type of vision prescription.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkvisionprescription/prescriptiontype
func (h_ HKVisionPrescription) SetPrescriptionType(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setPrescriptionType:"), value)
}




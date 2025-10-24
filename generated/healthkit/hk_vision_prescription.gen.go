// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKVisionPrescription */


/* debug [class_header]: Header for HKVisionPrescription */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKVisionPrescription */
// An interface definition for the [HKVisionPrescription] class.
type IHKVisionPrescription interface {
	IHKSample
	
/* debug [class_interface_properties]: Properties for HKVisionPrescription */
	// properties:
	DateIssued() objc.IObject /* cross-framework: NSDate */
	ExpirationDate() objc.IObject /* cross-framework: NSDate */
	PrescriptionType() HKVisionPrescriptionType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKVisionPrescription */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKVisionPrescription */
// Alloc allocates a new instance without initialization.
func (hc _HKVisionPrescriptionClass) Alloc() HKVisionPrescription {
	rv := objc.Send[HKVisionPrescription](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKVisionPrescription */
// A sample that stores a vision prescription.
//
// Use this class to create an image-only prescription. Here, you attach the prescription as an image or PDF to a simple sample. The sample contains only basic information about the prescription, such as the issue and expiration dates. To see the prescription data, people must view the attached image or PDF. To create an image-only prescription, start by creating an sample object. Next, save the sample to the HealthKit store. Then, you can attach the image or PDF to the sample. For more information about adding images or pdfs as attachments, see . To create a vision prescription sample that contains the full data for the prescription, use or instead.


// A sample that stores a vision prescription.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKVisionPrescription */

// Creates a new vision prescription sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVisionPrescription/init(type:dateIssued:expirationDate:device:metadata:)
func NewHKVisionPrescriptionWithTypeDateIssuedExpirationDateDeviceMetadata(type_ HKVisionPrescriptionType, dateIssued objc.IObject /* cross-framework: NSDate */, expirationDate objc.IObject /* cross-framework: NSDate */, device IHKDevice, metadata foundation.IDictionary) HKVisionPrescription {
	rv := objc.Send[HKVisionPrescription](objc.ID(getHKVisionPrescriptionClass().class), objc.Sel("prescriptionWithType:dateIssued:expirationDate:device:metadata:"), type_, dateIssued, expirationDate, device, metadata)
	return rv
}/* debug [class_init_methods/constructor]: NewHKVisionPrescriptionWithTypeDateIssuedExpirationDateDeviceMetadata */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKVisionPrescription */

// Creates a new vision prescription sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVisionPrescription/init(type:dateIssued:expirationDate:device:metadata:)
func (hc _HKVisionPrescriptionClass) PrescriptionWithTypeDateIssuedExpirationDateDeviceMetadata(type_ HKVisionPrescriptionType, dateIssued objc.IObject /* cross-framework: NSDate */, expirationDate objc.IObject /* cross-framework: NSDate */, device IHKDevice, metadata foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("prescriptionWithType:dateIssued:expirationDate:device:metadata:"), type_, dateIssued, expirationDate, device, metadata)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PrescriptionWithTypeDateIssuedExpirationDateDeviceMetadata) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKVisionPrescription */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKVisionPrescription */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKVisionPrescription */

// The date when the doctor issued the prescription.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVisionPrescription/dateIssued
func (h_ HKVisionPrescription) DateIssued() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](h_.ID, objc.Sel("dateIssued"))
	return rv
}/* debug [instance_properties/getter]: dateIssued */


// The date when the prescription expires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVisionPrescription/expirationDate
func (h_ HKVisionPrescription) ExpirationDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](h_.ID, objc.Sel("expirationDate"))
	return rv
}/* debug [instance_properties/getter]: expirationDate */


// The type of vision prescription.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVisionPrescription/prescriptionType
func (h_ HKVisionPrescription) PrescriptionType() HKVisionPrescriptionType {
	rv := objc.Send[HKVisionPrescriptionType](h_.ID, objc.Sel("prescriptionType"))
	return rv
}/* debug [instance_properties/getter]: prescriptionType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKVisionPrescription */



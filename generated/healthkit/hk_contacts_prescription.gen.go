// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKContactsPrescription */


/* debug [class_header]: Header for HKContactsPrescription */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKContactsPrescription */
// An interface definition for the [HKContactsPrescription] class.
type IHKContactsPrescription interface {
	IHKVisionPrescription
	
/* debug [class_interface_properties]: Properties for HKContactsPrescription */
	// properties:
	Brand() objc.IObject /* cross-framework: NSString */
	LeftEye() IHKContactsLensSpecification
	RightEye() IHKContactsLensSpecification
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKContactsPrescription */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKContactsPrescription */
// Alloc allocates a new instance without initialization.
func (hc _HKContactsPrescriptionClass) Alloc() HKContactsPrescription {
	rv := objc.Send[HKContactsPrescription](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKContactsPrescription */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKContactsPrescription */

// Creates a new glasses prescription sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKContactsPrescription/init(rightEyeSpecification:leftEyeSpecification:brand:dateIssued:expirationDate:device:metadata:)
func NewHKContactsPrescriptionWithRightEyeSpecificationLeftEyeSpecificationBrandDateIssuedExpirationDateDeviceMetadata(rightEyeSpecification IHKContactsLensSpecification, leftEyeSpecification IHKContactsLensSpecification, brand objc.IObject /* cross-framework: NSString */, dateIssued objc.IObject /* cross-framework: NSDate */, expirationDate objc.IObject /* cross-framework: NSDate */, device IHKDevice, metadata foundation.IDictionary) HKContactsPrescription {
	rv := objc.Send[HKContactsPrescription](objc.ID(getHKContactsPrescriptionClass().class), objc.Sel("prescriptionWithRightEyeSpecification:leftEyeSpecification:brand:dateIssued:expirationDate:device:metadata:"), rightEyeSpecification, leftEyeSpecification, brand, dateIssued, expirationDate, device, metadata)
	return rv
}/* debug [class_init_methods/constructor]: NewHKContactsPrescriptionWithRightEyeSpecificationLeftEyeSpecificationBrandDateIssuedExpirationDateDeviceMetadata */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKContactsPrescription */

// Creates a new glasses prescription sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKContactsPrescription/init(rightEyeSpecification:leftEyeSpecification:brand:dateIssued:expirationDate:device:metadata:)
func (hc _HKContactsPrescriptionClass) PrescriptionWithRightEyeSpecificationLeftEyeSpecificationBrandDateIssuedExpirationDateDeviceMetadata(rightEyeSpecification IHKContactsLensSpecification, leftEyeSpecification IHKContactsLensSpecification, brand objc.IObject /* cross-framework: NSString */, dateIssued objc.IObject /* cross-framework: NSDate */, expirationDate objc.IObject /* cross-framework: NSDate */, device IHKDevice, metadata foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("prescriptionWithRightEyeSpecification:leftEyeSpecification:brand:dateIssued:expirationDate:device:metadata:"), rightEyeSpecification, leftEyeSpecification, brand, dateIssued, expirationDate, device, metadata)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PrescriptionWithRightEyeSpecificationLeftEyeSpecificationBrandDateIssuedExpirationDateDeviceMetadata) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKContactsPrescription */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKContactsPrescription */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKContactsPrescription */

// The name of the prescribed brand, based on the contact lens fitting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKContactsPrescription/brand
func (h_ HKContactsPrescription) Brand() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("brand"))
	return rv
}/* debug [instance_properties/getter]: brand */


// The lens specification for the left eye.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKContactsPrescription/leftEye
func (h_ HKContactsPrescription) LeftEye() IHKContactsLensSpecification {
	rv := objc.Send[HKContactsLensSpecification](h_.ID, objc.Sel("leftEye"))
	return rv
}/* debug [instance_properties/getter]: leftEye */


// The lens specification for the right eye.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKContactsPrescription/rightEye
func (h_ HKContactsPrescription) RightEye() IHKContactsLensSpecification {
	rv := objc.Send[HKContactsLensSpecification](h_.ID, objc.Sel("rightEye"))
	return rv
}/* debug [instance_properties/getter]: rightEye */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKContactsPrescription */



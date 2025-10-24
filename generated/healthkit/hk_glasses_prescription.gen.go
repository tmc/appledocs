// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKGlassesPrescription */


/* debug [class_header]: Header for HKGlassesPrescription */
// The class instance for the [HKGlassesPrescription] class.
var (
	HKGlassesPrescriptionClass     _HKGlassesPrescriptionClass
	HKGlassesPrescriptionClassOnce sync.Once
)

func getHKGlassesPrescriptionClass() _HKGlassesPrescriptionClass {
	HKGlassesPrescriptionClassOnce.Do(func() {
		HKGlassesPrescriptionClass = _HKGlassesPrescriptionClass{objc.GetClass("HKGlassesPrescription")}
	})
	return HKGlassesPrescriptionClass
}

type _HKGlassesPrescriptionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKGlassesPrescription */
// An interface definition for the [HKGlassesPrescription] class.
type IHKGlassesPrescription interface {
	IHKVisionPrescription
	
/* debug [class_interface_properties]: Properties for HKGlassesPrescription */
	// properties:
	LeftEye() IHKGlassesLensSpecification
	RightEye() IHKGlassesLensSpecification
	HKMetadataKeyGlassesPrescriptionDescription() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKGlassesPrescription */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKGlassesPrescription */
// Alloc allocates a new instance without initialization.
func (hc _HKGlassesPrescriptionClass) Alloc() HKGlassesPrescription {
	rv := objc.Send[HKGlassesPrescription](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKGlassesPrescriptionClass) New() HKGlassesPrescription {
	rv := objc.Send[HKGlassesPrescription](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKGlassesPrescription) Init() HKGlassesPrescription {
	rv := objc.Send[HKGlassesPrescription](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKGlassesPrescription) Autorelease() HKGlassesPrescription {
	rv := objc.Send[HKGlassesPrescription](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKGlassesPrescription creates a new HKGlassesPrescription instance.
func NewHKGlassesPrescription() HKGlassesPrescription {
	return getHKGlassesPrescriptionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKGlassesPrescription */
// A sample that stores a prescription for glasses.
//
// To create a sample that stores a glasses prescription, start by defining a specification for each eye. Each lens specification object requires a parameter. This measures the lens’s strength for correcting either nearsightedness or farsightedness (measured in units). Next, create values for any of the prescription’s optional parameters. For example, if the prescription corrects for astigmatism, create the and values. The value uses units, while the uses . To add a multifocal correction for reading, create an value using units. To add a correction for eye alignment, create an object. To add information about the distance between the eye and the back of the lens, or the pupil and the center of the nose, create , , and values. All of these use millimeters. Then you can create the lens specification. After you create your lens specifications, you can create an sample. Then save the sample to the HealthKit store. Finally, add an image or PDF of the prescription to the sample as an attachment.


// A sample that stores a prescription for glasses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGlassesPrescription
type HKGlassesPrescription struct {
	HKVisionPrescription
}

// HKGlassesPrescriptionFrom constructs a [HKGlassesPrescription] from an unsafe.Pointer.
//
// A sample that stores a prescription for glasses.
func HKGlassesPrescriptionFrom(ptr unsafe.Pointer) HKGlassesPrescription {
	return HKGlassesPrescription{
		HKVisionPrescription: HKVisionPrescriptionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKGlassesPrescription */

// Creates a new glasses prescription sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGlassesPrescription/init(rightEyeSpecification:leftEyeSpecification:dateIssued:expirationDate:device:metadata:)
func NewHKGlassesPrescriptionWithRightEyeSpecificationLeftEyeSpecificationDateIssuedExpirationDateDeviceMetadata(rightEyeSpecification IHKGlassesLensSpecification, leftEyeSpecification IHKGlassesLensSpecification, dateIssued objc.IObject /* cross-framework: NSDate */, expirationDate objc.IObject /* cross-framework: NSDate */, device IHKDevice, metadata foundation.IDictionary) HKGlassesPrescription {
	rv := objc.Send[HKGlassesPrescription](objc.ID(getHKGlassesPrescriptionClass().class), objc.Sel("prescriptionWithRightEyeSpecification:leftEyeSpecification:dateIssued:expirationDate:device:metadata:"), rightEyeSpecification, leftEyeSpecification, dateIssued, expirationDate, device, metadata)
	return rv
}/* debug [class_init_methods/constructor]: NewHKGlassesPrescriptionWithRightEyeSpecificationLeftEyeSpecificationDateIssuedExpirationDateDeviceMetadata */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKGlassesPrescription */

// Creates a new glasses prescription sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGlassesPrescription/init(rightEyeSpecification:leftEyeSpecification:dateIssued:expirationDate:device:metadata:)
func (hc _HKGlassesPrescriptionClass) PrescriptionWithRightEyeSpecificationLeftEyeSpecificationDateIssuedExpirationDateDeviceMetadata(rightEyeSpecification IHKGlassesLensSpecification, leftEyeSpecification IHKGlassesLensSpecification, dateIssued objc.IObject /* cross-framework: NSDate */, expirationDate objc.IObject /* cross-framework: NSDate */, device IHKDevice, metadata foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("prescriptionWithRightEyeSpecification:leftEyeSpecification:dateIssued:expirationDate:device:metadata:"), rightEyeSpecification, leftEyeSpecification, dateIssued, expirationDate, device, metadata)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PrescriptionWithRightEyeSpecificationLeftEyeSpecificationDateIssuedExpirationDateDeviceMetadata) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKGlassesPrescription */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKGlassesPrescription */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKGlassesPrescription */

// The lens specification for the left eye.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGlassesPrescription/leftEye
func (h_ HKGlassesPrescription) LeftEye() IHKGlassesLensSpecification {
	rv := objc.Send[HKGlassesLensSpecification](h_.ID, objc.Sel("leftEye"))
	return rv
}/* debug [instance_properties/getter]: leftEye */


// The lens specification for the right eye.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGlassesPrescription/rightEye
func (h_ HKGlassesPrescription) RightEye() IHKGlassesLensSpecification {
	rv := objc.Send[HKGlassesLensSpecification](h_.ID, objc.Sel("rightEye"))
	return rv
}/* debug [instance_properties/getter]: rightEye */


// A description of the glasses prescription.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkmetadatakeyglassesprescriptiondescription
func (h_ HKGlassesPrescription) HKMetadataKeyGlassesPrescriptionDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKMetadataKeyGlassesPrescriptionDescription"))
	return rv
}/* debug [instance_properties/getter]: HKMetadataKeyGlassesPrescriptionDescription */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKGlassesPrescription */



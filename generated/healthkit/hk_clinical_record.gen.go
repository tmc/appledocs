// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class HKClinicalRecord */


/* debug [class_header]: Header for HKClinicalRecord */
// The class instance for the [HKClinicalRecord] class.
var (
	HKClinicalRecordClass     _HKClinicalRecordClass
	HKClinicalRecordClassOnce sync.Once
)

func getHKClinicalRecordClass() _HKClinicalRecordClass {
	HKClinicalRecordClassOnce.Do(func() {
		HKClinicalRecordClass = _HKClinicalRecordClass{objc.GetClass("HKClinicalRecord")}
	})
	return HKClinicalRecordClass
}

type _HKClinicalRecordClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKClinicalRecord */
// An interface definition for the [HKClinicalRecord] class.
type IHKClinicalRecord interface {
	IHKSample
	
/* debug [class_interface_properties]: Properties for HKClinicalRecord */
	// properties:
	ClinicalType() IHKClinicalType
	DisplayName() objc.IObject /* cross-framework: NSString */
	FHIRResource() IHKFHIRResource
	EndDate() foundation.Date
	SetEndDate(value foundation.Date)
	StartDate() foundation.Date
	SetStartDate(value foundation.Date)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKClinicalRecord */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKClinicalRecord */
// Alloc allocates a new instance without initialization.
func (hc _HKClinicalRecordClass) Alloc() HKClinicalRecord {
	rv := objc.Send[HKClinicalRecord](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKClinicalRecordClass) New() HKClinicalRecord {
	rv := objc.Send[HKClinicalRecord](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKClinicalRecord) Init() HKClinicalRecord {
	rv := objc.Send[HKClinicalRecord](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKClinicalRecord) Autorelease() HKClinicalRecord {
	rv := objc.Send[HKClinicalRecord](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKClinicalRecord creates a new HKClinicalRecord instance.
func NewHKClinicalRecord() HKClinicalRecord {
	return getHKClinicalRecordClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKClinicalRecord */
// A sample that stores a clinical record.
//
// The clinical record stores information about a single condition, procedure, or result. While the record’s properties expose some high-level information, the property contains the underlying data from the user’s healthcare institution. Note that the record inherits the class’s and properties. However, the system does not populate these properties with information from the FHIR data; instead, the and reflect the time and date when the system downloaded the FHIR data to the device.


// A sample that stores a clinical record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKClinicalRecord
type HKClinicalRecord struct {
	HKSample
}

// HKClinicalRecordFrom constructs a [HKClinicalRecord] from an unsafe.Pointer.
//
// A sample that stores a clinical record.
func HKClinicalRecordFrom(ptr unsafe.Pointer) HKClinicalRecord {
	return HKClinicalRecord{
		HKSample: HKSampleFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKClinicalRecord *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKClinicalRecord */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKClinicalRecord */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKClinicalRecord */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKClinicalRecord */

// An identifier that indicates the type of record, such as an allergic reaction, a lab result, or a medical procedure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKClinicalRecord/clinicalType
func (h_ HKClinicalRecord) ClinicalType() IHKClinicalType {
	rv := objc.Send[HKClinicalType](h_.ID, objc.Sel("clinicalType"))
	return rv
}/* debug [instance_properties/getter]: clinicalType */


// The primary display name as shown in the Health app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKClinicalRecord/displayName
func (h_ HKClinicalRecord) DisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("displayName"))
	return rv
}/* debug [instance_properties/getter]: displayName */


// The Fast Healthcare Interoperability Resources (FHIR) data for this record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKClinicalRecord/fhirResource
func (h_ HKClinicalRecord) FHIRResource() IHKFHIRResource {
	rv := objc.Send[HKFHIRResource](h_.ID, objc.Sel("FHIRResource"))
	return rv
}/* debug [instance_properties/getter]: FHIRResource */


// The sample’s end date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksample/enddate
func (h_ HKClinicalRecord) EndDate() foundation.Date {
	rv := objc.Send[foundation.Date](h_.ID, objc.Sel("endDate"))
	return rv
}/* debug [instance_properties/getter]: endDate */


// The sample’s end date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksample/enddate
func (h_ HKClinicalRecord) SetEndDate(value foundation.Date) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setEndDate:"), value)
}/* debug [instance_properties/setter]: endDate */


// The sample’s start date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksample/startdate
func (h_ HKClinicalRecord) StartDate() foundation.Date {
	rv := objc.Send[foundation.Date](h_.ID, objc.Sel("startDate"))
	return rv
}/* debug [instance_properties/getter]: startDate */


// The sample’s start date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksample/startdate
func (h_ HKClinicalRecord) SetStartDate(value foundation.Date) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setStartDate:"), value)
}/* debug [instance_properties/setter]: startDate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKClinicalRecord */




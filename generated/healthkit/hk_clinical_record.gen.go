// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [HKClinicalRecord] class.
type IHKClinicalRecord interface {
	IHKSample
	// properties:
	ClinicalType() IHKClinicalType
	SetClinicalType(value IHKClinicalType)
	DisplayName() objc.IObject /* cross-framework: NSString */
	SetDisplayName(value objc.IObject /* cross-framework: NSString */)
	FhirResource() IHKFHIRResource
	SetFhirResource(value IHKFHIRResource)
	EndDate() objc.IObject /* cross-framework: Date */
	SetEndDate(value objc.IObject /* cross-framework: Date */)
	StartDate() objc.IObject /* cross-framework: Date */
	SetStartDate(value objc.IObject /* cross-framework: Date */)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (hc _HKClinicalRecordClass) Alloc() HKClinicalRecord {
	rv := objc.Send[HKClinicalRecord](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// An identifier that indicates the type of record, such as an allergic reaction, a lab result, or a medical procedure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkclinicalrecord/clinicaltype
func (h_ HKClinicalRecord) ClinicalType() IHKClinicalType {
	rv := objc.Send[HKClinicalType](h_.ID, objc.Sel("clinicalType"))
	return rv
}


// An identifier that indicates the type of record, such as an allergic reaction, a lab result, or a medical procedure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkclinicalrecord/clinicaltype
func (h_ HKClinicalRecord) SetClinicalType(value IHKClinicalType) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setClinicalType:"), value)
}


// The primary display name as shown in the Health app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkclinicalrecord/displayname
func (h_ HKClinicalRecord) DisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("displayName"))
	return rv
}


// The primary display name as shown in the Health app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkclinicalrecord/displayname
func (h_ HKClinicalRecord) SetDisplayName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setDisplayName:"), value)
}


// The Fast Healthcare Interoperability Resources (FHIR) data for this record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkclinicalrecord/fhirresource
func (h_ HKClinicalRecord) FhirResource() IHKFHIRResource {
	rv := objc.Send[HKFHIRResource](h_.ID, objc.Sel("fhirResource"))
	return rv
}


// The Fast Healthcare Interoperability Resources (FHIR) data for this record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkclinicalrecord/fhirresource
func (h_ HKClinicalRecord) SetFhirResource(value IHKFHIRResource) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setFhirResource:"), value)
}


// The sample’s end date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksample/enddate
func (h_ HKClinicalRecord) EndDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](h_.ID, objc.Sel("endDate"))
	return rv
}


// The sample’s end date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksample/enddate
func (h_ HKClinicalRecord) SetEndDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setEndDate:"), value)
}


// The sample’s start date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksample/startdate
func (h_ HKClinicalRecord) StartDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](h_.ID, objc.Sel("startDate"))
	return rv
}


// The sample’s start date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksample/startdate
func (h_ HKClinicalRecord) SetStartDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setStartDate:"), value)
}




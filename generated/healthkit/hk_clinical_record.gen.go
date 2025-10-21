// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
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
}

// A sample that stores a clinical record.
//
// The clinical record stores information about a single condition, procedure, or result. While the record’s properties expose some high-level information, the property contains the underlying data from the user’s healthcare institution. Note that the record inherits the class’s and properties. However, the system does not populate these properties with information from the FHIR data; instead, the and reflect the time and date when the system downloaded the FHIR data to the device.
//
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
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKClinicalRecord/clinicalType
func (h_ HKClinicalRecord) ClinicalType() HKClinicalType {
	rv := objc.Send[HKClinicalType](h_.ID, objc.Sel("clinicalType"))
	return rv
}

// The primary display name as shown in the Health app.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKClinicalRecord/displayName
func (h_ HKClinicalRecord) DisplayName() appkit.string {
	rv := objc.Send[appkit.string](h_.ID, objc.Sel("displayName"))
	return rv
}

// The Fast Healthcare Interoperability Resources (FHIR) data for this record.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKClinicalRecord/fhirResource
func (h_ HKClinicalRecord) FHIRResource() HKFHIRResource {
	rv := objc.Send[HKFHIRResource](h_.ID, objc.Sel("FHIRResource"))
	return rv
}

// The sample’s end date.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksample/enddate
func (h_ HKClinicalRecord) EndDate() foundation.Date {
	rv := objc.Send[foundation.Date](h_.ID, objc.Sel("endDate"))
	return rv
}


// SetEndDate sets the value of the endDate property.
// The sample’s end date.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksample/enddate
func (h_ HKClinicalRecord) SetEndDate(value foundation.IDate) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setEndDate:"), value)
}

// The sample’s start date.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksample/startdate
func (h_ HKClinicalRecord) StartDate() foundation.Date {
	rv := objc.Send[foundation.Date](h_.ID, objc.Sel("startDate"))
	return rv
}


// SetStartDate sets the value of the startDate property.
// The sample’s start date.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hksample/startdate
func (h_ HKClinicalRecord) SetStartDate(value foundation.IDate) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setStartDate:"), value)
}




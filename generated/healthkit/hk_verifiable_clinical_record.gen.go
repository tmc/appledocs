// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKVerifiableClinicalRecord] class.
var (
	HKVerifiableClinicalRecordClass     _HKVerifiableClinicalRecordClass
	HKVerifiableClinicalRecordClassOnce sync.Once
)

func getHKVerifiableClinicalRecordClass() _HKVerifiableClinicalRecordClass {
	HKVerifiableClinicalRecordClassOnce.Do(func() {
		HKVerifiableClinicalRecordClass = _HKVerifiableClinicalRecordClass{objc.GetClass("HKVerifiableClinicalRecord")}
	})
	return HKVerifiableClinicalRecordClass
}

type _HKVerifiableClinicalRecordClass struct {
	class objc.Class
}

// An interface definition for the [HKVerifiableClinicalRecord] class.
type IHKVerifiableClinicalRecord interface {
	IHKSample
}

// A sample that represents the contents of a SMART Health Card or EU Digital COVID Certificate.
//
// samples contain data from a SMART Health Card or EU Digital COVID Certificate. Verifiable clinical records combine information about the user’s identity with clinical data, like an immunization record or a lab test result. The organization that produced the data cryptographically signs the bundle. Apps that use verifiable clinical records can use the cryptographic signature to verify the authenticity of the contents. To verify the card: Access the card’s raw payload using the clinical record’s property. Unzip the payload and parse out the value, which contains a URL that identifies the organization that issued the card. Get the public key from the issuer. Verify the payload’s signature. For more information, see and . You can download example SMART cards for testing and development from .
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVerifiableClinicalRecord
type HKVerifiableClinicalRecord struct {
	HKSample
}

// HKVerifiableClinicalRecordFrom constructs a [HKVerifiableClinicalRecord] from an unsafe.Pointer.
//
// A sample that represents the contents of a SMART Health Card or EU Digital COVID Certificate.
func HKVerifiableClinicalRecordFrom(ptr unsafe.Pointer) HKVerifiableClinicalRecord {
	return HKVerifiableClinicalRecord{
		HKSample: HKSampleFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKVerifiableClinicalRecordClass) Alloc() HKVerifiableClinicalRecord {
	rv := objc.Send[HKVerifiableClinicalRecord](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKVerifiableClinicalRecordClass) New() HKVerifiableClinicalRecord {
	rv := objc.Send[HKVerifiableClinicalRecord](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKVerifiableClinicalRecord) Init() HKVerifiableClinicalRecord {
	rv := objc.Send[HKVerifiableClinicalRecord](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKVerifiableClinicalRecord) Autorelease() HKVerifiableClinicalRecord {
	rv := objc.Send[HKVerifiableClinicalRecord](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKVerifiableClinicalRecord creates a new HKVerifiableClinicalRecord instance.
func NewHKVerifiableClinicalRecord() HKVerifiableClinicalRecord {
	return getHKVerifiableClinicalRecordClass().New()
}


// A raw representation of the record’s data.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVerifiableClinicalRecord/dataRepresentation
func (h_ HKVerifiableClinicalRecord) DataRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("dataRepresentation"))
	return rv
}




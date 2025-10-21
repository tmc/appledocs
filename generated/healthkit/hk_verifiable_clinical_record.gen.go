// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
func (h_ HKVerifiableClinicalRecord) DataRepresentation() foundation.NSData {
	rv := objc.Send[foundation.NSData](h_.ID, objc.Sel("dataRepresentation"))
	return rv
}

// A date relevant to this record, such as when the issuer administered a vaccine or performed a test.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVerifiableClinicalRecord/relevantDate
func (h_ HKVerifiableClinicalRecord) RelevantDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](h_.ID, objc.Sel("relevantDate"))
	return rv
}

// The date when the card expires.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/expirationdate
func (h_ HKVerifiableClinicalRecord) ExpirationDate() foundation.Date {
	rv := objc.Send[foundation.Date](h_.ID, objc.Sel("expirationDate"))
	return rv
}


// SetExpirationDate sets the value of the expirationDate property.
// The date when the card expires.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/expirationdate
func (h_ HKVerifiableClinicalRecord) SetExpirationDate(value foundation.IDate) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setExpirationDate:"), value)
}

// The date when the issuer created the card.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/issueddate
func (h_ HKVerifiableClinicalRecord) IssuedDate() foundation.Date {
	rv := objc.Send[foundation.Date](h_.ID, objc.Sel("issuedDate"))
	return rv
}


// SetIssuedDate sets the value of the issuedDate property.
// The date when the issuer created the card.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/issueddate
func (h_ HKVerifiableClinicalRecord) SetIssuedDate(value foundation.IDate) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIssuedDate:"), value)
}

// An identifier that represents the card’s issuer.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/issueridentifier
func (h_ HKVerifiableClinicalRecord) IssuerIdentifier() appkit.string {
	rv := objc.Send[appkit.string](h_.ID, objc.Sel("issuerIdentifier"))
	return rv
}


// SetIssuerIdentifier sets the value of the issuerIdentifier property.
// An identifier that represents the card’s issuer.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/issueridentifier
func (h_ HKVerifiableClinicalRecord) SetIssuerIdentifier(value appkit.string) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIssuerIdentifier:"), value)
}

// A human-readable description of the card’s contents.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/itemnames
func (h_ HKVerifiableClinicalRecord) ItemNames() appkit.string {
	rv := objc.Send[appkit.string](h_.ID, objc.Sel("itemNames"))
	return rv
}


// SetItemNames sets the value of the itemNames property.
// A human-readable description of the card’s contents.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/itemnames
func (h_ HKVerifiableClinicalRecord) SetItemNames(value appkit.string) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setItemNames:"), value)
}

// A raw representation of the SMART Health Card’s contents.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/jwsrepresentation
func (h_ HKVerifiableClinicalRecord) JwsRepresentation() foundation.Data {
	rv := objc.Send[foundation.Data](h_.ID, objc.Sel("jwsRepresentation"))
	return rv
}


// SetJwsRepresentation sets the value of the jwsRepresentation property.
// A raw representation of the SMART Health Card’s contents.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/jwsrepresentation
func (h_ HKVerifiableClinicalRecord) SetJwsRepresentation(value foundation.IData) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setJwsRepresentation:"), value)
}

// An array of strings representing the types of records contained in the card.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/recordtypes
func (h_ HKVerifiableClinicalRecord) RecordTypes() appkit.string {
	rv := objc.Send[appkit.string](h_.ID, objc.Sel("recordTypes"))
	return rv
}


// SetRecordTypes sets the value of the recordTypes property.
// An array of strings representing the types of records contained in the card.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/recordtypes
func (h_ HKVerifiableClinicalRecord) SetRecordTypes(value appkit.string) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setRecordTypes:"), value)
}

// The source for the verifiable clinical record
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/sourcetype
func (h_ HKVerifiableClinicalRecord) SourceType() HKVerifiableClinicalRecordSourceType {
	rv := objc.Send[HKVerifiableClinicalRecordSourceType](h_.ID, objc.Sel("sourceType"))
	return rv
}


// SetSourceType sets the value of the sourceType property.
// The source for the verifiable clinical record

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/sourcetype
func (h_ HKVerifiableClinicalRecord) SetSourceType(value HKVerifiableClinicalRecordSourceType) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSourceType:"), value)
}

// Data about the person whose clinical data the card contains.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/subject
func (h_ HKVerifiableClinicalRecord) Subject() HKVerifiableClinicalRecordSubject {
	rv := objc.Send[HKVerifiableClinicalRecordSubject](h_.ID, objc.Sel("subject"))
	return rv
}


// SetSubject sets the value of the subject property.
// Data about the person whose clinical data the card contains.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/subject
func (h_ HKVerifiableClinicalRecord) SetSubject(value IHKVerifiableClinicalRecordSubject) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSubject:"), value)
}




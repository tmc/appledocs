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
	// properties:
	DataRepresentation() objc.IObject /* cross-framework: Data */
	SetDataRepresentation(value objc.IObject /* cross-framework: Data */)
	ExpirationDate() objc.IObject /* cross-framework: Date */
	SetExpirationDate(value objc.IObject /* cross-framework: Date */)
	IssuedDate() objc.IObject /* cross-framework: Date */
	SetIssuedDate(value objc.IObject /* cross-framework: Date */)
	IssuerIdentifier() objc.IObject /* cross-framework: NSString */
	SetIssuerIdentifier(value objc.IObject /* cross-framework: NSString */)
	ItemNames() objc.IObject /* cross-framework: NSString */
	SetItemNames(value objc.IObject /* cross-framework: NSString */)
	JwsRepresentation() objc.IObject /* cross-framework: Data */
	SetJwsRepresentation(value objc.IObject /* cross-framework: Data */)
	RecordTypes() objc.IObject /* cross-framework: NSString */
	SetRecordTypes(value objc.IObject /* cross-framework: NSString */)
	RelevantDate() objc.IObject /* cross-framework: Date */
	SetRelevantDate(value objc.IObject /* cross-framework: Date */)
	SourceType() HKVerifiableClinicalRecordSourceType /* typedef */
	SetSourceType(value HKVerifiableClinicalRecordSourceType /* typedef */)
	Subject() IHKVerifiableClinicalRecordSubject
	SetSubject(value IHKVerifiableClinicalRecordSubject)
	// methods:
}

// A sample that represents the contents of a SMART Health Card or EU Digital COVID Certificate.
//
// samples contain data from a SMART Health Card or EU Digital COVID Certificate. Verifiable clinical records combine information about the user’s identity with clinical data, like an immunization record or a lab test result. The organization that produced the data cryptographically signs the bundle. Apps that use verifiable clinical records can use the cryptographic signature to verify the authenticity of the contents. To verify the card: Access the card’s raw payload using the clinical record’s property. Unzip the payload and parse out the value, which contains a URL that identifies the organization that issued the card. Get the public key from the issuer. Verify the payload’s signature. For more information, see and . You can download example SMART cards for testing and development from .


// A sample that represents the contents of a SMART Health Card or EU Digital COVID Certificate.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/datarepresentation
func (h_ HKVerifiableClinicalRecord) DataRepresentation() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](h_.ID, objc.Sel("dataRepresentation"))
	return rv
}


// A raw representation of the record’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/datarepresentation
func (h_ HKVerifiableClinicalRecord) SetDataRepresentation(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setDataRepresentation:"), value)
}


// The date when the card expires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/expirationdate
func (h_ HKVerifiableClinicalRecord) ExpirationDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](h_.ID, objc.Sel("expirationDate"))
	return rv
}


// The date when the card expires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/expirationdate
func (h_ HKVerifiableClinicalRecord) SetExpirationDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setExpirationDate:"), value)
}


// The date when the issuer created the card.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/issueddate
func (h_ HKVerifiableClinicalRecord) IssuedDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](h_.ID, objc.Sel("issuedDate"))
	return rv
}


// The date when the issuer created the card.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/issueddate
func (h_ HKVerifiableClinicalRecord) SetIssuedDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIssuedDate:"), value)
}


// An identifier that represents the card’s issuer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/issueridentifier
func (h_ HKVerifiableClinicalRecord) IssuerIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("issuerIdentifier"))
	return rv
}


// An identifier that represents the card’s issuer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/issueridentifier
func (h_ HKVerifiableClinicalRecord) SetIssuerIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setIssuerIdentifier:"), value)
}


// A human-readable description of the card’s contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/itemnames
func (h_ HKVerifiableClinicalRecord) ItemNames() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("itemNames"))
	return rv
}


// A human-readable description of the card’s contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/itemnames
func (h_ HKVerifiableClinicalRecord) SetItemNames(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setItemNames:"), value)
}


// A raw representation of the SMART Health Card’s contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/jwsrepresentation
func (h_ HKVerifiableClinicalRecord) JwsRepresentation() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](h_.ID, objc.Sel("jwsRepresentation"))
	return rv
}


// A raw representation of the SMART Health Card’s contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/jwsrepresentation
func (h_ HKVerifiableClinicalRecord) SetJwsRepresentation(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setJwsRepresentation:"), value)
}


// An array of strings representing the types of records contained in the card.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/recordtypes
func (h_ HKVerifiableClinicalRecord) RecordTypes() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("recordTypes"))
	return rv
}


// An array of strings representing the types of records contained in the card.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/recordtypes
func (h_ HKVerifiableClinicalRecord) SetRecordTypes(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setRecordTypes:"), value)
}


// A date relevant to this record, such as when the issuer administered a vaccine or performed a test.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/relevantdate
func (h_ HKVerifiableClinicalRecord) RelevantDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](h_.ID, objc.Sel("relevantDate"))
	return rv
}


// A date relevant to this record, such as when the issuer administered a vaccine or performed a test.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/relevantdate
func (h_ HKVerifiableClinicalRecord) SetRelevantDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setRelevantDate:"), value)
}


// The source for the verifiable clinical record
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/sourcetype
func (h_ HKVerifiableClinicalRecord) SourceType() HKVerifiableClinicalRecordSourceType /* typedef */ {
	rv := objc.Send[HKVerifiableClinicalRecordSourceType](h_.ID, objc.Sel("sourceType"))
	return rv
}


// The source for the verifiable clinical record
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/sourcetype
func (h_ HKVerifiableClinicalRecord) SetSourceType(value HKVerifiableClinicalRecordSourceType /* typedef */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSourceType:"), value)
}


// Data about the person whose clinical data the card contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/subject
func (h_ HKVerifiableClinicalRecord) Subject() IHKVerifiableClinicalRecordSubject {
	rv := objc.Send[HKVerifiableClinicalRecordSubject](h_.ID, objc.Sel("subject"))
	return rv
}


// Data about the person whose clinical data the card contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkverifiableclinicalrecord/subject
func (h_ HKVerifiableClinicalRecord) SetSubject(value IHKVerifiableClinicalRecordSubject) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSubject:"), value)
}




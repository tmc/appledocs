// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class HKVerifiableClinicalRecord */


/* debug [class_header]: Header for HKVerifiableClinicalRecord */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKVerifiableClinicalRecord */
// An interface definition for the [HKVerifiableClinicalRecord] class.
type IHKVerifiableClinicalRecord interface {
	IHKSample
	
/* debug [class_interface_properties]: Properties for HKVerifiableClinicalRecord */
	// properties:
	DataRepresentation() objc.IObject /* cross-framework: NSData */
	ExpirationDate() objc.IObject /* cross-framework: NSDate */
	IssuedDate() objc.IObject /* cross-framework: NSDate */
	IssuerIdentifier() objc.IObject /* cross-framework: NSString */
	ItemNames() []string
	JWSRepresentation() objc.IObject /* cross-framework: NSData */
	RecordTypes() []string
	RelevantDate() objc.IObject /* cross-framework: NSDate */
	SourceType() HKVerifiableClinicalRecordSourceType /* typedef */
	Subject() IHKVerifiableClinicalRecordSubject
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKVerifiableClinicalRecord */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKVerifiableClinicalRecord */
// Alloc allocates a new instance without initialization.
func (hc _HKVerifiableClinicalRecordClass) Alloc() HKVerifiableClinicalRecord {
	rv := objc.Send[HKVerifiableClinicalRecord](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKVerifiableClinicalRecord */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKVerifiableClinicalRecord *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKVerifiableClinicalRecord */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKVerifiableClinicalRecord */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKVerifiableClinicalRecord */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKVerifiableClinicalRecord */

// A raw representation of the record’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVerifiableClinicalRecord/dataRepresentation
func (h_ HKVerifiableClinicalRecord) DataRepresentation() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](h_.ID, objc.Sel("dataRepresentation"))
	return rv
}/* debug [instance_properties/getter]: dataRepresentation */


// The date when the card expires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVerifiableClinicalRecord/expirationDate
func (h_ HKVerifiableClinicalRecord) ExpirationDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](h_.ID, objc.Sel("expirationDate"))
	return rv
}/* debug [instance_properties/getter]: expirationDate */


// The date when the issuer created the card.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVerifiableClinicalRecord/issuedDate
func (h_ HKVerifiableClinicalRecord) IssuedDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](h_.ID, objc.Sel("issuedDate"))
	return rv
}/* debug [instance_properties/getter]: issuedDate */


// An identifier that represents the card’s issuer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVerifiableClinicalRecord/issuerIdentifier
func (h_ HKVerifiableClinicalRecord) IssuerIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("issuerIdentifier"))
	return rv
}/* debug [instance_properties/getter]: issuerIdentifier */


// A human-readable description of the card’s contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVerifiableClinicalRecord/itemNames
func (h_ HKVerifiableClinicalRecord) ItemNames() []string {
	rv := objc.Send[[]string](h_.ID, objc.Sel("itemNames"))
	return rv
}/* debug [instance_properties/getter]: itemNames */


// A raw representation of the SMART Health Card’s contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVerifiableClinicalRecord/jwsRepresentation
func (h_ HKVerifiableClinicalRecord) JWSRepresentation() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](h_.ID, objc.Sel("JWSRepresentation"))
	return rv
}/* debug [instance_properties/getter]: JWSRepresentation */


// An array of strings representing the types of records contained in the card.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVerifiableClinicalRecord/recordTypes
func (h_ HKVerifiableClinicalRecord) RecordTypes() []string {
	rv := objc.Send[[]string](h_.ID, objc.Sel("recordTypes"))
	return rv
}/* debug [instance_properties/getter]: recordTypes */


// A date relevant to this record, such as when the issuer administered a vaccine or performed a test.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVerifiableClinicalRecord/relevantDate
func (h_ HKVerifiableClinicalRecord) RelevantDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](h_.ID, objc.Sel("relevantDate"))
	return rv
}/* debug [instance_properties/getter]: relevantDate */


// The source for the verifiable clinical record
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVerifiableClinicalRecord/sourceType
func (h_ HKVerifiableClinicalRecord) SourceType() HKVerifiableClinicalRecordSourceType /* typedef */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("sourceType"))
	return rv
}/* debug [instance_properties/getter]: sourceType */


// Data about the person whose clinical data the card contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKVerifiableClinicalRecord/subject
func (h_ HKVerifiableClinicalRecord) Subject() IHKVerifiableClinicalRecordSubject {
	rv := objc.Send[HKVerifiableClinicalRecordSubject](h_.ID, objc.Sel("subject"))
	return rv
}/* debug [instance_properties/getter]: subject */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKVerifiableClinicalRecord */




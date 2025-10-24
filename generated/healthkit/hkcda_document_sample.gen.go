// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKCDADocumentSample */


/* debug [class_header]: Header for HKCDADocumentSample */
// The class instance for the [HKCDADocumentSample] class.
var (
	HKCDADocumentSampleClass     _HKCDADocumentSampleClass
	HKCDADocumentSampleClassOnce sync.Once
)

func getHKCDADocumentSampleClass() _HKCDADocumentSampleClass {
	HKCDADocumentSampleClassOnce.Do(func() {
		HKCDADocumentSampleClass = _HKCDADocumentSampleClass{objc.GetClass("HKCDADocumentSample")}
	})
	return HKCDADocumentSampleClass
}

type _HKCDADocumentSampleClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKCDADocumentSample */
// An interface definition for the [HKCDADocumentSample] class.
type IHKCDADocumentSample interface {
	IHKDocumentSample
	
/* debug [class_interface_properties]: Properties for HKCDADocumentSample */
	// properties:
	Document() IHKCDADocument
	HKDetailedCDAValidationErrorKey() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathCDAAuthorName() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathCDACustodianName() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathCDAPatientName() objc.IObject /* cross-framework: NSString */
	HKPredicateKeyPathCDATitle() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKCDADocumentSample */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKCDADocumentSample */
// Alloc allocates a new instance without initialization.
func (hc _HKCDADocumentSampleClass) Alloc() HKCDADocumentSample {
	rv := objc.Send[HKCDADocumentSample](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKCDADocumentSampleClass) New() HKCDADocumentSample {
	rv := objc.Send[HKCDADocumentSample](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKCDADocumentSample) Init() HKCDADocumentSample {
	rv := objc.Send[HKCDADocumentSample](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKCDADocumentSample) Autorelease() HKCDADocumentSample {
	rv := objc.Send[HKCDADocumentSample](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKCDADocumentSample creates a new HKCDADocumentSample instance.
func NewHKCDADocumentSample() HKCDADocumentSample {
	return getHKCDADocumentSampleClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKCDADocumentSample */
// A Clinical Document Architecture (CDA) sample that stores a single document.
//
// The sample’s property contains an object, representing the underlying XML document. The class is a concrete subclass of the class. Document samples are immutable. HealthKit assigns the document’s properties when the sample is created. They cannot change. If you need to update a document in HealthKit, create a new document sample with the updated CDA document.


// A Clinical Document Architecture (CDA) sample that stores a single document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCDADocumentSample
type HKCDADocumentSample struct {
	HKDocumentSample
}

// HKCDADocumentSampleFrom constructs a [HKCDADocumentSample] from an unsafe.Pointer.
//
// A Clinical Document Architecture (CDA) sample that stores a single document.
func HKCDADocumentSampleFrom(ptr unsafe.Pointer) HKCDADocumentSample {
	return HKCDADocumentSample{
		HKDocumentSample: HKDocumentSampleFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKCDADocumentSample */

// Returns a CDA document sample containing the provided XML document and metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCDADocumentSample/init(data:start:end:metadata:)
func NewHKCDADocumentSampleWithDataStartDateEndDateMetadataValidationError(documentData objc.IObject /* cross-framework: NSData */, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, metadata foundation.IDictionary, validationError objectivec.IObject) HKCDADocumentSample {
	rv := objc.Send[HKCDADocumentSample](objc.ID(getHKCDADocumentSampleClass().class), objc.Sel("CDADocumentSampleWithData:startDate:endDate:metadata:validationError:"), documentData, startDate, endDate, metadata, validationError)
	return rv
}/* debug [class_init_methods/constructor]: NewHKCDADocumentSampleWithDataStartDateEndDateMetadataValidationError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKCDADocumentSample */

// Returns a CDA document sample containing the provided XML document and metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCDADocumentSample/init(data:start:end:metadata:)
func (hc _HKCDADocumentSampleClass) CDADocumentSampleWithDataStartDateEndDateMetadataValidationError(documentData objc.IObject /* cross-framework: NSData */, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */, metadata foundation.IDictionary, validationError objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("CDADocumentSampleWithData:startDate:endDate:metadata:validationError:"), documentData, startDate, endDate, metadata, validationError)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CDADocumentSampleWithDataStartDateEndDateMetadataValidationError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKCDADocumentSample */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKCDADocumentSample */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKCDADocumentSample */

// The CDA document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCDADocumentSample/document
func (h_ HKCDADocumentSample) Document() IHKCDADocument {
	rv := objc.Send[HKCDADocument](h_.ID, objc.Sel("document"))
	return rv
}/* debug [instance_properties/getter]: document */


// A key for accessing validation error information from an error object’s user information dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdetailedcdavalidationerrorkey
func (h_ HKCDADocumentSample) HKDetailedCDAValidationErrorKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKDetailedCDAValidationErrorKey"))
	return rv
}/* debug [instance_properties/getter]: HKDetailedCDAValidationErrorKey */


// The key path for accessing the author’s name inside a predicate format string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathcdaauthorname
func (h_ HKCDADocumentSample) HKPredicateKeyPathCDAAuthorName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathCDAAuthorName"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathCDAAuthorName */


// The key path for accessing the custodian’s name inside a predicate format string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathcdacustodianname
func (h_ HKCDADocumentSample) HKPredicateKeyPathCDACustodianName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathCDACustodianName"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathCDACustodianName */


// The key path for accessing the patient’s name inside a predicate format string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathcdapatientname
func (h_ HKCDADocumentSample) HKPredicateKeyPathCDAPatientName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathCDAPatientName"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathCDAPatientName */


// The key path for accessing the document’s title inside a predicate format string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathcdatitle
func (h_ HKCDADocumentSample) HKPredicateKeyPathCDATitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("HKPredicateKeyPathCDATitle"))
	return rv
}/* debug [instance_properties/getter]: HKPredicateKeyPathCDATitle */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKCDADocumentSample */



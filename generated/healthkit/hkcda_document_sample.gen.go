// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [HKCDADocumentSample] class.
type IHKCDADocumentSample interface {
	IHKDocumentSample
}

// A Clinical Document Architecture (CDA) sample that stores a single document.
//
// The sample’s property contains an object, representing the underlying XML document. The class is a concrete subclass of the class. Document samples are immutable. HealthKit assigns the document’s properties when the sample is created. They cannot change. If you need to update a document in HealthKit, create a new document sample with the updated CDA document.
//
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

// Alloc allocates a new instance without initialization.
func (hc _HKCDADocumentSampleClass) Alloc() HKCDADocumentSample {
	rv := objc.Send[HKCDADocumentSample](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The CDA document.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCDADocumentSample/document
func (h_ HKCDADocumentSample) Document() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("document"))
	return rv
}

// A key for accessing validation error information from an error object’s user information dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdetailedcdavalidationerrorkey
func (h_ HKCDADocumentSample) HKDetailedCDAValidationErrorKey() appkit.string {
	rv := objc.Send[appkit.string](h_.ID, objc.Sel("HKDetailedCDAValidationErrorKey"))
	return rv
}

// The key path for accessing the author’s name inside a predicate format string.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathcdaauthorname
func (h_ HKCDADocumentSample) HKPredicateKeyPathCDAAuthorName() appkit.string {
	rv := objc.Send[appkit.string](h_.ID, objc.Sel("HKPredicateKeyPathCDAAuthorName"))
	return rv
}

// The key path for accessing the custodian’s name inside a predicate format string.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathcdacustodianname
func (h_ HKCDADocumentSample) HKPredicateKeyPathCDACustodianName() appkit.string {
	rv := objc.Send[appkit.string](h_.ID, objc.Sel("HKPredicateKeyPathCDACustodianName"))
	return rv
}

// The key path for accessing the patient’s name inside a predicate format string.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathcdapatientname
func (h_ HKCDADocumentSample) HKPredicateKeyPathCDAPatientName() appkit.string {
	rv := objc.Send[appkit.string](h_.ID, objc.Sel("HKPredicateKeyPathCDAPatientName"))
	return rv
}

// The key path for accessing the document’s title inside a predicate format string.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkpredicatekeypathcdatitle
func (h_ HKCDADocumentSample) HKPredicateKeyPathCDATitle() appkit.string {
	rv := objc.Send[appkit.string](h_.ID, objc.Sel("HKPredicateKeyPathCDATitle"))
	return rv
}




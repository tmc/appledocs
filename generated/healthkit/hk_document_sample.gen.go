// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKDocumentSample] class.
var (
	HKDocumentSampleClass     _HKDocumentSampleClass
	HKDocumentSampleClassOnce sync.Once
)

func getHKDocumentSampleClass() _HKDocumentSampleClass {
	HKDocumentSampleClassOnce.Do(func() {
		HKDocumentSampleClass = _HKDocumentSampleClass{objc.GetClass("HKDocumentSample")}
	})
	return HKDocumentSampleClass
}

type _HKDocumentSampleClass struct {
	class objc.Class
}

// An interface definition for the [HKDocumentSample] class.
type IHKDocumentSample interface {
	IHKSample
}

// An abstract class that represents a health document in the HealthKit store.
//
// You should never instantiate an object directly. Instead, you always work with a concrete subclass. In iOS 10 and watchOS 3, the only concrete class is the class. Document samples are immutable: You set the sample’s properties when you create it, and they cannot change.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDocumentSample
type HKDocumentSample struct {
	HKSample
}

// HKDocumentSampleFrom constructs a [HKDocumentSample] from an unsafe.Pointer.
//
// An abstract class that represents a health document in the HealthKit store.
func HKDocumentSampleFrom(ptr unsafe.Pointer) HKDocumentSample {
	return HKDocumentSample{
		HKSample: HKSampleFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKDocumentSampleClass) Alloc() HKDocumentSample {
	rv := objc.Send[HKDocumentSample](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKDocumentSampleClass) New() HKDocumentSample {
	rv := objc.Send[HKDocumentSample](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKDocumentSample) Init() HKDocumentSample {
	rv := objc.Send[HKDocumentSample](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKDocumentSample) Autorelease() HKDocumentSample {
	rv := objc.Send[HKDocumentSample](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKDocumentSample creates a new HKDocumentSample instance.
func NewHKDocumentSample() HKDocumentSample {
	return getHKDocumentSampleClass().New()
}


// The type of document represented by the sample.
//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdocumentsample/documenttype
func (h_ HKDocumentSample) DocumentType() HKDocumentType {
	rv := objc.Send[HKDocumentType](h_.ID, objc.Sel("documentType"))
	return rv
}


// SetDocumentType sets the value of the documentType property.
// The type of document represented by the sample.

//
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkdocumentsample/documenttype
func (h_ HKDocumentSample) SetDocumentType(value HKDocumentType) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setDocumentType:"), value)
}




// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKDocumentType] class.
var (
	HKDocumentTypeClass     _HKDocumentTypeClass
	HKDocumentTypeClassOnce sync.Once
)

func getHKDocumentTypeClass() _HKDocumentTypeClass {
	HKDocumentTypeClassOnce.Do(func() {
		HKDocumentTypeClass = _HKDocumentTypeClass{objc.GetClass("HKDocumentType")}
	})
	return HKDocumentTypeClass
}

type _HKDocumentTypeClass struct {
	class objc.Class
}

// An interface definition for the [HKDocumentType] class.
type IHKDocumentType interface {
	IHKSampleType
}

// A sample type used to create queries for documents.
//
// To create a document type instance, use the class’s convenience method.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKDocumentType
type HKDocumentType struct {
	HKSampleType
}

// HKDocumentTypeFrom constructs a [HKDocumentType] from an unsafe.Pointer.
//
// A sample type used to create queries for documents.
func HKDocumentTypeFrom(ptr unsafe.Pointer) HKDocumentType {
	return HKDocumentType{
		HKSampleType: HKSampleTypeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKDocumentTypeClass) Alloc() HKDocumentType {
	rv := objc.Send[HKDocumentType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKDocumentTypeClass) New() HKDocumentType {
	rv := objc.Send[HKDocumentType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKDocumentType) Init() HKDocumentType {
	rv := objc.Send[HKDocumentType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKDocumentType) Autorelease() HKDocumentType {
	rv := objc.Send[HKDocumentType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKDocumentType creates a new HKDocumentType instance.
func NewHKDocumentType() HKDocumentType {
	return getHKDocumentTypeClass().New()
}





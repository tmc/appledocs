// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKCategoryType] class.
var (
	HKCategoryTypeClass     _HKCategoryTypeClass
	HKCategoryTypeClassOnce sync.Once
)

func getHKCategoryTypeClass() _HKCategoryTypeClass {
	HKCategoryTypeClassOnce.Do(func() {
		HKCategoryTypeClass = _HKCategoryTypeClass{objc.GetClass("HKCategoryType")}
	})
	return HKCategoryTypeClass
}

type _HKCategoryTypeClass struct {
	class objc.Class
}

// An interface definition for the [HKCategoryType] class.
type IHKCategoryType interface {
	IHKSampleType
}

// A type that identifies samples that contain a value from a small set of possible values.
//
// The class is a concrete subclass of the HKObjectType class. To create a category type instance, use the convenience method. For example, the following code creates a category sample type for handwashing events. Use category types to: Request permission to read or write matching category samples. Create and share matching category samples. Query for matching category samples. For a complete list of category types, refer to .
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryType
type HKCategoryType struct {
	HKSampleType
}

// HKCategoryTypeFrom constructs a [HKCategoryType] from an unsafe.Pointer.
//
// A type that identifies samples that contain a value from a small set of possible values.
func HKCategoryTypeFrom(ptr unsafe.Pointer) HKCategoryType {
	return HKCategoryType{
		HKSampleType: HKSampleTypeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKCategoryTypeClass) Alloc() HKCategoryType {
	rv := objc.Send[HKCategoryType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKCategoryTypeClass) New() HKCategoryType {
	rv := objc.Send[HKCategoryType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKCategoryType) Init() HKCategoryType {
	rv := objc.Send[HKCategoryType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKCategoryType) Autorelease() HKCategoryType {
	rv := objc.Send[HKCategoryType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKCategoryType creates a new HKCategoryType instance.
func NewHKCategoryType() HKCategoryType {
	return getHKCategoryTypeClass().New()
}





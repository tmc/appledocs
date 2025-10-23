// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKActivitySummaryType] class.
var (
	HKActivitySummaryTypeClass     _HKActivitySummaryTypeClass
	HKActivitySummaryTypeClassOnce sync.Once
)

func getHKActivitySummaryTypeClass() _HKActivitySummaryTypeClass {
	HKActivitySummaryTypeClassOnce.Do(func() {
		HKActivitySummaryTypeClass = _HKActivitySummaryTypeClass{objc.GetClass("HKActivitySummaryType")}
	})
	return HKActivitySummaryTypeClass
}

type _HKActivitySummaryTypeClass struct {
	class objc.Class
}

// An interface definition for the [HKActivitySummaryType] class.
type IHKActivitySummaryType interface {
	IHKObjectType
}

// A type that identifies activity summary objects.
//
// Use the activity summary type to request permission to read objects from the HealthKit store. To create an activity summary type, use the class’s convenience method. The class is a concrete subclass of the class. Like many HealthKit classes, activity summary types aren’t extensible and you shouldn’t subclass them.


// A type that identifies activity summary objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKActivitySummaryType
type HKActivitySummaryType struct {
	HKObjectType
}

// HKActivitySummaryTypeFrom constructs a [HKActivitySummaryType] from an unsafe.Pointer.
//
// A type that identifies activity summary objects.
func HKActivitySummaryTypeFrom(ptr unsafe.Pointer) HKActivitySummaryType {
	return HKActivitySummaryType{
		HKObjectType: HKObjectTypeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKActivitySummaryTypeClass) Alloc() HKActivitySummaryType {
	rv := objc.Send[HKActivitySummaryType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKActivitySummaryTypeClass) New() HKActivitySummaryType {
	rv := objc.Send[HKActivitySummaryType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKActivitySummaryType) Init() HKActivitySummaryType {
	rv := objc.Send[HKActivitySummaryType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKActivitySummaryType) Autorelease() HKActivitySummaryType {
	rv := objc.Send[HKActivitySummaryType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKActivitySummaryType creates a new HKActivitySummaryType instance.
func NewHKActivitySummaryType() HKActivitySummaryType {
	return getHKActivitySummaryTypeClass().New()
}





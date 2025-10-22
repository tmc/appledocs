// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKStateOfMindType] class.
var (
	HKStateOfMindTypeClass     _HKStateOfMindTypeClass
	HKStateOfMindTypeClassOnce sync.Once
)

func getHKStateOfMindTypeClass() _HKStateOfMindTypeClass {
	HKStateOfMindTypeClassOnce.Do(func() {
		HKStateOfMindTypeClass = _HKStateOfMindTypeClass{objc.GetClass("HKStateOfMindType")}
	})
	return HKStateOfMindTypeClass
}

type _HKStateOfMindTypeClass struct {
	class objc.Class
}

// An interface definition for the [HKStateOfMindType] class.
type IHKStateOfMindType interface {
	IHKSampleType
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMindType

type HKStateOfMindType struct {
	HKSampleType
}

// HKStateOfMindTypeFrom constructs a [HKStateOfMindType] from an unsafe.Pointer.
func HKStateOfMindTypeFrom(ptr unsafe.Pointer) HKStateOfMindType {
	return HKStateOfMindType{
		HKSampleType: HKSampleTypeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKStateOfMindTypeClass) Alloc() HKStateOfMindType {
	rv := objc.Send[HKStateOfMindType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKStateOfMindTypeClass) New() HKStateOfMindType {
	rv := objc.Send[HKStateOfMindType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKStateOfMindType) Init() HKStateOfMindType {
	rv := objc.Send[HKStateOfMindType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKStateOfMindType) Autorelease() HKStateOfMindType {
	rv := objc.Send[HKStateOfMindType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKStateOfMindType creates a new HKStateOfMindType instance.
func NewHKStateOfMindType() HKStateOfMindType {
	return getHKStateOfMindTypeClass().New()
}





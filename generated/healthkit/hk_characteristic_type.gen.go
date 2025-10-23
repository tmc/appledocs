// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKCharacteristicType] class.
var (
	HKCharacteristicTypeClass     _HKCharacteristicTypeClass
	HKCharacteristicTypeClassOnce sync.Once
)

func getHKCharacteristicTypeClass() _HKCharacteristicTypeClass {
	HKCharacteristicTypeClassOnce.Do(func() {
		HKCharacteristicTypeClass = _HKCharacteristicTypeClass{objc.GetClass("HKCharacteristicType")}
	})
	return HKCharacteristicTypeClass
}

type _HKCharacteristicTypeClass struct {
	class objc.Class
}

// An interface definition for the [HKCharacteristicType] class.
type IHKCharacteristicType interface {
	IHKObjectType
}

// A type that represents data that doesn’t typically change over time.
//
// The class is a concrete subclass of the class. To create a characteristic type instance, use the object type’s convenience method. Unlike the other object types, characteristic types cannot be used to create and save new HealthKit objects. Instead, users must enter and edit their characteristic data using the Health app. Similarly, you cannot create queries for characteristic types. Instead, use the HealthKit store to access the data (see Reading characteristic data). HealthKit provides five characteristic types: biological sex, blood type, birthdate, Fitzpatrick skin type, and wheelchair use. These types are used only when asking for permission to read data from the HealthKit store.


// A type that represents data that doesn’t typically change over time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCharacteristicType
type HKCharacteristicType struct {
	HKObjectType
}

// HKCharacteristicTypeFrom constructs a [HKCharacteristicType] from an unsafe.Pointer.
//
// A type that represents data that doesn’t typically change over time.
func HKCharacteristicTypeFrom(ptr unsafe.Pointer) HKCharacteristicType {
	return HKCharacteristicType{
		HKObjectType: HKObjectTypeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKCharacteristicTypeClass) Alloc() HKCharacteristicType {
	rv := objc.Send[HKCharacteristicType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKCharacteristicTypeClass) New() HKCharacteristicType {
	rv := objc.Send[HKCharacteristicType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKCharacteristicType) Init() HKCharacteristicType {
	rv := objc.Send[HKCharacteristicType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKCharacteristicType) Autorelease() HKCharacteristicType {
	rv := objc.Send[HKCharacteristicType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKCharacteristicType creates a new HKCharacteristicType instance.
func NewHKCharacteristicType() HKCharacteristicType {
	return getHKCharacteristicTypeClass().New()
}





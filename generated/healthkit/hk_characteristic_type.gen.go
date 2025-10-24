// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class HKCharacteristicType */


/* debug [class_header]: Header for HKCharacteristicType */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKCharacteristicType */
// An interface definition for the [HKCharacteristicType] class.
type IHKCharacteristicType interface {
	IHKObjectType
	
/* debug [class_interface_properties]: Properties for HKCharacteristicType */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKCharacteristicType */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKCharacteristicType */
// Alloc allocates a new instance without initialization.
func (hc _HKCharacteristicTypeClass) Alloc() HKCharacteristicType {
	rv := objc.Send[HKCharacteristicType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKCharacteristicType */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKCharacteristicType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKCharacteristicType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKCharacteristicType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKCharacteristicType */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKCharacteristicType */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKCharacteristicType */




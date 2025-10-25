// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSUnitInformationStorage */


/* debug [class_header]: Header for NSUnitInformationStorage */
// The class instance for the [UnitInformationStorage] class.
var (
	UnitInformationStorageClass     _UnitInformationStorageClass
	UnitInformationStorageClassOnce sync.Once
)

func getUnitInformationStorageClass() _UnitInformationStorageClass {
	UnitInformationStorageClassOnce.Do(func() {
		UnitInformationStorageClass = _UnitInformationStorageClass{objc.GetClass("NSUnitInformationStorage")}
	})
	return UnitInformationStorageClass
}

type _UnitInformationStorageClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UnitInformationStorage */
// An interface definition for the [UnitInformationStorage] class.
type IUnitInformationStorage interface {
	IDimension
	
/* debug [class_interface_properties]: Properties for UnitInformationStorage */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UnitInformationStorage */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UnitInformationStorage */
// Alloc allocates a new instance without initialization.
func (uc _UnitInformationStorageClass) Alloc() UnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UnitInformationStorageClass) New() UnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitInformationStorage) Init() UnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitInformationStorage) Autorelease() UnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitInformationStorage creates a new UnitInformationStorage instance.
func NewUnitInformationStorage() UnitInformationStorage {
	return getUnitInformationStorageClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UnitInformationStorage */
// A unit of measure for quantities of information.
//
// Use instances of to represent quantities of information using the class. The base unit of measure for information is the bit, with a nibble representing four bits and a byte representing eight bits. Larger units of information expand on bits and bytes by orders of magnitude in both decimal and binary forms.


// A unit of measure for quantities of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage
type UnitInformationStorage struct {
	Dimension
}

// UnitInformationStorageFrom constructs a [UnitInformationStorage] from an unsafe.Pointer.
//
// A unit of measure for quantities of information.
func UnitInformationStorageFrom(ptr unsafe.Pointer) UnitInformationStorage {
	return UnitInformationStorage{
		Dimension: DimensionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UnitInformationStorage *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UnitInformationStorage */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UnitInformationStorage */

// The gibibytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/gibibytes
func (uc _UnitInformationStorageClass) Gibibytes() UnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](objc.ID(uc.class), objc.Sel("gibibytes"))
	return rv
}/* debug [class_properties_class/property]: gibibytes */

// The gigabits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/gigabits
func (uc _UnitInformationStorageClass) Gigabits() UnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](objc.ID(uc.class), objc.Sel("gigabits"))
	return rv
}/* debug [class_properties_class/property]: gigabits */

// The mebibits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/mebibits
func (uc _UnitInformationStorageClass) Mebibits() UnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](objc.ID(uc.class), objc.Sel("mebibits"))
	return rv
}/* debug [class_properties_class/property]: mebibits */

// The megabytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/megabytes
func (uc _UnitInformationStorageClass) Megabytes() UnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](objc.ID(uc.class), objc.Sel("megabytes"))
	return rv
}/* debug [class_properties_class/property]: megabytes */

// The nibbles unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/nibbles
func (uc _UnitInformationStorageClass) Nibbles() UnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](objc.ID(uc.class), objc.Sel("nibbles"))
	return rv
}/* debug [class_properties_class/property]: nibbles */

// The pebibits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/pebibits
func (uc _UnitInformationStorageClass) Pebibits() UnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](objc.ID(uc.class), objc.Sel("pebibits"))
	return rv
}/* debug [class_properties_class/property]: pebibits */

// The pebibytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/pebibytes
func (uc _UnitInformationStorageClass) Pebibytes() UnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](objc.ID(uc.class), objc.Sel("pebibytes"))
	return rv
}/* debug [class_properties_class/property]: pebibytes */

// The yobibits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/yobibits
func (uc _UnitInformationStorageClass) Yobibits() UnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](objc.ID(uc.class), objc.Sel("yobibits"))
	return rv
}/* debug [class_properties_class/property]: yobibits */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UnitInformationStorage */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UnitInformationStorage */

// The gibibytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/gibibytes
func (u_ UnitInformationStorage) Gibibytes() IUnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](u_.ID, objc.Sel("gibibytes"))
	return rv
}/* debug [instance_properties/getter]: gibibytes */


// The gigabits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/gigabits
func (u_ UnitInformationStorage) Gigabits() IUnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](u_.ID, objc.Sel("gigabits"))
	return rv
}/* debug [instance_properties/getter]: gigabits */


// The mebibits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/mebibits
func (u_ UnitInformationStorage) Mebibits() IUnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](u_.ID, objc.Sel("mebibits"))
	return rv
}/* debug [instance_properties/getter]: mebibits */


// The megabytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/megabytes
func (u_ UnitInformationStorage) Megabytes() IUnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](u_.ID, objc.Sel("megabytes"))
	return rv
}/* debug [instance_properties/getter]: megabytes */


// The nibbles unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/nibbles
func (u_ UnitInformationStorage) Nibbles() IUnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](u_.ID, objc.Sel("nibbles"))
	return rv
}/* debug [instance_properties/getter]: nibbles */


// The pebibits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/pebibits
func (u_ UnitInformationStorage) Pebibits() IUnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](u_.ID, objc.Sel("pebibits"))
	return rv
}/* debug [instance_properties/getter]: pebibits */


// The pebibytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/pebibytes
func (u_ UnitInformationStorage) Pebibytes() IUnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](u_.ID, objc.Sel("pebibytes"))
	return rv
}/* debug [instance_properties/getter]: pebibytes */


// The yobibits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/yobibits
func (u_ UnitInformationStorage) Yobibits() IUnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](u_.ID, objc.Sel("yobibits"))
	return rv
}/* debug [instance_properties/getter]: yobibits */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUnitInformationStorage */




// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





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





// An interface definition for the [UnitInformationStorage] class.
type IUnitInformationStorage interface {
	IDimension
	

	// properties:


	

	// methods:


}





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















// The gibibytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/gibibytes
func (uc _UnitInformationStorageClass) Gibibytes() UnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](objc.ID(uc.class), objc.Sel("gibibytes"))
	return rv
}

// The gigabits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/gigabits
func (uc _UnitInformationStorageClass) Gigabits() UnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](objc.ID(uc.class), objc.Sel("gigabits"))
	return rv
}

// The mebibits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/mebibits
func (uc _UnitInformationStorageClass) Mebibits() UnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](objc.ID(uc.class), objc.Sel("mebibits"))
	return rv
}

// The megabytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/megabytes
func (uc _UnitInformationStorageClass) Megabytes() UnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](objc.ID(uc.class), objc.Sel("megabytes"))
	return rv
}

// The nibbles unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/nibbles
func (uc _UnitInformationStorageClass) Nibbles() UnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](objc.ID(uc.class), objc.Sel("nibbles"))
	return rv
}

// The pebibits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/pebibits
func (uc _UnitInformationStorageClass) Pebibits() UnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](objc.ID(uc.class), objc.Sel("pebibits"))
	return rv
}

// The pebibytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/pebibytes
func (uc _UnitInformationStorageClass) Pebibytes() UnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](objc.ID(uc.class), objc.Sel("pebibytes"))
	return rv
}

// The yobibits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/yobibits
func (uc _UnitInformationStorageClass) Yobibits() UnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](objc.ID(uc.class), objc.Sel("yobibits"))
	return rv
}











// The gibibytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/gibibytes
func (u_ UnitInformationStorage) Gibibytes() IUnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](u_.ID, objc.Sel("gibibytes"))
	return rv
}


// The gigabits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/gigabits
func (u_ UnitInformationStorage) Gigabits() IUnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](u_.ID, objc.Sel("gigabits"))
	return rv
}


// The mebibits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/mebibits
func (u_ UnitInformationStorage) Mebibits() IUnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](u_.ID, objc.Sel("mebibits"))
	return rv
}


// The megabytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/megabytes
func (u_ UnitInformationStorage) Megabytes() IUnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](u_.ID, objc.Sel("megabytes"))
	return rv
}


// The nibbles unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/nibbles
func (u_ UnitInformationStorage) Nibbles() IUnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](u_.ID, objc.Sel("nibbles"))
	return rv
}


// The pebibits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/pebibits
func (u_ UnitInformationStorage) Pebibits() IUnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](u_.ID, objc.Sel("pebibits"))
	return rv
}


// The pebibytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/pebibytes
func (u_ UnitInformationStorage) Pebibytes() IUnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](u_.ID, objc.Sel("pebibytes"))
	return rv
}


// The yobibits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/yobibits
func (u_ UnitInformationStorage) Yobibits() IUnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](u_.ID, objc.Sel("yobibits"))
	return rv
}









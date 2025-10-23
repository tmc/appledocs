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

// Alloc allocates a new instance without initialization.
func (uc _UnitInformationStorageClass) Alloc() UnitInformationStorage {
	rv := objc.Send[UnitInformationStorage](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The bits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/bits
func (uc _UnitInformationStorageClass) Bits() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("bits"))
	return rv
}

// The bytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/bytes
func (uc _UnitInformationStorageClass) Bytes() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("bytes"))
	return rv
}

// The exabits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/exabits
func (uc _UnitInformationStorageClass) Exabits() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("exabits"))
	return rv
}

// The exabytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/exabytes
func (uc _UnitInformationStorageClass) Exabytes() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("exabytes"))
	return rv
}

// The exbibits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/exbibits
func (uc _UnitInformationStorageClass) Exbibits() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("exbibits"))
	return rv
}

// The exbibytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/exbibytes
func (uc _UnitInformationStorageClass) Exbibytes() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("exbibytes"))
	return rv
}

// The gibibits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/gibibits
func (uc _UnitInformationStorageClass) Gibibits() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("gibibits"))
	return rv
}

// The gibibytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/gibibytes
func (uc _UnitInformationStorageClass) Gibibytes() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("gibibytes"))
	return rv
}

// The gigabits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/gigabits
func (uc _UnitInformationStorageClass) Gigabits() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("gigabits"))
	return rv
}

// The gigabytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/gigabytes
func (uc _UnitInformationStorageClass) Gigabytes() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("gigabytes"))
	return rv
}

// The kibibits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/kibibits
func (uc _UnitInformationStorageClass) Kibibits() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("kibibits"))
	return rv
}

// The kibibytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/kibibytes
func (uc _UnitInformationStorageClass) Kibibytes() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("kibibytes"))
	return rv
}

// The kilobits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/kilobits
func (uc _UnitInformationStorageClass) Kilobits() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("kilobits"))
	return rv
}

// The kilobytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/kilobytes
func (uc _UnitInformationStorageClass) Kilobytes() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("kilobytes"))
	return rv
}

// The mebibits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/mebibits
func (uc _UnitInformationStorageClass) Mebibits() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("mebibits"))
	return rv
}

// The mebibytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/mebibytes
func (uc _UnitInformationStorageClass) Mebibytes() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("mebibytes"))
	return rv
}

// The megabits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/megabits
func (uc _UnitInformationStorageClass) Megabits() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("megabits"))
	return rv
}

// The megabytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/megabytes
func (uc _UnitInformationStorageClass) Megabytes() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("megabytes"))
	return rv
}

// The nibbles unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/nibbles
func (uc _UnitInformationStorageClass) Nibbles() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("nibbles"))
	return rv
}

// The pebibits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/pebibits
func (uc _UnitInformationStorageClass) Pebibits() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("pebibits"))
	return rv
}

// The pebibytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/pebibytes
func (uc _UnitInformationStorageClass) Pebibytes() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("pebibytes"))
	return rv
}

// The petabits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/petabits
func (uc _UnitInformationStorageClass) Petabits() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("petabits"))
	return rv
}

// The petabytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/petabytes
func (uc _UnitInformationStorageClass) Petabytes() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("petabytes"))
	return rv
}

// The tebibits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/tebibits
func (uc _UnitInformationStorageClass) Tebibits() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("tebibits"))
	return rv
}

// The tebibytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/tebibytes
func (uc _UnitInformationStorageClass) Tebibytes() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("tebibytes"))
	return rv
}

// The terabits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/terabits
func (uc _UnitInformationStorageClass) Terabits() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("terabits"))
	return rv
}

// The terrabytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/terabytes
func (uc _UnitInformationStorageClass) Terabytes() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("terabytes"))
	return rv
}

// The yobibits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/yobibits
func (uc _UnitInformationStorageClass) Yobibits() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("yobibits"))
	return rv
}

// The yobibytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/yobibytes
func (uc _UnitInformationStorageClass) Yobibytes() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("yobibytes"))
	return rv
}

// The yottabits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/yottabits
func (uc _UnitInformationStorageClass) Yottabits() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("yottabits"))
	return rv
}

// The yottabytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/yottabytes
func (uc _UnitInformationStorageClass) Yottabytes() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("yottabytes"))
	return rv
}

// The zebibits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/zebibits
func (uc _UnitInformationStorageClass) Zebibits() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("zebibits"))
	return rv
}

// The zebibytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/zebibytes
func (uc _UnitInformationStorageClass) Zebibytes() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("zebibytes"))
	return rv
}

// The zettabits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/zettabits
func (uc _UnitInformationStorageClass) Zettabits() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("zettabits"))
	return rv
}

// The zettabytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/zettabytes
func (uc _UnitInformationStorageClass) Zettabytes() UnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](objc.ID(uc.class), objc.Sel("zettabytes"))
	return rv
}

// The bits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/bits
func (u_ UnitInformationStorage) Bits() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("bits"))
	return rv
}


// The bytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/bytes
func (u_ UnitInformationStorage) Bytes() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("bytes"))
	return rv
}


// The exabits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/exabits
func (u_ UnitInformationStorage) Exabits() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("exabits"))
	return rv
}


// The exabytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/exabytes
func (u_ UnitInformationStorage) Exabytes() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("exabytes"))
	return rv
}


// The exbibits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/exbibits
func (u_ UnitInformationStorage) Exbibits() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("exbibits"))
	return rv
}


// The exbibytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/exbibytes
func (u_ UnitInformationStorage) Exbibytes() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("exbibytes"))
	return rv
}


// The gibibits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/gibibits
func (u_ UnitInformationStorage) Gibibits() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("gibibits"))
	return rv
}


// The gibibytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/gibibytes
func (u_ UnitInformationStorage) Gibibytes() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("gibibytes"))
	return rv
}


// The gigabits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/gigabits
func (u_ UnitInformationStorage) Gigabits() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("gigabits"))
	return rv
}


// The gigabytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/gigabytes
func (u_ UnitInformationStorage) Gigabytes() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("gigabytes"))
	return rv
}


// The kibibits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/kibibits
func (u_ UnitInformationStorage) Kibibits() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("kibibits"))
	return rv
}


// The kibibytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/kibibytes
func (u_ UnitInformationStorage) Kibibytes() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("kibibytes"))
	return rv
}


// The kilobits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/kilobits
func (u_ UnitInformationStorage) Kilobits() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("kilobits"))
	return rv
}


// The kilobytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/kilobytes
func (u_ UnitInformationStorage) Kilobytes() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("kilobytes"))
	return rv
}


// The mebibits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/mebibits
func (u_ UnitInformationStorage) Mebibits() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("mebibits"))
	return rv
}


// The mebibytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/mebibytes
func (u_ UnitInformationStorage) Mebibytes() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("mebibytes"))
	return rv
}


// The megabits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/megabits
func (u_ UnitInformationStorage) Megabits() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("megabits"))
	return rv
}


// The megabytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/megabytes
func (u_ UnitInformationStorage) Megabytes() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("megabytes"))
	return rv
}


// The nibbles unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/nibbles
func (u_ UnitInformationStorage) Nibbles() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("nibbles"))
	return rv
}


// The pebibits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/pebibits
func (u_ UnitInformationStorage) Pebibits() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("pebibits"))
	return rv
}


// The pebibytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/pebibytes
func (u_ UnitInformationStorage) Pebibytes() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("pebibytes"))
	return rv
}


// The petabits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/petabits
func (u_ UnitInformationStorage) Petabits() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("petabits"))
	return rv
}


// The petabytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/petabytes
func (u_ UnitInformationStorage) Petabytes() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("petabytes"))
	return rv
}


// The tebibits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/tebibits
func (u_ UnitInformationStorage) Tebibits() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("tebibits"))
	return rv
}


// The tebibytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/tebibytes
func (u_ UnitInformationStorage) Tebibytes() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("tebibytes"))
	return rv
}


// The terabits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/terabits
func (u_ UnitInformationStorage) Terabits() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("terabits"))
	return rv
}


// The terrabytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/terabytes
func (u_ UnitInformationStorage) Terabytes() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("terabytes"))
	return rv
}


// The yobibits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/yobibits
func (u_ UnitInformationStorage) Yobibits() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("yobibits"))
	return rv
}


// The yobibytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/yobibytes
func (u_ UnitInformationStorage) Yobibytes() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("yobibytes"))
	return rv
}


// The yottabits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/yottabits
func (u_ UnitInformationStorage) Yottabits() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("yottabits"))
	return rv
}


// The yottabytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/yottabytes
func (u_ UnitInformationStorage) Yottabytes() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("yottabytes"))
	return rv
}


// The zebibits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/zebibits
func (u_ UnitInformationStorage) Zebibits() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("zebibits"))
	return rv
}


// The zebibytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/zebibytes
func (u_ UnitInformationStorage) Zebibytes() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("zebibytes"))
	return rv
}


// The zettabits unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/zettabits
func (u_ UnitInformationStorage) Zettabits() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("zettabits"))
	return rv
}


// The zettabytes unit of information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitInformationStorage/zettabytes
func (u_ UnitInformationStorage) Zettabytes() IUnitInformationStorage {
	rv := objc.Send[NSUnitInformationStorage](u_.ID, objc.Sel("zettabytes"))
	return rv
}




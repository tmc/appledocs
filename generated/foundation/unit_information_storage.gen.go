// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitInformationStorage] class.
var (
	unitInformationStorageClass     _UnitInformationStorageClass
	unitInformationStorageClassOnce sync.Once
)

func getUnitInformationStorageClass() _UnitInformationStorageClass {
	unitInformationStorageClassOnce.Do(func() {
		unitInformationStorageClass = _UnitInformationStorageClass{objc.GetClass("NSUnitInformationStorage")}
	})
	return unitInformationStorageClass
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





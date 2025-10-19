// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UnitConverter] class.
var (
	unitConverterClass     _UnitConverterClass
	unitConverterClassOnce sync.Once
)

func getUnitConverterClass() _UnitConverterClass {
	unitConverterClassOnce.Do(func() {
		unitConverterClass = _UnitConverterClass{objc.GetClass("NSUnitConverter")}
	})
	return unitConverterClass
}

type _UnitConverterClass struct {
	class objc.Class
}

// An interface definition for the [UnitConverter] class.
type IUnitConverter interface {
	objectivec.IObject
}

// An abstract class that provides a description of how to convert a unit to and from the base unit of its dimension.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitConverter
type UnitConverter struct {
	objectivec.Object
}

// UnitConverterFrom constructs a [UnitConverter] from an unsafe.Pointer.
//
// An abstract class that provides a description of how to convert a unit to and from the base unit of its dimension.
func UnitConverterFrom(ptr unsafe.Pointer) UnitConverter {
	return UnitConverter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _UnitConverterClass) Alloc() UnitConverter {
	rv := objc.Send[UnitConverter](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UnitConverterClass) New() UnitConverter {
	rv := objc.Send[UnitConverter](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitConverter) Init() UnitConverter {
	rv := objc.Send[UnitConverter](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitConverter) Autorelease() UnitConverter {
	rv := objc.Send[UnitConverter](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitConverter creates a new UnitConverter instance.
func NewUnitConverter() UnitConverter {
	return getUnitConverterClass().New()
}





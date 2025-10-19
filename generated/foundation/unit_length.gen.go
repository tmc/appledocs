// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitLength] class.
var (
	unitLengthClass     _UnitLengthClass
	unitLengthClassOnce sync.Once
)

func getUnitLengthClass() _UnitLengthClass {
	unitLengthClassOnce.Do(func() {
		unitLengthClass = _UnitLengthClass{objc.GetClass("NSUnitLength")}
	})
	return unitLengthClass
}

type _UnitLengthClass struct {
	class objc.Class
}

// An interface definition for the [UnitLength] class.
type IUnitLength interface {
	IDimension
}

// A unit of measure for length. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength
type UnitLength struct {
	Dimension
}

// UnitLengthFrom constructs a [UnitLength] from an unsafe.Pointer.
//
// A unit of measure for length.
func UnitLengthFrom(ptr unsafe.Pointer) UnitLength {
	return UnitLength{
		Dimension: DimensionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UnitLengthClass) Alloc() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UnitLengthClass) New() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitLength) Init() UnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitLength) Autorelease() UnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitLength creates a new UnitLength instance.
func NewUnitLength() UnitLength {
	return getUnitLengthClass().New()
}





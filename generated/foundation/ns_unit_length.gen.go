// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitLength] class.
var (
	UnitLengthClass     _UnitLengthClass
	UnitLengthClassOnce sync.Once
)

func getUnitLengthClass() _UnitLengthClass {
	UnitLengthClassOnce.Do(func() {
		UnitLengthClass = _UnitLengthClass{objc.GetClass("NSUnitLength")}
	})
	return UnitLengthClass
}

type _UnitLengthClass struct {
	class objc.Class
}

// An interface definition for the [UnitLength] class.
type IUnitLength interface {
	IDimension
}

// A unit of measure for length.
//
// You typically use instances of to represent specific quantities of length using the class.
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


// The hectometers unit of length.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/hectometers
func (uc _UnitLengthClass) Hectometers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("hectometers"))
	return rv
}
// The hectometers unit of length.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/hectometers
func (u_ UnitLength) Hectometers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("hectometers"))
	return rv
}




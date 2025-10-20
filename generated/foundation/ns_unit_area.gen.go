// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitArea] class.
var (
	UnitAreaClass     _UnitAreaClass
	UnitAreaClassOnce sync.Once
)

func getUnitAreaClass() _UnitAreaClass {
	UnitAreaClassOnce.Do(func() {
		UnitAreaClass = _UnitAreaClass{objc.GetClass("NSUnitArea")}
	})
	return UnitAreaClass
}

type _UnitAreaClass struct {
	class objc.Class
}

// An interface definition for the [UnitArea] class.
type IUnitArea interface {
	IDimension
}

// A unit of measure for area.
//
// You typically use instances of to represent specific quantities of area using the class.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea
type UnitArea struct {
	Dimension
}

// UnitAreaFrom constructs a [UnitArea] from an unsafe.Pointer.
//
// A unit of measure for area.
func UnitAreaFrom(ptr unsafe.Pointer) UnitArea {
	return UnitArea{
		Dimension: DimensionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UnitAreaClass) Alloc() UnitArea {
	rv := objc.Send[UnitArea](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UnitAreaClass) New() UnitArea {
	rv := objc.Send[UnitArea](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitArea) Init() UnitArea {
	rv := objc.Send[UnitArea](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitArea) Autorelease() UnitArea {
	rv := objc.Send[UnitArea](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitArea creates a new UnitArea instance.
func NewUnitArea() UnitArea {
	return getUnitAreaClass().New()
}





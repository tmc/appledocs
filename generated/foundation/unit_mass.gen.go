// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitMass] class.
var (
	unitMassClass     _UnitMassClass
	unitMassClassOnce sync.Once
)

func getUnitMassClass() _UnitMassClass {
	unitMassClassOnce.Do(func() {
		unitMassClass = _UnitMassClass{objc.GetClass("NSUnitMass")}
	})
	return unitMassClass
}

type _UnitMassClass struct {
	class objc.Class
}

// An interface definition for the [UnitMass] class.
type IUnitMass interface {
	IDimension
}

// A unit of measure for mass.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass
type UnitMass struct {
	Dimension
}

// UnitMassFrom constructs a [UnitMass] from an unsafe.Pointer.
//
// A unit of measure for mass.
func UnitMassFrom(ptr unsafe.Pointer) UnitMass {
	return UnitMass{
		Dimension: DimensionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UnitMassClass) Alloc() UnitMass {
	rv := objc.Send[UnitMass](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UnitMassClass) New() UnitMass {
	rv := objc.Send[UnitMass](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitMass) Init() UnitMass {
	rv := objc.Send[UnitMass](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitMass) Autorelease() UnitMass {
	rv := objc.Send[UnitMass](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitMass creates a new UnitMass instance.
func NewUnitMass() UnitMass {
	return getUnitMassClass().New()
}





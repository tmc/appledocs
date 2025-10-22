// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitAngle] class.
var (
	UnitAngleClass     _UnitAngleClass
	UnitAngleClassOnce sync.Once
)

func getUnitAngleClass() _UnitAngleClass {
	UnitAngleClassOnce.Do(func() {
		UnitAngleClass = _UnitAngleClass{objc.GetClass("NSUnitAngle")}
	})
	return UnitAngleClass
}

type _UnitAngleClass struct {
	class objc.Class
}

// An interface definition for the [UnitAngle] class.
type IUnitAngle interface {
	IDimension
}

// A unit of measure for planar angle and rotation.
//
// You typically use instances of to represent specific quantities of planar angle using the class.


// A unit of measure for planar angle and rotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitAngle

type UnitAngle struct {
	Dimension
}

// UnitAngleFrom constructs a [UnitAngle] from an unsafe.Pointer.
//
// A unit of measure for planar angle and rotation.
func UnitAngleFrom(ptr unsafe.Pointer) UnitAngle {
	return UnitAngle{
		Dimension: DimensionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UnitAngleClass) Alloc() UnitAngle {
	rv := objc.Send[UnitAngle](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UnitAngleClass) New() UnitAngle {
	rv := objc.Send[UnitAngle](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitAngle) Init() UnitAngle {
	rv := objc.Send[UnitAngle](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitAngle) Autorelease() UnitAngle {
	rv := objc.Send[UnitAngle](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitAngle creates a new UnitAngle instance.
func NewUnitAngle() UnitAngle {
	return getUnitAngleClass().New()
}



// The arc seconds unit of angle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitAngle/arcSeconds

func (uc _UnitAngleClass) ArcSeconds() UnitAngle {
	rv := objc.Send[NSUnitAngle](objc.ID(uc.class), objc.Sel("arcSeconds"))
	return rv
}

// The degrees unit of angle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitAngle/degrees

func (uc _UnitAngleClass) Degrees() UnitAngle {
	rv := objc.Send[NSUnitAngle](objc.ID(uc.class), objc.Sel("degrees"))
	return rv
}

// The arc seconds unit of angle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitAngle/arcSeconds

func (u_ UnitAngle) ArcSeconds() NSUnitAngle {
	rv := objc.Send[NSUnitAngle](u_.ID, objc.Sel("arcSeconds"))
	return rv
}


// The degrees unit of angle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitAngle/degrees

func (u_ UnitAngle) Degrees() NSUnitAngle {
	rv := objc.Send[NSUnitAngle](u_.ID, objc.Sel("degrees"))
	return rv
}




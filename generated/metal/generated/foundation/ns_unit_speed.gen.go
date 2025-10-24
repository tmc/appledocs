// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitSpeed] class.
var (
	UnitSpeedClass     _UnitSpeedClass
	UnitSpeedClassOnce sync.Once
)

func getUnitSpeedClass() _UnitSpeedClass {
	UnitSpeedClassOnce.Do(func() {
		UnitSpeedClass = _UnitSpeedClass{objc.GetClass("NSUnitSpeed")}
	})
	return UnitSpeedClass
}

type _UnitSpeedClass struct {
	class objc.Class
}

// An interface definition for the [UnitSpeed] class.
type IUnitSpeed interface {
	IDimension
	// properties:
	// methods:
}

// A unit of measure for speed.
//
// You typically use instances of to represent specific quantities of speed using the class.


// A unit of measure for speed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitSpeed
type UnitSpeed struct {
	Dimension
}

// UnitSpeedFrom constructs a [UnitSpeed] from an unsafe.Pointer.
//
// A unit of measure for speed.
func UnitSpeedFrom(ptr unsafe.Pointer) UnitSpeed {
	return UnitSpeed{
		Dimension: DimensionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UnitSpeedClass) Alloc() UnitSpeed {
	rv := objc.Send[UnitSpeed](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UnitSpeedClass) New() UnitSpeed {
	rv := objc.Send[UnitSpeed](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitSpeed) Init() UnitSpeed {
	rv := objc.Send[UnitSpeed](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitSpeed) Autorelease() UnitSpeed {
	rv := objc.Send[UnitSpeed](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitSpeed creates a new UnitSpeed instance.
func NewUnitSpeed() UnitSpeed {
	return getUnitSpeedClass().New()
}





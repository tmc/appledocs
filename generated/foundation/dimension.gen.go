// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Dimension] class.
var (
	dimensionClass     _DimensionClass
	dimensionClassOnce sync.Once
)

func getDimensionClass() _DimensionClass {
	dimensionClassOnce.Do(func() {
		dimensionClass = _DimensionClass{objc.GetClass("NSDimension")}
	})
	return dimensionClass
}

type _DimensionClass struct {
	class objc.Class
}

// An interface definition for the [Dimension] class.
type IDimension interface {
	IUnit
}

// An abstract class representing a dimensional unit of measure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Dimension
type Dimension struct {
	Unit
}

// DimensionFrom constructs a [Dimension] from an unsafe.Pointer.
//
// An abstract class representing a dimensional unit of measure.
func DimensionFrom(ptr unsafe.Pointer) Dimension {
	return Dimension{
		Unit: UnitFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DimensionClass) Alloc() Dimension {
	rv := objc.Send[Dimension](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DimensionClass) New() Dimension {
	rv := objc.Send[Dimension](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ Dimension) Init() Dimension {
	rv := objc.Send[Dimension](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ Dimension) Autorelease() Dimension {
	rv := objc.Send[Dimension](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDimension creates a new Dimension instance.
func NewDimension() Dimension {
	return getDimensionClass().New()
}





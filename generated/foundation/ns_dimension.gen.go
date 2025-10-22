// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Dimension] class.
var (
	DimensionClass     _DimensionClass
	DimensionClassOnce sync.Once
)

func getDimensionClass() _DimensionClass {
	DimensionClassOnce.Do(func() {
		DimensionClass = _DimensionClass{objc.GetClass("NSDimension")}
	})
	return DimensionClass
}

type _DimensionClass struct {
	class objc.Class
}

// An interface definition for the [Dimension] class.
type IDimension interface {
	IUnit
	Converter() NSUnitConverter
	Coefficient() float64
	SetCoefficient(value float64)
}

// An abstract class representing a dimensional unit of measure.
//
// The Foundation framework provides concrete subclasses for many of the most common types of physical units. Table 1: subclasses. Each instance of a subclass has a , which represents the unit in terms of the dimension’s . For example, the class uses as its base unit. The system defines the predefined unit by a with a of , which corresponds to the conversion ratio of miles to meters (1 mi = 1609.34 m); the system defines the predefined unit by a with a of because it’s the base unit. You typically use an subclass in conjunction with the class to represent specific quantities of a particular unit.


// An abstract class representing a dimensional unit of measure.
//
// [Full Topic]
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




// Initializes a dimensional unit with the symbol and unit converter you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Dimension/init(symbol:converter:)

func NewDimensionWithSymbolConverter(symbol string, converter IUnitConverter) Dimension {
	instance := getDimensionClass().Alloc()
	rv := objc.Send[Dimension](instance.ID, objc.Sel("initWithSymbol:converter:"), objc.String(symbol), converter)
	rv.Autorelease()
	return rv
}



// Returns the base unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Dimension/baseUnit()

func (dc _DimensionClass) BaseUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("baseUnit"))
	return rv
}


// The unit converter that represents the unit in terms of the dimension’s base unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Dimension/converter

func (d_ Dimension) Converter() NSUnitConverter {
	rv := objc.Send[NSUnitConverter](d_.ID, objc.Sel("converter"))
	return rv
}


// The coefficient to use in the linear unit conversion calculation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/unitconverterlinear/coefficient

func (d_ Dimension) Coefficient() float64 {
	rv := objc.Send[float64](d_.ID, objc.Sel("coefficient"))
	return rv
}


// The coefficient to use in the linear unit conversion calculation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/unitconverterlinear/coefficient

func (d_ Dimension) SetCoefficient(value float64) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCoefficient:"), value)
}



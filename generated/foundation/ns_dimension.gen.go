// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSDimension */


/* debug [class_header]: Header for NSDimension */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Dimension */
// An interface definition for the [Dimension] class.
type IDimension interface {
	IUnit
	
/* debug [class_interface_properties]: Properties for Dimension */
	// properties:
	Converter() IUnitConverter
	Coefficient() float64
	SetCoefficient(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Dimension */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Dimension */
// Alloc allocates a new instance without initialization.
func (dc _DimensionClass) Alloc() Dimension {
	rv := objc.Send[Dimension](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Dimension */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Dimension */

// Initializes a dimensional unit with the symbol and unit converter you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Dimension/init(symbol:converter:)
func NewDimensionWithSymbolConverter(symbol IString, converter IUnitConverter) Dimension {
	instance := getDimensionClass().Alloc()
	rv := objc.Send[Dimension](instance.ID, objc.Sel("initWithSymbol:converter:"), symbol, converter)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDimensionWithSymbolConverter */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Dimension */

// Returns the base unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Dimension/baseUnit()
func (dc _DimensionClass) BaseUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(dc.class), objc.Sel("baseUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=BaseUnit) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Dimension */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Dimension */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Dimension */

// The unit converter that represents the unit in terms of the dimension’s base unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Dimension/converter
func (d_ Dimension) Converter() IUnitConverter {
	rv := objc.Send[UnitConverter](d_.ID, objc.Sel("converter"))
	return rv
}/* debug [instance_properties/getter]: converter */


// The coefficient to use in the linear unit conversion calculation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/unitconverterlinear/coefficient
func (d_ Dimension) Coefficient() float64 {
	rv := objc.Send[float64](d_.ID, objc.Sel("coefficient"))
	return rv
}/* debug [instance_properties/getter]: coefficient */


// The coefficient to use in the linear unit conversion calculation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/unitconverterlinear/coefficient
func (d_ Dimension) SetCoefficient(value float64) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCoefficient:"), value)
}/* debug [instance_properties/setter]: coefficient */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSDimension */



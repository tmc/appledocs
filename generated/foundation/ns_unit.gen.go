// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSUnit */


/* debug [class_header]: Header for NSUnit */
// The class instance for the [Unit] class.
var (
	UnitClass     _UnitClass
	UnitClassOnce sync.Once
)

func getUnitClass() _UnitClass {
	UnitClassOnce.Do(func() {
		UnitClass = _UnitClass{objc.GetClass("NSUnit")}
	})
	return UnitClass
}

type _UnitClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Unit */
// An interface definition for the [Unit] class.
type IUnit interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Unit */
	// properties:
	Symbol() IString
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Unit */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Unit */
// Alloc allocates a new instance without initialization.
func (uc _UnitClass) Alloc() Unit {
	rv := objc.Send[Unit](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UnitClass) New() Unit {
	rv := objc.Send[Unit](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ Unit) Init() Unit {
	rv := objc.Send[Unit](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ Unit) Autorelease() Unit {
	rv := objc.Send[Unit](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnit creates a new Unit instance.
func NewUnit() Unit {
	return getUnitClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Unit */
// An abstract class representing a unit of measure.
//
// Each instance of an subclass consists of a , which can be used to create string representations of objects with the class. The subclass is an abstract class that represents a dimensional unit, which can be converted into different units of the same type. The Foundation framework provides several concrete subclasses to represent the most common physical quantities, including mass, length, duration, and speed.


// An abstract class representing a unit of measure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Unit
type Unit struct {
	objectivec.Object
}

// UnitFrom constructs a [Unit] from an unsafe.Pointer.
//
// An abstract class representing a unit of measure.
func UnitFrom(ptr unsafe.Pointer) Unit {
	return Unit{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Unit */

// Initializes a new unit with the specified symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Unit/init(symbol:)
func NewUnitWithSymbol(symbol IString) Unit {
	instance := getUnitClass().Alloc()
	rv := objc.Send[Unit](instance.ID, objc.Sel("initWithSymbol:"), symbol)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewUnitWithSymbol */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Unit */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Unit */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Unit */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Unit */

// The symbolic representation of the unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Unit/symbol
func (u_ Unit) Symbol() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("symbol"))
	return rv
}/* debug [instance_properties/getter]: symbol */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUnit */



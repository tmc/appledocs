// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSUnitArea */


/* debug [class_header]: Header for NSUnitArea */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UnitArea */
// An interface definition for the [UnitArea] class.
type IUnitArea interface {
	IDimension
	
/* debug [class_interface_properties]: Properties for UnitArea */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UnitArea */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UnitArea */
// Alloc allocates a new instance without initialization.
func (uc _UnitAreaClass) Alloc() UnitArea {
	rv := objc.Send[UnitArea](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UnitArea */
// A unit of measure for area.
//
// You typically use instances of to represent specific quantities of area using the class.


// A unit of measure for area.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UnitArea *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UnitArea */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UnitArea */

// The acres unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/acres
func (uc _UnitAreaClass) Acres() UnitArea {
	rv := objc.Send[UnitArea](objc.ID(uc.class), objc.Sel("acres"))
	return rv
}/* debug [class_properties_class/property]: acres */

// The ares unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/ares
func (uc _UnitAreaClass) Ares() UnitArea {
	rv := objc.Send[UnitArea](objc.ID(uc.class), objc.Sel("ares"))
	return rv
}/* debug [class_properties_class/property]: ares */

// The hectares unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/hectares
func (uc _UnitAreaClass) Hectares() UnitArea {
	rv := objc.Send[UnitArea](objc.ID(uc.class), objc.Sel("hectares"))
	return rv
}/* debug [class_properties_class/property]: hectares */

// The square miles unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareMiles
func (uc _UnitAreaClass) SquareMiles() UnitArea {
	rv := objc.Send[UnitArea](objc.ID(uc.class), objc.Sel("squareMiles"))
	return rv
}/* debug [class_properties_class/property]: squareMiles */

// The square yards unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareYards
func (uc _UnitAreaClass) SquareYards() UnitArea {
	rv := objc.Send[UnitArea](objc.ID(uc.class), objc.Sel("squareYards"))
	return rv
}/* debug [class_properties_class/property]: squareYards */

// The square centimeters unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareCentimeters
func (uc _UnitAreaClass) SquareCentimeters() UnitArea {
	rv := objc.Send[UnitArea](objc.ID(uc.class), objc.Sel("squareCentimeters"))
	return rv
}/* debug [class_properties_class/property]: squareCentimeters */

// The square feet unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareFeet
func (uc _UnitAreaClass) SquareFeet() UnitArea {
	rv := objc.Send[UnitArea](objc.ID(uc.class), objc.Sel("squareFeet"))
	return rv
}/* debug [class_properties_class/property]: squareFeet */

// The square inches unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareInches
func (uc _UnitAreaClass) SquareInches() UnitArea {
	rv := objc.Send[UnitArea](objc.ID(uc.class), objc.Sel("squareInches"))
	return rv
}/* debug [class_properties_class/property]: squareInches */

// The square kilometers unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareKilometers
func (uc _UnitAreaClass) SquareKilometers() UnitArea {
	rv := objc.Send[UnitArea](objc.ID(uc.class), objc.Sel("squareKilometers"))
	return rv
}/* debug [class_properties_class/property]: squareKilometers */

// The square megameters unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareMegameters
func (uc _UnitAreaClass) SquareMegameters() UnitArea {
	rv := objc.Send[UnitArea](objc.ID(uc.class), objc.Sel("squareMegameters"))
	return rv
}/* debug [class_properties_class/property]: squareMegameters */

// The square meters unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareMeters
func (uc _UnitAreaClass) SquareMeters() UnitArea {
	rv := objc.Send[UnitArea](objc.ID(uc.class), objc.Sel("squareMeters"))
	return rv
}/* debug [class_properties_class/property]: squareMeters */

// The square micrometers unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareMicrometers
func (uc _UnitAreaClass) SquareMicrometers() UnitArea {
	rv := objc.Send[UnitArea](objc.ID(uc.class), objc.Sel("squareMicrometers"))
	return rv
}/* debug [class_properties_class/property]: squareMicrometers */

// The square millimeters unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareMillimeters
func (uc _UnitAreaClass) SquareMillimeters() UnitArea {
	rv := objc.Send[UnitArea](objc.ID(uc.class), objc.Sel("squareMillimeters"))
	return rv
}/* debug [class_properties_class/property]: squareMillimeters */

// The square nanometers unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareNanometers
func (uc _UnitAreaClass) SquareNanometers() UnitArea {
	rv := objc.Send[UnitArea](objc.ID(uc.class), objc.Sel("squareNanometers"))
	return rv
}/* debug [class_properties_class/property]: squareNanometers */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UnitArea */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UnitArea */

// The acres unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/acres
func (u_ UnitArea) Acres() IUnitArea {
	rv := objc.Send[UnitArea](u_.ID, objc.Sel("acres"))
	return rv
}/* debug [instance_properties/getter]: acres */


// The ares unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/ares
func (u_ UnitArea) Ares() IUnitArea {
	rv := objc.Send[UnitArea](u_.ID, objc.Sel("ares"))
	return rv
}/* debug [instance_properties/getter]: ares */


// The hectares unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/hectares
func (u_ UnitArea) Hectares() IUnitArea {
	rv := objc.Send[UnitArea](u_.ID, objc.Sel("hectares"))
	return rv
}/* debug [instance_properties/getter]: hectares */


// The square miles unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareMiles
func (u_ UnitArea) SquareMiles() IUnitArea {
	rv := objc.Send[UnitArea](u_.ID, objc.Sel("squareMiles"))
	return rv
}/* debug [instance_properties/getter]: squareMiles */


// The square yards unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareYards
func (u_ UnitArea) SquareYards() IUnitArea {
	rv := objc.Send[UnitArea](u_.ID, objc.Sel("squareYards"))
	return rv
}/* debug [instance_properties/getter]: squareYards */


// The square centimeters unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareCentimeters
func (u_ UnitArea) SquareCentimeters() IUnitArea {
	rv := objc.Send[UnitArea](u_.ID, objc.Sel("squareCentimeters"))
	return rv
}/* debug [instance_properties/getter]: squareCentimeters */


// The square feet unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareFeet
func (u_ UnitArea) SquareFeet() IUnitArea {
	rv := objc.Send[UnitArea](u_.ID, objc.Sel("squareFeet"))
	return rv
}/* debug [instance_properties/getter]: squareFeet */


// The square inches unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareInches
func (u_ UnitArea) SquareInches() IUnitArea {
	rv := objc.Send[UnitArea](u_.ID, objc.Sel("squareInches"))
	return rv
}/* debug [instance_properties/getter]: squareInches */


// The square kilometers unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareKilometers
func (u_ UnitArea) SquareKilometers() IUnitArea {
	rv := objc.Send[UnitArea](u_.ID, objc.Sel("squareKilometers"))
	return rv
}/* debug [instance_properties/getter]: squareKilometers */


// The square megameters unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareMegameters
func (u_ UnitArea) SquareMegameters() IUnitArea {
	rv := objc.Send[UnitArea](u_.ID, objc.Sel("squareMegameters"))
	return rv
}/* debug [instance_properties/getter]: squareMegameters */


// The square meters unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareMeters
func (u_ UnitArea) SquareMeters() IUnitArea {
	rv := objc.Send[UnitArea](u_.ID, objc.Sel("squareMeters"))
	return rv
}/* debug [instance_properties/getter]: squareMeters */


// The square micrometers unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareMicrometers
func (u_ UnitArea) SquareMicrometers() IUnitArea {
	rv := objc.Send[UnitArea](u_.ID, objc.Sel("squareMicrometers"))
	return rv
}/* debug [instance_properties/getter]: squareMicrometers */


// The square millimeters unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareMillimeters
func (u_ UnitArea) SquareMillimeters() IUnitArea {
	rv := objc.Send[UnitArea](u_.ID, objc.Sel("squareMillimeters"))
	return rv
}/* debug [instance_properties/getter]: squareMillimeters */


// The square nanometers unit of area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitArea/squareNanometers
func (u_ UnitArea) SquareNanometers() IUnitArea {
	rv := objc.Send[UnitArea](u_.ID, objc.Sel("squareNanometers"))
	return rv
}/* debug [instance_properties/getter]: squareNanometers */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUnitArea */




// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSUnitLength */


/* debug [class_header]: Header for NSUnitLength */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UnitLength */
// An interface definition for the [UnitLength] class.
type IUnitLength interface {
	IDimension
	
/* debug [class_interface_properties]: Properties for UnitLength */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UnitLength */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UnitLength */
// Alloc allocates a new instance without initialization.
func (uc _UnitLengthClass) Alloc() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UnitLength */
// A unit of measure for length.
//
// You typically use instances of to represent specific quantities of length using the class.


// A unit of measure for length.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UnitLength *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UnitLength */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UnitLength */

// The centimeters unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/centimeters
func (uc _UnitLengthClass) Centimeters() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("centimeters"))
	return rv
}/* debug [class_properties_class/property]: centimeters */

// The decameters unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/decameters
func (uc _UnitLengthClass) Decameters() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("decameters"))
	return rv
}/* debug [class_properties_class/property]: decameters */

// The decimeters unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/decimeters
func (uc _UnitLengthClass) Decimeters() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("decimeters"))
	return rv
}/* debug [class_properties_class/property]: decimeters */

// The feet unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/feet
func (uc _UnitLengthClass) Feet() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("feet"))
	return rv
}/* debug [class_properties_class/property]: feet */

// The hectometers unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/hectometers
func (uc _UnitLengthClass) Hectometers() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("hectometers"))
	return rv
}/* debug [class_properties_class/property]: hectometers */

// The inches unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/inches
func (uc _UnitLengthClass) Inches() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("inches"))
	return rv
}/* debug [class_properties_class/property]: inches */

// The kilometers unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/kilometers
func (uc _UnitLengthClass) Kilometers() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("kilometers"))
	return rv
}/* debug [class_properties_class/property]: kilometers */

// The megameters unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/megameters
func (uc _UnitLengthClass) Megameters() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("megameters"))
	return rv
}/* debug [class_properties_class/property]: megameters */

// The meters unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/meters
func (uc _UnitLengthClass) Meters() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("meters"))
	return rv
}/* debug [class_properties_class/property]: meters */

// The micrometers unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/micrometers
func (uc _UnitLengthClass) Micrometers() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("micrometers"))
	return rv
}/* debug [class_properties_class/property]: micrometers */

// The miles unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/miles
func (uc _UnitLengthClass) Miles() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("miles"))
	return rv
}/* debug [class_properties_class/property]: miles */

// The millimeters unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/millimeters
func (uc _UnitLengthClass) Millimeters() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("millimeters"))
	return rv
}/* debug [class_properties_class/property]: millimeters */

// The nanometers unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/nanometers
func (uc _UnitLengthClass) Nanometers() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("nanometers"))
	return rv
}/* debug [class_properties_class/property]: nanometers */

// The picometers unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/picometers
func (uc _UnitLengthClass) Picometers() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("picometers"))
	return rv
}/* debug [class_properties_class/property]: picometers */

// The Scandinavian miles unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/scandinavianMiles
func (uc _UnitLengthClass) ScandinavianMiles() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("scandinavianMiles"))
	return rv
}/* debug [class_properties_class/property]: scandinavianMiles */

// The yards unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/yards
func (uc _UnitLengthClass) Yards() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("yards"))
	return rv
}/* debug [class_properties_class/property]: yards */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UnitLength */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UnitLength */

// The centimeters unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/centimeters
func (u_ UnitLength) Centimeters() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("centimeters"))
	return rv
}/* debug [instance_properties/getter]: centimeters */


// The decameters unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/decameters
func (u_ UnitLength) Decameters() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("decameters"))
	return rv
}/* debug [instance_properties/getter]: decameters */


// The decimeters unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/decimeters
func (u_ UnitLength) Decimeters() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("decimeters"))
	return rv
}/* debug [instance_properties/getter]: decimeters */


// The feet unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/feet
func (u_ UnitLength) Feet() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("feet"))
	return rv
}/* debug [instance_properties/getter]: feet */


// The hectometers unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/hectometers
func (u_ UnitLength) Hectometers() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("hectometers"))
	return rv
}/* debug [instance_properties/getter]: hectometers */


// The inches unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/inches
func (u_ UnitLength) Inches() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("inches"))
	return rv
}/* debug [instance_properties/getter]: inches */


// The kilometers unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/kilometers
func (u_ UnitLength) Kilometers() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("kilometers"))
	return rv
}/* debug [instance_properties/getter]: kilometers */


// The megameters unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/megameters
func (u_ UnitLength) Megameters() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("megameters"))
	return rv
}/* debug [instance_properties/getter]: megameters */


// The meters unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/meters
func (u_ UnitLength) Meters() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("meters"))
	return rv
}/* debug [instance_properties/getter]: meters */


// The micrometers unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/micrometers
func (u_ UnitLength) Micrometers() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("micrometers"))
	return rv
}/* debug [instance_properties/getter]: micrometers */


// The miles unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/miles
func (u_ UnitLength) Miles() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("miles"))
	return rv
}/* debug [instance_properties/getter]: miles */


// The millimeters unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/millimeters
func (u_ UnitLength) Millimeters() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("millimeters"))
	return rv
}/* debug [instance_properties/getter]: millimeters */


// The nanometers unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/nanometers
func (u_ UnitLength) Nanometers() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("nanometers"))
	return rv
}/* debug [instance_properties/getter]: nanometers */


// The picometers unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/picometers
func (u_ UnitLength) Picometers() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("picometers"))
	return rv
}/* debug [instance_properties/getter]: picometers */


// The Scandinavian miles unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/scandinavianMiles
func (u_ UnitLength) ScandinavianMiles() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("scandinavianMiles"))
	return rv
}/* debug [instance_properties/getter]: scandinavianMiles */


// The yards unit of length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitLength/yards
func (u_ UnitLength) Yards() IUnitLength {
	rv := objc.Send[UnitLength](u_.ID, objc.Sel("yards"))
	return rv
}/* debug [instance_properties/getter]: yards */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUnitLength */




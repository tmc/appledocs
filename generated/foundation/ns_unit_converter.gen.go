// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [UnitConverter] class.
var (
	UnitConverterClass     _UnitConverterClass
	UnitConverterClassOnce sync.Once
)

func getUnitConverterClass() _UnitConverterClass {
	UnitConverterClassOnce.Do(func() {
		UnitConverterClass = _UnitConverterClass{objc.GetClass("NSUnitConverter")}
	})
	return UnitConverterClass
}

type _UnitConverterClass struct {
	class objc.Class
}





// An interface definition for the [UnitConverter] class.
type IUnitConverter interface {
	objectivec.IObject
	

	// properties:


	

	// methods:
	BaseUnitValueFromValue(value float64) float64
	ValueFromBaseUnitValue(baseUnitValue float64) float64


}





// Alloc allocates a new instance without initialization.
func (uc _UnitConverterClass) Alloc() UnitConverter {
	rv := objc.Send[UnitConverter](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UnitConverterClass) New() UnitConverter {
	rv := objc.Send[UnitConverter](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitConverter) Init() UnitConverter {
	rv := objc.Send[UnitConverter](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitConverter) Autorelease() UnitConverter {
	rv := objc.Send[UnitConverter](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitConverter creates a new UnitConverter instance.
func NewUnitConverter() UnitConverter {
	return getUnitConverterClass().New()
}





// An abstract class that provides a description of how to convert a unit to and from the base unit of its dimension.
//
// For units that can be converted by a scale factor or linear equation, use the concrete subclass .


// An abstract class that provides a description of how to convert a unit to and from the base unit of its dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitConverter
type UnitConverter struct {
	objectivec.Object
}

// UnitConverterFrom constructs a [UnitConverter] from an unsafe.Pointer.
//
// An abstract class that provides a description of how to convert a unit to and from the base unit of its dimension.
func UnitConverterFrom(ptr unsafe.Pointer) UnitConverter {
	return UnitConverter{objectivec.Object{objc.ID(ptr)}}
}




















// For a given unit, returns the specified value of that unit in terms of the base unit of its dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitConverter/baseUnitValue(fromValue:)
func (u_ UnitConverter) BaseUnitValueFromValue(value float64) float64 {
	rv := objc.Send[float64](u_.ID, objc.Sel("baseUnitValueFromValue:"), value)
	return rv
}


// For a given unit, returns the specified value of the base unit in terms of that unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitConverter/value(fromBaseUnitValue:)
func (u_ UnitConverter) ValueFromBaseUnitValue(baseUnitValue float64) float64 {
	rv := objc.Send[float64](u_.ID, objc.Sel("valueFromBaseUnitValue:"), baseUnitValue)
	return rv
}














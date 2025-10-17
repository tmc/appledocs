// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UnitConverter] class.
var unitConverterClass = _UnitConverterClass{objc.GetClass("NSUnitConverter")}

type _UnitConverterClass struct {
	class objc.Class
}

// An interface definition for the [UnitConverter] class.
type IUnitConverter interface {
	objectivec.IObject
}

// An abstract class that provides a description of how to convert a unit to and from the base unit of its dimension. [Full Topic]
//
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




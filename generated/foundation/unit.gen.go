// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Unit] class.
var unitClass = _UnitClass{objc.GetClass("NSUnit")}

type _UnitClass struct {
	class objc.Class
}

// An interface definition for the [Unit] class.
type IUnit interface {
	objectivec.IObject
}

// An abstract class representing a unit of measure. [Full Topic]
//
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




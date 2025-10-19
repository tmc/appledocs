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
// Alloc allocates a new instance without initialization.
func (uc _UnitClass) Alloc() Unit {
	rv := objc.Send[Unit](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
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
	return unitClass.New()
}





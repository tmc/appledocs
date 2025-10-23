// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [Unit] class.
type IUnit interface {
	objectivec.IObject
	// properties:
	Symbol() IString
	SetSymbol(value IString)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (uc _UnitClass) Alloc() Unit {
	rv := objc.Send[Unit](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The symbolic representation of the unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/unit/symbol
func (u_ Unit) Symbol() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("symbol"))
	return rv
}


// The symbolic representation of the unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/unit/symbol
func (u_ Unit) SetSymbol(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSymbol:"), value)
}




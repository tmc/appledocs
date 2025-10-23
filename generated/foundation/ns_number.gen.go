// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Number] class.
var (
	NumberClass     _NumberClass
	NumberClassOnce sync.Once
)

func getNumberClass() _NumberClass {
	NumberClassOnce.Do(func() {
		NumberClass = _NumberClass{objc.GetClass("NSNumber")}
	})
	return NumberClass
}

type _NumberClass struct {
	class objc.Class
}

// An interface definition for the [Number] class.
type INumber interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other Foundation classes.


// A parent class referenced by other Foundation classes. [Full Topic]
type Number struct {
	objectivec.Object
}

// NumberFrom constructs a [Number] from an unsafe.Pointer.
//
// A parent class referenced by other Foundation classes.
func NumberFrom(ptr unsafe.Pointer) Number {
	return Number{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NumberClass) Alloc() Number {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NumberClass) New() Number {
	rv := objc.Send[Number](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ Number) Init() Number {
	rv := objc.Send[Number](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ Number) Autorelease() Number {
	rv := objc.Send[Number](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNumber creates a new Number instance.
func NewNumber() Number {
	return getNumberClass().New()
}





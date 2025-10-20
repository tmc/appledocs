// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Number] class.
var (
	numberClass     _NumberClass
	numberClassOnce sync.Once
)

func getNumberClass() _NumberClass {
	numberClassOnce.Do(func() {
		numberClass = _NumberClass{objc.GetClass("NSNumber")}
	})
	return numberClass
}

type _NumberClass struct {
	class objc.Class
}

// An interface definition for the [Number] class.
type INumber interface {
	IValue
	DescriptionWithLocale(locale objc.ID) unsafe.Pointer
}

// An object wrapper for primitive scalar numeric values.
//
// is a subclass of that offers a value as any C scalar (numeric) type. It defines a set of methods specifically for setting and accessing the value as a signed or unsigned , , , , , , or or as a . (Note that number objects do not necessarily preserve the type they are created with.) It also defines a method to determine the ordering of two objects. is “toll-free bridged” with its Core Foundation counterparts: for integer and floating point values, and for Boolean values. See for more information on toll-free bridging.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber
type Number struct {
	Value
}

// NumberFrom constructs a [Number] from an unsafe.Pointer.
//
// An object wrapper for primitive scalar numeric values.
func NumberFrom(ptr unsafe.Pointer) Number {
	return Number{
		Value: ValueFrom(ptr),
	}
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


// Returns a string that represents the contents of the number object for a given locale.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/description(withLocale:)
func (n_ Number) DescriptionWithLocale(locale objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("descriptionWithLocale:"), locale)
	return rv
}

// The number object’s value expressed as a Boolean value.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNumber/boolValue
func (n_ Number) BoolValue() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("boolValue"))
	return rv
}




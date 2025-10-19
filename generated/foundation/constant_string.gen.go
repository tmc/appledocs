// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ConstantString] class.
var (
	constantStringClass     _ConstantStringClass
	constantStringClassOnce sync.Once
)

func getConstantStringClass() _ConstantStringClass {
	constantStringClassOnce.Do(func() {
		constantStringClass = _ConstantStringClass{objc.GetClass("NSConstantString")}
	})
	return constantStringClass
}

type _ConstantStringClass struct {
	class objc.Class
}

// An interface definition for the [ConstantString] class.
type IConstantString interface {
	ISimpleCString
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConstantString
type ConstantString struct {
	SimpleCString
}

// ConstantStringFrom constructs a [ConstantString] from an unsafe.Pointer.
func ConstantStringFrom(ptr unsafe.Pointer) ConstantString {
	return ConstantString{
		SimpleCString: SimpleCStringFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _ConstantStringClass) Alloc() ConstantString {
	rv := objc.Send[ConstantString](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ConstantStringClass) New() ConstantString {
	rv := objc.Send[ConstantString](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ConstantString) Init() ConstantString {
	rv := objc.Send[ConstantString](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ConstantString) Autorelease() ConstantString {
	rv := objc.Send[ConstantString](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewConstantString creates a new ConstantString instance.
func NewConstantString() ConstantString {
	return getConstantStringClass().New()
}





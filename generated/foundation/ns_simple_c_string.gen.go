// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SimpleCString] class.
var (
	SimpleCStringClass     _SimpleCStringClass
	SimpleCStringClassOnce sync.Once
)

func getSimpleCStringClass() _SimpleCStringClass {
	SimpleCStringClassOnce.Do(func() {
		SimpleCStringClass = _SimpleCStringClass{objc.GetClass("NSSimpleCString")}
	})
	return SimpleCStringClass
}

type _SimpleCStringClass struct {
	class objc.Class
}

// An interface definition for the [SimpleCString] class.
type ISimpleCString interface {
	IString
}



//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSimpleCString

type SimpleCString struct {
	string
}

// SimpleCStringFrom constructs a [SimpleCString] from an unsafe.Pointer.
func SimpleCStringFrom(ptr unsafe.Pointer) SimpleCString {
	return SimpleCString{
		String: stringFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SimpleCStringClass) Alloc() SimpleCString {
	rv := objc.Send[SimpleCString](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SimpleCStringClass) New() SimpleCString {
	rv := objc.Send[SimpleCString](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SimpleCString) Init() SimpleCString {
	rv := objc.Send[SimpleCString](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SimpleCString) Autorelease() SimpleCString {
	rv := objc.Send[SimpleCString](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSimpleCString creates a new SimpleCString instance.
func NewSimpleCString() SimpleCString {
	return getSimpleCStringClass().New()
}





// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Null] class.
var (
	NullClass     _NullClass
	NullClassOnce sync.Once
)

func getNullClass() _NullClass {
	NullClassOnce.Do(func() {
		NullClass = _NullClass{objc.GetClass("NSNull")}
	})
	return NullClass
}

type _NullClass struct {
	class objc.Class
}

// An interface definition for the [Null] class.
type INull interface {
	objectivec.IObject
}

// A singleton object used to represent null values in collection objects that don’t allow values.
//
// is “toll-free bridged” with its Core Foundation counterpart, . See for more information on toll-free bridging.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNull
type Null struct {
	objectivec.Object
}

// NullFrom constructs a [Null] from an unsafe.Pointer.
//
// A singleton object used to represent null values in collection objects that don’t allow values.
func NullFrom(ptr unsafe.Pointer) Null {
	return Null{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NullClass) Alloc() Null {
	rv := objc.Send[Null](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NullClass) New() Null {
	rv := objc.Send[Null](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ Null) Init() Null {
	rv := objc.Send[Null](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ Null) Autorelease() Null {
	rv := objc.Send[Null](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNull creates a new Null instance.
func NewNull() Null {
	return getNullClass().New()
}


// Returns the singleton instance of .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNull/null
func (nc _NullClass) Null() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("null"))
	return rv
}




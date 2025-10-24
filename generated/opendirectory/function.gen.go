// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [function] class.
var (
	FunctionClass     _functionClass
	FunctionClassOnce sync.Once
)

func getfunctionClass() _functionClass {
	FunctionClassOnce.Do(func() {
		FunctionClass = _functionClass{objc.GetClass("function")}
	})
	return FunctionClass
}

type _functionClass struct {
	class objc.Class
}

// An interface definition for the [function] class.
type Ifunction interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/function-c.ivar
type function struct {
	objectivec.Object
}

// functionFrom constructs a [function] from an unsafe.Pointer.
func functionFrom(ptr unsafe.Pointer) function {
	return function{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _functionClass) Alloc() function {
	rv := objc.Send[function](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _functionClass) New() function {
	rv := objc.Send[function](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ function) Init() function {
	rv := objc.Send[function](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ function) Autorelease() function {
	rv := objc.Send[function](f_.ID, objc.Sel("autorelease"))
	return rv
}

// Newfunction creates a new function instance.
func Newfunction() function {
	return getfunctionClass().New()
}





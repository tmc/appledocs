// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Invocation] class.
var (
	invocationClass     _InvocationClass
	invocationClassOnce sync.Once
)

func getInvocationClass() _InvocationClass {
	invocationClassOnce.Do(func() {
		invocationClass = _InvocationClass{objc.GetClass("NSInvocation")}
	})
	return invocationClass
}

type _InvocationClass struct {
	class objc.Class
}

// An interface definition for the [Invocation] class.
type IInvocation interface {
	objectivec.IObject
}

// An Objective-C message rendered as an object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInvocation
type Invocation struct {
	objectivec.Object
}

// InvocationFrom constructs a [Invocation] from an unsafe.Pointer.
//
// An Objective-C message rendered as an object.
func InvocationFrom(ptr unsafe.Pointer) Invocation {
	return Invocation{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _InvocationClass) Alloc() Invocation {
	rv := objc.Send[Invocation](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _InvocationClass) New() Invocation {
	rv := objc.Send[Invocation](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ Invocation) Init() Invocation {
	rv := objc.Send[Invocation](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ Invocation) Autorelease() Invocation {
	rv := objc.Send[Invocation](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInvocation creates a new Invocation instance.
func NewInvocation() Invocation {
	return getInvocationClass().New()
}





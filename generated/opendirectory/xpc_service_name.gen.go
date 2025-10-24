// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [xpcServiceName] class.
var (
	XpcServiceNameClass     _xpcServiceNameClass
	XpcServiceNameClassOnce sync.Once
)

func getxpcServiceNameClass() _xpcServiceNameClass {
	XpcServiceNameClassOnce.Do(func() {
		XpcServiceNameClass = _xpcServiceNameClass{objc.GetClass("xpcServiceName")}
	})
	return XpcServiceNameClass
}

type _xpcServiceNameClass struct {
	class objc.Class
}

// An interface definition for the [xpcServiceName] class.
type IxpcServiceName interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/xpcServiceName-c.ivar
type xpcServiceName struct {
	objectivec.Object
}

// xpcServiceNameFrom constructs a [xpcServiceName] from an unsafe.Pointer.
func xpcServiceNameFrom(ptr unsafe.Pointer) xpcServiceName {
	return xpcServiceName{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (xc _xpcServiceNameClass) Alloc() xpcServiceName {
	rv := objc.Send[xpcServiceName](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (xc _xpcServiceNameClass) New() xpcServiceName {
	rv := objc.Send[xpcServiceName](objc.ID(xc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (x_ xpcServiceName) Init() xpcServiceName {
	rv := objc.Send[xpcServiceName](x_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x_ xpcServiceName) Autorelease() xpcServiceName {
	rv := objc.Send[xpcServiceName](x_.ID, objc.Sel("autorelease"))
	return rv
}

// NewxpcServiceName creates a new xpcServiceName instance.
func NewxpcServiceName() xpcServiceName {
	return getxpcServiceNameClass().New()
}





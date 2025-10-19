// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [XPCInterface] class.
var (
	xPCInterfaceClass     _XPCInterfaceClass
	xPCInterfaceClassOnce sync.Once
)

func getXPCInterfaceClass() _XPCInterfaceClass {
	xPCInterfaceClassOnce.Do(func() {
		xPCInterfaceClass = _XPCInterfaceClass{objc.GetClass("NSXPCInterface")}
	})
	return xPCInterfaceClass
}

type _XPCInterfaceClass struct {
	class objc.Class
}

// An interface definition for the [XPCInterface] class.
type IXPCInterface interface {
	objectivec.IObject
}

// An interface that may be sent to an exported object or remote object proxy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCInterface
type XPCInterface struct {
	objectivec.Object
}

// XPCInterfaceFrom constructs a [XPCInterface] from an unsafe.Pointer.
//
// An interface that may be sent to an exported object or remote object proxy.
func XPCInterfaceFrom(ptr unsafe.Pointer) XPCInterface {
	return XPCInterface{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (xc _XPCInterfaceClass) Alloc() XPCInterface {
	rv := objc.Send[XPCInterface](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (xc _XPCInterfaceClass) New() XPCInterface {
	rv := objc.Send[XPCInterface](objc.ID(xc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (x_ XPCInterface) Init() XPCInterface {
	rv := objc.Send[XPCInterface](x_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x_ XPCInterface) Autorelease() XPCInterface {
	rv := objc.Send[XPCInterface](x_.ID, objc.Sel("autorelease"))
	return rv
}

// NewXPCInterface creates a new XPCInterface instance.
func NewXPCInterface() XPCInterface {
	return getXPCInterfaceClass().New()
}





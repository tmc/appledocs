// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [XPCCoder] class.
var (
	xPCCoderClass     _XPCCoderClass
	xPCCoderClassOnce sync.Once
)

func getXPCCoderClass() _XPCCoderClass {
	xPCCoderClassOnce.Do(func() {
		xPCCoderClass = _XPCCoderClass{objc.GetClass("NSXPCCoder")}
	})
	return xPCCoderClass
}

type _XPCCoderClass struct {
	class objc.Class
}

// An interface definition for the [XPCCoder] class.
type IXPCCoder interface {
	ICoder
}

// A coder that encodes and decodes objects that your app sends over an XPC connection.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCCoder
type XPCCoder struct {
	Coder
}

// XPCCoderFrom constructs a [XPCCoder] from an unsafe.Pointer.
//
// A coder that encodes and decodes objects that your app sends over an XPC connection.
func XPCCoderFrom(ptr unsafe.Pointer) XPCCoder {
	return XPCCoder{
		Coder: CoderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (xc _XPCCoderClass) Alloc() XPCCoder {
	rv := objc.Send[XPCCoder](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (xc _XPCCoderClass) New() XPCCoder {
	rv := objc.Send[XPCCoder](objc.ID(xc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (x_ XPCCoder) Init() XPCCoder {
	rv := objc.Send[XPCCoder](x_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x_ XPCCoder) Autorelease() XPCCoder {
	rv := objc.Send[XPCCoder](x_.ID, objc.Sel("autorelease"))
	return rv
}

// NewXPCCoder creates a new XPCCoder instance.
func NewXPCCoder() XPCCoder {
	return getXPCCoderClass().New()
}





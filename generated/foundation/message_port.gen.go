// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MessagePort] class.
var messagePortClass = _MessagePortClass{objc.GetClass("NSMessagePort")}

type _MessagePortClass struct {
	class objc.Class
}

// An interface definition for the [MessagePort] class.
type IMessagePort interface {
	IPort
}

// A port that can be used as an endpoint for distributed object connections (or raw messaging). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MessagePort

type MessagePort struct {
	Port
}

// MessagePortFrom constructs a [MessagePort] from an unsafe.Pointer.
//
// A port that can be used as an endpoint for distributed object connections (or raw messaging).
func MessagePortFrom(ptr unsafe.Pointer) MessagePort {
	return MessagePort{
		Port: PortFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (mc _MessagePortClass) Alloc() MessagePort {
	rv := objc.Send[MessagePort](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (mc _MessagePortClass) New() MessagePort {
	rv := objc.Send[MessagePort](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MessagePort) Init() MessagePort {
	rv := objc.Send[MessagePort](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MessagePort) Autorelease() MessagePort {
	rv := objc.Send[MessagePort](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMessagePort creates a new MessagePort instance.
func NewMessagePort() MessagePort {
	return messagePortClass.New()
}





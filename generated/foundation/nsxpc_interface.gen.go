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
	XPCInterfaceClass     _XPCInterfaceClass
	XPCInterfaceClassOnce sync.Once
)

func getXPCInterfaceClass() _XPCInterfaceClass {
	XPCInterfaceClassOnce.Do(func() {
		XPCInterfaceClass = _XPCInterfaceClass{objc.GetClass("NSXPCInterface")}
	})
	return XPCInterfaceClass
}

type _XPCInterfaceClass struct {
	class objc.Class
}

// An interface definition for the [XPCInterface] class.
type IXPCInterface interface {
	objectivec.IObject
	ClassesForSelectorArgumentIndexOfReply(sel objc.SEL, arg uint, ofReply bool) unsafe.Pointer
	InterfaceForSelectorArgumentIndexOfReply(sel objc.SEL, arg uint, ofReply bool) unsafe.Pointer
	SetClassesForSelectorArgumentIndexOfReply(classes unsafe.Pointer, sel objc.SEL, arg uint, ofReply bool)
	SetInterfaceForSelectorArgumentIndexOfReply(ifc unsafe.Pointer, sel objc.SEL, arg uint, ofReply bool)
	SetXPCTypeForSelectorArgumentIndexOfReply(type_ unsafe.Pointer, sel objc.SEL, arg uint, ofReply bool)
}

// An interface that may be sent to an exported object or remote object proxy.
//
// This object holds all information about the interface of an exported object or remote object proxy. It describes what messages are allowed, what kinds of objects are allowed as arguments, what the signature of any reply blocks are, and information about additional proxy objects.
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




// Returns an NSXPCInterface instance for a given protocol.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCInterface/init(with:)
func NewXPCInterfaceWithProtocol(protocol unsafe.Pointer) XPCInterface {
	rv := objc.Send[XPCInterface](objc.ID(getXPCInterfaceClass().class), objc.Sel("interfaceWithProtocol:"), protocol)
	return rv
}


// Returns an NSXPCInterface instance for a given protocol.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCInterface/init(with:)
func (xc _XPCInterfaceClass) InterfaceWithProtocol(protocol unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(xc.class), objc.Sel("interfaceWithProtocol:"), protocol)
	return rv
}

// Returns the current list of allowed classes that can appear within the specified collection object argument to the specified method.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCInterface/classes(for:argumentIndex:ofReply:)
func (x_ XPCInterface) ClassesForSelectorArgumentIndexOfReply(sel objc.SEL, arg uint, ofReply bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("classesForSelector:argumentIndex:ofReply:"), sel, arg, ofReply)
	return rv
}

// Returns the interface previously set for the specified selector and parameter.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCInterface/forSelector(_:argumentIndex:ofReply:)
func (x_ XPCInterface) InterfaceForSelectorArgumentIndexOfReply(sel objc.SEL, arg uint, ofReply bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("interfaceForSelector:argumentIndex:ofReply:"), sel, arg, ofReply)
	return rv
}

// Sets the classes that can appear within the (numerically) specified collection object argument to the specified method.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCInterface/setClasses(_:for:argumentIndex:ofReply:)
func (x_ XPCInterface) SetClassesForSelectorArgumentIndexOfReply(classes unsafe.Pointer, sel objc.SEL, arg uint, ofReply bool) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setClasses:forSelector:argumentIndex:ofReply:"), classes, sel, arg, ofReply)
}

// Configures a specific parameter of a method to be sent as a proxy object instead of copied.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCInterface/setInterface(_:for:argumentIndex:ofReply:)
func (x_ XPCInterface) SetInterfaceForSelectorArgumentIndexOfReply(ifc unsafe.Pointer, sel objc.SEL, arg uint, ofReply bool) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setInterface:forSelector:argumentIndex:ofReply:"), ifc, sel, arg, ofReply)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCInterface/setXPCType(_:for:argumentIndex:ofReply:)
func (x_ XPCInterface) SetXPCTypeForSelectorArgumentIndexOfReply(type_ unsafe.Pointer, sel objc.SEL, arg uint, ofReply bool) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setXPCType:forSelector:argumentIndex:ofReply:"), type_, sel, arg, ofReply)
}

// The Objective-C protocol that this interface is based on.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCInterface/protocol
func (x_ XPCInterface) Protocol() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("protocol"))
	return rv
}


// SetProtocol sets the value of the protocol property.
// The Objective-C protocol that this interface is based on.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCInterface/protocol
func (x_ XPCInterface) SetProtocol(value unsafe.Pointer) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setProtocol:"), value)
}



// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSXPCInterface */


/* debug [class_header]: Header for NSXPCInterface */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for XPCInterface */
// An interface definition for the [XPCInterface] class.
type IXPCInterface interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for XPCInterface */
	// properties:
	Protocol() objectivec.Protocol
	SetProtocol(value objectivec.Protocol)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for XPCInterface */
	// methods:
	ClassesForSelectorArgumentIndexOfReply(sel objc.SEL, arg uint, ofReply bool) unsafe.Pointer
	InterfaceForSelectorArgumentIndexOfReply(sel objc.SEL, arg uint, ofReply bool) IXPCInterface
	SetClassesForSelectorArgumentIndexOfReply(classes unsafe.Pointer, sel objc.SEL, arg uint, ofReply bool)
	SetInterfaceForSelectorArgumentIndexOfReply(ifc IXPCInterface, sel objc.SEL, arg uint, ofReply bool)
	SetXPCTypeForSelectorArgumentIndexOfReply(type_ objectivec.IObject, sel objc.SEL, arg uint, ofReply bool)
	XPCTypeForSelectorArgumentIndexOfReply(sel objc.SEL, arg uint, ofReply bool) objectivec.IObject
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for XPCInterface */
// Alloc allocates a new instance without initialization.
func (xc _XPCInterfaceClass) Alloc() XPCInterface {
	rv := objc.Send[XPCInterface](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for XPCInterface */
// An interface that may be sent to an exported object or remote object proxy.
//
// This object holds all information about the interface of an exported object or remote object proxy. It describes what messages are allowed, what kinds of objects are allowed as arguments, what the signature of any reply blocks are, and information about additional proxy objects.


// An interface that may be sent to an exported object or remote object proxy.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for XPCInterface */

// Returns an NSXPCInterface instance for a given protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCInterface/init(with:)
func NewXPCInterfaceWithProtocol(protocol_ objectivec.Protocol) XPCInterface {
	rv := objc.Send[XPCInterface](objc.ID(getXPCInterfaceClass().class), objc.Sel("interfaceWithProtocol:"), protocol_)
	return rv
}/* debug [class_init_methods/constructor]: NewXPCInterfaceWithProtocol */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for XPCInterface */

// Returns an NSXPCInterface instance for a given protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCInterface/init(with:)
func (xc _XPCInterfaceClass) InterfaceWithProtocol(protocol_ objectivec.Protocol) IXPCInterface {
	rv := objc.Send[XPCInterface](objc.ID(xc.class), objc.Sel("interfaceWithProtocol:"), protocol_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=InterfaceWithProtocol) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for XPCInterface */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for XPCInterface */

// Returns the current list of allowed classes that can appear within the specified collection object argument to the specified method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCInterface/classes(for:argumentIndex:ofReply:)
func (x_ XPCInterface) ClassesForSelectorArgumentIndexOfReply(sel objc.SEL, arg uint, ofReply bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](x_.ID, objc.Sel("classesForSelector:argumentIndex:ofReply:"), sel, arg, ofReply)
	return rv
}/* debug [instance_methods/method]: ClassesForSelectorArgumentIndexOfReply */


// Returns the interface previously set for the specified selector and parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCInterface/forSelector(_:argumentIndex:ofReply:)
func (x_ XPCInterface) InterfaceForSelectorArgumentIndexOfReply(sel objc.SEL, arg uint, ofReply bool) IXPCInterface {
	rv := objc.Send[XPCInterface](x_.ID, objc.Sel("interfaceForSelector:argumentIndex:ofReply:"), sel, arg, ofReply)
	return rv
}/* debug [instance_methods/method]: InterfaceForSelectorArgumentIndexOfReply */


// Sets the classes that can appear within the (numerically) specified collection object argument to the specified method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCInterface/setClasses(_:for:argumentIndex:ofReply:)
func (x_ XPCInterface) SetClassesForSelectorArgumentIndexOfReply(classes unsafe.Pointer, sel objc.SEL, arg uint, ofReply bool) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setClasses:forSelector:argumentIndex:ofReply:"), classes, sel, arg, ofReply)
}/* debug [instance_methods/method]: SetClassesForSelectorArgumentIndexOfReply */


// Configures a specific parameter of a method to be sent as a proxy object instead of copied.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCInterface/setInterface(_:for:argumentIndex:ofReply:)
func (x_ XPCInterface) SetInterfaceForSelectorArgumentIndexOfReply(ifc IXPCInterface, sel objc.SEL, arg uint, ofReply bool) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setInterface:forSelector:argumentIndex:ofReply:"), ifc, sel, arg, ofReply)
}/* debug [instance_methods/method]: SetInterfaceForSelectorArgumentIndexOfReply */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCInterface/setXPCType(_:for:argumentIndex:ofReply:)
func (x_ XPCInterface) SetXPCTypeForSelectorArgumentIndexOfReply(type_ objectivec.IObject, sel objc.SEL, arg uint, ofReply bool) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setXPCType:forSelector:argumentIndex:ofReply:"), type_, sel, arg, ofReply)
}/* debug [instance_methods/method]: SetXPCTypeForSelectorArgumentIndexOfReply */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCInterface/xpcType(for:argumentIndex:ofReply:)
func (x_ XPCInterface) XPCTypeForSelectorArgumentIndexOfReply(sel objc.SEL, arg uint, ofReply bool) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](x_.ID, objc.Sel("XPCTypeForSelector:argumentIndex:ofReply:"), sel, arg, ofReply)
	return rv
}/* debug [instance_methods/method]: XPCTypeForSelectorArgumentIndexOfReply */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for XPCInterface */

// The Objective-C protocol that this interface is based on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCInterface/protocol
func (x_ XPCInterface) Protocol() objectivec.Protocol {
	rv := objc.Send[objectivec.Protocol](x_.ID, objc.Sel("protocol"))
	return rv
}/* debug [instance_properties/getter]: protocol */


// The Objective-C protocol that this interface is based on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXPCInterface/protocol
func (x_ XPCInterface) SetProtocol(value objectivec.Protocol) {
	objc.Send[objc.ID](x_.ID, objc.Sel("setProtocol:"), value)
}/* debug [instance_properties/setter]: protocol */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSXPCInterface */



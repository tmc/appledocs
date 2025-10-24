// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSMessagePort */


/* debug [class_header]: Header for NSMessagePort */
// The class instance for the [MessagePort] class.
var (
	MessagePortClass     _MessagePortClass
	MessagePortClassOnce sync.Once
)

func getMessagePortClass() _MessagePortClass {
	MessagePortClassOnce.Do(func() {
		MessagePortClass = _MessagePortClass{objc.GetClass("NSMessagePort")}
	})
	return MessagePortClass
}

type _MessagePortClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MessagePort */
// An interface definition for the [MessagePort] class.
type IMessagePort interface {
	IPort
	
/* debug [class_interface_properties]: Properties for MessagePort */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MessagePort */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MessagePort */
// Alloc allocates a new instance without initialization.
func (mc _MessagePortClass) Alloc() MessagePort {
	rv := objc.Send[MessagePort](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
	return getMessagePortClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MessagePort */
// A port that can be used as an endpoint for distributed object connections (or raw messaging).
//
// is a subclass of that allows for local (on the same machine) communication only. A companion class, , allows for both local and remote communication, but may be more expensive than for the local case. defines no additional methods over those already defined by .


// A port that can be used as an endpoint for distributed object connections (or raw messaging).
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MessagePort *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MessagePort */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MessagePort */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MessagePort */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MessagePort */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSMessagePort */




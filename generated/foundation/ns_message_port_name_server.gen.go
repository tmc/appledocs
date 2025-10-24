// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSMessagePortNameServer */


/* debug [class_header]: Header for NSMessagePortNameServer */
// The class instance for the [MessagePortNameServer] class.
var (
	MessagePortNameServerClass     _MessagePortNameServerClass
	MessagePortNameServerClassOnce sync.Once
)

func getMessagePortNameServerClass() _MessagePortNameServerClass {
	MessagePortNameServerClassOnce.Do(func() {
		MessagePortNameServerClass = _MessagePortNameServerClass{objc.GetClass("NSMessagePortNameServer")}
	})
	return MessagePortNameServerClass
}

type _MessagePortNameServerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MessagePortNameServer */
// An interface definition for the [MessagePortNameServer] class.
type IMessagePortNameServer interface {
	IPortNameServer
	
/* debug [class_interface_properties]: Properties for MessagePortNameServer */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MessagePortNameServer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MessagePortNameServer */
// Alloc allocates a new instance without initialization.
func (mc _MessagePortNameServerClass) Alloc() MessagePortNameServer {
	rv := objc.Send[MessagePortNameServer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MessagePortNameServerClass) New() MessagePortNameServer {
	rv := objc.Send[MessagePortNameServer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MessagePortNameServer) Init() MessagePortNameServer {
	rv := objc.Send[MessagePortNameServer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MessagePortNameServer) Autorelease() MessagePortNameServer {
	rv := objc.Send[MessagePortNameServer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMessagePortNameServer creates a new MessagePortNameServer instance.
func NewMessagePortNameServer() MessagePortNameServer {
	return getMessagePortNameServerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MessagePortNameServer */
// A server takes and returns message ports.
//
// This port name server takes and returns instances of . Port removal functionality is not supported in ; if you want to cancel a service, you have to destroy the port (invalidate the object given to ).


// A server takes and returns message ports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMessagePortNameServer
type MessagePortNameServer struct {
	PortNameServer
}

// MessagePortNameServerFrom constructs a [MessagePortNameServer] from an unsafe.Pointer.
//
// A server takes and returns message ports.
func MessagePortNameServerFrom(ptr unsafe.Pointer) MessagePortNameServer {
	return MessagePortNameServer{
		PortNameServer: PortNameServerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MessagePortNameServer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MessagePortNameServer */

// Returns the singleton instance of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMessagePortNameServer/sharedInstance
func (mc _MessagePortNameServerClass) SharedInstance() objc.ID {
	rv := objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("sharedInstance"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedInstance) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MessagePortNameServer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MessagePortNameServer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MessagePortNameServer */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSMessagePortNameServer */




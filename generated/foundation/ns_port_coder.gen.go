// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSPortCoder */


/* debug [class_header]: Header for NSPortCoder */
// The class instance for the [PortCoder] class.
var (
	PortCoderClass     _PortCoderClass
	PortCoderClassOnce sync.Once
)

func getPortCoderClass() _PortCoderClass {
	PortCoderClassOnce.Do(func() {
		PortCoderClass = _PortCoderClass{objc.GetClass("NSPortCoder")}
	})
	return PortCoderClass
}

type _PortCoderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PortCoder */
// An interface definition for the [PortCoder] class.
type IPortCoder interface {
	ICoder
	
/* debug [class_interface_properties]: Properties for PortCoder */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PortCoder */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PortCoder */
// Alloc allocates a new instance without initialization.
func (pc _PortCoderClass) Alloc() PortCoder {
	rv := objc.Send[PortCoder](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PortCoderClass) New() PortCoder {
	rv := objc.Send[PortCoder](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PortCoder) Init() PortCoder {
	rv := objc.Send[PortCoder](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PortCoder) Autorelease() PortCoder {
	rv := objc.Send[PortCoder](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPortCoder creates a new PortCoder instance.
func NewPortCoder() PortCoder {
	return getPortCoderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PortCoder */
// A coder used to transmit object proxies (and sometimes objects themselves) between connections.
//
// is a concrete subclass of used in the distributed objects system to transmit object proxies (and sometimes objects themselves) between objects. An instance is always created and used by an object; you should never need to explicitly create or use one directly yourself.


// A coder used to transmit object proxies (and sometimes objects themselves) between connections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPortCoder
type PortCoder struct {
	Coder
}

// PortCoderFrom constructs a [PortCoder] from an unsafe.Pointer.
//
// A coder used to transmit object proxies (and sometimes objects themselves) between connections.
func PortCoderFrom(ptr unsafe.Pointer) PortCoder {
	return PortCoder{
		Coder: CoderFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PortCoder */

// Initializes and returns an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPortCoder/initWithReceivePort:sendPort:components:
func NewPortCoderWithReceivePortSendPortComponents(rcvPort IPort, sndPort IPort, comps IArray) PortCoder {
	instance := getPortCoderClass().Alloc()
	rv := objc.Send[PortCoder](instance.ID, objc.Sel("initWithReceivePort:sendPort:components:"), rcvPort, sndPort, comps)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPortCoderWithReceivePortSendPortComponents */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PortCoder */

// Creates and returns a new object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPortCoder/portCoderWithReceivePort:sendPort:components:
func (pc _PortCoderClass) PortCoderWithReceivePortSendPortComponents(rcvPort IPort, sndPort IPort, comps IArray) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("portCoderWithReceivePort:sendPort:components:"), rcvPort, sndPort, comps)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PortCoderWithReceivePortSendPortComponents) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PortCoder */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PortCoder */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PortCoder */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPortCoder */



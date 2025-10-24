// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSPort */


/* debug [class_header]: Header for NSPort */
// The class instance for the [Port] class.
var (
	PortClass     _PortClass
	PortClassOnce sync.Once
)

func getPortClass() _PortClass {
	PortClassOnce.Do(func() {
		PortClass = _PortClass{objc.GetClass("NSPort")}
	})
	return PortClass
}

type _PortClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Port */
// An interface definition for the [Port] class.
type IPort interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Port */
	// properties:
	Valid() bool
	IsValid() bool
	SetIsValid(value bool)
	ReservedSpaceLength() int
	SetReservedSpaceLength(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Port */
	// methods:
	Invalidate()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Port */
// Alloc allocates a new instance without initialization.
func (pc _PortClass) Alloc() Port {
	rv := objc.Send[Port](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PortClass) New() Port {
	rv := objc.Send[Port](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Port) Init() Port {
	rv := objc.Send[Port](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Port) Autorelease() Port {
	rv := objc.Send[Port](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPort creates a new Port instance.
func NewPort() Port {
	return getPortClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Port */
// An abstract class that represents a communication channel.
//
// Communication occurs between objects, which typically reside in different threads or tasks. The distributed objects system uses objects to send objects back and forth. Implement interapplication communication using distributed objects whenever possible and use objects only when necessary. To receive incoming messages, add objects to an instance of as input sources. objects automatically add their receive port when initialized. When the object receives a port message, it forwards the message to its delegate in a or message. The delegate should implement only one of these methods to process the incoming message in whatever form desired. provides a message as a raw Mach message beginning with a structure. provides a message as an instance of , which is an object-oriented wrapper for a Mach message. If a delegate has not been set, the object handles the message itself. When you are finished using a port object, you must explicitly invalidate the port object prior to sending it a message. Similarly, if your application uses garbage collection, you must invalidate the port object before removing any strong references to it. If you do not invalidate the port, the resulting port object may linger and create a memory leak. To invalidate the port object, invoke its method. Foundation defines three concrete subclasses of . and allow local (on the same machine) communication only. allows for both local and remote communication, but may be more expensive than the others for the local case. When creating an object, using doc:nsport/1807189-allocwithzone or , an object is created instead. For backward compatibility on Mach, returns an instance of the class when sent to this class. Otherwise, it returns an instance of a concrete subclass that can be used for messaging between threads or processes on the local machine, or, in the case of , between processes on separate machines.


// An abstract class that represents a communication channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Port
type Port struct {
	objectivec.Object
}

// PortFrom constructs a [Port] from an unsafe.Pointer.
//
// An abstract class that represents a communication channel.
func PortFrom(ptr unsafe.Pointer) Port {
	return Port{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Port *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Port */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Port */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Port */

// Marks the receiver as invalid and posts an to the default notification center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Port/invalidate()
func (p_ Port) Invalidate() {
	objc.Send[objc.ID](p_.ID, objc.Sel("invalidate"))
}/* debug [instance_methods/method]: Invalidate */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Port */

// A Boolean value that indicates whether the receiver is valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Port/isValid
func (p_ Port) Valid() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("valid"))
	return rv
}/* debug [instance_properties/getter]: valid */


// A Boolean value that indicates whether the receiver is valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/port/isvalid
func (p_ Port) IsValid() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isValid"))
	return rv
}/* debug [instance_properties/getter]: isValid */


// A Boolean value that indicates whether the receiver is valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/port/isvalid
func (p_ Port) SetIsValid(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsValid:"), value)
}/* debug [instance_properties/setter]: isValid */


// The number of bytes of space reserved by the receiver for sending data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/port/reservedspacelength
func (p_ Port) ReservedSpaceLength() int {
	rv := objc.Send[int](p_.ID, objc.Sel("reservedSpaceLength"))
	return rv
}/* debug [instance_properties/getter]: reservedSpaceLength */


// The number of bytes of space reserved by the receiver for sending data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/port/reservedspacelength
func (p_ Port) SetReservedSpaceLength(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setReservedSpaceLength:"), value)
}/* debug [instance_properties/setter]: reservedSpaceLength */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPort */




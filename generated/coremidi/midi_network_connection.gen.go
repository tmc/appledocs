// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MIDINetworkConnection */


/* debug [class_header]: Header for MIDINetworkConnection */
// The class instance for the [MIDINetworkConnection] class.
var (
	MIDINetworkConnectionClass     _MIDINetworkConnectionClass
	MIDINetworkConnectionClassOnce sync.Once
)

func getMIDINetworkConnectionClass() _MIDINetworkConnectionClass {
	MIDINetworkConnectionClassOnce.Do(func() {
		MIDINetworkConnectionClass = _MIDINetworkConnectionClass{objc.GetClass("MIDINetworkConnection")}
	})
	return MIDINetworkConnectionClass
}

type _MIDINetworkConnectionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MIDINetworkConnection */
// An interface definition for the [MIDINetworkConnection] class.
type IMIDINetworkConnection interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MIDINetworkConnection */
	// properties:
	Host() IMIDINetworkHost
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MIDINetworkConnection */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MIDINetworkConnection */
// Alloc allocates a new instance without initialization.
func (mc _MIDINetworkConnectionClass) Alloc() MIDINetworkConnection {
	rv := objc.Send[MIDINetworkConnection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MIDINetworkConnectionClass) New() MIDINetworkConnection {
	rv := objc.Send[MIDINetworkConnection](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDINetworkConnection) Init() MIDINetworkConnection {
	rv := objc.Send[MIDINetworkConnection](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDINetworkConnection) Autorelease() MIDINetworkConnection {
	rv := objc.Send[MIDINetworkConnection](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDINetworkConnection creates a new MIDINetworkConnection instance.
func NewMIDINetworkConnection() MIDINetworkConnection {
	return getMIDINetworkConnectionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MIDINetworkConnection */
// An object that connects a session to a host.


// An object that connects a session to a host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkConnection
type MIDINetworkConnection struct {
	objectivec.Object
}

// MIDINetworkConnectionFrom constructs a [MIDINetworkConnection] from an unsafe.Pointer.
//
// An object that connects a session to a host.
func MIDINetworkConnectionFrom(ptr unsafe.Pointer) MIDINetworkConnection {
	return MIDINetworkConnection{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MIDINetworkConnection */

// Creates a connection to the specified host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkConnection/init(host:)
func NewMIDINetworkConnectionWithHost(host IMIDINetworkHost) MIDINetworkConnection {
	rv := objc.Send[MIDINetworkConnection](objc.ID(getMIDINetworkConnectionClass().class), objc.Sel("connectionWithHost:"), host)
	return rv
}/* debug [class_init_methods/constructor]: NewMIDINetworkConnectionWithHost */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MIDINetworkConnection */

// Creates a connection to the specified host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkConnection/init(host:)
func (mc _MIDINetworkConnectionClass) ConnectionWithHost(host IMIDINetworkHost) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("connectionWithHost:"), host)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ConnectionWithHost) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MIDINetworkConnection */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MIDINetworkConnection */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MIDINetworkConnection */

// The host connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDINetworkConnection/host
func (m_ MIDINetworkConnection) Host() IMIDINetworkHost {
	rv := objc.Send[MIDINetworkHost](m_.ID, objc.Sel("host"))
	return rv
}/* debug [instance_properties/getter]: host */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MIDINetworkConnection */



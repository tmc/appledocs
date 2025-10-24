// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CBL2CAPChannel */


/* debug [class_header]: Header for CBL2CAPChannel */
// The class instance for the [CBL2CAPChannel] class.
var (
	CBL2CAPChannelClass     _CBL2CAPChannelClass
	CBL2CAPChannelClassOnce sync.Once
)

func getCBL2CAPChannelClass() _CBL2CAPChannelClass {
	CBL2CAPChannelClassOnce.Do(func() {
		CBL2CAPChannelClass = _CBL2CAPChannelClass{objc.GetClass("CBL2CAPChannel")}
	})
	return CBL2CAPChannelClass
}

type _CBL2CAPChannelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CBL2CAPChannel */
// An interface definition for the [CBL2CAPChannel] class.
type ICBL2CAPChannel interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CBL2CAPChannel */
	// properties:
	InputStream() foundation.InputStream
	OutputStream() foundation.OutputStream
	Peer() ICBPeer
	PSM() CBL2CAPPSM /* typedef */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CBL2CAPChannel */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CBL2CAPChannel */
// Alloc allocates a new instance without initialization.
func (cc _CBL2CAPChannelClass) Alloc() CBL2CAPChannel {
	rv := objc.Send[CBL2CAPChannel](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CBL2CAPChannelClass) New() CBL2CAPChannel {
	rv := objc.Send[CBL2CAPChannel](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBL2CAPChannel) Init() CBL2CAPChannel {
	rv := objc.Send[CBL2CAPChannel](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBL2CAPChannel) Autorelease() CBL2CAPChannel {
	rv := objc.Send[CBL2CAPChannel](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBL2CAPChannel creates a new CBL2CAPChannel instance.
func NewCBL2CAPChannel() CBL2CAPChannel {
	return getCBL2CAPChannelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CBL2CAPChannel */
// A live L2CAP connection to a remote device.


// A live L2CAP connection to a remote device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBL2CAPChannel
type CBL2CAPChannel struct {
	objectivec.Object
}

// CBL2CAPChannelFrom constructs a [CBL2CAPChannel] from an unsafe.Pointer.
//
// A live L2CAP connection to a remote device.
func CBL2CAPChannelFrom(ptr unsafe.Pointer) CBL2CAPChannel {
	return CBL2CAPChannel{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CBL2CAPChannel *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CBL2CAPChannel */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CBL2CAPChannel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CBL2CAPChannel */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CBL2CAPChannel */

// The stream used for reading data from the remote peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBL2CAPChannel/inputStream
func (c_ CBL2CAPChannel) InputStream() foundation.InputStream {
	rv := objc.Send[foundation.InputStream](c_.ID, objc.Sel("inputStream"))
	return rv
}/* debug [instance_properties/getter]: inputStream */


// The stream used for writing data to the peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBL2CAPChannel/outputStream
func (c_ CBL2CAPChannel) OutputStream() foundation.OutputStream {
	rv := objc.Send[foundation.OutputStream](c_.ID, objc.Sel("outputStream"))
	return rv
}/* debug [instance_properties/getter]: outputStream */


// The peer connected to the channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBL2CAPChannel/peer
func (c_ CBL2CAPChannel) Peer() ICBPeer {
	rv := objc.Send[CBPeer](c_.ID, objc.Sel("peer"))
	return rv
}/* debug [instance_properties/getter]: peer */


// The PSM of the channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBL2CAPChannel/psm
func (c_ CBL2CAPChannel) PSM() CBL2CAPPSM /* typedef */ {
	rv := objc.Send[uint16](c_.ID, objc.Sel("PSM"))
	return rv
}/* debug [instance_properties/getter]: PSM */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CBL2CAPChannel */






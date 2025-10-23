// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [CBL2CAPChannel] class.
type ICBL2CAPChannel interface {
	objectivec.IObject
	// properties:
	InputStream() InputStream /* not a class type */
	SetInputStream(value InputStream /* not a class type */)
	OutputStream() objc.IObject /* cross-framework: OutputStream */
	SetOutputStream(value objc.IObject /* cross-framework: OutputStream */)
	Peer() ICBPeer
	SetPeer(value ICBPeer)
	Psm() CBL2CAPPSM /* typedef */
	SetPsm(value CBL2CAPPSM /* typedef */)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (cc _CBL2CAPChannelClass) Alloc() CBL2CAPChannel {
	rv := objc.Send[CBL2CAPChannel](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The stream used for reading data from the remote peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corebluetooth/cbl2capchannel/inputstream
func (c_ CBL2CAPChannel) InputStream() InputStream /* not a class type */ {
	rv := objc.Send[InputStream](c_.ID, objc.Sel("inputStream"))
	return rv
}


// The stream used for reading data from the remote peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corebluetooth/cbl2capchannel/inputstream
func (c_ CBL2CAPChannel) SetInputStream(value InputStream /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInputStream:"), value)
}


// The stream used for writing data to the peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corebluetooth/cbl2capchannel/outputstream
func (c_ CBL2CAPChannel) OutputStream() objc.IObject /* cross-framework: OutputStream */ {
	rv := objc.Send[OutputStream](c_.ID, objc.Sel("outputStream"))
	return rv
}


// The stream used for writing data to the peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corebluetooth/cbl2capchannel/outputstream
func (c_ CBL2CAPChannel) SetOutputStream(value objc.IObject /* cross-framework: OutputStream */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOutputStream:"), value)
}


// The peer connected to the channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corebluetooth/cbl2capchannel/peer
func (c_ CBL2CAPChannel) Peer() ICBPeer {
	rv := objc.Send[CBPeer](c_.ID, objc.Sel("peer"))
	return rv
}


// The peer connected to the channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corebluetooth/cbl2capchannel/peer
func (c_ CBL2CAPChannel) SetPeer(value ICBPeer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPeer:"), value)
}


// The PSM of the channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corebluetooth/cbl2capchannel/psm
func (c_ CBL2CAPChannel) Psm() CBL2CAPPSM /* typedef */ {
	rv := objc.Send[CBL2CAPPSM](c_.ID, objc.Sel("psm"))
	return rv
}


// The PSM of the channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corebluetooth/cbl2capchannel/psm
func (c_ CBL2CAPChannel) SetPsm(value CBL2CAPPSM /* typedef */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPsm:"), value)
}




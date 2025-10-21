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
}

// A live L2CAP connection to a remote device.
//
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





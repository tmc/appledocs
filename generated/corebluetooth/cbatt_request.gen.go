// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CBATTRequest] class.
var (
	CBATTRequestClass     _CBATTRequestClass
	CBATTRequestClassOnce sync.Once
)

func getCBATTRequestClass() _CBATTRequestClass {
	CBATTRequestClassOnce.Do(func() {
		CBATTRequestClass = _CBATTRequestClass{objc.GetClass("CBATTRequest")}
	})
	return CBATTRequestClass
}

type _CBATTRequestClass struct {
	class objc.Class
}

// An interface definition for the [CBATTRequest] class.
type ICBATTRequest interface {
	objectivec.IObject
}

// A request that uses the Attribute Protocol (ATT).
//
// The class represents Attribute Protocol (ATT) read and write requests from remote central devices (represented by objects). Remote centrals use these ATT requests to read and write characteristic values on local peripherals (represented by objects). Local peripherals, on the other hand, use the properties of objects to respond to the read and write requests appropriately, using the method of the class.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTRequest
type CBATTRequest struct {
	objectivec.Object
}

// CBATTRequestFrom constructs a [CBATTRequest] from an unsafe.Pointer.
//
// A request that uses the Attribute Protocol (ATT).
func CBATTRequestFrom(ptr unsafe.Pointer) CBATTRequest {
	return CBATTRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CBATTRequestClass) Alloc() CBATTRequest {
	rv := objc.Send[CBATTRequest](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CBATTRequestClass) New() CBATTRequest {
	rv := objc.Send[CBATTRequest](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBATTRequest) Init() CBATTRequest {
	rv := objc.Send[CBATTRequest](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBATTRequest) Autorelease() CBATTRequest {
	rv := objc.Send[CBATTRequest](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBATTRequest creates a new CBATTRequest instance.
func NewCBATTRequest() CBATTRequest {
	return getCBATTRequestClass().New()
}


// The remote central device that originated the request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTRequest/central
func (c_ CBATTRequest) Central() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("central"))
	return rv
}

// The characteristic to read or write the value of.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTRequest/characteristic
func (c_ CBATTRequest) Characteristic() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("characteristic"))
	return rv
}

// The zero-based index of the first byte for the read or write request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTRequest/offset
func (c_ CBATTRequest) Offset() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("offset"))
	return rv
}

// The data that the central reads from or writes to the peripheral.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTRequest/value
func (c_ CBATTRequest) Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
// The data that the central reads from or writes to the peripheral.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTRequest/value
func (c_ CBATTRequest) SetValue(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setValue:"), value)
}




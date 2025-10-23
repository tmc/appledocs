// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [BluetoothSDPUUID] class.
var (
	BluetoothSDPUUIDClass     _BluetoothSDPUUIDClass
	BluetoothSDPUUIDClassOnce sync.Once
)

func getBluetoothSDPUUIDClass() _BluetoothSDPUUIDClass {
	BluetoothSDPUUIDClassOnce.Do(func() {
		BluetoothSDPUUIDClass = _BluetoothSDPUUIDClass{objc.GetClass("IOBluetoothSDPUUID")}
	})
	return BluetoothSDPUUIDClass
}

type _BluetoothSDPUUIDClass struct {
	class objc.Class
}

// An interface definition for the [BluetoothSDPUUID] class.
type IBluetoothSDPUUID interface {
	foundation.IData
	// properties:
	// methods:
	ClassForArchiver() objc.Class
	ClassForCoder() objc.Class
	ClassForPortCoder() objc.Class
	GetUUIDWithLength(newLength unsafe.Pointer) unsafe.Pointer
	IsEqualToUUID(otherUUID BluetoothSDPUUID /* already interface */) bool /* primitive/slice/pointer. */
}

// An NSData subclass that represents a UUID as defined in the Bluetooth SDP spec.
//
// The IOBluetoothSDPUUID class can represent a UUID of any valid size (16, 32 or 128 bits). It provides the ability to compare two UUIDs no matter what their size as well as the ability to promote the size of a UUID to a larger one.


// An NSData subclass that represents a UUID as defined in the Bluetooth SDP spec.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID
type BluetoothSDPUUID struct {
	foundation.Data
}

// BluetoothSDPUUIDFrom constructs a [BluetoothSDPUUID] from an unsafe.Pointer.
//
// An NSData subclass that represents a UUID as defined in the Bluetooth SDP spec.
func BluetoothSDPUUIDFrom(ptr unsafe.Pointer) BluetoothSDPUUID {
	return BluetoothSDPUUID{
		Data: foundation.DataFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BluetoothSDPUUIDClass) Alloc() BluetoothSDPUUID {
	rv := objc.Send[BluetoothSDPUUID](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BluetoothSDPUUIDClass) New() BluetoothSDPUUID {
	rv := objc.Send[BluetoothSDPUUID](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BluetoothSDPUUID) Init() BluetoothSDPUUID {
	rv := objc.Send[BluetoothSDPUUID](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BluetoothSDPUUID) Autorelease() BluetoothSDPUUID {
	rv := objc.Send[BluetoothSDPUUID](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBluetoothSDPUUID creates a new BluetoothSDPUUID instance.
func NewBluetoothSDPUUID() BluetoothSDPUUID {
	return getBluetoothSDPUUIDClass().New()
}



// Creates a new IOBluetoothSDPUUID object with the given bytes of the given length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID/init(bytes:length:)
func NewBluetoothSDPUUIDUuidWithBytesLength(bytes unsafe.Pointer, length unsafe.Pointer) BluetoothSDPUUID {
	rv := objc.Send[BluetoothSDPUUID](objc.ID(getBluetoothSDPUUIDClass().class), objc.Sel("uuidWithBytes:length:"), bytes, length)
	return rv
}


// Creates a new IOBluetoothSDPUUID object from the given NSData.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID/init(data:)
func NewBluetoothSDPUUIDUuidWithData(data foundation.objc.IObject /* cross-framework NSData */) BluetoothSDPUUID {
	rv := objc.Send[BluetoothSDPUUID](objc.ID(getBluetoothSDPUUIDClass().class), objc.Sel("uuidWithData:"), data)
	return rv
}


// Initializes a new 16-bit IOBluetoothSDPUUID with the given UUID16
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID/init(uuid16:)
func NewBluetoothSDPUUIDWithUUID16(uuid16 BluetoothSDPUUID16 /* typedef */) BluetoothSDPUUID {
	instance := getBluetoothSDPUUIDClass().Alloc()
	rv := objc.Send[BluetoothSDPUUID](instance.ID, objc.Sel("initWithUUID16:"), uuid16)
	rv.Autorelease()
	return rv
}


// Creates a new 32-bit IOBluetoothSDPUUID with the given UUID32
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID/init(uuid32:)
func NewBluetoothSDPUUIDWithUUID32(uuid32 BluetoothSDPUUID32 /* typedef */) BluetoothSDPUUID {
	instance := getBluetoothSDPUUIDClass().Alloc()
	rv := objc.Send[BluetoothSDPUUID](instance.ID, objc.Sel("initWithUUID32:"), uuid32)
	rv.Autorelease()
	return rv
}



// Creates a new IOBluetoothSDPUUID object with the given bytes of the given length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID/init(bytes:length:)
func (bc _BluetoothSDPUUIDClass) UuidWithBytesLength(bytes unsafe.Pointer, length unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("uuidWithBytes:length:"), bytes, length)
	return rv
}


// Creates a new IOBluetoothSDPUUID object from the given NSData.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID/init(data:)
func (bc _BluetoothSDPUUIDClass) UuidWithData(data foundation.objc.IObject /* cross-framework NSData */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("uuidWithData:"), data)
	return rv
}


// Creates a new 16-bit IOBluetoothSDPUUID with the given UUID16
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID/uuid16(_:)
func (bc _BluetoothSDPUUIDClass) Uuid16(uuid16 BluetoothSDPUUID16 /* typedef */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("uuid16:"), uuid16)
	return rv
}


// Creates a new 32-bit IOBluetoothSDPUUID with the given UUID32
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID/uuid32(_:)
func (bc _BluetoothSDPUUIDClass) Uuid32(uuid32 BluetoothSDPUUID32 /* typedef */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("uuid32:"), uuid32)
	return rv
}


// Method call to convert an IOBluetoothSDPUUIDRef into an IOBluetoothSDPUUID *.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID/withSDPUUIDRef:
func (bc _BluetoothSDPUUIDClass) WithSDPUUIDRef(sdpUUIDRef objc.IObject /* cross-framework BluetoothSDPUUIDRef */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("withSDPUUIDRef:"), sdpUUIDRef)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID/classForArchiver()
func (b_ BluetoothSDPUUID) ClassForArchiver() objc.Class {
	rv := objc.Send[objc.Class](b_.ID, objc.Sel("classForArchiver"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID/classForCoder()
func (b_ BluetoothSDPUUID) ClassForCoder() objc.Class {
	rv := objc.Send[objc.Class](b_.ID, objc.Sel("classForCoder"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID/classForPortCoder()
func (b_ BluetoothSDPUUID) ClassForPortCoder() objc.Class {
	rv := objc.Send[objc.Class](b_.ID, objc.Sel("classForPortCoder"))
	return rv
}


// Returns an IOBluetoothSDPUUID object matching the target UUID, but with the given number of bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID/getWithLength(_:)
func (b_ BluetoothSDPUUID) GetUUIDWithLength(newLength unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getUUIDWithLength:"), newLength)
	return rv
}


// Compares the target IOBluetoothSDPUUID object with the given otherUUID object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPUUID/isEqual(to:)
func (b_ BluetoothSDPUUID) IsEqualToUUID(otherUUID BluetoothSDPUUID /* already interface */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("isEqualToUUID:"), otherUUID)
	return rv
}



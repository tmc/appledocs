// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CBUUID] class.
var (
	CBUUIDClass     _CBUUIDClass
	CBUUIDClassOnce sync.Once
)

func getCBUUIDClass() _CBUUIDClass {
	CBUUIDClassOnce.Do(func() {
		CBUUIDClass = _CBUUIDClass{objc.GetClass("CBUUID")}
	})
	return CBUUIDClass
}

type _CBUUIDClass struct {
	class objc.Class
}

// An interface definition for the [CBUUID] class.
type ICBUUID interface {
	objectivec.IObject
}

// A universally unique identifier, as defined by Bluetooth standards.
//
// Instances of the class represent the 128-bit universally unique identifiers (UUIDs) of attributes used in Bluetooth low energy communication, such as a peripheral’s services, characteristics, and descriptors. This class provides a number of factory methods for dealing with long UUIDs when developing your app. For example, instead of passing around the string representation of a 128-bit Bluetooth low energy attribute in your code, you can create a object that represents it, and pass that around instead. The Bluetooth Special Interest Group (SIG) publishes a list of commonly-used UUIDs, many of which are 16- or 32-bits for convenience. The class provides methods that automatically transform these predefined shorter UUIDs into their 128-bit equivalent UUIDs. When you create a object from a predefined 16- or 32-bit UUID, Core Bluetooth pre-fills the rest of the 128-bit UUID with the Bluetooth base UUID, as defined in the Bluetooth 4.0 specification, Volume 3, Part F, Section 3.2.1. In addition to providing methods for creating objects, this class defines constants that represent the UUIDs of the Bluetooth-defined characteristic descriptors, as defined in the Bluetooth 4.0 specification, Volume 3, Part G, Section 3.3.3.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBUUID
type CBUUID struct {
	objectivec.Object
}

// CBUUIDFrom constructs a [CBUUID] from an unsafe.Pointer.
//
// A universally unique identifier, as defined by Bluetooth standards.
func CBUUIDFrom(ptr unsafe.Pointer) CBUUID {
	return CBUUID{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CBUUIDClass) Alloc() CBUUID {
	rv := objc.Send[CBUUID](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CBUUIDClass) New() CBUUID {
	rv := objc.Send[CBUUID](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBUUID) Init() CBUUID {
	rv := objc.Send[CBUUID](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBUUID) Autorelease() CBUUID {
	rv := objc.Send[CBUUID](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBUUID creates a new CBUUID instance.
func NewCBUUID() CBUUID {
	return getCBUUIDClass().New()
}




// Creates a Core Bluetooth UUID object from a Core Foundation UUID object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBUUID/init(cfuuid:)
func NewCBUUIDWithCFUUID(theUUID unsafe.Pointer) CBUUID {
	rv := objc.Send[CBUUID](objc.ID(getCBUUIDClass().class), objc.Sel("UUIDWithCFUUID:"), theUUID)
	return rv
}



// Creates a Core Bluetooth UUID object from a 16-, 32-, or 128-bit UUID data container.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBUUID/init(data:)
func NewCBUUIDWithData(theData foundation.IData) CBUUID {
	rv := objc.Send[CBUUID](objc.ID(getCBUUIDClass().class), objc.Sel("UUIDWithData:"), theData)
	return rv
}



// Creates a Core Bluetooth UUID object from a Foundation UUID object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBUUID/init(nsuuid:)
func NewCBUUIDWithNSUUID(theUUID foundation.IUUID) CBUUID {
	rv := objc.Send[CBUUID](objc.ID(getCBUUIDClass().class), objc.Sel("UUIDWithNSUUID:"), theUUID)
	return rv
}



// Creates a Core Bluetooth UUID object from a 16-, 32-, or 128-bit UUID string.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBUUID/init(string:)
func NewCBUUIDWithString(theString appkit.string) CBUUID {
	rv := objc.Send[CBUUID](objc.ID(getCBUUIDClass().class), objc.Sel("UUIDWithString:"), theString)
	return rv
}


// Creates a Core Bluetooth UUID object from a Core Foundation UUID object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBUUID/init(cfuuid:)
func (cc _CBUUIDClass) UUIDWithCFUUID(theUUID unsafe.Pointer) CBUUID {
	rv := objc.Send[CBUUID](objc.ID(cc.class), objc.Sel("UUIDWithCFUUID:"), theUUID)
	return rv
}

// Creates a Core Bluetooth UUID object from a 16-, 32-, or 128-bit UUID data container.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBUUID/init(data:)
func (cc _CBUUIDClass) UUIDWithData(theData foundation.IData) CBUUID {
	rv := objc.Send[CBUUID](objc.ID(cc.class), objc.Sel("UUIDWithData:"), theData)
	return rv
}

// Creates a Core Bluetooth UUID object from a Foundation UUID object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBUUID/init(nsuuid:)
func (cc _CBUUIDClass) UUIDWithNSUUID(theUUID foundation.IUUID) CBUUID {
	rv := objc.Send[CBUUID](objc.ID(cc.class), objc.Sel("UUIDWithNSUUID:"), theUUID)
	return rv
}

// Creates a Core Bluetooth UUID object from a 16-, 32-, or 128-bit UUID string.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBUUID/init(string:)
func (cc _CBUUIDClass) UUIDWithString(theString appkit.string) CBUUID {
	rv := objc.Send[CBUUID](objc.ID(cc.class), objc.Sel("UUIDWithString:"), theString)
	return rv
}

// The data of the UUID.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBUUID/data
func (c_ CBUUID) Data() foundation.NSData {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("data"))
	return rv
}

// The UUID represented as a string.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBUUID/uuidString
func (c_ CBUUID) UUIDString() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("UUIDString"))
	return rv
}



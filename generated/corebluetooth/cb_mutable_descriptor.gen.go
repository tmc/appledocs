// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CBMutableDescriptor] class.
var (
	CBMutableDescriptorClass     _CBMutableDescriptorClass
	CBMutableDescriptorClassOnce sync.Once
)

func getCBMutableDescriptorClass() _CBMutableDescriptorClass {
	CBMutableDescriptorClassOnce.Do(func() {
		CBMutableDescriptorClass = _CBMutableDescriptorClass{objc.GetClass("CBMutableDescriptor")}
	})
	return CBMutableDescriptorClass
}

type _CBMutableDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [CBMutableDescriptor] class.
type ICBMutableDescriptor interface {
	ICBDescriptor
	CBUUIDCharacteristicFormatString() string
	CBUUIDCharacteristicUserDescriptionString() string
}

// An object that provides additional information about a local peripheral’s characteristic.
//
// You use the class to create a local characteristic descriptor. After you create a descriptor and associate it with a local characteristic, you can publish it to the peripheral’s local database using the method of the class. This also publishes the characteristic and local service to which the descriptor belongs. After you publish a local descriptor, Core Bluetooth caches the descriptor and you can no longer make changes to it. details predefined descriptor types and their corresponding value types. That said, only two of these are currently supported when creating local, mutable descriptors: the characteristic user description descriptor and the characteristic format descriptor. declares these as the constants and , respectively. The system automatically creates the extended properties descriptor and the client configuration descriptor, depending on the properties of the characteristic to which the descriptor belongs.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBMutableDescriptor
type CBMutableDescriptor struct {
	CBDescriptor
}

// CBMutableDescriptorFrom constructs a [CBMutableDescriptor] from an unsafe.Pointer.
//
// An object that provides additional information about a local peripheral’s characteristic.
func CBMutableDescriptorFrom(ptr unsafe.Pointer) CBMutableDescriptor {
	return CBMutableDescriptor{
		CBDescriptor: CBDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CBMutableDescriptorClass) Alloc() CBMutableDescriptor {
	rv := objc.Send[CBMutableDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CBMutableDescriptorClass) New() CBMutableDescriptor {
	rv := objc.Send[CBMutableDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBMutableDescriptor) Init() CBMutableDescriptor {
	rv := objc.Send[CBMutableDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBMutableDescriptor) Autorelease() CBMutableDescriptor {
	rv := objc.Send[CBMutableDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBMutableDescriptor creates a new CBMutableDescriptor instance.
func NewCBMutableDescriptor() CBMutableDescriptor {
	return getCBMutableDescriptorClass().New()
}




// Creates a mutable descriptor with a specified value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBMutableDescriptor/init(type:value:)
func NewCBMutableDescriptorWithTypeValue(UUID ICBUUID, value objectivec.IObject) CBMutableDescriptor {
	instance := getCBMutableDescriptorClass().Alloc()
	rv := objc.Send[CBMutableDescriptor](instance.ID, objc.Sel("initWithType:value:"), UUID, value)
	rv.Autorelease()
	return rv
}


// The UUID for the Presentation Format descriptor, as a string.
//
// [Full Topic]: https://developer.apple.com/documentation/corebluetooth/cbuuidcharacteristicformatstring
func (c_ CBMutableDescriptor) CBUUIDCharacteristicFormatString() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CBUUIDCharacteristicFormatString"))
	return rv
}

// The UUID for the User Description descriptor, as a string.
//
// [Full Topic]: https://developer.apple.com/documentation/corebluetooth/cbuuidcharacteristicuserdescriptionstring
func (c_ CBMutableDescriptor) CBUUIDCharacteristicUserDescriptionString() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CBUUIDCharacteristicUserDescriptionString"))
	return rv
}



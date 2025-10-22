// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [VZVirtioConsoleDeviceSerialPortConfiguration] class.
var (
	VZVirtioConsoleDeviceSerialPortConfigurationClass     _VZVirtioConsoleDeviceSerialPortConfigurationClass
	VZVirtioConsoleDeviceSerialPortConfigurationClassOnce sync.Once
)

func getVZVirtioConsoleDeviceSerialPortConfigurationClass() _VZVirtioConsoleDeviceSerialPortConfigurationClass {
	VZVirtioConsoleDeviceSerialPortConfigurationClassOnce.Do(func() {
		VZVirtioConsoleDeviceSerialPortConfigurationClass = _VZVirtioConsoleDeviceSerialPortConfigurationClass{objc.GetClass("VZVirtioConsoleDeviceSerialPortConfiguration")}
	})
	return VZVirtioConsoleDeviceSerialPortConfigurationClass
}

type _VZVirtioConsoleDeviceSerialPortConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [VZVirtioConsoleDeviceSerialPortConfiguration] class.
type IVZVirtioConsoleDeviceSerialPortConfiguration interface {
	IVZSerialPortConfiguration
	Attachment() VZSerialPortAttachment
	SetAttachment(value IVZSerialPortAttachment)
}

// A configuration object that requests the creation of a console device to communicate with the guest system.
//
// A object enables serial communication between the guest operating system and host computer through the Virtio interface. After you create this configuration object, configure its inherited property with an object that defines the type of serial communication you want to enable. Use a object to enable two-way communication between the guest and host, and use a object to enable one-way communication from the guest to the file you designate.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDeviceSerialPortConfiguration
type VZVirtioConsoleDeviceSerialPortConfiguration struct {
	VZSerialPortConfiguration
}

// VZVirtioConsoleDeviceSerialPortConfigurationFrom constructs a [VZVirtioConsoleDeviceSerialPortConfiguration] from an unsafe.Pointer.
//
// A configuration object that requests the creation of a console device to communicate with the guest system.
func VZVirtioConsoleDeviceSerialPortConfigurationFrom(ptr unsafe.Pointer) VZVirtioConsoleDeviceSerialPortConfiguration {
	return VZVirtioConsoleDeviceSerialPortConfiguration{
		VZSerialPortConfiguration: VZSerialPortConfigurationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZVirtioConsoleDeviceSerialPortConfigurationClass) Alloc() VZVirtioConsoleDeviceSerialPortConfiguration {
	rv := objc.Send[VZVirtioConsoleDeviceSerialPortConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZVirtioConsoleDeviceSerialPortConfigurationClass) New() VZVirtioConsoleDeviceSerialPortConfiguration {
	rv := objc.Send[VZVirtioConsoleDeviceSerialPortConfiguration](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioConsoleDeviceSerialPortConfiguration) Init() VZVirtioConsoleDeviceSerialPortConfiguration {
	rv := objc.Send[VZVirtioConsoleDeviceSerialPortConfiguration](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioConsoleDeviceSerialPortConfiguration) Autorelease() VZVirtioConsoleDeviceSerialPortConfiguration {
	rv := objc.Send[VZVirtioConsoleDeviceSerialPortConfiguration](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioConsoleDeviceSerialPortConfiguration creates a new VZVirtioConsoleDeviceSerialPortConfiguration instance.
func NewVZVirtioConsoleDeviceSerialPortConfiguration() VZVirtioConsoleDeviceSerialPortConfiguration {
	return getVZVirtioConsoleDeviceSerialPortConfigurationClass().New()
}



// The object that defines how the configuration of the virtual machine’s serial port interfaces.
//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzserialportconfiguration/attachment
func (v_ VZVirtioConsoleDeviceSerialPortConfiguration) Attachment() VZSerialPortAttachment {
	rv := objc.Send[VZSerialPortAttachment](v_.ID, objc.Sel("attachment"))
	return rv
}


// SetAttachment sets the value of the attachment property.
// The object that defines how the configuration of the virtual machine’s serial port interfaces.

//
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzserialportconfiguration/attachment
func (v_ VZVirtioConsoleDeviceSerialPortConfiguration) SetAttachment(value IVZSerialPortAttachment) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAttachment:"), value)
}



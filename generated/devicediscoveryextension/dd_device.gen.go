// Code generated from Apple documentation for DeviceDiscoveryExtension. DO NOT EDIT.

package devicediscoveryextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/uniformtypeidentifiers"
)

// The class instance for the [DDDevice] class.
var (
	DDDeviceClass     _DDDeviceClass
	DDDeviceClassOnce sync.Once
)

func getDDDeviceClass() _DDDeviceClass {
	DDDeviceClassOnce.Do(func() {
		DDDeviceClass = _DDDeviceClass{objc.GetClass("DDDevice")}
	})
	return DDDeviceClass
}

type _DDDeviceClass struct {
	class objc.Class
}

// An interface definition for the [DDDevice] class.
type IDDDevice interface {
	objectivec.IObject
	// properties:
	TxtRecord() unsafe.Pointer
	SetTxtRecord(value unsafe.Pointer)
	// methods:
}

// An object that describes a discovered device of interest.
//
// The extension creates an instance of this class for a discovered device of interest and passes it to the system for display in the device picker UI ( ). The extension discovers devices through either Core Bluetooth or the local network (that is, using ). For device discovery extensions of third-party media receivers, an instance of this class corresponds to the media receiver of interest. The extension reports the status of discovered devices to the system using the function, and it receives status updates about the device from the system by implementing .


// An object that describes a discovered device of interest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice
type DDDevice struct {
	objectivec.Object
}

// DDDeviceFrom constructs a [DDDevice] from an unsafe.Pointer.
//
// An object that describes a discovered device of interest.
func DDDeviceFrom(ptr unsafe.Pointer) DDDevice {
	return DDDevice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DDDeviceClass) Alloc() DDDevice {
	rv := objc.Send[DDDevice](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DDDeviceClass) New() DDDevice {
	rv := objc.Send[DDDevice](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DDDevice) Init() DDDevice {
	rv := objc.Send[DDDevice](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DDDevice) Autorelease() DDDevice {
	rv := objc.Send[DDDevice](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDDDevice creates a new DDDevice instance.
func NewDDDevice() DDDevice {
	return getDDDeviceClass().New()
}



// Creates an object that describes a discovered device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/init(displayName:category:protocolType:identifier:)
func NewDDDeviceWithDisplayNameCategoryProtocolTypeIdentifier(displayName objc.IObject /* cross-framework: NSString */, category DDDeviceCategory, protocolType objc.IObject /* cross-framework: UTType */, identifier objc.IObject /* cross-framework: NSString */) DDDevice {
	instance := getDDDeviceClass().Alloc()
	rv := objc.Send[DDDevice](instance.ID, objc.Sel("initWithDisplayName:category:protocolType:identifier:"), displayName, category, protocolType, identifier)
	rv.Autorelease()
	return rv
}



// A dictionary of metadata for the device that the extension communicates with over the local network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/devicediscoveryextension/dddevice/txtrecord
func (d_ DDDevice) TxtRecord() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("txtRecord"))
	return rv
}


// A dictionary of metadata for the device that the extension communicates with over the local network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/devicediscoveryextension/dddevice/txtrecord
func (d_ DDDevice) SetTxtRecord(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTxtRecord:"), value)
}



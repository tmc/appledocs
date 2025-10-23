// Code generated from Apple documentation for ExternalAccessory. DO NOT EDIT.

package externalaccessory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [EAAccessory] class.
var (
	EAAccessoryClass     _EAAccessoryClass
	EAAccessoryClassOnce sync.Once
)

func getEAAccessoryClass() _EAAccessoryClass {
	EAAccessoryClassOnce.Do(func() {
		EAAccessoryClass = _EAAccessoryClass{objc.GetClass("EAAccessory")}
	})
	return EAAccessoryClass
}

type _EAAccessoryClass struct {
	class objc.Class
}

// An interface definition for the [EAAccessory] class.
type IEAAccessory interface {
	objectivec.IObject
	ConnectionID() uint
	DockType() string
	FirmwareRevision() string
	HardwareRevision() string
	Connected() bool
	Manufacturer() string
	ModelNumber() string
	Name() string
	ProtocolStrings() []string
	SerialNumber() string
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	IsConnected() bool
	SetIsConnected(value bool)
}

// An object that contains information about a single, connected hardware accessory.
//
// An object gives your app information about a single connected hardware accessory. You can use the information in this class to determine whether your app is able to open a session to a given accessory. After you have an open session, you can also associate a custom delegate with the accessory object to be notified to changes in the accessory state. Your delegate must adopt the protocol. You use an accessory object to create an object, which itself provides the communications channel to and from the accessory hardware. The accessory object provides information about the communications protocols the accessory supports, along with information about current hardware and firmware revisions. When deciding whether to connect to an accessory, you should always first check the accessory’s declared protocols in the array. This list indicates the types of data the accessory is capable of processing at that moment, which may not be the full list of protocols for which the accessory is designed. For example, an accessory that is connected but not yet authenticated will report no supported protocols until authentication is successful. Don’t connect to the accessory unless and until the list includes the protocol you intend to use. Accessories can be physically connected to the device through the Lightning connector (or through the 30-pin connector on older devices) or wirelessly using Bluetooth.


// An object that contains information about a single, connected hardware accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAAccessory
type EAAccessory struct {
	objectivec.Object
}

// EAAccessoryFrom constructs a [EAAccessory] from an unsafe.Pointer.
//
// An object that contains information about a single, connected hardware accessory.
func EAAccessoryFrom(ptr unsafe.Pointer) EAAccessory {
	return EAAccessory{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _EAAccessoryClass) Alloc() EAAccessory {
	rv := objc.Send[EAAccessory](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EAAccessoryClass) New() EAAccessory {
	rv := objc.Send[EAAccessory](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EAAccessory) Init() EAAccessory {
	rv := objc.Send[EAAccessory](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EAAccessory) Autorelease() EAAccessory {
	rv := objc.Send[EAAccessory](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEAAccessory creates a new EAAccessory instance.
func NewEAAccessory() EAAccessory {
	return getEAAccessoryClass().New()
}



// The accessory’s unique ID for connecting to the iOS-based device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAAccessory/connectionID
func (e_ EAAccessory) ConnectionID() uint {
	rv := objc.Send[uint](e_.ID, objc.Sel("connectionID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAAccessory/dockType
func (e_ EAAccessory) DockType() string {
	rv := objc.Send[string](e_.ID, objc.Sel("dockType"))
	return rv
}


// The current firmware version for the accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAAccessory/firmwareRevision
func (e_ EAAccessory) FirmwareRevision() string {
	rv := objc.Send[string](e_.ID, objc.Sel("firmwareRevision"))
	return rv
}


// The hardware version of the accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAAccessory/hardwareRevision
func (e_ EAAccessory) HardwareRevision() string {
	rv := objc.Send[string](e_.ID, objc.Sel("hardwareRevision"))
	return rv
}


// A Boolean value indicating whether the accessory is currently connected to the iOS-based device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAAccessory/isConnected
func (e_ EAAccessory) Connected() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("connected"))
	return rv
}


// The name of the accessory’s manufacturer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAAccessory/manufacturer
func (e_ EAAccessory) Manufacturer() string {
	rv := objc.Send[string](e_.ID, objc.Sel("manufacturer"))
	return rv
}


// The model information for the accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAAccessory/modelNumber
func (e_ EAAccessory) ModelNumber() string {
	rv := objc.Send[string](e_.ID, objc.Sel("modelNumber"))
	return rv
}


// The display name of the accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAAccessory/name
func (e_ EAAccessory) Name() string {
	rv := objc.Send[string](e_.ID, objc.Sel("name"))
	return rv
}


// The communication protocols supported by the accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAAccessory/protocolStrings
func (e_ EAAccessory) ProtocolStrings() []string {
	rv := objc.Send[[]string](e_.ID, objc.Sel("protocolStrings"))
	return rv
}


// The serial number of the accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAAccessory/serialNumber
func (e_ EAAccessory) SerialNumber() string {
	rv := objc.Send[string](e_.ID, objc.Sel("serialNumber"))
	return rv
}


// The object that acts as the delegate of the accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/externalaccessory/eaaccessory/delegate
func (e_ EAAccessory) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("delegate"))
	return rv
}


// The object that acts as the delegate of the accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/externalaccessory/eaaccessory/delegate
func (e_ EAAccessory) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setDelegate:"), value)
}


// A Boolean value indicating whether the accessory is currently connected to the iOS-based device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/externalaccessory/eaaccessory/isconnected
func (e_ EAAccessory) IsConnected() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isConnected"))
	return rv
}


// A Boolean value indicating whether the accessory is currently connected to the iOS-based device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/externalaccessory/eaaccessory/isconnected
func (e_ EAAccessory) SetIsConnected(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setIsConnected:"), value)
}




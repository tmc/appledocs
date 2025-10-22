// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CWWiFiClient] class.
var (
	CWWiFiClientClass     _CWWiFiClientClass
	CWWiFiClientClassOnce sync.Once
)

func getCWWiFiClientClass() _CWWiFiClientClass {
	CWWiFiClientClassOnce.Do(func() {
		CWWiFiClientClass = _CWWiFiClientClass{objc.GetClass("CWWiFiClient")}
	})
	return CWWiFiClientClass
}

type _CWWiFiClientClass struct {
	class objc.Class
}

// An interface definition for the [CWWiFiClient] class.
type ICWWiFiClient interface {
	objectivec.IObject
	Interface() CWInterface
	InterfaceWithName(interfaceName string) CWInterface
	InterfaceNames() []string
	Interfaces() []CWInterface
	StartMonitoringEventWithTypeError(type_ CWEventType, error_ unsafe.Pointer) bool
	StopMonitoringAllEventsAndReturnError(error_ unsafe.Pointer) bool
	StopMonitoringEventWithTypeError(type_ CWEventType, error_ unsafe.Pointer) bool
	Delegate() objc.ID
	SetDelegate(value objc.ID)
}

// A wrapper around the entire Wi-Fi subsystem that you use to access interfaces and set up event notifications.
//
// Wi-Fi client objects are heavy. Therefore, it’s more efficient to use a single, long-running client instance, rather than creating several short-lived instances. For convenience, you can use the singleton instance returned by the class method. Instead of instantiating objects directly, use the ones provided by the instance methods of this class. For example, the method returns the default Wi-Fi interface.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient
type CWWiFiClient struct {
	objectivec.Object
}

// CWWiFiClientFrom constructs a [CWWiFiClient] from an unsafe.Pointer.
//
// A wrapper around the entire Wi-Fi subsystem that you use to access interfaces and set up event notifications.
func CWWiFiClientFrom(ptr unsafe.Pointer) CWWiFiClient {
	return CWWiFiClient{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CWWiFiClientClass) Alloc() CWWiFiClient {
	rv := objc.Send[CWWiFiClient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CWWiFiClientClass) New() CWWiFiClient {
	rv := objc.Send[CWWiFiClient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CWWiFiClient) Init() CWWiFiClient {
	rv := objc.Send[CWWiFiClient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CWWiFiClient) Autorelease() CWWiFiClient {
	rv := objc.Send[CWWiFiClient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCWWiFiClient creates a new CWWiFiClient instance.
func NewCWWiFiClient() CWWiFiClient {
	return getCWWiFiClientClass().New()
}



// Returns the list of the names of available Wi-Fi interfaces.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/interfaceNames()-swift.type.method
func (cc _CWWiFiClientClass) InterfaceNames() []string {
	rv := objc.Send[[]string](objc.ID(cc.class), objc.Sel("interfaceNames"))
	return rv
}

// The shared Wi-Fi client object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/shared()
func (cc _CWWiFiClientClass) SharedWiFiClient() CWWiFiClient {
	rv := objc.Send[CWWiFiClient](objc.ID(cc.class), objc.Sel("sharedWiFiClient"))
	return rv
}

// Returns the default Wi-Fi interface.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/interface()
func (c_ CWWiFiClient) Interface() CWInterface {
	rv := objc.Send[CWInterface](c_.ID, objc.Sel("interface"))
	return rv
}

// Returns the Wi-Fi interface with the given name.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/interface(withName:)
func (c_ CWWiFiClient) InterfaceWithName(interfaceName string) CWInterface {
	rv := objc.Send[CWInterface](c_.ID, objc.Sel("interfaceWithName:"), objc.String(interfaceName))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/interfaceNames()-swift.method
func (c_ CWWiFiClient) InterfaceNames() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("interfaceNames"))
	return rv
}

// Returns all available Wi-Fi interfaces.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/interfaces()
func (c_ CWWiFiClient) Interfaces() []CWInterface {
	rv := objc.Send[[]CWInterface](c_.ID, objc.Sel("interfaces"))
	return rv
}

// Register for specific Wi-Fi event notifications.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/startMonitoringEvent(with:)
func (c_ CWWiFiClient) StartMonitoringEventWithTypeError(type_ CWEventType, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("startMonitoringEventWithType:error:"), type_, error_)
	return rv
}

// Unregister for all Wi-Fi event notifications.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/stopMonitoringAllEvents()
func (c_ CWWiFiClient) StopMonitoringAllEventsAndReturnError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("stopMonitoringAllEventsAndReturnError:"), error_)
	return rv
}

// Unregister for specific Wi-Fi event notifications.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/stopMonitoringEvent(with:)
func (c_ CWWiFiClient) StopMonitoringEventWithTypeError(type_ CWEventType, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("stopMonitoringEventWithType:error:"), type_, error_)
	return rv
}

// An object that provides Wi-Fi event handling.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/delegate
func (c_ CWWiFiClient) Delegate() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// An object that provides Wi-Fi event handling.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/delegate
func (c_ CWWiFiClient) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}



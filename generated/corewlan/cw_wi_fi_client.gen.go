// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CWWiFiClient */


/* debug [class_header]: Header for CWWiFiClient */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CWWiFiClient */
// An interface definition for the [CWWiFiClient] class.
type ICWWiFiClient interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CWWiFiClient */
	// properties:
	Delegate() objc.ID
	SetDelegate(value objc.ID)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CWWiFiClient */
	// methods:
	Interface() ICWInterface
	InterfaceWithName(interfaceName objc.IObject /* cross-framework: NSString */) ICWInterface
	InterfaceNames() []string
	Interfaces() []CWInterface
	StartMonitoringEventWithTypeError(type_ CWEventType, error_ unsafe.Pointer) bool
	StopMonitoringAllEventsAndReturnError(error_ unsafe.Pointer) bool
	StopMonitoringEventWithTypeError(type_ CWEventType, error_ unsafe.Pointer) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CWWiFiClient */
// Alloc allocates a new instance without initialization.
func (cc _CWWiFiClientClass) Alloc() CWWiFiClient {
	rv := objc.Send[CWWiFiClient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CWWiFiClient */
// A wrapper around the entire Wi-Fi subsystem that you use to access interfaces and set up event notifications.
//
// Wi-Fi client objects are heavy. Therefore, it’s more efficient to use a single, long-running client instance, rather than creating several short-lived instances. For convenience, you can use the singleton instance returned by the class method. Instead of instantiating objects directly, use the ones provided by the instance methods of this class. For example, the method returns the default Wi-Fi interface.


// A wrapper around the entire Wi-Fi subsystem that you use to access interfaces and set up event notifications.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CWWiFiClient */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CWWiFiClient */

// Returns the list of the names of available Wi-Fi interfaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/interfaceNames()-swift.type.method
func (cc _CWWiFiClientClass) InterfaceNames() []string {
	rv := objc.Send[[]string](objc.ID(cc.class), objc.Sel("interfaceNames"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=InterfaceNames) */


// The shared Wi-Fi client object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/shared()
func (cc _CWWiFiClientClass) SharedWiFiClient() CWWiFiClient {
	rv := objc.Send[CWWiFiClient](objc.ID(cc.class), objc.Sel("sharedWiFiClient"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedWiFiClient) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CWWiFiClient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CWWiFiClient */

// Returns the default Wi-Fi interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/interface()
func (c_ CWWiFiClient) Interface() ICWInterface {
	rv := objc.Send[CWInterface](c_.ID, objc.Sel("interface"))
	return rv
}/* debug [instance_methods/method]: Interface */


// Returns the Wi-Fi interface with the given name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/interface(withName:)
func (c_ CWWiFiClient) InterfaceWithName(interfaceName objc.IObject /* cross-framework: NSString */) ICWInterface {
	rv := objc.Send[CWInterface](c_.ID, objc.Sel("interfaceWithName:"), interfaceName)
	return rv
}/* debug [instance_methods/method]: InterfaceWithName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/interfaceNames()-swift.method
func (c_ CWWiFiClient) InterfaceNames() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("interfaceNames"))
	return rv
}/* debug [instance_methods/method]: InterfaceNames */


// Returns all available Wi-Fi interfaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/interfaces()
func (c_ CWWiFiClient) Interfaces() []CWInterface {
	rv := objc.Send[[]CWInterface](c_.ID, objc.Sel("interfaces"))
	return rv
}/* debug [instance_methods/method]: Interfaces */


// Register for specific Wi-Fi event notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/startMonitoringEvent(with:)
func (c_ CWWiFiClient) StartMonitoringEventWithTypeError(type_ CWEventType, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("startMonitoringEventWithType:error:"), type_, error_)
	return rv
}/* debug [instance_methods/method]: StartMonitoringEventWithTypeError */


// Unregister for all Wi-Fi event notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/stopMonitoringAllEvents()
func (c_ CWWiFiClient) StopMonitoringAllEventsAndReturnError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("stopMonitoringAllEventsAndReturnError:"), error_)
	return rv
}/* debug [instance_methods/method]: StopMonitoringAllEventsAndReturnError */


// Unregister for specific Wi-Fi event notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/stopMonitoringEvent(with:)
func (c_ CWWiFiClient) StopMonitoringEventWithTypeError(type_ CWEventType, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("stopMonitoringEventWithType:error:"), type_, error_)
	return rv
}/* debug [instance_methods/method]: StopMonitoringEventWithTypeError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CWWiFiClient */

// An object that provides Wi-Fi event handling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/delegate
func (c_ CWWiFiClient) Delegate() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// An object that provides Wi-Fi event handling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWWiFiClient/delegate
func (c_ CWWiFiClient) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CWWiFiClient */



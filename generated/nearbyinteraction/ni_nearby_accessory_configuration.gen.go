// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class NINearbyAccessoryConfiguration */


/* debug [class_header]: Header for NINearbyAccessoryConfiguration */
// The class instance for the [NINearbyAccessoryConfiguration] class.
var (
	NINearbyAccessoryConfigurationClass     _NINearbyAccessoryConfigurationClass
	NINearbyAccessoryConfigurationClassOnce sync.Once
)

func getNINearbyAccessoryConfigurationClass() _NINearbyAccessoryConfigurationClass {
	NINearbyAccessoryConfigurationClassOnce.Do(func() {
		NINearbyAccessoryConfigurationClass = _NINearbyAccessoryConfigurationClass{objc.GetClass("NINearbyAccessoryConfiguration")}
	})
	return NINearbyAccessoryConfigurationClass
}

type _NINearbyAccessoryConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NINearbyAccessoryConfiguration */
// An interface definition for the [NINearbyAccessoryConfiguration] class.
type ININearbyAccessoryConfiguration interface {
	INIConfiguration
	
/* debug [class_interface_properties]: Properties for NINearbyAccessoryConfiguration */
	// properties:
	IsCameraAssistanceEnabled() bool
	SetIsCameraAssistanceEnabled(value bool)
	Distance() float32
	SetDistance(value float32)
	Delegate() objc.IObject /* cross-framework: NISessionDelegate */
	SetDelegate(value objc.IObject /* cross-framework: NISessionDelegate */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NINearbyAccessoryConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NINearbyAccessoryConfiguration */
// Alloc allocates a new instance without initialization.
func (nc _NINearbyAccessoryConfigurationClass) Alloc() NINearbyAccessoryConfiguration {
	rv := objc.Send[NINearbyAccessoryConfiguration](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NINearbyAccessoryConfigurationClass) New() NINearbyAccessoryConfiguration {
	rv := objc.Send[NINearbyAccessoryConfiguration](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NINearbyAccessoryConfiguration) Init() NINearbyAccessoryConfiguration {
	rv := objc.Send[NINearbyAccessoryConfiguration](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NINearbyAccessoryConfiguration) Autorelease() NINearbyAccessoryConfiguration {
	rv := objc.Send[NINearbyAccessoryConfiguration](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNINearbyAccessoryConfiguration creates a new NINearbyAccessoryConfiguration instance.
func NewNINearbyAccessoryConfiguration() NINearbyAccessoryConfiguration {
	return getNINearbyAccessoryConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NINearbyAccessoryConfiguration */
// A configuration that enables interaction between iPhone and third-party accessories.
//
// Use this class to interact with a third-party accessory that you partner with or develop. For an example app that demonstrates this configuration, see .


// A configuration that enables interaction between iPhone and third-party accessories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyAccessoryConfiguration
type NINearbyAccessoryConfiguration struct {
	NIConfiguration
}

// NINearbyAccessoryConfigurationFrom constructs a [NINearbyAccessoryConfiguration] from an unsafe.Pointer.
//
// A configuration that enables interaction between iPhone and third-party accessories.
func NINearbyAccessoryConfigurationFrom(ptr unsafe.Pointer) NINearbyAccessoryConfiguration {
	return NINearbyAccessoryConfiguration{
		NIConfiguration: NIConfigurationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NINearbyAccessoryConfiguration */

// Creates a configuration for an accessory with the given Bluetooth peer identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyAccessoryConfiguration/init(accessoryData:bluetoothPeerIdentifier:)
func NewNINearbyAccessoryConfigurationWithAccessoryDataBluetoothPeerIdentifierError(accessoryData objc.IObject /* cross-framework: NSData */, identifier foundation.UUID, error_ unsafe.Pointer) NINearbyAccessoryConfiguration {
	instance := getNINearbyAccessoryConfigurationClass().Alloc()
	rv := objc.Send[NINearbyAccessoryConfiguration](instance.ID, objc.Sel("initWithAccessoryData:bluetoothPeerIdentifier:error:"), accessoryData, identifier, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNINearbyAccessoryConfigurationWithAccessoryDataBluetoothPeerIdentifierError */


// Creates a configuration for interaction between iPhone and third-party accessories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyAccessoryConfiguration/init(data:)
func NewNINearbyAccessoryConfigurationWithDataError(data objc.IObject /* cross-framework: NSData */, error_ unsafe.Pointer) NINearbyAccessoryConfiguration {
	instance := getNINearbyAccessoryConfigurationClass().Alloc()
	rv := objc.Send[NINearbyAccessoryConfiguration](instance.ID, objc.Sel("initWithData:error:"), data, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNINearbyAccessoryConfigurationWithDataError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NINearbyAccessoryConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NINearbyAccessoryConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NINearbyAccessoryConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NINearbyAccessoryConfiguration */

// A Boolean value that combines the spatial awareness of ARKit with Nearby Interaction to improve the accuracy of a nearby object’s position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/ninearbyaccessoryconfiguration/iscameraassistanceenabled
func (n_ NINearbyAccessoryConfiguration) IsCameraAssistanceEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isCameraAssistanceEnabled"))
	return rv
}/* debug [instance_properties/getter]: isCameraAssistanceEnabled */


// A Boolean value that combines the spatial awareness of ARKit with Nearby Interaction to improve the accuracy of a nearby object’s position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/ninearbyaccessoryconfiguration/iscameraassistanceenabled
func (n_ NINearbyAccessoryConfiguration) SetIsCameraAssistanceEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsCameraAssistanceEnabled:"), value)
}/* debug [instance_properties/setter]: isCameraAssistanceEnabled */


// The distance from the user’s device to the peer device in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/ninearbyobject/distance-676dm
func (n_ NINearbyAccessoryConfiguration) Distance() float32 {
	rv := objc.Send[float32](n_.ID, objc.Sel("distance"))
	return rv
}/* debug [instance_properties/getter]: distance */


// The distance from the user’s device to the peer device in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/ninearbyobject/distance-676dm
func (n_ NINearbyAccessoryConfiguration) SetDistance(value float32) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDistance:"), value)
}/* debug [instance_properties/setter]: distance */


// An object that the framework notifies of session events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nisession/delegate
func (n_ NINearbyAccessoryConfiguration) Delegate() objc.IObject /* cross-framework: NISessionDelegate */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// An object that the framework notifies of session events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nisession/delegate
func (n_ NINearbyAccessoryConfiguration) SetDelegate(value objc.IObject /* cross-framework: NISessionDelegate */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NINearbyAccessoryConfiguration */



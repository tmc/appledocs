// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NINearbyPeerConfiguration */


/* debug [class_header]: Header for NINearbyPeerConfiguration */
// The class instance for the [NINearbyPeerConfiguration] class.
var (
	NINearbyPeerConfigurationClass     _NINearbyPeerConfigurationClass
	NINearbyPeerConfigurationClassOnce sync.Once
)

func getNINearbyPeerConfigurationClass() _NINearbyPeerConfigurationClass {
	NINearbyPeerConfigurationClassOnce.Do(func() {
		NINearbyPeerConfigurationClass = _NINearbyPeerConfigurationClass{objc.GetClass("NINearbyPeerConfiguration")}
	})
	return NINearbyPeerConfigurationClass
}

type _NINearbyPeerConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NINearbyPeerConfiguration */
// An interface definition for the [NINearbyPeerConfiguration] class.
type ININearbyPeerConfiguration interface {
	INIConfiguration
	
/* debug [class_interface_properties]: Properties for NINearbyPeerConfiguration */
	// properties:
	IsCameraAssistanceEnabled() bool
	SetIsCameraAssistanceEnabled(value bool)
	IsExtendedDistanceMeasurementEnabled() bool
	SetIsExtendedDistanceMeasurementEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NINearbyPeerConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NINearbyPeerConfiguration */
// Alloc allocates a new instance without initialization.
func (nc _NINearbyPeerConfigurationClass) Alloc() NINearbyPeerConfiguration {
	rv := objc.Send[NINearbyPeerConfiguration](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NINearbyPeerConfigurationClass) New() NINearbyPeerConfiguration {
	rv := objc.Send[NINearbyPeerConfiguration](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NINearbyPeerConfiguration) Init() NINearbyPeerConfiguration {
	rv := objc.Send[NINearbyPeerConfiguration](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NINearbyPeerConfiguration) Autorelease() NINearbyPeerConfiguration {
	rv := objc.Send[NINearbyPeerConfiguration](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNINearbyPeerConfiguration creates a new NINearbyPeerConfiguration instance.
func NewNINearbyPeerConfiguration() NINearbyPeerConfiguration {
	return getNINearbyPeerConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NINearbyPeerConfiguration */
// A configuration that enables interaction between iPhone or Apple Watch devices.
//
// A peer interaction session enables two Apple devices to share their respective distance and direction through the device’s Ultra Wideband (UWB) chip. To start a peer interaction session, create a instance and pass it to an instance with the function. For an example app that demonstrates this configuration, see .


// A configuration that enables interaction between iPhone or Apple Watch devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyPeerConfiguration
type NINearbyPeerConfiguration struct {
	NIConfiguration
}

// NINearbyPeerConfigurationFrom constructs a [NINearbyPeerConfiguration] from an unsafe.Pointer.
//
// A configuration that enables interaction between iPhone or Apple Watch devices.
func NINearbyPeerConfigurationFrom(ptr unsafe.Pointer) NINearbyPeerConfiguration {
	return NINearbyPeerConfiguration{
		NIConfiguration: NIConfigurationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NINearbyPeerConfiguration */

// Creates a configuration for interaction between devices, including iPhone and Apple Watch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyPeerConfiguration/init(peerToken:)
func NewNINearbyPeerConfigurationWithPeerToken(peerToken INIDiscoveryToken) NINearbyPeerConfiguration {
	instance := getNINearbyPeerConfigurationClass().Alloc()
	rv := objc.Send[NINearbyPeerConfiguration](instance.ID, objc.Sel("initWithPeerToken:"), peerToken)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNINearbyPeerConfigurationWithPeerToken */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NINearbyPeerConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NINearbyPeerConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NINearbyPeerConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NINearbyPeerConfiguration */

// A Boolean value that combines the spatial awareness of ARKit with Nearby Interaction to improve the accuracy of a nearby object’s position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/ninearbypeerconfiguration/iscameraassistanceenabled
func (n_ NINearbyPeerConfiguration) IsCameraAssistanceEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isCameraAssistanceEnabled"))
	return rv
}/* debug [instance_properties/getter]: isCameraAssistanceEnabled */


// A Boolean value that combines the spatial awareness of ARKit with Nearby Interaction to improve the accuracy of a nearby object’s position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/ninearbypeerconfiguration/iscameraassistanceenabled
func (n_ NINearbyPeerConfiguration) SetIsCameraAssistanceEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsCameraAssistanceEnabled:"), value)
}/* debug [instance_properties/setter]: isCameraAssistanceEnabled */


// A Boolean value that indicates whether both peers can use extended distance measurement for this Nearby Interaction session instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/ninearbypeerconfiguration/isextendeddistancemeasurementenabled
func (n_ NINearbyPeerConfiguration) IsExtendedDistanceMeasurementEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isExtendedDistanceMeasurementEnabled"))
	return rv
}/* debug [instance_properties/getter]: isExtendedDistanceMeasurementEnabled */


// A Boolean value that indicates whether both peers can use extended distance measurement for this Nearby Interaction session instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/ninearbypeerconfiguration/isextendeddistancemeasurementenabled
func (n_ NINearbyPeerConfiguration) SetIsExtendedDistanceMeasurementEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsExtendedDistanceMeasurementEnabled:"), value)
}/* debug [instance_properties/setter]: isExtendedDistanceMeasurementEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NINearbyPeerConfiguration */



// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [NINearbyPeerConfiguration] class.
type ININearbyPeerConfiguration interface {
	INIConfiguration
}

// A configuration that enables interaction between iPhone or Apple Watch devices.
//
// A peer interaction session enables two Apple devices to share their respective distance and direction through the device’s Ultra Wideband (UWB) chip. To start a peer interaction session, create a instance and pass it to an instance with the function. For an example app that demonstrates this configuration, see .
//
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

// Alloc allocates a new instance without initialization.
func (nc _NINearbyPeerConfigurationClass) Alloc() NINearbyPeerConfiguration {
	rv := objc.Send[NINearbyPeerConfiguration](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// A Boolean value that combines the spatial awareness of ARKit with Nearby Interaction to improve the accuracy of a nearby object’s position.
//
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyPeerConfiguration/isCameraAssistanceEnabled
func (n_ NINearbyPeerConfiguration) CameraAssistanceEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("cameraAssistanceEnabled"))
	return rv
}


// SetCameraAssistanceEnabled sets the value of the cameraAssistanceEnabled property.
// A Boolean value that combines the spatial awareness of ARKit with Nearby Interaction to improve the accuracy of a nearby object’s position.

//
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyPeerConfiguration/isCameraAssistanceEnabled
func (n_ NINearbyPeerConfiguration) SetCameraAssistanceEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setCameraAssistanceEnabled:"), value)
}

// A Boolean value that indicates whether both peers can use extended distance measurement for this Nearby Interaction session instance.
//
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyPeerConfiguration/isExtendedDistanceMeasurementEnabled
func (n_ NINearbyPeerConfiguration) ExtendedDistanceMeasurementEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("extendedDistanceMeasurementEnabled"))
	return rv
}


// SetExtendedDistanceMeasurementEnabled sets the value of the extendedDistanceMeasurementEnabled property.
// A Boolean value that indicates whether both peers can use extended distance measurement for this Nearby Interaction session instance.

//
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyPeerConfiguration/isExtendedDistanceMeasurementEnabled
func (n_ NINearbyPeerConfiguration) SetExtendedDistanceMeasurementEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setExtendedDistanceMeasurementEnabled:"), value)
}

// A value that uniquely identifies the other peer in the interaction session.
//
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyPeerConfiguration/peerDiscoveryToken
func (n_ NINearbyPeerConfiguration) PeerDiscoveryToken() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("peerDiscoveryToken"))
	return rv
}

// A Boolean value that combines the spatial awareness of ARKit with Nearby Interaction to improve the accuracy of a nearby object’s position.
//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/ninearbypeerconfiguration/iscameraassistanceenabled
func (n_ NINearbyPeerConfiguration) IsCameraAssistanceEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isCameraAssistanceEnabled"))
	return rv
}


// SetIsCameraAssistanceEnabled sets the value of the isCameraAssistanceEnabled property.
// A Boolean value that combines the spatial awareness of ARKit with Nearby Interaction to improve the accuracy of a nearby object’s position.

//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/ninearbypeerconfiguration/iscameraassistanceenabled
func (n_ NINearbyPeerConfiguration) SetIsCameraAssistanceEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsCameraAssistanceEnabled:"), value)
}

// A Boolean value that indicates whether both peers can use extended distance measurement for this Nearby Interaction session instance.
//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/ninearbypeerconfiguration/isextendeddistancemeasurementenabled
func (n_ NINearbyPeerConfiguration) IsExtendedDistanceMeasurementEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isExtendedDistanceMeasurementEnabled"))
	return rv
}


// SetIsExtendedDistanceMeasurementEnabled sets the value of the isExtendedDistanceMeasurementEnabled property.
// A Boolean value that indicates whether both peers can use extended distance measurement for this Nearby Interaction session instance.

//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/ninearbypeerconfiguration/isextendeddistancemeasurementenabled
func (n_ NINearbyPeerConfiguration) SetIsExtendedDistanceMeasurementEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsExtendedDistanceMeasurementEnabled:"), value)
}




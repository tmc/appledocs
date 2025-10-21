// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [NINearbyAccessoryConfiguration] class.
type ININearbyAccessoryConfiguration interface {
	INIConfiguration
}

// A configuration that enables interaction between iPhone and third-party accessories.
//
// Use this class to interact with a third-party accessory that you partner with or develop. For an example app that demonstrates this configuration, see .
//
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

// Alloc allocates a new instance without initialization.
func (nc _NINearbyAccessoryConfigurationClass) Alloc() NINearbyAccessoryConfiguration {
	rv := objc.Send[NINearbyAccessoryConfiguration](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a configuration for an accessory with the given Bluetooth peer identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyAccessoryConfiguration/init(accessoryData:bluetoothPeerIdentifier:)
func NewNINearbyAccessoryConfigurationWithAccessoryDataBluetoothPeerIdentifierError(accessoryData unsafe.Pointer, identifier unsafe.Pointer, error_ unsafe.Pointer) NINearbyAccessoryConfiguration {
	instance := getNINearbyAccessoryConfigurationClass().Alloc()
	rv := objc.Send[NINearbyAccessoryConfiguration](instance.ID, objc.Sel("initWithAccessoryData:bluetoothPeerIdentifier:error:"), accessoryData, identifier, error_)
	rv.Autorelease()
	return rv
}



// Creates a configuration for interaction between iPhone and third-party accessories.
//
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyAccessoryConfiguration/init(data:)
func NewNINearbyAccessoryConfigurationWithDataError(data unsafe.Pointer, error_ unsafe.Pointer) NINearbyAccessoryConfiguration {
	instance := getNINearbyAccessoryConfigurationClass().Alloc()
	rv := objc.Send[NINearbyAccessoryConfiguration](instance.ID, objc.Sel("initWithData:error:"), data, error_)
	rv.Autorelease()
	return rv
}


// An object that the framework notifies of session events.
//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nisession/delegate
func (n_ NINearbyAccessoryConfiguration) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// An object that the framework notifies of session events.

//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nisession/delegate
func (n_ NINearbyAccessoryConfiguration) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that combines the spatial awareness of ARKit with Nearby Interaction to improve the accuracy of a nearby object’s position.
//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/ninearbyaccessoryconfiguration/iscameraassistanceenabled
func (n_ NINearbyAccessoryConfiguration) IsCameraAssistanceEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isCameraAssistanceEnabled"))
	return rv
}


// SetIsCameraAssistanceEnabled sets the value of the isCameraAssistanceEnabled property.
// A Boolean value that combines the spatial awareness of ARKit with Nearby Interaction to improve the accuracy of a nearby object’s position.

//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/ninearbyaccessoryconfiguration/iscameraassistanceenabled
func (n_ NINearbyAccessoryConfiguration) SetIsCameraAssistanceEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsCameraAssistanceEnabled:"), value)
}

// The distance from the user’s device to the peer device in meters.
//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/ninearbyobject/distance-676dm
func (n_ NINearbyAccessoryConfiguration) Distance() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("distance"))
	return rv
}


// SetDistance sets the value of the distance property.
// The distance from the user’s device to the peer device in meters.

//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/ninearbyobject/distance-676dm
func (n_ NINearbyAccessoryConfiguration) SetDistance(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDistance:"), value)
}

// An identifier for the accessory in a session.
//
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyAccessoryConfiguration/accessoryDiscoveryToken
func (n_ NINearbyAccessoryConfiguration) AccessoryDiscoveryToken() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("accessoryDiscoveryToken"))
	return rv
}

// A Boolean value that combines the spatial awareness of ARKit with Nearby Interaction to improve the accuracy of a nearby object’s position.
//
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyAccessoryConfiguration/isCameraAssistanceEnabled
func (n_ NINearbyAccessoryConfiguration) CameraAssistanceEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("cameraAssistanceEnabled"))
	return rv
}


// SetCameraAssistanceEnabled sets the value of the cameraAssistanceEnabled property.
// A Boolean value that combines the spatial awareness of ARKit with Nearby Interaction to improve the accuracy of a nearby object’s position.

//
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyAccessoryConfiguration/isCameraAssistanceEnabled
func (n_ NINearbyAccessoryConfiguration) SetCameraAssistanceEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setCameraAssistanceEnabled:"), value)
}



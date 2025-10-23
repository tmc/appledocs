// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CWConfiguration] class.
var (
	CWConfigurationClass     _CWConfigurationClass
	CWConfigurationClassOnce sync.Once
)

func getCWConfigurationClass() _CWConfigurationClass {
	CWConfigurationClassOnce.Do(func() {
		CWConfigurationClass = _CWConfigurationClass{objc.GetClass("CWConfiguration")}
	})
	return CWConfigurationClass
}

type _CWConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [CWConfiguration] class.
type ICWConfiguration interface {
	objectivec.IObject
	// properties:
	NetworkProfiles() unsafe.Pointer
	RememberJoinedNetworks() bool /* primitive/slice/pointer. */
	RequireAdministratorForAssociation() bool /* primitive/slice/pointer. */
	RequireAdministratorForIBSSMode() bool /* primitive/slice/pointer. */
	RequireAdministratorForPower() bool /* primitive/slice/pointer. */
	// methods:
	IsEqualToConfiguration(configuration ICWConfiguration) bool /* primitive/slice/pointer. */
}

// Encapsulates an immutable configuration for an AirPort WLAN interface.


// Encapsulates an immutable configuration for an AirPort WLAN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration
type CWConfiguration struct {
	objectivec.Object
}

// CWConfigurationFrom constructs a [CWConfiguration] from an unsafe.Pointer.
//
// Encapsulates an immutable configuration for an AirPort WLAN interface.
func CWConfigurationFrom(ptr unsafe.Pointer) CWConfiguration {
	return CWConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CWConfigurationClass) Alloc() CWConfiguration {
	rv := objc.Send[CWConfiguration](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CWConfigurationClass) New() CWConfiguration {
	rv := objc.Send[CWConfiguration](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CWConfiguration) Init() CWConfiguration {
	rv := objc.Send[CWConfiguration](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CWConfiguration) Autorelease() CWConfiguration {
	rv := objc.Send[CWConfiguration](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCWConfiguration creates a new CWConfiguration instance.
func NewCWConfiguration() CWConfiguration {
	return getCWConfigurationClass().New()
}



// Creates and returns a CWConfiguration object initialized with the given CWConfiguration object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration/init(configuration:)
func NewCWConfigurationWithConfiguration(configuration ICWConfiguration) CWConfiguration {
	instance := getCWConfigurationClass().Alloc()
	rv := objc.Send[CWConfiguration](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	rv.Autorelease()
	return rv
}



// Convenience method for getting an empty CWConfiguration object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration/configuration
func (cc _CWConfigurationClass) Configuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("configuration"))
	return rv
}


// Convenience method for getting a CWConfiguration object initialized with the given CWConfiguration object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration/configurationWithConfiguration:
func (cc _CWConfigurationClass) ConfigurationWithConfiguration(configuration ICWConfiguration) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("configurationWithConfiguration:"), configuration)
	return rv
}


// Determine CWConfiguration object equality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration/isEqual(to:)
func (c_ CWConfiguration) IsEqualToConfiguration(configuration ICWConfiguration) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEqualToConfiguration:"), configuration)
	return rv
}


// An array of remembered CWNetworkProfile objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration/networkProfiles
func (c_ CWConfiguration) NetworkProfiles() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("networkProfiles"))
	return rv
}


// AirPort client will remember all joined networks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration/rememberJoinedNetworks
func (c_ CWConfiguration) RememberJoinedNetworks() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("rememberJoinedNetworks"))
	return rv
}


// Require an administrator password to change networks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration/requireAdministratorForAssociation
func (c_ CWConfiguration) RequireAdministratorForAssociation() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("requireAdministratorForAssociation"))
	return rv
}


// Require an administrator password to create a computer-to-computer network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration/requireAdministratorForIBSSMode
func (c_ CWConfiguration) RequireAdministratorForIBSSMode() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("requireAdministratorForIBSSMode"))
	return rv
}


// Require an administrator password to change the interface power state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWConfiguration/requireAdministratorForPower
func (c_ CWConfiguration) RequireAdministratorForPower() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("requireAdministratorForPower"))
	return rv
}



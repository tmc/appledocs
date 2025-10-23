// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CWMutableConfiguration] class.
var (
	CWMutableConfigurationClass     _CWMutableConfigurationClass
	CWMutableConfigurationClassOnce sync.Once
)

func getCWMutableConfigurationClass() _CWMutableConfigurationClass {
	CWMutableConfigurationClassOnce.Do(func() {
		CWMutableConfigurationClass = _CWMutableConfigurationClass{objc.GetClass("CWMutableConfiguration")}
	})
	return CWMutableConfigurationClass
}

type _CWMutableConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [CWMutableConfiguration] class.
type ICWMutableConfiguration interface {
	ICWConfiguration
	NetworkProfiles() unsafe.Pointer
	SetNetworkProfiles(value unsafe.Pointer)
	RememberJoinedNetworks() bool
	SetRememberJoinedNetworks(value bool)
	RequireAdministratorForAssociation() bool
	SetRequireAdministratorForAssociation(value bool)
	RequireAdministratorForIBSSMode() bool
	SetRequireAdministratorForIBSSMode(value bool)
	RequireAdministratorForPower() bool
	SetRequireAdministratorForPower(value bool)
}

// Encapsulates a mutable configuration for an AirPort WLAN interface.
//
// Use this class to change configuration settings or the preferred networks list. To commit configuration changes, use .


// Encapsulates a mutable configuration for an AirPort WLAN interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWMutableConfiguration
type CWMutableConfiguration struct {
	CWConfiguration
}

// CWMutableConfigurationFrom constructs a [CWMutableConfiguration] from an unsafe.Pointer.
//
// Encapsulates a mutable configuration for an AirPort WLAN interface.
func CWMutableConfigurationFrom(ptr unsafe.Pointer) CWMutableConfiguration {
	return CWMutableConfiguration{
		CWConfiguration: CWConfigurationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CWMutableConfigurationClass) Alloc() CWMutableConfiguration {
	rv := objc.Send[CWMutableConfiguration](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CWMutableConfigurationClass) New() CWMutableConfiguration {
	rv := objc.Send[CWMutableConfiguration](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CWMutableConfiguration) Init() CWMutableConfiguration {
	rv := objc.Send[CWMutableConfiguration](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CWMutableConfiguration) Autorelease() CWMutableConfiguration {
	rv := objc.Send[CWMutableConfiguration](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCWMutableConfiguration creates a new CWMutableConfiguration instance.
func NewCWMutableConfiguration() CWMutableConfiguration {
	return getCWMutableConfigurationClass().New()
}



// The preferred networks list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWMutableConfiguration/networkProfiles
func (c_ CWMutableConfiguration) NetworkProfiles() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("networkProfiles"))
	return rv
}


// The preferred networks list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWMutableConfiguration/networkProfiles
func (c_ CWMutableConfiguration) SetNetworkProfiles(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNetworkProfiles:"), value)
}


// A Boolean value that determines whether to remember all joined Wi-Fi networks unless the user specifies otherwise when joining a particular Wi-Fi network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWMutableConfiguration/rememberJoinedNetworks
func (c_ CWMutableConfiguration) RememberJoinedNetworks() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("rememberJoinedNetworks"))
	return rv
}


// A Boolean value that determines whether to remember all joined Wi-Fi networks unless the user specifies otherwise when joining a particular Wi-Fi network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWMutableConfiguration/rememberJoinedNetworks
func (c_ CWMutableConfiguration) SetRememberJoinedNetworks(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRememberJoinedNetworks:"), value)
}


// A Boolean value that determines whether to require an administrator password to change networks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWMutableConfiguration/requireAdministratorForAssociation
func (c_ CWMutableConfiguration) RequireAdministratorForAssociation() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("requireAdministratorForAssociation"))
	return rv
}


// A Boolean value that determines whether to require an administrator password to change networks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWMutableConfiguration/requireAdministratorForAssociation
func (c_ CWMutableConfiguration) SetRequireAdministratorForAssociation(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRequireAdministratorForAssociation:"), value)
}


// A Boolean value that determines whether to require an administrator password to create a computer-to-computer network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWMutableConfiguration/requireAdministratorForIBSSMode
func (c_ CWMutableConfiguration) RequireAdministratorForIBSSMode() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("requireAdministratorForIBSSMode"))
	return rv
}


// A Boolean value that determines whether to require an administrator password to create a computer-to-computer network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWMutableConfiguration/requireAdministratorForIBSSMode
func (c_ CWMutableConfiguration) SetRequireAdministratorForIBSSMode(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRequireAdministratorForIBSSMode:"), value)
}


// A Boolean value that determines whether to require an administrator password to change the interface power state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWMutableConfiguration/requireAdministratorForPower
func (c_ CWMutableConfiguration) RequireAdministratorForPower() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("requireAdministratorForPower"))
	return rv
}


// A Boolean value that determines whether to require an administrator password to change the interface power state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWMutableConfiguration/requireAdministratorForPower
func (c_ CWMutableConfiguration) SetRequireAdministratorForPower(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRequireAdministratorForPower:"), value)
}




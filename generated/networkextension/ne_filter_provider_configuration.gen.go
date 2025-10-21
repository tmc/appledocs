// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEFilterProviderConfiguration] class.
var (
	NEFilterProviderConfigurationClass     _NEFilterProviderConfigurationClass
	NEFilterProviderConfigurationClassOnce sync.Once
)

func getNEFilterProviderConfigurationClass() _NEFilterProviderConfigurationClass {
	NEFilterProviderConfigurationClassOnce.Do(func() {
		NEFilterProviderConfigurationClass = _NEFilterProviderConfigurationClass{objc.GetClass("NEFilterProviderConfiguration")}
	})
	return NEFilterProviderConfigurationClass
}

type _NEFilterProviderConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [NEFilterProviderConfiguration] class.
type INEFilterProviderConfiguration interface {
	objectivec.IObject
}

// Configuration parameters for a content filter.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration
type NEFilterProviderConfiguration struct {
	objectivec.Object
}

// NEFilterProviderConfigurationFrom constructs a [NEFilterProviderConfiguration] from an unsafe.Pointer.
//
// Configuration parameters for a content filter.
func NEFilterProviderConfigurationFrom(ptr unsafe.Pointer) NEFilterProviderConfiguration {
	return NEFilterProviderConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEFilterProviderConfigurationClass) Alloc() NEFilterProviderConfiguration {
	rv := objc.Send[NEFilterProviderConfiguration](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEFilterProviderConfigurationClass) New() NEFilterProviderConfiguration {
	rv := objc.Send[NEFilterProviderConfiguration](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterProviderConfiguration) Init() NEFilterProviderConfiguration {
	rv := objc.Send[NEFilterProviderConfiguration](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterProviderConfiguration) Autorelease() NEFilterProviderConfiguration {
	rv := objc.Send[NEFilterProviderConfiguration](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterProviderConfiguration creates a new NEFilterProviderConfiguration instance.
func NewNEFilterProviderConfiguration() NEFilterProviderConfiguration {
	return getNEFilterProviderConfigurationClass().New()
}


// A Boolean value that indicates that the system applies the filter to flows of network data originated from WebKit browser objects.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/filterbrowsers
func (n_ NEFilterProviderConfiguration) FilterBrowsers() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("filterBrowsers"))
	return rv
}


// SetFilterBrowsers sets the value of the filterBrowsers property.
// A Boolean value that indicates that the system applies the filter to flows of network data originated from WebKit browser objects.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/filterbrowsers
func (n_ NEFilterProviderConfiguration) SetFilterBrowsers(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFilterBrowsers:"), value)
}

// The bundle identifier of the filter data provider system extension.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/filterdataproviderbundleidentifier
func (n_ NEFilterProviderConfiguration) FilterDataProviderBundleIdentifier() string {
	rv := objc.Send[string](n_.ID, objc.Sel("filterDataProviderBundleIdentifier"))
	return rv
}


// SetFilterDataProviderBundleIdentifier sets the value of the filterDataProviderBundleIdentifier property.
// The bundle identifier of the filter data provider system extension.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/filterdataproviderbundleidentifier
func (n_ NEFilterProviderConfiguration) SetFilterDataProviderBundleIdentifier(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFilterDataProviderBundleIdentifier:"), objc.String(value))
}

// The bundle identifier of the filter packet provider system extension.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/filterpacketproviderbundleidentifier
func (n_ NEFilterProviderConfiguration) FilterPacketProviderBundleIdentifier() string {
	rv := objc.Send[string](n_.ID, objc.Sel("filterPacketProviderBundleIdentifier"))
	return rv
}


// SetFilterPacketProviderBundleIdentifier sets the value of the filterPacketProviderBundleIdentifier property.
// The bundle identifier of the filter packet provider system extension.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/filterpacketproviderbundleidentifier
func (n_ NEFilterProviderConfiguration) SetFilterPacketProviderBundleIdentifier(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFilterPacketProviderBundleIdentifier:"), objc.String(value))
}

// A Boolean value that indicates that the system applies the filter to packets of network data.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/filterpackets
func (n_ NEFilterProviderConfiguration) FilterPackets() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("filterPackets"))
	return rv
}


// SetFilterPackets sets the value of the filterPackets property.
// A Boolean value that indicates that the system applies the filter to packets of network data.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/filterpackets
func (n_ NEFilterProviderConfiguration) SetFilterPackets(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFilterPackets:"), value)
}

// A Boolean value that indicates that the system applies the filter to flows of network data originated from sockets.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/filtersockets
func (n_ NEFilterProviderConfiguration) FilterSockets() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("filterSockets"))
	return rv
}


// SetFilterSockets sets the value of the filterSockets property.
// A Boolean value that indicates that the system applies the filter to flows of network data originated from sockets.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/filtersockets
func (n_ NEFilterProviderConfiguration) SetFilterSockets(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFilterSockets:"), value)
}

// A persistent reference to a keychain item containing a certificate and private key associated with the filter.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/identityreference
func (n_ NEFilterProviderConfiguration) IdentityReference() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("identityReference"))
	return rv
}


// SetIdentityReference sets the value of the identityReference property.
// A persistent reference to a keychain item containing a certificate and private key associated with the filter.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/identityreference
func (n_ NEFilterProviderConfiguration) SetIdentityReference(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIdentityReference:"), value)
}

// A string that identifies the organization that administers the filter.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/organization
func (n_ NEFilterProviderConfiguration) Organization() string {
	rv := objc.Send[string](n_.ID, objc.Sel("organization"))
	return rv
}


// SetOrganization sets the value of the organization property.
// A string that identifies the organization that administers the filter.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/organization
func (n_ NEFilterProviderConfiguration) SetOrganization(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setOrganization:"), objc.String(value))
}

// A persistent reference to a keychain item containing a password associated with the filter.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/passwordreference
func (n_ NEFilterProviderConfiguration) PasswordReference() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("passwordReference"))
	return rv
}


// SetPasswordReference sets the value of the passwordReference property.
// A persistent reference to a keychain item containing a password associated with the filter.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/passwordreference
func (n_ NEFilterProviderConfiguration) SetPasswordReference(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPasswordReference:"), value)
}

// The address of a server that the Filter Control Provider may contact for rules and other configuration information.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/serveraddress
func (n_ NEFilterProviderConfiguration) ServerAddress() string {
	rv := objc.Send[string](n_.ID, objc.Sel("serverAddress"))
	return rv
}


// SetServerAddress sets the value of the serverAddress property.
// The address of a server that the Filter Control Provider may contact for rules and other configuration information.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/serveraddress
func (n_ NEFilterProviderConfiguration) SetServerAddress(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setServerAddress:"), objc.String(value))
}

// A string that identifies the user.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/username
func (n_ NEFilterProviderConfiguration) Username() string {
	rv := objc.Send[string](n_.ID, objc.Sel("username"))
	return rv
}


// SetUsername sets the value of the username property.
// A string that identifies the user.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/username
func (n_ NEFilterProviderConfiguration) SetUsername(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUsername:"), objc.String(value))
}

// A dictionary of provider-specific configuration settings.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/vendorconfiguration
func (n_ NEFilterProviderConfiguration) VendorConfiguration() string {
	rv := objc.Send[string](n_.ID, objc.Sel("vendorConfiguration"))
	return rv
}


// SetVendorConfiguration sets the value of the vendorConfiguration property.
// A dictionary of provider-specific configuration settings.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/vendorconfiguration
func (n_ NEFilterProviderConfiguration) SetVendorConfiguration(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setVendorConfiguration:"), objc.String(value))
}




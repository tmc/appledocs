// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	FilterBrowsers() bool
	SetFilterBrowsers(value bool)
	FilterDataProviderBundleIdentifier() objc.IObject /* cross-framework: NSString */
	SetFilterDataProviderBundleIdentifier(value objc.IObject /* cross-framework: NSString */)
	FilterPacketProviderBundleIdentifier() objc.IObject /* cross-framework: NSString */
	SetFilterPacketProviderBundleIdentifier(value objc.IObject /* cross-framework: NSString */)
	FilterPackets() bool
	SetFilterPackets(value bool)
	FilterSockets() bool
	SetFilterSockets(value bool)
	IdentityReference() objc.IObject /* cross-framework: Data */
	SetIdentityReference(value objc.IObject /* cross-framework: Data */)
	Organization() objc.IObject /* cross-framework: NSString */
	SetOrganization(value objc.IObject /* cross-framework: NSString */)
	PasswordReference() objc.IObject /* cross-framework: Data */
	SetPasswordReference(value objc.IObject /* cross-framework: Data */)
	ServerAddress() objc.IObject /* cross-framework: NSString */
	SetServerAddress(value objc.IObject /* cross-framework: NSString */)
	Username() objc.IObject /* cross-framework: NSString */
	SetUsername(value objc.IObject /* cross-framework: NSString */)
	VendorConfiguration() objc.IObject /* cross-framework: NSString */
	SetVendorConfiguration(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// Configuration parameters for a content filter.


// Configuration parameters for a content filter.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/filterbrowsers
func (n_ NEFilterProviderConfiguration) FilterBrowsers() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("filterBrowsers"))
	return rv
}


// A Boolean value that indicates that the system applies the filter to flows of network data originated from WebKit browser objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/filterbrowsers
func (n_ NEFilterProviderConfiguration) SetFilterBrowsers(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFilterBrowsers:"), value)
}


// The bundle identifier of the filter data provider system extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/filterdataproviderbundleidentifier
func (n_ NEFilterProviderConfiguration) FilterDataProviderBundleIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("filterDataProviderBundleIdentifier"))
	return rv
}


// The bundle identifier of the filter data provider system extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/filterdataproviderbundleidentifier
func (n_ NEFilterProviderConfiguration) SetFilterDataProviderBundleIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFilterDataProviderBundleIdentifier:"), value)
}


// The bundle identifier of the filter packet provider system extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/filterpacketproviderbundleidentifier
func (n_ NEFilterProviderConfiguration) FilterPacketProviderBundleIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("filterPacketProviderBundleIdentifier"))
	return rv
}


// The bundle identifier of the filter packet provider system extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/filterpacketproviderbundleidentifier
func (n_ NEFilterProviderConfiguration) SetFilterPacketProviderBundleIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFilterPacketProviderBundleIdentifier:"), value)
}


// A Boolean value that indicates that the system applies the filter to packets of network data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/filterpackets
func (n_ NEFilterProviderConfiguration) FilterPackets() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("filterPackets"))
	return rv
}


// A Boolean value that indicates that the system applies the filter to packets of network data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/filterpackets
func (n_ NEFilterProviderConfiguration) SetFilterPackets(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFilterPackets:"), value)
}


// A Boolean value that indicates that the system applies the filter to flows of network data originated from sockets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/filtersockets
func (n_ NEFilterProviderConfiguration) FilterSockets() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("filterSockets"))
	return rv
}


// A Boolean value that indicates that the system applies the filter to flows of network data originated from sockets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/filtersockets
func (n_ NEFilterProviderConfiguration) SetFilterSockets(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFilterSockets:"), value)
}


// A persistent reference to a keychain item containing a certificate and private key associated with the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/identityreference
func (n_ NEFilterProviderConfiguration) IdentityReference() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](n_.ID, objc.Sel("identityReference"))
	return rv
}


// A persistent reference to a keychain item containing a certificate and private key associated with the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/identityreference
func (n_ NEFilterProviderConfiguration) SetIdentityReference(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIdentityReference:"), value)
}


// A string that identifies the organization that administers the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/organization
func (n_ NEFilterProviderConfiguration) Organization() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("organization"))
	return rv
}


// A string that identifies the organization that administers the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/organization
func (n_ NEFilterProviderConfiguration) SetOrganization(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setOrganization:"), value)
}


// A persistent reference to a keychain item containing a password associated with the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/passwordreference
func (n_ NEFilterProviderConfiguration) PasswordReference() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](n_.ID, objc.Sel("passwordReference"))
	return rv
}


// A persistent reference to a keychain item containing a password associated with the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/passwordreference
func (n_ NEFilterProviderConfiguration) SetPasswordReference(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPasswordReference:"), value)
}


// The address of a server that the Filter Control Provider may contact for rules and other configuration information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/serveraddress
func (n_ NEFilterProviderConfiguration) ServerAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("serverAddress"))
	return rv
}


// The address of a server that the Filter Control Provider may contact for rules and other configuration information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/serveraddress
func (n_ NEFilterProviderConfiguration) SetServerAddress(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setServerAddress:"), value)
}


// A string that identifies the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/username
func (n_ NEFilterProviderConfiguration) Username() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("username"))
	return rv
}


// A string that identifies the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/username
func (n_ NEFilterProviderConfiguration) SetUsername(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUsername:"), value)
}


// A dictionary of provider-specific configuration settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/vendorconfiguration
func (n_ NEFilterProviderConfiguration) VendorConfiguration() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("vendorConfiguration"))
	return rv
}


// A dictionary of provider-specific configuration settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterproviderconfiguration/vendorconfiguration
func (n_ NEFilterProviderConfiguration) SetVendorConfiguration(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setVendorConfiguration:"), value)
}




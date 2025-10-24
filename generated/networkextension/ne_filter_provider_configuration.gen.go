// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEFilterProviderConfiguration */


/* debug [class_header]: Header for NEFilterProviderConfiguration */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEFilterProviderConfiguration */
// An interface definition for the [NEFilterProviderConfiguration] class.
type INEFilterProviderConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEFilterProviderConfiguration */
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
	IdentityReference() objc.IObject /* cross-framework: NSData */
	SetIdentityReference(value objc.IObject /* cross-framework: NSData */)
	Organization() objc.IObject /* cross-framework: NSString */
	SetOrganization(value objc.IObject /* cross-framework: NSString */)
	PasswordReference() objc.IObject /* cross-framework: NSData */
	SetPasswordReference(value objc.IObject /* cross-framework: NSData */)
	ServerAddress() objc.IObject /* cross-framework: NSString */
	SetServerAddress(value objc.IObject /* cross-framework: NSString */)
	Username() objc.IObject /* cross-framework: NSString */
	SetUsername(value objc.IObject /* cross-framework: NSString */)
	VendorConfiguration() foundation.IDictionary
	SetVendorConfiguration(value foundation.IDictionary)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEFilterProviderConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEFilterProviderConfiguration */
// Alloc allocates a new instance without initialization.
func (nc _NEFilterProviderConfigurationClass) Alloc() NEFilterProviderConfiguration {
	rv := objc.Send[NEFilterProviderConfiguration](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEFilterProviderConfiguration */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEFilterProviderConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEFilterProviderConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEFilterProviderConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEFilterProviderConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEFilterProviderConfiguration */

// A Boolean value that indicates that the system applies the filter to flows of network data originated from WebKit browser objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/filterBrowsers
func (n_ NEFilterProviderConfiguration) FilterBrowsers() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("filterBrowsers"))
	return rv
}/* debug [instance_properties/getter]: filterBrowsers */


// A Boolean value that indicates that the system applies the filter to flows of network data originated from WebKit browser objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/filterBrowsers
func (n_ NEFilterProviderConfiguration) SetFilterBrowsers(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFilterBrowsers:"), value)
}/* debug [instance_properties/setter]: filterBrowsers */


// The bundle identifier of the filter data provider system extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/filterDataProviderBundleIdentifier
func (n_ NEFilterProviderConfiguration) FilterDataProviderBundleIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("filterDataProviderBundleIdentifier"))
	return rv
}/* debug [instance_properties/getter]: filterDataProviderBundleIdentifier */


// The bundle identifier of the filter data provider system extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/filterDataProviderBundleIdentifier
func (n_ NEFilterProviderConfiguration) SetFilterDataProviderBundleIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFilterDataProviderBundleIdentifier:"), value)
}/* debug [instance_properties/setter]: filterDataProviderBundleIdentifier */


// The bundle identifier of the filter packet provider system extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/filterPacketProviderBundleIdentifier
func (n_ NEFilterProviderConfiguration) FilterPacketProviderBundleIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("filterPacketProviderBundleIdentifier"))
	return rv
}/* debug [instance_properties/getter]: filterPacketProviderBundleIdentifier */


// The bundle identifier of the filter packet provider system extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/filterPacketProviderBundleIdentifier
func (n_ NEFilterProviderConfiguration) SetFilterPacketProviderBundleIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFilterPacketProviderBundleIdentifier:"), value)
}/* debug [instance_properties/setter]: filterPacketProviderBundleIdentifier */


// A Boolean value that indicates that the system applies the filter to packets of network data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/filterPackets
func (n_ NEFilterProviderConfiguration) FilterPackets() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("filterPackets"))
	return rv
}/* debug [instance_properties/getter]: filterPackets */


// A Boolean value that indicates that the system applies the filter to packets of network data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/filterPackets
func (n_ NEFilterProviderConfiguration) SetFilterPackets(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFilterPackets:"), value)
}/* debug [instance_properties/setter]: filterPackets */


// A Boolean value that indicates that the system applies the filter to flows of network data originated from sockets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/filterSockets
func (n_ NEFilterProviderConfiguration) FilterSockets() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("filterSockets"))
	return rv
}/* debug [instance_properties/getter]: filterSockets */


// A Boolean value that indicates that the system applies the filter to flows of network data originated from sockets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/filterSockets
func (n_ NEFilterProviderConfiguration) SetFilterSockets(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFilterSockets:"), value)
}/* debug [instance_properties/setter]: filterSockets */


// A persistent reference to a keychain item containing a certificate and private key associated with the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/identityReference
func (n_ NEFilterProviderConfiguration) IdentityReference() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("identityReference"))
	return rv
}/* debug [instance_properties/getter]: identityReference */


// A persistent reference to a keychain item containing a certificate and private key associated with the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/identityReference
func (n_ NEFilterProviderConfiguration) SetIdentityReference(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIdentityReference:"), value)
}/* debug [instance_properties/setter]: identityReference */


// A string that identifies the organization that administers the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/organization
func (n_ NEFilterProviderConfiguration) Organization() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("organization"))
	return rv
}/* debug [instance_properties/getter]: organization */


// A string that identifies the organization that administers the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/organization
func (n_ NEFilterProviderConfiguration) SetOrganization(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setOrganization:"), value)
}/* debug [instance_properties/setter]: organization */


// A persistent reference to a keychain item containing a password associated with the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/passwordReference
func (n_ NEFilterProviderConfiguration) PasswordReference() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("passwordReference"))
	return rv
}/* debug [instance_properties/getter]: passwordReference */


// A persistent reference to a keychain item containing a password associated with the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/passwordReference
func (n_ NEFilterProviderConfiguration) SetPasswordReference(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPasswordReference:"), value)
}/* debug [instance_properties/setter]: passwordReference */


// The address of a server that the Filter Control Provider may contact for rules and other configuration information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/serverAddress
func (n_ NEFilterProviderConfiguration) ServerAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("serverAddress"))
	return rv
}/* debug [instance_properties/getter]: serverAddress */


// The address of a server that the Filter Control Provider may contact for rules and other configuration information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/serverAddress
func (n_ NEFilterProviderConfiguration) SetServerAddress(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setServerAddress:"), value)
}/* debug [instance_properties/setter]: serverAddress */


// A string that identifies the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/username
func (n_ NEFilterProviderConfiguration) Username() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("username"))
	return rv
}/* debug [instance_properties/getter]: username */


// A string that identifies the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/username
func (n_ NEFilterProviderConfiguration) SetUsername(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUsername:"), value)
}/* debug [instance_properties/setter]: username */


// A dictionary of provider-specific configuration settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/vendorConfiguration
func (n_ NEFilterProviderConfiguration) VendorConfiguration() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](n_.ID, objc.Sel("vendorConfiguration"))
	return rv
}/* debug [instance_properties/getter]: vendorConfiguration */


// A dictionary of provider-specific configuration settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterProviderConfiguration/vendorConfiguration
func (n_ NEFilterProviderConfiguration) SetVendorConfiguration(value foundation.IDictionary) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setVendorConfiguration:"), value)
}/* debug [instance_properties/setter]: vendorConfiguration */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEFilterProviderConfiguration */




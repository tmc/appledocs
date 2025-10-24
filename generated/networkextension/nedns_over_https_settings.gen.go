// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class NEDNSOverHTTPSSettings */


/* debug [class_header]: Header for NEDNSOverHTTPSSettings */
// The class instance for the [NEDNSOverHTTPSSettings] class.
var (
	NEDNSOverHTTPSSettingsClass     _NEDNSOverHTTPSSettingsClass
	NEDNSOverHTTPSSettingsClassOnce sync.Once
)

func getNEDNSOverHTTPSSettingsClass() _NEDNSOverHTTPSSettingsClass {
	NEDNSOverHTTPSSettingsClassOnce.Do(func() {
		NEDNSOverHTTPSSettingsClass = _NEDNSOverHTTPSSettingsClass{objc.GetClass("NEDNSOverHTTPSSettings")}
	})
	return NEDNSOverHTTPSSettingsClass
}

type _NEDNSOverHTTPSSettingsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEDNSOverHTTPSSettings */
// An interface definition for the [NEDNSOverHTTPSSettings] class.
type INEDNSOverHTTPSSettings interface {
	INEDNSSettings
	
/* debug [class_interface_properties]: Properties for NEDNSOverHTTPSSettings */
	// properties:
	IdentityReference() objc.IObject /* cross-framework: NSData */
	SetIdentityReference(value objc.IObject /* cross-framework: NSData */)
	ServerURL() objc.IObject /* cross-framework: NSURL */
	SetServerURL(value objc.IObject /* cross-framework: NSURL */)
	MatchDomains() objc.IObject /* cross-framework: NSString */
	SetMatchDomains(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEDNSOverHTTPSSettings */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEDNSOverHTTPSSettings */
// Alloc allocates a new instance without initialization.
func (nc _NEDNSOverHTTPSSettingsClass) Alloc() NEDNSOverHTTPSSettings {
	rv := objc.Send[NEDNSOverHTTPSSettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEDNSOverHTTPSSettingsClass) New() NEDNSOverHTTPSSettings {
	rv := objc.Send[NEDNSOverHTTPSSettings](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEDNSOverHTTPSSettings) Init() NEDNSOverHTTPSSettings {
	rv := objc.Send[NEDNSOverHTTPSSettings](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEDNSOverHTTPSSettings) Autorelease() NEDNSOverHTTPSSettings {
	rv := objc.Send[NEDNSOverHTTPSSettings](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEDNSOverHTTPSSettings creates a new NEDNSOverHTTPSSettings instance.
func NewNEDNSOverHTTPSSettings() NEDNSOverHTTPSSettings {
	return getNEDNSOverHTTPSSettingsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEDNSOverHTTPSSettings */
// The DNS resolver settings for a DNS-over-HTTPS server.


// The DNS resolver settings for a DNS-over-HTTPS server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSOverHTTPSSettings
type NEDNSOverHTTPSSettings struct {
	NEDNSSettings
}

// NEDNSOverHTTPSSettingsFrom constructs a [NEDNSOverHTTPSSettings] from an unsafe.Pointer.
//
// The DNS resolver settings for a DNS-over-HTTPS server.
func NEDNSOverHTTPSSettingsFrom(ptr unsafe.Pointer) NEDNSOverHTTPSSettings {
	return NEDNSOverHTTPSSettings{
		NEDNSSettings: NEDNSSettingsFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEDNSOverHTTPSSettings *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEDNSOverHTTPSSettings */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEDNSOverHTTPSSettings */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEDNSOverHTTPSSettings */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEDNSOverHTTPSSettings */

// A persistent keychain reference to a keychain item containing the certificate and private key components of the DNS client credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSOverHTTPSSettings/identityReference
func (n_ NEDNSOverHTTPSSettings) IdentityReference() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("identityReference"))
	return rv
}/* debug [instance_properties/getter]: identityReference */


// A persistent keychain reference to a keychain item containing the certificate and private key components of the DNS client credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSOverHTTPSSettings/identityReference
func (n_ NEDNSOverHTTPSSettings) SetIdentityReference(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIdentityReference:"), value)
}/* debug [instance_properties/setter]: identityReference */


// The URL of a DNS-over-HTTPS server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSOverHTTPSSettings/serverURL
func (n_ NEDNSOverHTTPSSettings) ServerURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](n_.ID, objc.Sel("serverURL"))
	return rv
}/* debug [instance_properties/getter]: serverURL */


// The URL of a DNS-over-HTTPS server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSOverHTTPSSettings/serverURL
func (n_ NEDNSOverHTTPSSettings) SetServerURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setServerURL:"), value)
}/* debug [instance_properties/setter]: serverURL */


// A list of domain strings used to determine which DNS queries will use the DNS resolver settings contained in this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/matchdomains
func (n_ NEDNSOverHTTPSSettings) MatchDomains() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("matchDomains"))
	return rv
}/* debug [instance_properties/getter]: matchDomains */


// A list of domain strings used to determine which DNS queries will use the DNS resolver settings contained in this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/matchdomains
func (n_ NEDNSOverHTTPSSettings) SetMatchDomains(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchDomains:"), value)
}/* debug [instance_properties/setter]: matchDomains */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEDNSOverHTTPSSettings */




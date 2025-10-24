// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class NEDNSOverTLSSettings */


/* debug [class_header]: Header for NEDNSOverTLSSettings */
// The class instance for the [NEDNSOverTLSSettings] class.
var (
	NEDNSOverTLSSettingsClass     _NEDNSOverTLSSettingsClass
	NEDNSOverTLSSettingsClassOnce sync.Once
)

func getNEDNSOverTLSSettingsClass() _NEDNSOverTLSSettingsClass {
	NEDNSOverTLSSettingsClassOnce.Do(func() {
		NEDNSOverTLSSettingsClass = _NEDNSOverTLSSettingsClass{objc.GetClass("NEDNSOverTLSSettings")}
	})
	return NEDNSOverTLSSettingsClass
}

type _NEDNSOverTLSSettingsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEDNSOverTLSSettings */
// An interface definition for the [NEDNSOverTLSSettings] class.
type INEDNSOverTLSSettings interface {
	INEDNSSettings
	
/* debug [class_interface_properties]: Properties for NEDNSOverTLSSettings */
	// properties:
	IdentityReference() objc.IObject /* cross-framework: NSData */
	SetIdentityReference(value objc.IObject /* cross-framework: NSData */)
	ServerName() objc.IObject /* cross-framework: NSString */
	SetServerName(value objc.IObject /* cross-framework: NSString */)
	MatchDomains() objc.IObject /* cross-framework: NSString */
	SetMatchDomains(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEDNSOverTLSSettings */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEDNSOverTLSSettings */
// Alloc allocates a new instance without initialization.
func (nc _NEDNSOverTLSSettingsClass) Alloc() NEDNSOverTLSSettings {
	rv := objc.Send[NEDNSOverTLSSettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEDNSOverTLSSettingsClass) New() NEDNSOverTLSSettings {
	rv := objc.Send[NEDNSOverTLSSettings](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEDNSOverTLSSettings) Init() NEDNSOverTLSSettings {
	rv := objc.Send[NEDNSOverTLSSettings](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEDNSOverTLSSettings) Autorelease() NEDNSOverTLSSettings {
	rv := objc.Send[NEDNSOverTLSSettings](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEDNSOverTLSSettings creates a new NEDNSOverTLSSettings instance.
func NewNEDNSOverTLSSettings() NEDNSOverTLSSettings {
	return getNEDNSOverTLSSettingsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEDNSOverTLSSettings */
// The DNS resolver settings for a DNS-over-TLS server.


// The DNS resolver settings for a DNS-over-TLS server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSOverTLSSettings
type NEDNSOverTLSSettings struct {
	NEDNSSettings
}

// NEDNSOverTLSSettingsFrom constructs a [NEDNSOverTLSSettings] from an unsafe.Pointer.
//
// The DNS resolver settings for a DNS-over-TLS server.
func NEDNSOverTLSSettingsFrom(ptr unsafe.Pointer) NEDNSOverTLSSettings {
	return NEDNSOverTLSSettings{
		NEDNSSettings: NEDNSSettingsFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEDNSOverTLSSettings *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEDNSOverTLSSettings */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEDNSOverTLSSettings */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEDNSOverTLSSettings */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEDNSOverTLSSettings */

// A persistent keychain reference to a keychain item containing the certificate and private key components of the DNS client credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSOverTLSSettings/identityReference
func (n_ NEDNSOverTLSSettings) IdentityReference() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("identityReference"))
	return rv
}/* debug [instance_properties/getter]: identityReference */


// A persistent keychain reference to a keychain item containing the certificate and private key components of the DNS client credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSOverTLSSettings/identityReference
func (n_ NEDNSOverTLSSettings) SetIdentityReference(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIdentityReference:"), value)
}/* debug [instance_properties/setter]: identityReference */


// The TLS name of a DNS-over-TLS server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSOverTLSSettings/serverName
func (n_ NEDNSOverTLSSettings) ServerName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("serverName"))
	return rv
}/* debug [instance_properties/getter]: serverName */


// The TLS name of a DNS-over-TLS server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEDNSOverTLSSettings/serverName
func (n_ NEDNSOverTLSSettings) SetServerName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setServerName:"), value)
}/* debug [instance_properties/setter]: serverName */


// A list of domain strings used to determine which DNS queries will use the DNS resolver settings contained in this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/matchdomains
func (n_ NEDNSOverTLSSettings) MatchDomains() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("matchDomains"))
	return rv
}/* debug [instance_properties/getter]: matchDomains */


// A list of domain strings used to determine which DNS queries will use the DNS resolver settings contained in this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nednssettings/matchdomains
func (n_ NEDNSOverTLSSettings) SetMatchDomains(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchDomains:"), value)
}/* debug [instance_properties/setter]: matchDomains */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEDNSOverTLSSettings */




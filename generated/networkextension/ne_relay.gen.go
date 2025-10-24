// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NERelay */


/* debug [class_header]: Header for NERelay */
// The class instance for the [NERelay] class.
var (
	NERelayClass     _NERelayClass
	NERelayClassOnce sync.Once
)

func getNERelayClass() _NERelayClass {
	NERelayClassOnce.Do(func() {
		NERelayClass = _NERelayClass{objc.GetClass("NERelay")}
	})
	return NERelayClass
}

type _NERelayClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NERelay */
// An interface definition for the [NERelay] class.
type INERelay interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NERelay */
	// properties:
	AdditionalHTTPHeaderFields() foundation.IDictionary
	SetAdditionalHTTPHeaderFields(value foundation.IDictionary)
	DnsOverHTTPSURL() objc.IObject /* cross-framework: NSURL */
	SetDnsOverHTTPSURL(value objc.IObject /* cross-framework: NSURL */)
	HTTP2RelayURL() objc.IObject /* cross-framework: NSURL */
	SetHTTP2RelayURL(value objc.IObject /* cross-framework: NSURL */)
	HTTP3RelayURL() objc.IObject /* cross-framework: NSURL */
	SetHTTP3RelayURL(value objc.IObject /* cross-framework: NSURL */)
	IdentityData() objc.IObject /* cross-framework: NSData */
	SetIdentityData(value objc.IObject /* cross-framework: NSData */)
	IdentityDataPassword() objc.IObject /* cross-framework: NSString */
	SetIdentityDataPassword(value objc.IObject /* cross-framework: NSString */)
	RawPublicKeys() []foundation.Data
	SetRawPublicKeys(value []foundation.Data)
	SyntheticDNSAnswerIPv4Prefix() objc.IObject /* cross-framework: NSString */
	SetSyntheticDNSAnswerIPv4Prefix(value objc.IObject /* cross-framework: NSString */)
	SyntheticDNSAnswerIPv6Prefix() objc.IObject /* cross-framework: NSString */
	SetSyntheticDNSAnswerIPv6Prefix(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NERelay */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NERelay */
// Alloc allocates a new instance without initialization.
func (nc _NERelayClass) Alloc() NERelay {
	rv := objc.Send[NERelay](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NERelayClass) New() NERelay {
	rv := objc.Send[NERelay](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NERelay) Init() NERelay {
	rv := objc.Send[NERelay](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NERelay) Autorelease() NERelay {
	rv := objc.Send[NERelay](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNERelay creates a new NERelay instance.
func NewNERelay() NERelay {
	return getNERelayClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NERelay */
// A single relay server configuration that you can chain together with other relays.
//
// Relay servers are secure HTTP proxies that allow proxying TCP traffic using the method and UDP traffic using the protocol defined in .


// A single relay server configuration that you can chain together with other relays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay
type NERelay struct {
	objectivec.Object
}

// NERelayFrom constructs a [NERelay] from an unsafe.Pointer.
//
// A single relay server configuration that you can chain together with other relays.
func NERelayFrom(ptr unsafe.Pointer) NERelay {
	return NERelay{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NERelay *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NERelay */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NERelay */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NERelay */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NERelay */

// A dictionary of additional HTTP headers to send as part of requests to the relay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/additionalHTTPHeaderFields
func (n_ NERelay) AdditionalHTTPHeaderFields() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](n_.ID, objc.Sel("additionalHTTPHeaderFields"))
	return rv
}/* debug [instance_properties/getter]: additionalHTTPHeaderFields */


// A dictionary of additional HTTP headers to send as part of requests to the relay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/additionalHTTPHeaderFields
func (n_ NERelay) SetAdditionalHTTPHeaderFields(value foundation.IDictionary) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAdditionalHTTPHeaderFields:"), value)
}/* debug [instance_properties/setter]: additionalHTTPHeaderFields */


// The URL of a DNS-over-HTTPS (DoH) resolver accessible from the relay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/dnsOverHTTPSURL
func (n_ NERelay) DnsOverHTTPSURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](n_.ID, objc.Sel("dnsOverHTTPSURL"))
	return rv
}/* debug [instance_properties/getter]: dnsOverHTTPSURL */


// The URL of a DNS-over-HTTPS (DoH) resolver accessible from the relay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/dnsOverHTTPSURL
func (n_ NERelay) SetDnsOverHTTPSURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDnsOverHTTPSURL:"), value)
}/* debug [instance_properties/setter]: dnsOverHTTPSURL */


// A URL identifying the relay server accessible using HTTP/2.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/http2RelayURL
func (n_ NERelay) HTTP2RelayURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](n_.ID, objc.Sel("HTTP2RelayURL"))
	return rv
}/* debug [instance_properties/getter]: HTTP2RelayURL */


// A URL identifying the relay server accessible using HTTP/2.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/http2RelayURL
func (n_ NERelay) SetHTTP2RelayURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHTTP2RelayURL:"), value)
}/* debug [instance_properties/setter]: HTTP2RelayURL */


// A URL identifying the relay server accessible using HTTP/3.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/http3RelayURL
func (n_ NERelay) HTTP3RelayURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](n_.ID, objc.Sel("HTTP3RelayURL"))
	return rv
}/* debug [instance_properties/getter]: HTTP3RelayURL */


// A URL identifying the relay server accessible using HTTP/3.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/http3RelayURL
func (n_ NERelay) SetHTTP3RelayURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHTTP3RelayURL:"), value)
}/* debug [instance_properties/setter]: HTTP3RelayURL */


// The PKCS12 data for the relay client authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/identityData
func (n_ NERelay) IdentityData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("identityData"))
	return rv
}/* debug [instance_properties/getter]: identityData */


// The PKCS12 data for the relay client authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/identityData
func (n_ NERelay) SetIdentityData(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIdentityData:"), value)
}/* debug [instance_properties/setter]: identityData */


// The password the relay uses to decrypt the PKCS12 identity data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/identityDataPassword
func (n_ NERelay) IdentityDataPassword() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("identityDataPassword"))
	return rv
}/* debug [instance_properties/getter]: identityDataPassword */


// The password the relay uses to decrypt the PKCS12 identity data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/identityDataPassword
func (n_ NERelay) SetIdentityDataPassword(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIdentityDataPassword:"), value)
}/* debug [instance_properties/setter]: identityDataPassword */


// An array of TLS raw public keys that the relay server can present during the TLS handshake.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/rawPublicKeys
func (n_ NERelay) RawPublicKeys() []foundation.Data {
	rv := objc.Send[[]foundation.Data](n_.ID, objc.Sel("rawPublicKeys"))
	return rv
}/* debug [instance_properties/getter]: rawPublicKeys */


// An array of TLS raw public keys that the relay server can present during the TLS handshake.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/rawPublicKeys
func (n_ NERelay) SetRawPublicKeys(value []foundation.Data) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](n_.ID, objc.Sel("setRawPublicKeys:"), nsArray)
}/* debug [instance_properties/setter]: rawPublicKeys */


// An IPv4 address prefix the relay uses to handle address info requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/syntheticDNSAnswerIPv4Prefix
func (n_ NERelay) SyntheticDNSAnswerIPv4Prefix() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("syntheticDNSAnswerIPv4Prefix"))
	return rv
}/* debug [instance_properties/getter]: syntheticDNSAnswerIPv4Prefix */


// An IPv4 address prefix the relay uses to handle address info requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/syntheticDNSAnswerIPv4Prefix
func (n_ NERelay) SetSyntheticDNSAnswerIPv4Prefix(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSyntheticDNSAnswerIPv4Prefix:"), value)
}/* debug [instance_properties/setter]: syntheticDNSAnswerIPv4Prefix */


// An IPv6 address prefix the relay uses to handle address info requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/syntheticDNSAnswerIPv6Prefix
func (n_ NERelay) SyntheticDNSAnswerIPv6Prefix() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("syntheticDNSAnswerIPv6Prefix"))
	return rv
}/* debug [instance_properties/getter]: syntheticDNSAnswerIPv6Prefix */


// An IPv6 address prefix the relay uses to handle address info requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/syntheticDNSAnswerIPv6Prefix
func (n_ NERelay) SetSyntheticDNSAnswerIPv6Prefix(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSyntheticDNSAnswerIPv6Prefix:"), value)
}/* debug [instance_properties/setter]: syntheticDNSAnswerIPv6Prefix */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NERelay */




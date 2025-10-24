// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [NERelay] class.
type INERelay interface {
	objectivec.IObject
	// properties:
	HTTP3RelayURL() objc.IObject /* cross-framework: NSURL */
	SetHTTP3RelayURL(value objc.IObject /* cross-framework: NSURL */)
	AdditionalHTTPHeaderFields() objc.IObject /* cross-framework: NSString */
	SetAdditionalHTTPHeaderFields(value objc.IObject /* cross-framework: NSString */)
	DnsOverHTTPSURL() objc.IObject /* cross-framework: URL */
	SetDnsOverHTTPSURL(value objc.IObject /* cross-framework: URL */)
	Http2RelayURL() objc.IObject /* cross-framework: URL */
	SetHttp2RelayURL(value objc.IObject /* cross-framework: URL */)
	IdentityData() objc.IObject /* cross-framework: Data */
	SetIdentityData(value objc.IObject /* cross-framework: Data */)
	IdentityDataPassword() objc.IObject /* cross-framework: NSString */
	SetIdentityDataPassword(value objc.IObject /* cross-framework: NSString */)
	RawPublicKeys() objc.IObject /* cross-framework: Data */
	SetRawPublicKeys(value objc.IObject /* cross-framework: Data */)
	SyntheticDNSAnswerIPv4Prefix() objc.IObject /* cross-framework: NSString */
	SetSyntheticDNSAnswerIPv4Prefix(value objc.IObject /* cross-framework: NSString */)
	SyntheticDNSAnswerIPv6Prefix() objc.IObject /* cross-framework: NSString */
	SetSyntheticDNSAnswerIPv6Prefix(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (nc _NERelayClass) Alloc() NERelay {
	rv := objc.Send[NERelay](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A URL identifying the relay server accessible using HTTP/3.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/http3RelayURL
func (n_ NERelay) HTTP3RelayURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](n_.ID, objc.Sel("HTTP3RelayURL"))
	return rv
}


// A URL identifying the relay server accessible using HTTP/3.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/http3RelayURL
func (n_ NERelay) SetHTTP3RelayURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHTTP3RelayURL:"), value)
}


// A dictionary of additional HTTP headers to send as part of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelay/additionalhttpheaderfields
func (n_ NERelay) AdditionalHTTPHeaderFields() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("additionalHTTPHeaderFields"))
	return rv
}


// A dictionary of additional HTTP headers to send as part of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelay/additionalhttpheaderfields
func (n_ NERelay) SetAdditionalHTTPHeaderFields(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAdditionalHTTPHeaderFields:"), value)
}


// The URL of a DNS-over-HTTPS (DoH) resolver accessible from the relay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelay/dnsoverhttpsurl
func (n_ NERelay) DnsOverHTTPSURL() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](n_.ID, objc.Sel("dnsOverHTTPSURL"))
	return rv
}


// The URL of a DNS-over-HTTPS (DoH) resolver accessible from the relay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelay/dnsoverhttpsurl
func (n_ NERelay) SetDnsOverHTTPSURL(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDnsOverHTTPSURL:"), value)
}


// A URL identifying the relay server accessible using HTTP/2.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelay/http2relayurl
func (n_ NERelay) Http2RelayURL() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](n_.ID, objc.Sel("http2RelayURL"))
	return rv
}


// A URL identifying the relay server accessible using HTTP/2.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelay/http2relayurl
func (n_ NERelay) SetHttp2RelayURL(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHttp2RelayURL:"), value)
}


// The PKCS12 data for the relay client authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelay/identitydata
func (n_ NERelay) IdentityData() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](n_.ID, objc.Sel("identityData"))
	return rv
}


// The PKCS12 data for the relay client authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelay/identitydata
func (n_ NERelay) SetIdentityData(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIdentityData:"), value)
}


// The password the relay uses to decrypt the PKCS12 identity data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelay/identitydatapassword
func (n_ NERelay) IdentityDataPassword() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("identityDataPassword"))
	return rv
}


// The password the relay uses to decrypt the PKCS12 identity data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelay/identitydatapassword
func (n_ NERelay) SetIdentityDataPassword(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIdentityDataPassword:"), value)
}


// An array of TLS raw public keys that the relay server can present during the TLS handshake.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelay/rawpublickeys
func (n_ NERelay) RawPublicKeys() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](n_.ID, objc.Sel("rawPublicKeys"))
	return rv
}


// An array of TLS raw public keys that the relay server can present during the TLS handshake.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelay/rawpublickeys
func (n_ NERelay) SetRawPublicKeys(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setRawPublicKeys:"), value)
}


// An IPv4 address prefix the relay uses to handle address info requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelay/syntheticdnsansweripv4prefix
func (n_ NERelay) SyntheticDNSAnswerIPv4Prefix() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("syntheticDNSAnswerIPv4Prefix"))
	return rv
}


// An IPv4 address prefix the relay uses to handle address info requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelay/syntheticdnsansweripv4prefix
func (n_ NERelay) SetSyntheticDNSAnswerIPv4Prefix(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSyntheticDNSAnswerIPv4Prefix:"), value)
}


// An IPv6 address prefix the relay uses to handle address info requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelay/syntheticdnsansweripv6prefix
func (n_ NERelay) SyntheticDNSAnswerIPv6Prefix() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("syntheticDNSAnswerIPv6Prefix"))
	return rv
}


// An IPv6 address prefix the relay uses to handle address info requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nerelay/syntheticdnsansweripv6prefix
func (n_ NERelay) SetSyntheticDNSAnswerIPv6Prefix(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSyntheticDNSAnswerIPv6Prefix:"), value)
}




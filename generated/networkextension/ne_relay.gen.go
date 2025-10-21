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
}

// A single relay server configuration that you can chain together with other relays.
//
// Relay servers are secure HTTP proxies that allow proxying TCP traffic using the method and UDP traffic using the protocol defined in .
//
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


// A dictionary of additional HTTP headers to send as part of requests to the relay.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/additionalHTTPHeaderFields
func (n_ NERelay) AdditionalHTTPHeaderFields() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("additionalHTTPHeaderFields"))
	return rv
}


// SetAdditionalHTTPHeaderFields sets the value of the additionalHTTPHeaderFields property.
// A dictionary of additional HTTP headers to send as part of requests to the relay.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/additionalHTTPHeaderFields
func (n_ NERelay) SetAdditionalHTTPHeaderFields(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAdditionalHTTPHeaderFields:"), value)
}

// The URL of a DNS-over-HTTPS (DoH) resolver accessible from the relay.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/dnsOverHTTPSURL
func (n_ NERelay) DnsOverHTTPSURL() foundation.URL {
	rv := objc.Send[foundation.URL](n_.ID, objc.Sel("dnsOverHTTPSURL"))
	return rv
}


// SetDnsOverHTTPSURL sets the value of the dnsOverHTTPSURL property.
// The URL of a DNS-over-HTTPS (DoH) resolver accessible from the relay.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/dnsOverHTTPSURL
func (n_ NERelay) SetDnsOverHTTPSURL(value foundation.URL) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDnsOverHTTPSURL:"), value)
}

// A URL identifying the relay server accessible using HTTP/2.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/http2RelayURL
func (n_ NERelay) HTTP2RelayURL() foundation.URL {
	rv := objc.Send[foundation.URL](n_.ID, objc.Sel("HTTP2RelayURL"))
	return rv
}


// SetHTTP2RelayURL sets the value of the HTTP2RelayURL property.
// A URL identifying the relay server accessible using HTTP/2.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/http2RelayURL
func (n_ NERelay) SetHTTP2RelayURL(value foundation.URL) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHTTP2RelayURL:"), value)
}

// A URL identifying the relay server accessible using HTTP/3.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/http3RelayURL
func (n_ NERelay) HTTP3RelayURL() foundation.URL {
	rv := objc.Send[foundation.URL](n_.ID, objc.Sel("HTTP3RelayURL"))
	return rv
}


// SetHTTP3RelayURL sets the value of the HTTP3RelayURL property.
// A URL identifying the relay server accessible using HTTP/3.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/http3RelayURL
func (n_ NERelay) SetHTTP3RelayURL(value foundation.URL) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHTTP3RelayURL:"), value)
}

// The PKCS12 data for the relay client authentication.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/identityData
func (n_ NERelay) IdentityData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("identityData"))
	return rv
}


// SetIdentityData sets the value of the identityData property.
// The PKCS12 data for the relay client authentication.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/identityData
func (n_ NERelay) SetIdentityData(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIdentityData:"), value)
}

// The password the relay uses to decrypt the PKCS12 identity data.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/identityDataPassword
func (n_ NERelay) IdentityDataPassword() string {
	rv := objc.Send[string](n_.ID, objc.Sel("identityDataPassword"))
	return rv
}


// SetIdentityDataPassword sets the value of the identityDataPassword property.
// The password the relay uses to decrypt the PKCS12 identity data.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/identityDataPassword
func (n_ NERelay) SetIdentityDataPassword(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIdentityDataPassword:"), objc.String(value))
}

// An array of TLS raw public keys that the relay server can present during the TLS handshake.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/rawPublicKeys
func (n_ NERelay) RawPublicKeys() []unsafe.Pointer {
	rv := objc.Send[[]unsafe.Pointer](n_.ID, objc.Sel("rawPublicKeys"))
	return rv
}


// SetRawPublicKeys sets the value of the rawPublicKeys property.
// An array of TLS raw public keys that the relay server can present during the TLS handshake.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/rawPublicKeys
func (n_ NERelay) SetRawPublicKeys(value []unsafe.Pointer) {
	// Convert Go slice to NSArray
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
}

// An IPv4 address prefix the relay uses to handle address info requests.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/syntheticDNSAnswerIPv4Prefix
func (n_ NERelay) SyntheticDNSAnswerIPv4Prefix() string {
	rv := objc.Send[string](n_.ID, objc.Sel("syntheticDNSAnswerIPv4Prefix"))
	return rv
}


// SetSyntheticDNSAnswerIPv4Prefix sets the value of the syntheticDNSAnswerIPv4Prefix property.
// An IPv4 address prefix the relay uses to handle address info requests.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/syntheticDNSAnswerIPv4Prefix
func (n_ NERelay) SetSyntheticDNSAnswerIPv4Prefix(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSyntheticDNSAnswerIPv4Prefix:"), objc.String(value))
}

// An IPv6 address prefix the relay uses to handle address info requests.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/syntheticDNSAnswerIPv6Prefix
func (n_ NERelay) SyntheticDNSAnswerIPv6Prefix() string {
	rv := objc.Send[string](n_.ID, objc.Sel("syntheticDNSAnswerIPv6Prefix"))
	return rv
}


// SetSyntheticDNSAnswerIPv6Prefix sets the value of the syntheticDNSAnswerIPv6Prefix property.
// An IPv6 address prefix the relay uses to handle address info requests.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NERelay/syntheticDNSAnswerIPv6Prefix
func (n_ NERelay) SetSyntheticDNSAnswerIPv6Prefix(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSyntheticDNSAnswerIPv6Prefix:"), objc.String(value))
}




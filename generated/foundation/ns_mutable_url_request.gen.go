// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/networkextension"
)

// The class instance for the [MutableURLRequest] class.
var (
	MutableURLRequestClass     _MutableURLRequestClass
	MutableURLRequestClassOnce sync.Once
)

func getMutableURLRequestClass() _MutableURLRequestClass {
	MutableURLRequestClassOnce.Do(func() {
		MutableURLRequestClass = _MutableURLRequestClass{objc.GetClass("NSMutableURLRequest")}
	})
	return MutableURLRequestClass
}

type _MutableURLRequestClass struct {
	class objc.Class
}

// An interface definition for the [MutableURLRequest] class.
type IMutableURLRequest interface {
	IURLRequest
	BindToHotspotHelperCommand(command networkextension.INEHotspotHelperCommand)
}

// A mutable URL load request that is independent of protocol or URL scheme.
//
// In Swift, this object bridges to and you use when you need reference semantics or other Foundation-specific behavior. is a subclass of that allows you to change the request’s properties. only represents information about the request. Use other classes, such as , to send the request to a server. See and for an introduction to these techniques. Classes that create a network operation based on a request make a deep copy of that request. Thus, changing the request after creating a network operation has no effect on the ongoing operation. For example, if you use to create a data task from a request, and then later change the request, the data task continues using the original request.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest
type MutableURLRequest struct {
	URLRequest
}

// MutableURLRequestFrom constructs a [MutableURLRequest] from an unsafe.Pointer.
//
// A mutable URL load request that is independent of protocol or URL scheme.
func MutableURLRequestFrom(ptr unsafe.Pointer) MutableURLRequest {
	return MutableURLRequest{
		URLRequest: URLRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MutableURLRequestClass) Alloc() MutableURLRequest {
	rv := objc.Send[MutableURLRequest](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MutableURLRequestClass) New() MutableURLRequest {
	rv := objc.Send[MutableURLRequest](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableURLRequest) Init() MutableURLRequest {
	rv := objc.Send[MutableURLRequest](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableURLRequest) Autorelease() MutableURLRequest {
	rv := objc.Send[MutableURLRequest](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableURLRequest creates a new MutableURLRequest instance.
func NewMutableURLRequest() MutableURLRequest {
	return getMutableURLRequestClass().New()
}


// Binds a URL request to the network interface associated with the hotspot helper command instance.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/bind(to:)
func (m_ MutableURLRequest) BindToHotspotHelperCommand(command networkextension.INEHotspotHelperCommand) {
	objc.Send[objc.ID](m_.ID, objc.Sel("bindToHotspotHelperCommand:"), command)
}

// The request’s cache policy.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/cachePolicy
func (m_ MutableURLRequest) CachePolicy() URLRequestCachePolicy {
	rv := objc.Send[URLRequestCachePolicy](m_.ID, objc.Sel("cachePolicy"))
	return rv
}


// SetCachePolicy sets the value of the cachePolicy property.
// The request’s cache policy.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/cachePolicy
func (m_ MutableURLRequest) SetCachePolicy(value URLRequestCachePolicy) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCachePolicy:"), value)
}

// The HTTP request method.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/httpMethod
func (m_ MutableURLRequest) HTTPMethod() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("HTTPMethod"))
	return rv
}


// SetHTTPMethod sets the value of the HTTPMethod property.
// The HTTP request method.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/httpMethod
func (m_ MutableURLRequest) SetHTTPMethod(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHTTPMethod:"), value)
}

// The URL being requested.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/url
func (m_ MutableURLRequest) URL() URL {
	rv := objc.Send[URL](m_.ID, objc.Sel("URL"))
	return rv
}


// SetURL sets the value of the URL property.
// The URL being requested.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/url
func (m_ MutableURLRequest) SetURL(value IURL) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setURL:"), value)
}

// A dictionary containing all of the HTTP header fields for a request.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/allhttpheaderfields
func (m_ MutableURLRequest) AllHTTPHeaderFields() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("allHTTPHeaderFields"))
	return rv
}


// SetAllHTTPHeaderFields sets the value of the allHTTPHeaderFields property.
// A dictionary containing all of the HTTP header fields for a request.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/allhttpheaderfields
func (m_ MutableURLRequest) SetAllHTTPHeaderFields(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllHTTPHeaderFields:"), value)
}

// A Boolean value that indicates whether a connection can use the device’s cellular network (if present).
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/allowscellularaccess
func (m_ MutableURLRequest) AllowsCellularAccess() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsCellularAccess"))
	return rv
}


// SetAllowsCellularAccess sets the value of the allowsCellularAccess property.
// A Boolean value that indicates whether a connection can use the device’s cellular network (if present).

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/allowscellularaccess
func (m_ MutableURLRequest) SetAllowsCellularAccess(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsCellularAccess:"), value)
}

// A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/allowsconstrainednetworkaccess
func (m_ MutableURLRequest) AllowsConstrainedNetworkAccess() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsConstrainedNetworkAccess"))
	return rv
}


// SetAllowsConstrainedNetworkAccess sets the value of the allowsConstrainedNetworkAccess property.
// A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/allowsconstrainednetworkaccess
func (m_ MutableURLRequest) SetAllowsConstrainedNetworkAccess(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsConstrainedNetworkAccess:"), value)
}

// A Boolean value that indicates whether connections may use a network interface that the system considers expensive.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/allowsexpensivenetworkaccess
func (m_ MutableURLRequest) AllowsExpensiveNetworkAccess() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsExpensiveNetworkAccess"))
	return rv
}


// SetAllowsExpensiveNetworkAccess sets the value of the allowsExpensiveNetworkAccess property.
// A Boolean value that indicates whether connections may use a network interface that the system considers expensive.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/allowsexpensivenetworkaccess
func (m_ MutableURLRequest) SetAllowsExpensiveNetworkAccess(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsExpensiveNetworkAccess:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/allowspersistentdns
func (m_ MutableURLRequest) AllowsPersistentDNS() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsPersistentDNS"))
	return rv
}


// SetAllowsPersistentDNS sets the value of the allowsPersistentDNS property.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/allowspersistentdns
func (m_ MutableURLRequest) SetAllowsPersistentDNS(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsPersistentDNS:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/allowsultraconstrainednetworkaccess
func (m_ MutableURLRequest) AllowsUltraConstrainedNetworkAccess() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsUltraConstrainedNetworkAccess"))
	return rv
}


// SetAllowsUltraConstrainedNetworkAccess sets the value of the allowsUltraConstrainedNetworkAccess property.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/allowsultraconstrainednetworkaccess
func (m_ MutableURLRequest) SetAllowsUltraConstrainedNetworkAccess(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsUltraConstrainedNetworkAccess:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/assumeshttp3capable
func (m_ MutableURLRequest) AssumesHTTP3Capable() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("assumesHTTP3Capable"))
	return rv
}


// SetAssumesHTTP3Capable sets the value of the assumesHTTP3Capable property.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/assumeshttp3capable
func (m_ MutableURLRequest) SetAssumesHTTP3Capable(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAssumesHTTP3Capable:"), value)
}

// The entity that initiates the network request.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/attribution
func (m_ MutableURLRequest) Attribution() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("attribution"))
	return rv
}


// SetAttribution sets the value of the attribution property.
// The entity that initiates the network request.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/attribution
func (m_ MutableURLRequest) SetAttribution(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttribution:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/cookiepartitionidentifier
func (m_ MutableURLRequest) CookiePartitionIdentifier() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("cookiePartitionIdentifier"))
	return rv
}


// SetCookiePartitionIdentifier sets the value of the cookiePartitionIdentifier property.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/cookiepartitionidentifier
func (m_ MutableURLRequest) SetCookiePartitionIdentifier(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCookiePartitionIdentifier:"), value)
}

// The request body.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/httpbody
func (m_ MutableURLRequest) HttpBody() Data {
	rv := objc.Send[Data](m_.ID, objc.Sel("httpBody"))
	return rv
}


// SetHttpBody sets the value of the httpBody property.
// The request body.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/httpbody
func (m_ MutableURLRequest) SetHttpBody(value IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHttpBody:"), value)
}

// The request body as an input stream.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/httpbodystream
func (m_ MutableURLRequest) HttpBodyStream() NSInputStream {
	rv := objc.Send[NSInputStream](m_.ID, objc.Sel("httpBodyStream"))
	return rv
}


// SetHttpBodyStream sets the value of the httpBodyStream property.
// The request body as an input stream.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/httpbodystream
func (m_ MutableURLRequest) SetHttpBodyStream(value IInputStream) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHttpBodyStream:"), value)
}

// A Boolean value that indicates whether the request should use the default cookie handling for the request.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/httpshouldhandlecookies
func (m_ MutableURLRequest) HttpShouldHandleCookies() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("httpShouldHandleCookies"))
	return rv
}


// SetHttpShouldHandleCookies sets the value of the httpShouldHandleCookies property.
// A Boolean value that indicates whether the request should use the default cookie handling for the request.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/httpshouldhandlecookies
func (m_ MutableURLRequest) SetHttpShouldHandleCookies(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHttpShouldHandleCookies:"), value)
}

// A Boolean value that indicates whether the request can continue transmitting data before receiving a response from an earlier transmission.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/httpshouldusepipelining
func (m_ MutableURLRequest) HttpShouldUsePipelining() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("httpShouldUsePipelining"))
	return rv
}


// SetHttpShouldUsePipelining sets the value of the httpShouldUsePipelining property.
// A Boolean value that indicates whether the request can continue transmitting data before receiving a response from an earlier transmission.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/httpshouldusepipelining
func (m_ MutableURLRequest) SetHttpShouldUsePipelining(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHttpShouldUsePipelining:"), value)
}

// The main document URL.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/maindocumenturl
func (m_ MutableURLRequest) MainDocumentURL() URL {
	rv := objc.Send[URL](m_.ID, objc.Sel("mainDocumentURL"))
	return rv
}


// SetMainDocumentURL sets the value of the mainDocumentURL property.
// The main document URL.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/maindocumenturl
func (m_ MutableURLRequest) SetMainDocumentURL(value IURL) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMainDocumentURL:"), value)
}

// The network service type of the connection.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/networkservicetype
func (m_ MutableURLRequest) NetworkServiceType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("networkServiceType"))
	return rv
}


// SetNetworkServiceType sets the value of the networkServiceType property.
// The network service type of the connection.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/networkservicetype
func (m_ MutableURLRequest) SetNetworkServiceType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetworkServiceType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/requiresdnssecvalidation
func (m_ MutableURLRequest) RequiresDNSSECValidation() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("requiresDNSSECValidation"))
	return rv
}


// SetRequiresDNSSECValidation sets the value of the requiresDNSSECValidation property.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/requiresdnssecvalidation
func (m_ MutableURLRequest) SetRequiresDNSSECValidation(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequiresDNSSECValidation:"), value)
}

// The request’s timeout interval, in seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/timeoutinterval
func (m_ MutableURLRequest) TimeoutInterval() TimeInterval {
	rv := objc.Send[TimeInterval](m_.ID, objc.Sel("timeoutInterval"))
	return rv
}


// SetTimeoutInterval sets the value of the timeoutInterval property.
// The request’s timeout interval, in seconds.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/timeoutinterval
func (m_ MutableURLRequest) SetTimeoutInterval(value ITimeInterval) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeoutInterval:"), value)
}




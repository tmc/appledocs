// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLRequest] class.
var (
	URLRequestClass     _URLRequestClass
	URLRequestClassOnce sync.Once
)

func getURLRequestClass() _URLRequestClass {
	URLRequestClassOnce.Do(func() {
		URLRequestClass = _URLRequestClass{objc.GetClass("NSURLRequest")}
	})
	return URLRequestClass
}

type _URLRequestClass struct {
	class objc.Class
}

// An interface definition for the [URLRequest] class.
type IURLRequest interface {
	objectivec.IObject
	// properties:
	AllHTTPHeaderFields() string /* primitive/slice/pointer */
	SetAllHTTPHeaderFields(value string /* primitive/slice/pointer */)
	AllowsCellularAccess() bool /* primitive/slice/pointer */
	SetAllowsCellularAccess(value bool /* primitive/slice/pointer */)
	AllowsConstrainedNetworkAccess() bool /* primitive/slice/pointer */
	SetAllowsConstrainedNetworkAccess(value bool /* primitive/slice/pointer */)
	AllowsExpensiveNetworkAccess() bool /* primitive/slice/pointer */
	SetAllowsExpensiveNetworkAccess(value bool /* primitive/slice/pointer */)
	AllowsPersistentDNS() bool /* primitive/slice/pointer */
	SetAllowsPersistentDNS(value bool /* primitive/slice/pointer */)
	AllowsUltraConstrainedNetworkAccess() bool /* primitive/slice/pointer */
	SetAllowsUltraConstrainedNetworkAccess(value bool /* primitive/slice/pointer */)
	AssumesHTTP3Capable() bool /* primitive/slice/pointer */
	SetAssumesHTTP3Capable(value bool /* primitive/slice/pointer */)
	Attribution() unsafe.Pointer
	SetAttribution(value unsafe.Pointer)
	CachePolicy() unsafe.Pointer
	SetCachePolicy(value unsafe.Pointer)
	CookiePartitionIdentifier() string /* primitive/slice/pointer */
	SetCookiePartitionIdentifier(value string /* primitive/slice/pointer */)
	HttpBody() IData
	SetHttpBody(value IData)
	HttpBodyStream() IInputStream
	SetHttpBodyStream(value IInputStream)
	HttpMethod() string /* primitive/slice/pointer */
	SetHttpMethod(value string /* primitive/slice/pointer */)
	HttpShouldHandleCookies() bool /* primitive/slice/pointer */
	SetHttpShouldHandleCookies(value bool /* primitive/slice/pointer */)
	HttpShouldUsePipelining() bool /* primitive/slice/pointer */
	SetHttpShouldUsePipelining(value bool /* primitive/slice/pointer */)
	MainDocumentURL() IURL
	SetMainDocumentURL(value IURL)
	NetworkServiceType() unsafe.Pointer
	SetNetworkServiceType(value unsafe.Pointer)
	RequiresDNSSECValidation() bool /* primitive/slice/pointer */
	SetRequiresDNSSECValidation(value bool /* primitive/slice/pointer */)
	TimeoutInterval() TimeInterval /* foo */
	SetTimeoutInterval(value TimeInterval /* foo */)
	Url() IURL
	SetUrl(value IURL)
	// methods:
}

// A URL load request that is independent of protocol or URL scheme.
//
// Use this type in Swift when you need reference semantics or other Foundation-specific behavior. encapsulates two essential properties of a load request: the URL to load and the policies used to load it. In addition, for HTTP and HTTPS requests, includes the HTTP method ( , , and so on) and the HTTP headers. Finally, custom protocols can support custom properties as explained in . only represents information about the request. Use other classes, such as , to send the request to a server. See and for an introduction to these techniques. The mutable subclass of is .


// A URL load request that is independent of protocol or URL scheme.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest
type URLRequest struct {
	objectivec.Object
}

// URLRequestFrom constructs a [URLRequest] from an unsafe.Pointer.
//
// A URL load request that is independent of protocol or URL scheme.
func URLRequestFrom(ptr unsafe.Pointer) URLRequest {
	return URLRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLRequestClass) Alloc() URLRequest {
	rv := objc.Send[URLRequest](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLRequestClass) New() URLRequest {
	rv := objc.Send[URLRequest](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLRequest) Init() URLRequest {
	rv := objc.Send[URLRequest](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLRequest) Autorelease() URLRequest {
	rv := objc.Send[URLRequest](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLRequest creates a new URLRequest instance.
func NewURLRequest() URLRequest {
	return getURLRequestClass().New()
}



// A dictionary containing all of the HTTP header fields for a request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allhttpheaderfields
func (u_ URLRequest) AllHTTPHeaderFields() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("allHTTPHeaderFields"))
	return rv
}


// A dictionary containing all of the HTTP header fields for a request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allhttpheaderfields
func (u_ URLRequest) SetAllHTTPHeaderFields(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllHTTPHeaderFields:"), objc.String(value))
}


// A Boolean value that indicates whether the request is allowed to use the cellular radio (if present).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowscellularaccess
func (u_ URLRequest) AllowsCellularAccess() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsCellularAccess"))
	return rv
}


// A Boolean value that indicates whether the request is allowed to use the cellular radio (if present).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowscellularaccess
func (u_ URLRequest) SetAllowsCellularAccess(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsCellularAccess:"), value)
}


// A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowsconstrainednetworkaccess
func (u_ URLRequest) AllowsConstrainedNetworkAccess() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsConstrainedNetworkAccess"))
	return rv
}


// A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowsconstrainednetworkaccess
func (u_ URLRequest) SetAllowsConstrainedNetworkAccess(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsConstrainedNetworkAccess:"), value)
}


// A Boolean value that indicates whether connections may use a network interface that the system considers expensive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowsexpensivenetworkaccess
func (u_ URLRequest) AllowsExpensiveNetworkAccess() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsExpensiveNetworkAccess"))
	return rv
}


// A Boolean value that indicates whether connections may use a network interface that the system considers expensive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowsexpensivenetworkaccess
func (u_ URLRequest) SetAllowsExpensiveNetworkAccess(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsExpensiveNetworkAccess:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowspersistentdns
func (u_ URLRequest) AllowsPersistentDNS() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsPersistentDNS"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowspersistentdns
func (u_ URLRequest) SetAllowsPersistentDNS(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsPersistentDNS:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowsultraconstrainednetworkaccess
func (u_ URLRequest) AllowsUltraConstrainedNetworkAccess() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsUltraConstrainedNetworkAccess"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowsultraconstrainednetworkaccess
func (u_ URLRequest) SetAllowsUltraConstrainedNetworkAccess(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsUltraConstrainedNetworkAccess:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/assumeshttp3capable
func (u_ URLRequest) AssumesHTTP3Capable() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("assumesHTTP3Capable"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/assumeshttp3capable
func (u_ URLRequest) SetAssumesHTTP3Capable(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAssumesHTTP3Capable:"), value)
}


// The entity that initiates the network request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/attribution-swift.property
func (u_ URLRequest) Attribution() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("attribution"))
	return rv
}


// The entity that initiates the network request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/attribution-swift.property
func (u_ URLRequest) SetAttribution(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAttribution:"), value)
}


// The request’s cache policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/cachepolicy-swift.property
func (u_ URLRequest) CachePolicy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("cachePolicy"))
	return rv
}


// The request’s cache policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/cachepolicy-swift.property
func (u_ URLRequest) SetCachePolicy(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCachePolicy:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/cookiepartitionidentifier
func (u_ URLRequest) CookiePartitionIdentifier() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("cookiePartitionIdentifier"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/cookiepartitionidentifier
func (u_ URLRequest) SetCookiePartitionIdentifier(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCookiePartitionIdentifier:"), objc.String(value))
}


// The request body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/httpbody
func (u_ URLRequest) HttpBody() IData {
	rv := objc.Send[Data](u_.ID, objc.Sel("httpBody"))
	return rv
}


// The request body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/httpbody
func (u_ URLRequest) SetHttpBody(value IData) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHttpBody:"), value)
}


// The request body as an input stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/httpbodystream
func (u_ URLRequest) HttpBodyStream() IInputStream {
	rv := objc.Send[InputStream](u_.ID, objc.Sel("httpBodyStream"))
	return rv
}


// The request body as an input stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/httpbodystream
func (u_ URLRequest) SetHttpBodyStream(value IInputStream) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHttpBodyStream:"), value)
}


// The HTTP request method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/httpmethod
func (u_ URLRequest) HttpMethod() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](u_.ID, objc.Sel("httpMethod"))
	return rv
}


// The HTTP request method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/httpmethod
func (u_ URLRequest) SetHttpMethod(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHttpMethod:"), objc.String(value))
}


// A Boolean value that indicates whether the default cookie handling will be used for this request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/httpshouldhandlecookies
func (u_ URLRequest) HttpShouldHandleCookies() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("httpShouldHandleCookies"))
	return rv
}


// A Boolean value that indicates whether the default cookie handling will be used for this request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/httpshouldhandlecookies
func (u_ URLRequest) SetHttpShouldHandleCookies(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHttpShouldHandleCookies:"), value)
}


// A Boolean value that indicates whether the request should continue transmitting data before receiving a response from an earlier transmission.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/httpshouldusepipelining
func (u_ URLRequest) HttpShouldUsePipelining() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("httpShouldUsePipelining"))
	return rv
}


// A Boolean value that indicates whether the request should continue transmitting data before receiving a response from an earlier transmission.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/httpshouldusepipelining
func (u_ URLRequest) SetHttpShouldUsePipelining(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHttpShouldUsePipelining:"), value)
}


// The main document URL associated with the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/maindocumenturl
func (u_ URLRequest) MainDocumentURL() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("mainDocumentURL"))
	return rv
}


// The main document URL associated with the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/maindocumenturl
func (u_ URLRequest) SetMainDocumentURL(value IURL) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setMainDocumentURL:"), value)
}


// The network service type of the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/networkservicetype-swift.property
func (u_ URLRequest) NetworkServiceType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("networkServiceType"))
	return rv
}


// The network service type of the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/networkservicetype-swift.property
func (u_ URLRequest) SetNetworkServiceType(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNetworkServiceType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/requiresdnssecvalidation
func (u_ URLRequest) RequiresDNSSECValidation() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("requiresDNSSECValidation"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/requiresdnssecvalidation
func (u_ URLRequest) SetRequiresDNSSECValidation(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRequiresDNSSECValidation:"), value)
}


// The request’s timeout interval, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/timeoutinterval
func (u_ URLRequest) TimeoutInterval() TimeInterval /* foo */ {
	rv := objc.Send[TimeInterval](u_.ID, objc.Sel("timeoutInterval"))
	return rv
}


// The request’s timeout interval, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/timeoutinterval
func (u_ URLRequest) SetTimeoutInterval(value TimeInterval /* foo */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTimeoutInterval:"), value)
}


// The URL being requested.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/url
func (u_ URLRequest) Url() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("url"))
	return rv
}


// The URL being requested.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/url
func (u_ URLRequest) SetUrl(value IURL) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUrl:"), value)
}




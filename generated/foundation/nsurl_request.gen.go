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
	AllHTTPHeaderFields() IString
	SetAllHTTPHeaderFields(value IString)
	AllowsCellularAccess() bool
	SetAllowsCellularAccess(value bool)
	AllowsConstrainedNetworkAccess() bool
	SetAllowsConstrainedNetworkAccess(value bool)
	AllowsExpensiveNetworkAccess() bool
	SetAllowsExpensiveNetworkAccess(value bool)
	AllowsPersistentDNS() bool
	SetAllowsPersistentDNS(value bool)
	AllowsUltraConstrainedNetworkAccess() bool
	SetAllowsUltraConstrainedNetworkAccess(value bool)
	AssumesHTTP3Capable() bool
	SetAssumesHTTP3Capable(value bool)
	Attribution() unsafe.Pointer
	SetAttribution(value unsafe.Pointer)
	CachePolicy() unsafe.Pointer
	SetCachePolicy(value unsafe.Pointer)
	CookiePartitionIdentifier() IString
	SetCookiePartitionIdentifier(value IString)
	HttpBody() IData
	SetHttpBody(value IData)
	HttpBodyStream() IInputStream
	SetHttpBodyStream(value IInputStream)
	HttpMethod() IString
	SetHttpMethod(value IString)
	HttpShouldHandleCookies() bool
	SetHttpShouldHandleCookies(value bool)
	HttpShouldUsePipelining() bool
	SetHttpShouldUsePipelining(value bool)
	MainDocumentURL() IURL
	SetMainDocumentURL(value IURL)
	NetworkServiceType() unsafe.Pointer
	SetNetworkServiceType(value unsafe.Pointer)
	RequiresDNSSECValidation() bool
	SetRequiresDNSSECValidation(value bool)
	TimeoutInterval() float64
	SetTimeoutInterval(value float64)
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



// Creates a URL request for a specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/init(url:)
func NewURLRequestWithURL(URL IURL) URLRequest {
	instance := getURLRequestClass().Alloc()
	rv := objc.Send[URLRequest](instance.ID, objc.Sel("initWithURL:"), URL)
	rv.Autorelease()
	return rv
}



// A dictionary containing all of the HTTP header fields for a request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allhttpheaderfields
func (u_ URLRequest) AllHTTPHeaderFields() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("allHTTPHeaderFields"))
	return rv
}


// A dictionary containing all of the HTTP header fields for a request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allhttpheaderfields
func (u_ URLRequest) SetAllHTTPHeaderFields(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllHTTPHeaderFields:"), value)
}


// A Boolean value that indicates whether the request is allowed to use the cellular radio (if present).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowscellularaccess
func (u_ URLRequest) AllowsCellularAccess() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsCellularAccess"))
	return rv
}


// A Boolean value that indicates whether the request is allowed to use the cellular radio (if present).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowscellularaccess
func (u_ URLRequest) SetAllowsCellularAccess(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsCellularAccess:"), value)
}


// A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowsconstrainednetworkaccess
func (u_ URLRequest) AllowsConstrainedNetworkAccess() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsConstrainedNetworkAccess"))
	return rv
}


// A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowsconstrainednetworkaccess
func (u_ URLRequest) SetAllowsConstrainedNetworkAccess(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsConstrainedNetworkAccess:"), value)
}


// A Boolean value that indicates whether connections may use a network interface that the system considers expensive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowsexpensivenetworkaccess
func (u_ URLRequest) AllowsExpensiveNetworkAccess() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsExpensiveNetworkAccess"))
	return rv
}


// A Boolean value that indicates whether connections may use a network interface that the system considers expensive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowsexpensivenetworkaccess
func (u_ URLRequest) SetAllowsExpensiveNetworkAccess(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsExpensiveNetworkAccess:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowspersistentdns
func (u_ URLRequest) AllowsPersistentDNS() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsPersistentDNS"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowspersistentdns
func (u_ URLRequest) SetAllowsPersistentDNS(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsPersistentDNS:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowsultraconstrainednetworkaccess
func (u_ URLRequest) AllowsUltraConstrainedNetworkAccess() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsUltraConstrainedNetworkAccess"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowsultraconstrainednetworkaccess
func (u_ URLRequest) SetAllowsUltraConstrainedNetworkAccess(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsUltraConstrainedNetworkAccess:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/assumeshttp3capable
func (u_ URLRequest) AssumesHTTP3Capable() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("assumesHTTP3Capable"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/assumeshttp3capable
func (u_ URLRequest) SetAssumesHTTP3Capable(value bool) {
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
func (u_ URLRequest) CookiePartitionIdentifier() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("cookiePartitionIdentifier"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/cookiepartitionidentifier
func (u_ URLRequest) SetCookiePartitionIdentifier(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCookiePartitionIdentifier:"), value)
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
func (u_ URLRequest) HttpMethod() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("httpMethod"))
	return rv
}


// The HTTP request method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/httpmethod
func (u_ URLRequest) SetHttpMethod(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHttpMethod:"), value)
}


// A Boolean value that indicates whether the default cookie handling will be used for this request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/httpshouldhandlecookies
func (u_ URLRequest) HttpShouldHandleCookies() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("httpShouldHandleCookies"))
	return rv
}


// A Boolean value that indicates whether the default cookie handling will be used for this request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/httpshouldhandlecookies
func (u_ URLRequest) SetHttpShouldHandleCookies(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHttpShouldHandleCookies:"), value)
}


// A Boolean value that indicates whether the request should continue transmitting data before receiving a response from an earlier transmission.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/httpshouldusepipelining
func (u_ URLRequest) HttpShouldUsePipelining() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("httpShouldUsePipelining"))
	return rv
}


// A Boolean value that indicates whether the request should continue transmitting data before receiving a response from an earlier transmission.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/httpshouldusepipelining
func (u_ URLRequest) SetHttpShouldUsePipelining(value bool) {
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
func (u_ URLRequest) RequiresDNSSECValidation() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("requiresDNSSECValidation"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/requiresdnssecvalidation
func (u_ URLRequest) SetRequiresDNSSECValidation(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRequiresDNSSECValidation:"), value)
}


// The request’s timeout interval, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/timeoutinterval
func (u_ URLRequest) TimeoutInterval() float64 {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("timeoutInterval"))
	return rv
}


// The request’s timeout interval, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/timeoutinterval
func (u_ URLRequest) SetTimeoutInterval(value float64) {
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



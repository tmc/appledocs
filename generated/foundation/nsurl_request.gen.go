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
	AllowsConstrainedNetworkAccess() bool
	AllowsPersistentDNS() bool
	CookiePartitionIdentifier() string
	HTTPShouldHandleCookies() bool
	NetworkServiceType() URLRequestNetworkServiceType
	AllHTTPHeaderFields() string
	SetAllHTTPHeaderFields(value string)
	AllowsCellularAccess() bool
	SetAllowsCellularAccess(value bool)
	AllowsExpensiveNetworkAccess() bool
	SetAllowsExpensiveNetworkAccess(value bool)
	AllowsUltraConstrainedNetworkAccess() bool
	SetAllowsUltraConstrainedNetworkAccess(value bool)
	AssumesHTTP3Capable() bool
	SetAssumesHTTP3Capable(value bool)
	Attribution() unsafe.Pointer
	SetAttribution(value unsafe.Pointer)
	CachePolicy() unsafe.Pointer
	SetCachePolicy(value unsafe.Pointer)
	HttpBody() Data
	SetHttpBody(value IData)
	HttpBodyStream() NSInputStream
	SetHttpBodyStream(value IInputStream)
	HttpMethod() string
	SetHttpMethod(value string)
	HttpShouldUsePipelining() bool
	SetHttpShouldUsePipelining(value bool)
	MainDocumentURL() URL
	SetMainDocumentURL(value IURL)
	RequiresDNSSECValidation() bool
	SetRequiresDNSSECValidation(value bool)
	TimeoutInterval() TimeInterval
	SetTimeoutInterval(value ITimeInterval)
	Url() URL
	SetUrl(value IURL)
}

// A URL load request that is independent of protocol or URL scheme.
//
// Use this type in Swift when you need reference semantics or other Foundation-specific behavior. encapsulates two essential properties of a load request: the URL to load and the policies used to load it. In addition, for HTTP and HTTPS requests, includes the HTTP method ( , , and so on) and the HTTP headers. Finally, custom protocols can support custom properties as explained in . only represents information about the request. Use other classes, such as , to send the request to a server. See and for an introduction to these techniques. The mutable subclass of is .
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


// A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/allowsConstrainedNetworkAccess
func (u_ URLRequest) AllowsConstrainedNetworkAccess() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsConstrainedNetworkAccess"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/allowsPersistentDNS
func (u_ URLRequest) AllowsPersistentDNS() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsPersistentDNS"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/cookiePartitionIdentifier
func (u_ URLRequest) CookiePartitionIdentifier() string {
	rv := objc.Send[string](u_.ID, objc.Sel("cookiePartitionIdentifier"))
	return rv
}

// A Boolean value that indicates whether the default cookie handling will be used for this request.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/httpShouldHandleCookies
func (u_ URLRequest) HTTPShouldHandleCookies() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("HTTPShouldHandleCookies"))
	return rv
}

// The network service type of the request.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/networkServiceType-swift.property
func (u_ URLRequest) NetworkServiceType() URLRequestNetworkServiceType {
	rv := objc.Send[URLRequestNetworkServiceType](u_.ID, objc.Sel("networkServiceType"))
	return rv
}

// A dictionary containing all of the HTTP header fields for a request.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allhttpheaderfields
func (u_ URLRequest) AllHTTPHeaderFields() string {
	rv := objc.Send[string](u_.ID, objc.Sel("allHTTPHeaderFields"))
	return rv
}


// SetAllHTTPHeaderFields sets the value of the allHTTPHeaderFields property.
// A dictionary containing all of the HTTP header fields for a request.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allhttpheaderfields
func (u_ URLRequest) SetAllHTTPHeaderFields(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllHTTPHeaderFields:"), objc.String(value))
}

// A Boolean value that indicates whether the request is allowed to use the cellular radio (if present).
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowscellularaccess
func (u_ URLRequest) AllowsCellularAccess() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsCellularAccess"))
	return rv
}


// SetAllowsCellularAccess sets the value of the allowsCellularAccess property.
// A Boolean value that indicates whether the request is allowed to use the cellular radio (if present).

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowscellularaccess
func (u_ URLRequest) SetAllowsCellularAccess(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsCellularAccess:"), value)
}

// A Boolean value that indicates whether connections may use a network interface that the system considers expensive.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowsexpensivenetworkaccess
func (u_ URLRequest) AllowsExpensiveNetworkAccess() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsExpensiveNetworkAccess"))
	return rv
}


// SetAllowsExpensiveNetworkAccess sets the value of the allowsExpensiveNetworkAccess property.
// A Boolean value that indicates whether connections may use a network interface that the system considers expensive.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowsexpensivenetworkaccess
func (u_ URLRequest) SetAllowsExpensiveNetworkAccess(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsExpensiveNetworkAccess:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowsultraconstrainednetworkaccess
func (u_ URLRequest) AllowsUltraConstrainedNetworkAccess() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsUltraConstrainedNetworkAccess"))
	return rv
}


// SetAllowsUltraConstrainedNetworkAccess sets the value of the allowsUltraConstrainedNetworkAccess property.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowsultraconstrainednetworkaccess
func (u_ URLRequest) SetAllowsUltraConstrainedNetworkAccess(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsUltraConstrainedNetworkAccess:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/assumeshttp3capable
func (u_ URLRequest) AssumesHTTP3Capable() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("assumesHTTP3Capable"))
	return rv
}


// SetAssumesHTTP3Capable sets the value of the assumesHTTP3Capable property.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/assumeshttp3capable
func (u_ URLRequest) SetAssumesHTTP3Capable(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAssumesHTTP3Capable:"), value)
}

// The entity that initiates the network request.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/attribution-swift.property
func (u_ URLRequest) Attribution() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("attribution"))
	return rv
}


// SetAttribution sets the value of the attribution property.
// The entity that initiates the network request.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/attribution-swift.property
func (u_ URLRequest) SetAttribution(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAttribution:"), value)
}

// The request’s cache policy.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/cachepolicy-swift.property
func (u_ URLRequest) CachePolicy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("cachePolicy"))
	return rv
}


// SetCachePolicy sets the value of the cachePolicy property.
// The request’s cache policy.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/cachepolicy-swift.property
func (u_ URLRequest) SetCachePolicy(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCachePolicy:"), value)
}

// The request body.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/httpbody
func (u_ URLRequest) HttpBody() Data {
	rv := objc.Send[Data](u_.ID, objc.Sel("httpBody"))
	return rv
}


// SetHttpBody sets the value of the httpBody property.
// The request body.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/httpbody
func (u_ URLRequest) SetHttpBody(value IData) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHttpBody:"), value)
}

// The request body as an input stream.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/httpbodystream
func (u_ URLRequest) HttpBodyStream() NSInputStream {
	rv := objc.Send[NSInputStream](u_.ID, objc.Sel("httpBodyStream"))
	return rv
}


// SetHttpBodyStream sets the value of the httpBodyStream property.
// The request body as an input stream.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/httpbodystream
func (u_ URLRequest) SetHttpBodyStream(value IInputStream) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHttpBodyStream:"), value)
}

// The HTTP request method.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/httpmethod
func (u_ URLRequest) HttpMethod() string {
	rv := objc.Send[string](u_.ID, objc.Sel("httpMethod"))
	return rv
}


// SetHttpMethod sets the value of the httpMethod property.
// The HTTP request method.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/httpmethod
func (u_ URLRequest) SetHttpMethod(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHttpMethod:"), objc.String(value))
}

// A Boolean value that indicates whether the request should continue transmitting data before receiving a response from an earlier transmission.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/httpshouldusepipelining
func (u_ URLRequest) HttpShouldUsePipelining() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("httpShouldUsePipelining"))
	return rv
}


// SetHttpShouldUsePipelining sets the value of the httpShouldUsePipelining property.
// A Boolean value that indicates whether the request should continue transmitting data before receiving a response from an earlier transmission.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/httpshouldusepipelining
func (u_ URLRequest) SetHttpShouldUsePipelining(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHttpShouldUsePipelining:"), value)
}

// The main document URL associated with the request.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/maindocumenturl
func (u_ URLRequest) MainDocumentURL() URL {
	rv := objc.Send[URL](u_.ID, objc.Sel("mainDocumentURL"))
	return rv
}


// SetMainDocumentURL sets the value of the mainDocumentURL property.
// The main document URL associated with the request.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/maindocumenturl
func (u_ URLRequest) SetMainDocumentURL(value IURL) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setMainDocumentURL:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/requiresdnssecvalidation
func (u_ URLRequest) RequiresDNSSECValidation() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("requiresDNSSECValidation"))
	return rv
}


// SetRequiresDNSSECValidation sets the value of the requiresDNSSECValidation property.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/requiresdnssecvalidation
func (u_ URLRequest) SetRequiresDNSSECValidation(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRequiresDNSSECValidation:"), value)
}

// The request’s timeout interval, in seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/timeoutinterval
func (u_ URLRequest) TimeoutInterval() TimeInterval {
	rv := objc.Send[TimeInterval](u_.ID, objc.Sel("timeoutInterval"))
	return rv
}


// SetTimeoutInterval sets the value of the timeoutInterval property.
// The request’s timeout interval, in seconds.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/timeoutinterval
func (u_ URLRequest) SetTimeoutInterval(value ITimeInterval) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTimeoutInterval:"), value)
}

// The URL being requested.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/url
func (u_ URLRequest) Url() URL {
	rv := objc.Send[URL](u_.ID, objc.Sel("url"))
	return rv
}


// SetUrl sets the value of the url property.
// The URL being requested.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/url
func (u_ URLRequest) SetUrl(value IURL) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUrl:"), value)
}




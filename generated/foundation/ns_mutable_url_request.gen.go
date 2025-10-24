// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	

	// properties:
	AllowsCellularAccess() bool
	SetAllowsCellularAccess(value bool)
	Attribution() URLRequestAttribution /* not a class type */
	SetAttribution(value URLRequestAttribution /* not a class type */)
	HTTPShouldHandleCookies() bool
	SetHTTPShouldHandleCookies(value bool)
	HTTPShouldUsePipelining() bool
	SetHTTPShouldUsePipelining(value bool)
	NetworkServiceType() URLRequestNetworkServiceType
	SetNetworkServiceType(value URLRequestNetworkServiceType)
	TimeoutInterval() float64
	SetTimeoutInterval(value float64)
	AllHTTPHeaderFields() IString
	SetAllHTTPHeaderFields(value IString)
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
	CachePolicy() objectivec.IObject
	SetCachePolicy(value objectivec.IObject)
	CookiePartitionIdentifier() IString
	SetCookiePartitionIdentifier(value IString)
	HttpBody() IData
	SetHttpBody(value IData)
	HttpBodyStream() IInputStream
	SetHttpBodyStream(value IInputStream)
	HttpMethod() IString
	SetHttpMethod(value IString)
	MainDocumentURL() IURL
	SetMainDocumentURL(value IURL)
	RequiresDNSSECValidation() bool
	SetRequiresDNSSECValidation(value bool)
	Url() IURL
	SetUrl(value IURL)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MutableURLRequestClass) Alloc() MutableURLRequest {
	rv := objc.Send[MutableURLRequest](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// A mutable URL load request that is independent of protocol or URL scheme.
//
// In Swift, this object bridges to and you use when you need reference semantics or other Foundation-specific behavior. is a subclass of that allows you to change the request’s properties. only represents information about the request. Use other classes, such as , to send the request to a server. See and for an introduction to these techniques. Classes that create a network operation based on a request make a deep copy of that request. Thus, changing the request after creating a network operation has no effect on the ongoing operation. For example, if you use to create a data task from a request, and then later change the request, the data task continues using the original request.


// A mutable URL load request that is independent of protocol or URL scheme.
//
// [Full Topic]
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

























// A Boolean value that indicates whether a connection can use the device’s cellular network (if present).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/allowsCellularAccess
func (m_ MutableURLRequest) AllowsCellularAccess() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsCellularAccess"))
	return rv
}


// A Boolean value that indicates whether a connection can use the device’s cellular network (if present).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/allowsCellularAccess
func (m_ MutableURLRequest) SetAllowsCellularAccess(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsCellularAccess:"), value)
}


// The entity that initiates the network request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/attribution
func (m_ MutableURLRequest) Attribution() URLRequestAttribution /* not a class type */ {
	rv := objc.Send[URLRequestAttribution](m_.ID, objc.Sel("attribution"))
	return rv
}


// The entity that initiates the network request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/attribution
func (m_ MutableURLRequest) SetAttribution(value URLRequestAttribution /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttribution:"), value)
}


// A Boolean value that indicates whether the request should use the default cookie handling for the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/httpShouldHandleCookies
func (m_ MutableURLRequest) HTTPShouldHandleCookies() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("HTTPShouldHandleCookies"))
	return rv
}


// A Boolean value that indicates whether the request should use the default cookie handling for the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/httpShouldHandleCookies
func (m_ MutableURLRequest) SetHTTPShouldHandleCookies(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHTTPShouldHandleCookies:"), value)
}


// A Boolean value that indicates whether the request can continue transmitting data before receiving a response from an earlier transmission.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/httpShouldUsePipelining
func (m_ MutableURLRequest) HTTPShouldUsePipelining() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("HTTPShouldUsePipelining"))
	return rv
}


// A Boolean value that indicates whether the request can continue transmitting data before receiving a response from an earlier transmission.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/httpShouldUsePipelining
func (m_ MutableURLRequest) SetHTTPShouldUsePipelining(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHTTPShouldUsePipelining:"), value)
}


// The network service type of the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/networkServiceType
func (m_ MutableURLRequest) NetworkServiceType() URLRequestNetworkServiceType {
	rv := objc.Send[URLRequestNetworkServiceType](m_.ID, objc.Sel("networkServiceType"))
	return rv
}


// The network service type of the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/networkServiceType
func (m_ MutableURLRequest) SetNetworkServiceType(value URLRequestNetworkServiceType) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetworkServiceType:"), value)
}


// The request’s timeout interval, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/timeoutInterval
func (m_ MutableURLRequest) TimeoutInterval() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("timeoutInterval"))
	return rv
}


// The request’s timeout interval, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/timeoutInterval
func (m_ MutableURLRequest) SetTimeoutInterval(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeoutInterval:"), value)
}


// A dictionary containing all of the HTTP header fields for a request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/allhttpheaderfields
func (m_ MutableURLRequest) AllHTTPHeaderFields() IString {
	rv := objc.Send[String](m_.ID, objc.Sel("allHTTPHeaderFields"))
	return rv
}


// A dictionary containing all of the HTTP header fields for a request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/allhttpheaderfields
func (m_ MutableURLRequest) SetAllHTTPHeaderFields(value IString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllHTTPHeaderFields:"), value)
}


// A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/allowsconstrainednetworkaccess
func (m_ MutableURLRequest) AllowsConstrainedNetworkAccess() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsConstrainedNetworkAccess"))
	return rv
}


// A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/allowsconstrainednetworkaccess
func (m_ MutableURLRequest) SetAllowsConstrainedNetworkAccess(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsConstrainedNetworkAccess:"), value)
}


// A Boolean value that indicates whether connections may use a network interface that the system considers expensive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/allowsexpensivenetworkaccess
func (m_ MutableURLRequest) AllowsExpensiveNetworkAccess() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsExpensiveNetworkAccess"))
	return rv
}


// A Boolean value that indicates whether connections may use a network interface that the system considers expensive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/allowsexpensivenetworkaccess
func (m_ MutableURLRequest) SetAllowsExpensiveNetworkAccess(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsExpensiveNetworkAccess:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/allowspersistentdns
func (m_ MutableURLRequest) AllowsPersistentDNS() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsPersistentDNS"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/allowspersistentdns
func (m_ MutableURLRequest) SetAllowsPersistentDNS(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsPersistentDNS:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/allowsultraconstrainednetworkaccess
func (m_ MutableURLRequest) AllowsUltraConstrainedNetworkAccess() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsUltraConstrainedNetworkAccess"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/allowsultraconstrainednetworkaccess
func (m_ MutableURLRequest) SetAllowsUltraConstrainedNetworkAccess(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsUltraConstrainedNetworkAccess:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/assumeshttp3capable
func (m_ MutableURLRequest) AssumesHTTP3Capable() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("assumesHTTP3Capable"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/assumeshttp3capable
func (m_ MutableURLRequest) SetAssumesHTTP3Capable(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAssumesHTTP3Capable:"), value)
}


// The request’s cache policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/cachepolicy
func (m_ MutableURLRequest) CachePolicy() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("cachePolicy"))
	return rv
}


// The request’s cache policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/cachepolicy
func (m_ MutableURLRequest) SetCachePolicy(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCachePolicy:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/cookiepartitionidentifier
func (m_ MutableURLRequest) CookiePartitionIdentifier() IString {
	rv := objc.Send[String](m_.ID, objc.Sel("cookiePartitionIdentifier"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/cookiepartitionidentifier
func (m_ MutableURLRequest) SetCookiePartitionIdentifier(value IString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCookiePartitionIdentifier:"), value)
}


// The request body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/httpbody
func (m_ MutableURLRequest) HttpBody() IData {
	rv := objc.Send[Data](m_.ID, objc.Sel("httpBody"))
	return rv
}


// The request body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/httpbody
func (m_ MutableURLRequest) SetHttpBody(value IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHttpBody:"), value)
}


// The request body as an input stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/httpbodystream
func (m_ MutableURLRequest) HttpBodyStream() IInputStream {
	rv := objc.Send[InputStream](m_.ID, objc.Sel("httpBodyStream"))
	return rv
}


// The request body as an input stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/httpbodystream
func (m_ MutableURLRequest) SetHttpBodyStream(value IInputStream) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHttpBodyStream:"), value)
}


// The HTTP request method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/httpmethod
func (m_ MutableURLRequest) HttpMethod() IString {
	rv := objc.Send[String](m_.ID, objc.Sel("httpMethod"))
	return rv
}


// The HTTP request method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/httpmethod
func (m_ MutableURLRequest) SetHttpMethod(value IString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHttpMethod:"), value)
}


// The main document URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/maindocumenturl
func (m_ MutableURLRequest) MainDocumentURL() IURL {
	rv := objc.Send[URL](m_.ID, objc.Sel("mainDocumentURL"))
	return rv
}


// The main document URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/maindocumenturl
func (m_ MutableURLRequest) SetMainDocumentURL(value IURL) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMainDocumentURL:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/requiresdnssecvalidation
func (m_ MutableURLRequest) RequiresDNSSECValidation() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("requiresDNSSECValidation"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/requiresdnssecvalidation
func (m_ MutableURLRequest) SetRequiresDNSSECValidation(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequiresDNSSECValidation:"), value)
}


// The URL being requested.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/url
func (m_ MutableURLRequest) Url() IURL {
	rv := objc.Send[URL](m_.ID, objc.Sel("url"))
	return rv
}


// The URL being requested.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableurlrequest/url
func (m_ MutableURLRequest) SetUrl(value IURL) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUrl:"), value)
}









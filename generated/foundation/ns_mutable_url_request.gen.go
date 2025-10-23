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
	AllHTTPHeaderFields() IDictionary /* already interface */
	SetAllHTTPHeaderFields(value IDictionary /* already interface */)
	AllowsCellularAccess() bool /* primitive/slice/pointer. */
	SetAllowsCellularAccess(value bool /* primitive/slice/pointer. */)
	AllowsConstrainedNetworkAccess() bool /* primitive/slice/pointer. */
	SetAllowsConstrainedNetworkAccess(value bool /* primitive/slice/pointer. */)
	AllowsExpensiveNetworkAccess() bool /* primitive/slice/pointer. */
	SetAllowsExpensiveNetworkAccess(value bool /* primitive/slice/pointer. */)
	AllowsPersistentDNS() bool /* primitive/slice/pointer. */
	SetAllowsPersistentDNS(value bool /* primitive/slice/pointer. */)
	AllowsUltraConstrainedNetworkAccess() bool /* primitive/slice/pointer. */
	SetAllowsUltraConstrainedNetworkAccess(value bool /* primitive/slice/pointer. */)
	AssumesHTTP3Capable() bool /* primitive/slice/pointer. */
	SetAssumesHTTP3Capable(value bool /* primitive/slice/pointer. */)
	Attribution() URLRequestAttribution
	SetAttribution(value URLRequestAttribution)
	CachePolicy() URLRequestCachePolicy
	SetCachePolicy(value URLRequestCachePolicy)
	CookiePartitionIdentifier() IString
	SetCookiePartitionIdentifier(value IString)
	HTTPBody() IData
	SetHTTPBody(value IData)
	HTTPBodyStream() IInputStream
	SetHTTPBodyStream(value IInputStream)
	HTTPMethod() IString
	SetHTTPMethod(value IString)
	HTTPShouldHandleCookies() bool /* primitive/slice/pointer. */
	SetHTTPShouldHandleCookies(value bool /* primitive/slice/pointer. */)
	HTTPShouldUsePipelining() bool /* primitive/slice/pointer. */
	SetHTTPShouldUsePipelining(value bool /* primitive/slice/pointer. */)
	MainDocumentURL() IURL
	SetMainDocumentURL(value IURL)
	NetworkServiceType() URLRequestNetworkServiceType
	SetNetworkServiceType(value URLRequestNetworkServiceType)
	RequiresDNSSECValidation() bool /* primitive/slice/pointer. */
	SetRequiresDNSSECValidation(value bool /* primitive/slice/pointer. */)
	TimeoutInterval() objc.IObject /* cross-framework: TimeInterval */
	SetTimeoutInterval(value objc.IObject /* cross-framework: TimeInterval */)
	URL() IURL
	SetURL(value IURL)
	// methods:
	AddValueForHTTPHeaderField(value IString, field IString)
	BindToHotspotHelperCommand(command objectivec.IObject)
	SetValueForHTTPHeaderField(value IString, field IString)
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



// Adds a value to the header field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/addValue(_:forHTTPHeaderField:)
func (m_ MutableURLRequest) AddValueForHTTPHeaderField(value IString, field IString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addValue:forHTTPHeaderField:"), value, field)
}


// Binds a URL request to the network interface associated with the hotspot helper command instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/bind(to:)
func (m_ MutableURLRequest) BindToHotspotHelperCommand(command objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("bindToHotspotHelperCommand:"), command)
}


// Sets a value for the header field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/setValue(_:forHTTPHeaderField:)
func (m_ MutableURLRequest) SetValueForHTTPHeaderField(value IString, field IString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:forHTTPHeaderField:"), value, field)
}


// A dictionary containing all of the HTTP header fields for a request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/allHTTPHeaderFields
func (m_ MutableURLRequest) AllHTTPHeaderFields() IDictionary /* already interface */ {
	rv := objc.Send[IDictionary](m_.ID, objc.Sel("allHTTPHeaderFields"))
	return rv
}


// A dictionary containing all of the HTTP header fields for a request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/allHTTPHeaderFields
func (m_ MutableURLRequest) SetAllHTTPHeaderFields(value IDictionary /* already interface */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllHTTPHeaderFields:"), value)
}


// A Boolean value that indicates whether a connection can use the device’s cellular network (if present).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/allowsCellularAccess
func (m_ MutableURLRequest) AllowsCellularAccess() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsCellularAccess"))
	return rv
}


// A Boolean value that indicates whether a connection can use the device’s cellular network (if present).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/allowsCellularAccess
func (m_ MutableURLRequest) SetAllowsCellularAccess(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsCellularAccess:"), value)
}


// A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/allowsConstrainedNetworkAccess
func (m_ MutableURLRequest) AllowsConstrainedNetworkAccess() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsConstrainedNetworkAccess"))
	return rv
}


// A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/allowsConstrainedNetworkAccess
func (m_ MutableURLRequest) SetAllowsConstrainedNetworkAccess(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsConstrainedNetworkAccess:"), value)
}


// A Boolean value that indicates whether connections may use a network interface that the system considers expensive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/allowsExpensiveNetworkAccess
func (m_ MutableURLRequest) AllowsExpensiveNetworkAccess() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsExpensiveNetworkAccess"))
	return rv
}


// A Boolean value that indicates whether connections may use a network interface that the system considers expensive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/allowsExpensiveNetworkAccess
func (m_ MutableURLRequest) SetAllowsExpensiveNetworkAccess(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsExpensiveNetworkAccess:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/allowsPersistentDNS
func (m_ MutableURLRequest) AllowsPersistentDNS() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsPersistentDNS"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/allowsPersistentDNS
func (m_ MutableURLRequest) SetAllowsPersistentDNS(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsPersistentDNS:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/allowsUltraConstrainedNetworkAccess
func (m_ MutableURLRequest) AllowsUltraConstrainedNetworkAccess() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsUltraConstrainedNetworkAccess"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/allowsUltraConstrainedNetworkAccess
func (m_ MutableURLRequest) SetAllowsUltraConstrainedNetworkAccess(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsUltraConstrainedNetworkAccess:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/assumesHTTP3Capable
func (m_ MutableURLRequest) AssumesHTTP3Capable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("assumesHTTP3Capable"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/assumesHTTP3Capable
func (m_ MutableURLRequest) SetAssumesHTTP3Capable(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAssumesHTTP3Capable:"), value)
}


// The entity that initiates the network request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/attribution
func (m_ MutableURLRequest) Attribution() URLRequestAttribution {
	rv := objc.Send[URLRequestAttribution](m_.ID, objc.Sel("attribution"))
	return rv
}


// The entity that initiates the network request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/attribution
func (m_ MutableURLRequest) SetAttribution(value URLRequestAttribution) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttribution:"), value)
}


// The request’s cache policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/cachePolicy
func (m_ MutableURLRequest) CachePolicy() URLRequestCachePolicy {
	rv := objc.Send[URLRequestCachePolicy](m_.ID, objc.Sel("cachePolicy"))
	return rv
}


// The request’s cache policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/cachePolicy
func (m_ MutableURLRequest) SetCachePolicy(value URLRequestCachePolicy) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCachePolicy:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/cookiePartitionIdentifier
func (m_ MutableURLRequest) CookiePartitionIdentifier() IString {
	rv := objc.Send[String](m_.ID, objc.Sel("cookiePartitionIdentifier"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/cookiePartitionIdentifier
func (m_ MutableURLRequest) SetCookiePartitionIdentifier(value IString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCookiePartitionIdentifier:"), value)
}


// The request body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/httpBody
func (m_ MutableURLRequest) HTTPBody() IData {
	rv := objc.Send[Data](m_.ID, objc.Sel("HTTPBody"))
	return rv
}


// The request body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/httpBody
func (m_ MutableURLRequest) SetHTTPBody(value IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHTTPBody:"), value)
}


// The request body as an input stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/httpBodyStream
func (m_ MutableURLRequest) HTTPBodyStream() IInputStream {
	rv := objc.Send[InputStream](m_.ID, objc.Sel("HTTPBodyStream"))
	return rv
}


// The request body as an input stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/httpBodyStream
func (m_ MutableURLRequest) SetHTTPBodyStream(value IInputStream) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHTTPBodyStream:"), value)
}


// The HTTP request method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/httpMethod
func (m_ MutableURLRequest) HTTPMethod() IString {
	rv := objc.Send[String](m_.ID, objc.Sel("HTTPMethod"))
	return rv
}


// The HTTP request method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/httpMethod
func (m_ MutableURLRequest) SetHTTPMethod(value IString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHTTPMethod:"), value)
}


// A Boolean value that indicates whether the request should use the default cookie handling for the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/httpShouldHandleCookies
func (m_ MutableURLRequest) HTTPShouldHandleCookies() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("HTTPShouldHandleCookies"))
	return rv
}


// A Boolean value that indicates whether the request should use the default cookie handling for the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/httpShouldHandleCookies
func (m_ MutableURLRequest) SetHTTPShouldHandleCookies(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHTTPShouldHandleCookies:"), value)
}


// A Boolean value that indicates whether the request can continue transmitting data before receiving a response from an earlier transmission.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/httpShouldUsePipelining
func (m_ MutableURLRequest) HTTPShouldUsePipelining() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("HTTPShouldUsePipelining"))
	return rv
}


// A Boolean value that indicates whether the request can continue transmitting data before receiving a response from an earlier transmission.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/httpShouldUsePipelining
func (m_ MutableURLRequest) SetHTTPShouldUsePipelining(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHTTPShouldUsePipelining:"), value)
}


// The main document URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/mainDocumentURL
func (m_ MutableURLRequest) MainDocumentURL() IURL {
	rv := objc.Send[URL](m_.ID, objc.Sel("mainDocumentURL"))
	return rv
}


// The main document URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/mainDocumentURL
func (m_ MutableURLRequest) SetMainDocumentURL(value IURL) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMainDocumentURL:"), value)
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


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/requiresDNSSECValidation
func (m_ MutableURLRequest) RequiresDNSSECValidation() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("requiresDNSSECValidation"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/requiresDNSSECValidation
func (m_ MutableURLRequest) SetRequiresDNSSECValidation(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequiresDNSSECValidation:"), value)
}


// The request’s timeout interval, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/timeoutInterval
func (m_ MutableURLRequest) TimeoutInterval() objc.IObject /* cross-framework: TimeInterval */ {
	rv := objc.Send[TimeInterval](m_.ID, objc.Sel("timeoutInterval"))
	return rv
}


// The request’s timeout interval, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/timeoutInterval
func (m_ MutableURLRequest) SetTimeoutInterval(value objc.IObject /* cross-framework: TimeInterval */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeoutInterval:"), value)
}


// The URL being requested.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/url
func (m_ MutableURLRequest) URL() IURL {
	rv := objc.Send[URL](m_.ID, objc.Sel("URL"))
	return rv
}


// The URL being requested.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/url
func (m_ MutableURLRequest) SetURL(value IURL) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setURL:"), value)
}




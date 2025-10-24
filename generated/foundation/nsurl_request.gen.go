// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSURLRequest */


/* debug [class_header]: Header for NSURLRequest */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for URLRequest */
// An interface definition for the [URLRequest] class.
type IURLRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for URLRequest */
	// properties:
	AllowsCellularAccess() bool
	HTTPShouldHandleCookies() bool
	HTTPShouldUsePipelining() bool
	TimeoutInterval() float64
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
	Attribution() objectivec.IObject
	SetAttribution(value objectivec.IObject)
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
	NetworkServiceType() objectivec.IObject
	SetNetworkServiceType(value objectivec.IObject)
	RequiresDNSSECValidation() bool
	SetRequiresDNSSECValidation(value bool)
	Url() IURL
	SetUrl(value IURL)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for URLRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for URLRequest */
// Alloc allocates a new instance without initialization.
func (uc _URLRequestClass) Alloc() URLRequest {
	rv := objc.Send[URLRequest](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for URLRequest */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for URLRequest */

// Creates a URL request for a specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/init(url:)
func NewURLRequestWithURL(URL IURL) URLRequest {
	instance := getURLRequestClass().Alloc()
	rv := objc.Send[URLRequest](instance.ID, objc.Sel("initWithURL:"), URL)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewURLRequestWithURL */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for URLRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for URLRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for URLRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for URLRequest */

// A Boolean value that indicates whether the request is allowed to use the cellular radio (if present).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/allowsCellularAccess
func (u_ URLRequest) AllowsCellularAccess() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsCellularAccess"))
	return rv
}/* debug [instance_properties/getter]: allowsCellularAccess */


// A Boolean value that indicates whether the default cookie handling will be used for this request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/httpShouldHandleCookies
func (u_ URLRequest) HTTPShouldHandleCookies() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("HTTPShouldHandleCookies"))
	return rv
}/* debug [instance_properties/getter]: HTTPShouldHandleCookies */


// A Boolean value that indicates whether the request should continue transmitting data before receiving a response from an earlier transmission.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/httpShouldUsePipelining
func (u_ URLRequest) HTTPShouldUsePipelining() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("HTTPShouldUsePipelining"))
	return rv
}/* debug [instance_properties/getter]: HTTPShouldUsePipelining */


// The request’s timeout interval, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLRequest/timeoutInterval
func (u_ URLRequest) TimeoutInterval() float64 {
	rv := objc.Send[float64](u_.ID, objc.Sel("timeoutInterval"))
	return rv
}/* debug [instance_properties/getter]: timeoutInterval */


// A dictionary containing all of the HTTP header fields for a request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allhttpheaderfields
func (u_ URLRequest) AllHTTPHeaderFields() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("allHTTPHeaderFields"))
	return rv
}/* debug [instance_properties/getter]: allHTTPHeaderFields */


// A dictionary containing all of the HTTP header fields for a request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allhttpheaderfields
func (u_ URLRequest) SetAllHTTPHeaderFields(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllHTTPHeaderFields:"), value)
}/* debug [instance_properties/setter]: allHTTPHeaderFields */


// A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowsconstrainednetworkaccess
func (u_ URLRequest) AllowsConstrainedNetworkAccess() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsConstrainedNetworkAccess"))
	return rv
}/* debug [instance_properties/getter]: allowsConstrainedNetworkAccess */


// A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowsconstrainednetworkaccess
func (u_ URLRequest) SetAllowsConstrainedNetworkAccess(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsConstrainedNetworkAccess:"), value)
}/* debug [instance_properties/setter]: allowsConstrainedNetworkAccess */


// A Boolean value that indicates whether connections may use a network interface that the system considers expensive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowsexpensivenetworkaccess
func (u_ URLRequest) AllowsExpensiveNetworkAccess() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsExpensiveNetworkAccess"))
	return rv
}/* debug [instance_properties/getter]: allowsExpensiveNetworkAccess */


// A Boolean value that indicates whether connections may use a network interface that the system considers expensive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowsexpensivenetworkaccess
func (u_ URLRequest) SetAllowsExpensiveNetworkAccess(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsExpensiveNetworkAccess:"), value)
}/* debug [instance_properties/setter]: allowsExpensiveNetworkAccess */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowspersistentdns
func (u_ URLRequest) AllowsPersistentDNS() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsPersistentDNS"))
	return rv
}/* debug [instance_properties/getter]: allowsPersistentDNS */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowspersistentdns
func (u_ URLRequest) SetAllowsPersistentDNS(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsPersistentDNS:"), value)
}/* debug [instance_properties/setter]: allowsPersistentDNS */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowsultraconstrainednetworkaccess
func (u_ URLRequest) AllowsUltraConstrainedNetworkAccess() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsUltraConstrainedNetworkAccess"))
	return rv
}/* debug [instance_properties/getter]: allowsUltraConstrainedNetworkAccess */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/allowsultraconstrainednetworkaccess
func (u_ URLRequest) SetAllowsUltraConstrainedNetworkAccess(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsUltraConstrainedNetworkAccess:"), value)
}/* debug [instance_properties/setter]: allowsUltraConstrainedNetworkAccess */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/assumeshttp3capable
func (u_ URLRequest) AssumesHTTP3Capable() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("assumesHTTP3Capable"))
	return rv
}/* debug [instance_properties/getter]: assumesHTTP3Capable */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/assumeshttp3capable
func (u_ URLRequest) SetAssumesHTTP3Capable(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAssumesHTTP3Capable:"), value)
}/* debug [instance_properties/setter]: assumesHTTP3Capable */


// The entity that initiates the network request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/attribution-swift.property
func (u_ URLRequest) Attribution() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](u_.ID, objc.Sel("attribution"))
	return rv
}/* debug [instance_properties/getter]: attribution */


// The entity that initiates the network request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/attribution-swift.property
func (u_ URLRequest) SetAttribution(value objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAttribution:"), value)
}/* debug [instance_properties/setter]: attribution */


// The request’s cache policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/cachepolicy-swift.property
func (u_ URLRequest) CachePolicy() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](u_.ID, objc.Sel("cachePolicy"))
	return rv
}/* debug [instance_properties/getter]: cachePolicy */


// The request’s cache policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/cachepolicy-swift.property
func (u_ URLRequest) SetCachePolicy(value objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCachePolicy:"), value)
}/* debug [instance_properties/setter]: cachePolicy */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/cookiepartitionidentifier
func (u_ URLRequest) CookiePartitionIdentifier() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("cookiePartitionIdentifier"))
	return rv
}/* debug [instance_properties/getter]: cookiePartitionIdentifier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/cookiepartitionidentifier
func (u_ URLRequest) SetCookiePartitionIdentifier(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCookiePartitionIdentifier:"), value)
}/* debug [instance_properties/setter]: cookiePartitionIdentifier */


// The request body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/httpbody
func (u_ URLRequest) HttpBody() IData {
	rv := objc.Send[Data](u_.ID, objc.Sel("httpBody"))
	return rv
}/* debug [instance_properties/getter]: httpBody */


// The request body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/httpbody
func (u_ URLRequest) SetHttpBody(value IData) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHttpBody:"), value)
}/* debug [instance_properties/setter]: httpBody */


// The request body as an input stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/httpbodystream
func (u_ URLRequest) HttpBodyStream() IInputStream {
	rv := objc.Send[InputStream](u_.ID, objc.Sel("httpBodyStream"))
	return rv
}/* debug [instance_properties/getter]: httpBodyStream */


// The request body as an input stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/httpbodystream
func (u_ URLRequest) SetHttpBodyStream(value IInputStream) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHttpBodyStream:"), value)
}/* debug [instance_properties/setter]: httpBodyStream */


// The HTTP request method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/httpmethod
func (u_ URLRequest) HttpMethod() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("httpMethod"))
	return rv
}/* debug [instance_properties/getter]: httpMethod */


// The HTTP request method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/httpmethod
func (u_ URLRequest) SetHttpMethod(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHttpMethod:"), value)
}/* debug [instance_properties/setter]: httpMethod */


// The main document URL associated with the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/maindocumenturl
func (u_ URLRequest) MainDocumentURL() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("mainDocumentURL"))
	return rv
}/* debug [instance_properties/getter]: mainDocumentURL */


// The main document URL associated with the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/maindocumenturl
func (u_ URLRequest) SetMainDocumentURL(value IURL) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setMainDocumentURL:"), value)
}/* debug [instance_properties/setter]: mainDocumentURL */


// The network service type of the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/networkservicetype-swift.property
func (u_ URLRequest) NetworkServiceType() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](u_.ID, objc.Sel("networkServiceType"))
	return rv
}/* debug [instance_properties/getter]: networkServiceType */


// The network service type of the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/networkservicetype-swift.property
func (u_ URLRequest) SetNetworkServiceType(value objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNetworkServiceType:"), value)
}/* debug [instance_properties/setter]: networkServiceType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/requiresdnssecvalidation
func (u_ URLRequest) RequiresDNSSECValidation() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("requiresDNSSECValidation"))
	return rv
}/* debug [instance_properties/getter]: requiresDNSSECValidation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/requiresdnssecvalidation
func (u_ URLRequest) SetRequiresDNSSECValidation(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRequiresDNSSECValidation:"), value)
}/* debug [instance_properties/setter]: requiresDNSSECValidation */


// The URL being requested.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/url
func (u_ URLRequest) Url() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("url"))
	return rv
}/* debug [instance_properties/getter]: url */


// The URL being requested.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlrequest/url
func (u_ URLRequest) SetUrl(value IURL) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUrl:"), value)
}/* debug [instance_properties/setter]: url */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSURLRequest */



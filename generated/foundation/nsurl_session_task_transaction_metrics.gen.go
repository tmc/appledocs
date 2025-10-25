// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSURLSessionTaskTransactionMetrics */


/* debug [class_header]: Header for NSURLSessionTaskTransactionMetrics */
// The class instance for the [URLSessionTaskTransactionMetrics] class.
var (
	URLSessionTaskTransactionMetricsClass     _URLSessionTaskTransactionMetricsClass
	URLSessionTaskTransactionMetricsClassOnce sync.Once
)

func getURLSessionTaskTransactionMetricsClass() _URLSessionTaskTransactionMetricsClass {
	URLSessionTaskTransactionMetricsClassOnce.Do(func() {
		URLSessionTaskTransactionMetricsClass = _URLSessionTaskTransactionMetricsClass{objc.GetClass("NSURLSessionTaskTransactionMetrics")}
	})
	return URLSessionTaskTransactionMetricsClass
}

type _URLSessionTaskTransactionMetricsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for URLSessionTaskTransactionMetrics */
// An interface definition for the [URLSessionTaskTransactionMetrics] class.
type IURLSessionTaskTransactionMetrics interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for URLSessionTaskTransactionMetrics */
	// properties:
	LocalPort() INumber
	NegotiatedTLSCipherSuite() INumber
	NegotiatedTLSProtocolVersion() INumber
	RemotePort() INumber
	ConnectEndDate() IDate
	ConnectStartDate() IDate
	CountOfRequestBodyBytesBeforeEncoding() int64
	CountOfRequestBodyBytesSent() int64
	CountOfRequestHeaderBytesSent() int64
	CountOfResponseBodyBytesAfterDecoding() int64
	CountOfResponseBodyBytesReceived() int64
	CountOfResponseHeaderBytesReceived() int64
	DomainLookupEndDate() IDate
	DomainLookupStartDate() IDate
	DomainResolutionProtocol() URLSessionTaskMetricsDomainResolutionProtocol
	FetchStartDate() IDate
	Cellular() bool
	Constrained() bool
	Expensive() bool
	Multipath() bool
	ProxyConnection() bool
	ReusedConnection() bool
	LocalAddress() IString
	NetworkProtocolName() IString
	RemoteAddress() IString
	Request() IURLRequest
	RequestEndDate() IDate
	RequestStartDate() IDate
	ResourceFetchType() URLSessionTaskMetricsResourceFetchType
	Response() IURLResponse
	ResponseEndDate() IDate
	ResponseStartDate() IDate
	SecureConnectionEndDate() IDate
	SecureConnectionStartDate() IDate
	RedirectCount() int
	SetRedirectCount(value int)
	TaskInterval() IDateInterval
	SetTaskInterval(value IDateInterval)
	TransactionMetrics() IURLSessionTaskTransactionMetrics
	SetTransactionMetrics(value IURLSessionTaskTransactionMetrics)
	IsCellular() bool
	SetIsCellular(value bool)
	IsConstrained() bool
	SetIsConstrained(value bool)
	IsExpensive() bool
	SetIsExpensive(value bool)
	IsMultipath() bool
	SetIsMultipath(value bool)
	IsProxyConnection() bool
	SetIsProxyConnection(value bool)
	IsReusedConnection() bool
	SetIsReusedConnection(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for URLSessionTaskTransactionMetrics */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for URLSessionTaskTransactionMetrics */
// Alloc allocates a new instance without initialization.
func (uc _URLSessionTaskTransactionMetricsClass) Alloc() URLSessionTaskTransactionMetrics {
	rv := objc.Send[URLSessionTaskTransactionMetrics](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _URLSessionTaskTransactionMetricsClass) New() URLSessionTaskTransactionMetrics {
	rv := objc.Send[URLSessionTaskTransactionMetrics](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLSessionTaskTransactionMetrics) Init() URLSessionTaskTransactionMetrics {
	rv := objc.Send[URLSessionTaskTransactionMetrics](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLSessionTaskTransactionMetrics) Autorelease() URLSessionTaskTransactionMetrics {
	rv := objc.Send[URLSessionTaskTransactionMetrics](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSessionTaskTransactionMetrics creates a new URLSessionTaskTransactionMetrics instance.
func NewURLSessionTaskTransactionMetrics() URLSessionTaskTransactionMetrics {
	return getURLSessionTaskTransactionMetricsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for URLSessionTaskTransactionMetrics */
// An object that encapsualtes the performance metrics collected by the URL Loading System during the execution of a session task.
//
// Each object consists of a and property, corresponding to the request and response of the corresponding task. It also contains temporal metrics, starting with and ending with , as well as other characteristics like and .


// An object that encapsualtes the performance metrics collected by the URL Loading System during the execution of a session task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics
type URLSessionTaskTransactionMetrics struct {
	objectivec.Object
}

// URLSessionTaskTransactionMetricsFrom constructs a [URLSessionTaskTransactionMetrics] from an unsafe.Pointer.
//
// An object that encapsualtes the performance metrics collected by the URL Loading System during the execution of a session task.
func URLSessionTaskTransactionMetricsFrom(ptr unsafe.Pointer) URLSessionTaskTransactionMetrics {
	return URLSessionTaskTransactionMetrics{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for URLSessionTaskTransactionMetrics */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for URLSessionTaskTransactionMetrics */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for URLSessionTaskTransactionMetrics */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for URLSessionTaskTransactionMetrics */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for URLSessionTaskTransactionMetrics */

// The port number of the local interface for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionTaskTransactionMetrics/localPort
func (u_ URLSessionTaskTransactionMetrics) LocalPort() INumber {
	rv := objc.Send[Number](u_.ID, objc.Sel("localPort"))
	return rv
}/* debug [instance_properties/getter]: localPort */


// The TLS cipher suite the task negotiated with the endpoint for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionTaskTransactionMetrics/negotiatedTLSCipherSuite
func (u_ URLSessionTaskTransactionMetrics) NegotiatedTLSCipherSuite() INumber {
	rv := objc.Send[Number](u_.ID, objc.Sel("negotiatedTLSCipherSuite"))
	return rv
}/* debug [instance_properties/getter]: negotiatedTLSCipherSuite */


// The TLS protocol version the task negotiated with the endpoint for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionTaskTransactionMetrics/negotiatedTLSProtocolVersion
func (u_ URLSessionTaskTransactionMetrics) NegotiatedTLSProtocolVersion() INumber {
	rv := objc.Send[Number](u_.ID, objc.Sel("negotiatedTLSProtocolVersion"))
	return rv
}/* debug [instance_properties/getter]: negotiatedTLSProtocolVersion */


// The port number of the remote interface for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionTaskTransactionMetrics/remotePort
func (u_ URLSessionTaskTransactionMetrics) RemotePort() INumber {
	rv := objc.Send[Number](u_.ID, objc.Sel("remotePort"))
	return rv
}/* debug [instance_properties/getter]: remotePort */


// The time immediately after the task finished establishing the connection to the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/connectEndDate
func (u_ URLSessionTaskTransactionMetrics) ConnectEndDate() IDate {
	rv := objc.Send[Date](u_.ID, objc.Sel("connectEndDate"))
	return rv
}/* debug [instance_properties/getter]: connectEndDate */


// The time immediately before the task started establishing a TCP connection to the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/connectStartDate
func (u_ URLSessionTaskTransactionMetrics) ConnectStartDate() IDate {
	rv := objc.Send[Date](u_.ID, objc.Sel("connectStartDate"))
	return rv
}/* debug [instance_properties/getter]: connectStartDate */


// The size of the upload body data, file, or stream, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/countOfRequestBodyBytesBeforeEncoding
func (u_ URLSessionTaskTransactionMetrics) CountOfRequestBodyBytesBeforeEncoding() int64 {
	rv := objc.Send[int64](u_.ID, objc.Sel("countOfRequestBodyBytesBeforeEncoding"))
	return rv
}/* debug [instance_properties/getter]: countOfRequestBodyBytesBeforeEncoding */


// The number of bytes transferred for the request body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/countOfRequestBodyBytesSent
func (u_ URLSessionTaskTransactionMetrics) CountOfRequestBodyBytesSent() int64 {
	rv := objc.Send[int64](u_.ID, objc.Sel("countOfRequestBodyBytesSent"))
	return rv
}/* debug [instance_properties/getter]: countOfRequestBodyBytesSent */


// The number of bytes transferred for the request header.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/countOfRequestHeaderBytesSent
func (u_ URLSessionTaskTransactionMetrics) CountOfRequestHeaderBytesSent() int64 {
	rv := objc.Send[int64](u_.ID, objc.Sel("countOfRequestHeaderBytesSent"))
	return rv
}/* debug [instance_properties/getter]: countOfRequestHeaderBytesSent */


// The size of data delivered to your delegate or completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/countOfResponseBodyBytesAfterDecoding
func (u_ URLSessionTaskTransactionMetrics) CountOfResponseBodyBytesAfterDecoding() int64 {
	rv := objc.Send[int64](u_.ID, objc.Sel("countOfResponseBodyBytesAfterDecoding"))
	return rv
}/* debug [instance_properties/getter]: countOfResponseBodyBytesAfterDecoding */


// The number of bytes transferred for the response body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/countOfResponseBodyBytesReceived
func (u_ URLSessionTaskTransactionMetrics) CountOfResponseBodyBytesReceived() int64 {
	rv := objc.Send[int64](u_.ID, objc.Sel("countOfResponseBodyBytesReceived"))
	return rv
}/* debug [instance_properties/getter]: countOfResponseBodyBytesReceived */


// The number of bytes transferred for the response header.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/countOfResponseHeaderBytesReceived
func (u_ URLSessionTaskTransactionMetrics) CountOfResponseHeaderBytesReceived() int64 {
	rv := objc.Send[int64](u_.ID, objc.Sel("countOfResponseHeaderBytesReceived"))
	return rv
}/* debug [instance_properties/getter]: countOfResponseHeaderBytesReceived */


// The time after the name lookup was completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/domainLookupEndDate
func (u_ URLSessionTaskTransactionMetrics) DomainLookupEndDate() IDate {
	rv := objc.Send[Date](u_.ID, objc.Sel("domainLookupEndDate"))
	return rv
}/* debug [instance_properties/getter]: domainLookupEndDate */


// The time immediately before the task started the name lookup for the resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/domainLookupStartDate
func (u_ URLSessionTaskTransactionMetrics) DomainLookupStartDate() IDate {
	rv := objc.Send[Date](u_.ID, objc.Sel("domainLookupStartDate"))
	return rv
}/* debug [instance_properties/getter]: domainLookupStartDate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/domainResolutionProtocol
func (u_ URLSessionTaskTransactionMetrics) DomainResolutionProtocol() URLSessionTaskMetricsDomainResolutionProtocol {
	rv := objc.Send[URLSessionTaskMetricsDomainResolutionProtocol](u_.ID, objc.Sel("domainResolutionProtocol"))
	return rv
}/* debug [instance_properties/getter]: domainResolutionProtocol */


// The time when the task started fetching the resource, from the server or locally.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/fetchStartDate
func (u_ URLSessionTaskTransactionMetrics) FetchStartDate() IDate {
	rv := objc.Send[Date](u_.ID, objc.Sel("fetchStartDate"))
	return rv
}/* debug [instance_properties/getter]: fetchStartDate */


// A Boolean value that indicates whether the connection operates over a cellular interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/isCellular
func (u_ URLSessionTaskTransactionMetrics) Cellular() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("cellular"))
	return rv
}/* debug [instance_properties/getter]: cellular */


// A Boolean value that indicates whether the connection operates over an interface marked as constrained.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/isConstrained
func (u_ URLSessionTaskTransactionMetrics) Constrained() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("constrained"))
	return rv
}/* debug [instance_properties/getter]: constrained */


// A Boolean value that indicates whether the connection operates over an expensive interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/isExpensive
func (u_ URLSessionTaskTransactionMetrics) Expensive() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("expensive"))
	return rv
}/* debug [instance_properties/getter]: expensive */


// A Boolean value that indicates whether the connection uses a successfully negotiated multipath protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/isMultipath
func (u_ URLSessionTaskTransactionMetrics) Multipath() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("multipath"))
	return rv
}/* debug [instance_properties/getter]: multipath */


// A Boolean value that indicastes whether the task used a proxy connection to fetch the resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/isProxyConnection
func (u_ URLSessionTaskTransactionMetrics) ProxyConnection() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("proxyConnection"))
	return rv
}/* debug [instance_properties/getter]: proxyConnection */


// A Boolean value that indicates whether the task used a persistent connection to fetch the resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/isReusedConnection
func (u_ URLSessionTaskTransactionMetrics) ReusedConnection() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("reusedConnection"))
	return rv
}/* debug [instance_properties/getter]: reusedConnection */


// The IP address string of the local interface for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/localAddress
func (u_ URLSessionTaskTransactionMetrics) LocalAddress() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("localAddress"))
	return rv
}/* debug [instance_properties/getter]: localAddress */


// The network protocol used to fetch the resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/networkProtocolName
func (u_ URLSessionTaskTransactionMetrics) NetworkProtocolName() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("networkProtocolName"))
	return rv
}/* debug [instance_properties/getter]: networkProtocolName */


// The IP address string of the remote interface for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/remoteAddress
func (u_ URLSessionTaskTransactionMetrics) RemoteAddress() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("remoteAddress"))
	return rv
}/* debug [instance_properties/getter]: remoteAddress */


// The transaction request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/request
func (u_ URLSessionTaskTransactionMetrics) Request() IURLRequest {
	rv := objc.Send[URLRequest](u_.ID, objc.Sel("request"))
	return rv
}/* debug [instance_properties/getter]: request */


// The time immediately after the task finished requesting the resource, regardless of whether it was retrieved from the server or local resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/requestEndDate
func (u_ URLSessionTaskTransactionMetrics) RequestEndDate() IDate {
	rv := objc.Send[Date](u_.ID, objc.Sel("requestEndDate"))
	return rv
}/* debug [instance_properties/getter]: requestEndDate */


// The time immediately before the task started requesting the resource, regardless of whether it is retrieved from the server or local resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/requestStartDate
func (u_ URLSessionTaskTransactionMetrics) RequestStartDate() IDate {
	rv := objc.Send[Date](u_.ID, objc.Sel("requestStartDate"))
	return rv
}/* debug [instance_properties/getter]: requestStartDate */


// A value that indicates whether the resource was loaded, pushed, or retrieved from the local cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/resourceFetchType
func (u_ URLSessionTaskTransactionMetrics) ResourceFetchType() URLSessionTaskMetricsResourceFetchType {
	rv := objc.Send[URLSessionTaskMetricsResourceFetchType](u_.ID, objc.Sel("resourceFetchType"))
	return rv
}/* debug [instance_properties/getter]: resourceFetchType */


// The transaction response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/response
func (u_ URLSessionTaskTransactionMetrics) Response() IURLResponse {
	rv := objc.Send[URLResponse](u_.ID, objc.Sel("response"))
	return rv
}/* debug [instance_properties/getter]: response */


// The time immediately after the task received the last byte of the resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/responseEndDate
func (u_ URLSessionTaskTransactionMetrics) ResponseEndDate() IDate {
	rv := objc.Send[Date](u_.ID, objc.Sel("responseEndDate"))
	return rv
}/* debug [instance_properties/getter]: responseEndDate */


// The time immediately after the task received the first byte of the response from the server or from local resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/responseStartDate
func (u_ URLSessionTaskTransactionMetrics) ResponseStartDate() IDate {
	rv := objc.Send[Date](u_.ID, objc.Sel("responseStartDate"))
	return rv
}/* debug [instance_properties/getter]: responseStartDate */


// The time immediately after the security handshake completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/secureConnectionEndDate
func (u_ URLSessionTaskTransactionMetrics) SecureConnectionEndDate() IDate {
	rv := objc.Send[Date](u_.ID, objc.Sel("secureConnectionEndDate"))
	return rv
}/* debug [instance_properties/getter]: secureConnectionEndDate */


// The time immediately before the task started the TLS security handshake to secure the current connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionTaskTransactionMetrics/secureConnectionStartDate
func (u_ URLSessionTaskTransactionMetrics) SecureConnectionStartDate() IDate {
	rv := objc.Send[Date](u_.ID, objc.Sel("secureConnectionStartDate"))
	return rv
}/* debug [instance_properties/getter]: secureConnectionStartDate */


// The number of redirects that occurred during the execution of the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontaskmetrics/redirectcount
func (u_ URLSessionTaskTransactionMetrics) RedirectCount() int {
	rv := objc.Send[int](u_.ID, objc.Sel("redirectCount"))
	return rv
}/* debug [instance_properties/getter]: redirectCount */


// The number of redirects that occurred during the execution of the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontaskmetrics/redirectcount
func (u_ URLSessionTaskTransactionMetrics) SetRedirectCount(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRedirectCount:"), value)
}/* debug [instance_properties/setter]: redirectCount */


// The time interval between when a task is instantiated and when the task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontaskmetrics/taskinterval
func (u_ URLSessionTaskTransactionMetrics) TaskInterval() IDateInterval {
	rv := objc.Send[DateInterval](u_.ID, objc.Sel("taskInterval"))
	return rv
}/* debug [instance_properties/getter]: taskInterval */


// The time interval between when a task is instantiated and when the task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontaskmetrics/taskinterval
func (u_ URLSessionTaskTransactionMetrics) SetTaskInterval(value IDateInterval) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTaskInterval:"), value)
}/* debug [instance_properties/setter]: taskInterval */


// An array of metrics for each individual request-response transaction made during the execution of the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontaskmetrics/transactionmetrics
func (u_ URLSessionTaskTransactionMetrics) TransactionMetrics() IURLSessionTaskTransactionMetrics {
	rv := objc.Send[URLSessionTaskTransactionMetrics](u_.ID, objc.Sel("transactionMetrics"))
	return rv
}/* debug [instance_properties/getter]: transactionMetrics */


// An array of metrics for each individual request-response transaction made during the execution of the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontaskmetrics/transactionmetrics
func (u_ URLSessionTaskTransactionMetrics) SetTransactionMetrics(value IURLSessionTaskTransactionMetrics) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTransactionMetrics:"), value)
}/* debug [instance_properties/setter]: transactionMetrics */


// A Boolean value that indicates whether the connection operates over a cellular interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/iscellular
func (u_ URLSessionTaskTransactionMetrics) IsCellular() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isCellular"))
	return rv
}/* debug [instance_properties/getter]: isCellular */


// A Boolean value that indicates whether the connection operates over a cellular interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/iscellular
func (u_ URLSessionTaskTransactionMetrics) SetIsCellular(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsCellular:"), value)
}/* debug [instance_properties/setter]: isCellular */


// A Boolean value that indicates whether the connection operates over an interface marked as constrained.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/isconstrained
func (u_ URLSessionTaskTransactionMetrics) IsConstrained() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isConstrained"))
	return rv
}/* debug [instance_properties/getter]: isConstrained */


// A Boolean value that indicates whether the connection operates over an interface marked as constrained.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/isconstrained
func (u_ URLSessionTaskTransactionMetrics) SetIsConstrained(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsConstrained:"), value)
}/* debug [instance_properties/setter]: isConstrained */


// A Boolean value that indicates whether the connection operates over an expensive interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/isexpensive
func (u_ URLSessionTaskTransactionMetrics) IsExpensive() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isExpensive"))
	return rv
}/* debug [instance_properties/getter]: isExpensive */


// A Boolean value that indicates whether the connection operates over an expensive interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/isexpensive
func (u_ URLSessionTaskTransactionMetrics) SetIsExpensive(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsExpensive:"), value)
}/* debug [instance_properties/setter]: isExpensive */


// A Boolean value that indicates whether the connection uses a successfully negotiated multipath protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/ismultipath
func (u_ URLSessionTaskTransactionMetrics) IsMultipath() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isMultipath"))
	return rv
}/* debug [instance_properties/getter]: isMultipath */


// A Boolean value that indicates whether the connection uses a successfully negotiated multipath protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/ismultipath
func (u_ URLSessionTaskTransactionMetrics) SetIsMultipath(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsMultipath:"), value)
}/* debug [instance_properties/setter]: isMultipath */


// A Boolean value that indicastes whether the task used a proxy connection to fetch the resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/isproxyconnection
func (u_ URLSessionTaskTransactionMetrics) IsProxyConnection() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isProxyConnection"))
	return rv
}/* debug [instance_properties/getter]: isProxyConnection */


// A Boolean value that indicastes whether the task used a proxy connection to fetch the resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/isproxyconnection
func (u_ URLSessionTaskTransactionMetrics) SetIsProxyConnection(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsProxyConnection:"), value)
}/* debug [instance_properties/setter]: isProxyConnection */


// A Boolean value that indicates whether the task used a persistent connection to fetch the resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/isreusedconnection
func (u_ URLSessionTaskTransactionMetrics) IsReusedConnection() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isReusedConnection"))
	return rv
}/* debug [instance_properties/getter]: isReusedConnection */


// A Boolean value that indicates whether the task used a persistent connection to fetch the resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/isreusedconnection
func (u_ URLSessionTaskTransactionMetrics) SetIsReusedConnection(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsReusedConnection:"), value)
}/* debug [instance_properties/setter]: isReusedConnection */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSURLSessionTaskTransactionMetrics */



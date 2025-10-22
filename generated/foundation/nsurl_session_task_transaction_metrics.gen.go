// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [URLSessionTaskTransactionMetrics] class.
type IURLSessionTaskTransactionMetrics interface {
	objectivec.IObject
	RedirectCount() int
	SetRedirectCount(value int)
	TaskInterval() DateInterval
	SetTaskInterval(value IDateInterval)
	TransactionMetrics() NSURLSessionTaskTransactionMetrics
	SetTransactionMetrics(value IURLSessionTaskTransactionMetrics)
	ConnectEndDate() Date
	SetConnectEndDate(value IDate)
	ConnectStartDate() Date
	SetConnectStartDate(value IDate)
	CountOfRequestBodyBytesBeforeEncoding() unsafe.Pointer
	SetCountOfRequestBodyBytesBeforeEncoding(value unsafe.Pointer)
	CountOfRequestBodyBytesSent() unsafe.Pointer
	SetCountOfRequestBodyBytesSent(value unsafe.Pointer)
	CountOfRequestHeaderBytesSent() unsafe.Pointer
	SetCountOfRequestHeaderBytesSent(value unsafe.Pointer)
	CountOfResponseBodyBytesAfterDecoding() unsafe.Pointer
	SetCountOfResponseBodyBytesAfterDecoding(value unsafe.Pointer)
	CountOfResponseBodyBytesReceived() unsafe.Pointer
	SetCountOfResponseBodyBytesReceived(value unsafe.Pointer)
	CountOfResponseHeaderBytesReceived() unsafe.Pointer
	SetCountOfResponseHeaderBytesReceived(value unsafe.Pointer)
	DomainLookupEndDate() Date
	SetDomainLookupEndDate(value IDate)
	DomainLookupStartDate() Date
	SetDomainLookupStartDate(value IDate)
	DomainResolutionProtocol() unsafe.Pointer
	SetDomainResolutionProtocol(value unsafe.Pointer)
	FetchStartDate() Date
	SetFetchStartDate(value IDate)
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
	LocalAddress() string
	SetLocalAddress(value string)
	LocalPort() int
	SetLocalPort(value int)
	NegotiatedTLSCipherSuite() unsafe.Pointer
	SetNegotiatedTLSCipherSuite(value unsafe.Pointer)
	NegotiatedTLSProtocolVersion() unsafe.Pointer
	SetNegotiatedTLSProtocolVersion(value unsafe.Pointer)
	NetworkProtocolName() string
	SetNetworkProtocolName(value string)
	RemoteAddress() string
	SetRemoteAddress(value string)
	RemotePort() int
	SetRemotePort(value int)
	Request() URLRequest
	SetRequest(value IURLRequest)
	RequestEndDate() Date
	SetRequestEndDate(value IDate)
	RequestStartDate() Date
	SetRequestStartDate(value IDate)
	ResourceFetchType() unsafe.Pointer
	SetResourceFetchType(value unsafe.Pointer)
	Response() URLResponse
	SetResponse(value IURLResponse)
	ResponseEndDate() Date
	SetResponseEndDate(value IDate)
	ResponseStartDate() Date
	SetResponseStartDate(value IDate)
	SecureConnectionEndDate() Date
	SetSecureConnectionEndDate(value IDate)
	SecureConnectionStartDate() Date
	SetSecureConnectionStartDate(value IDate)
}

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

// Alloc allocates a new instance without initialization.
func (uc _URLSessionTaskTransactionMetricsClass) Alloc() URLSessionTaskTransactionMetrics {
	rv := objc.Send[URLSessionTaskTransactionMetrics](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The number of redirects that occurred during the execution of the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontaskmetrics/redirectcount
func (u_ URLSessionTaskTransactionMetrics) RedirectCount() int {
	rv := objc.Send[int](u_.ID, objc.Sel("redirectCount"))
	return rv
}


// The number of redirects that occurred during the execution of the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontaskmetrics/redirectcount
func (u_ URLSessionTaskTransactionMetrics) SetRedirectCount(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRedirectCount:"), value)
}


// The time interval between when a task is instantiated and when the task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontaskmetrics/taskinterval
func (u_ URLSessionTaskTransactionMetrics) TaskInterval() DateInterval {
	rv := objc.Send[DateInterval](u_.ID, objc.Sel("taskInterval"))
	return rv
}


// The time interval between when a task is instantiated and when the task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontaskmetrics/taskinterval
func (u_ URLSessionTaskTransactionMetrics) SetTaskInterval(value IDateInterval) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTaskInterval:"), value)
}


// An array of metrics for each individual request-response transaction made during the execution of the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontaskmetrics/transactionmetrics
func (u_ URLSessionTaskTransactionMetrics) TransactionMetrics() NSURLSessionTaskTransactionMetrics {
	rv := objc.Send[NSURLSessionTaskTransactionMetrics](u_.ID, objc.Sel("transactionMetrics"))
	return rv
}


// An array of metrics for each individual request-response transaction made during the execution of the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontaskmetrics/transactionmetrics
func (u_ URLSessionTaskTransactionMetrics) SetTransactionMetrics(value IURLSessionTaskTransactionMetrics) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTransactionMetrics:"), value)
}


// The time immediately after the task finished establishing the connection to the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/connectenddate
func (u_ URLSessionTaskTransactionMetrics) ConnectEndDate() Date {
	rv := objc.Send[Date](u_.ID, objc.Sel("connectEndDate"))
	return rv
}


// The time immediately after the task finished establishing the connection to the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/connectenddate
func (u_ URLSessionTaskTransactionMetrics) SetConnectEndDate(value IDate) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setConnectEndDate:"), value)
}


// The time immediately before the task started establishing a TCP connection to the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/connectstartdate
func (u_ URLSessionTaskTransactionMetrics) ConnectStartDate() Date {
	rv := objc.Send[Date](u_.ID, objc.Sel("connectStartDate"))
	return rv
}


// The time immediately before the task started establishing a TCP connection to the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/connectstartdate
func (u_ URLSessionTaskTransactionMetrics) SetConnectStartDate(value IDate) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setConnectStartDate:"), value)
}


// The size of the upload body data, file, or stream, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/countofrequestbodybytesbeforeencoding
func (u_ URLSessionTaskTransactionMetrics) CountOfRequestBodyBytesBeforeEncoding() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("countOfRequestBodyBytesBeforeEncoding"))
	return rv
}


// The size of the upload body data, file, or stream, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/countofrequestbodybytesbeforeencoding
func (u_ URLSessionTaskTransactionMetrics) SetCountOfRequestBodyBytesBeforeEncoding(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCountOfRequestBodyBytesBeforeEncoding:"), value)
}


// The number of bytes transferred for the request body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/countofrequestbodybytessent
func (u_ URLSessionTaskTransactionMetrics) CountOfRequestBodyBytesSent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("countOfRequestBodyBytesSent"))
	return rv
}


// The number of bytes transferred for the request body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/countofrequestbodybytessent
func (u_ URLSessionTaskTransactionMetrics) SetCountOfRequestBodyBytesSent(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCountOfRequestBodyBytesSent:"), value)
}


// The number of bytes transferred for the request header.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/countofrequestheaderbytessent
func (u_ URLSessionTaskTransactionMetrics) CountOfRequestHeaderBytesSent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("countOfRequestHeaderBytesSent"))
	return rv
}


// The number of bytes transferred for the request header.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/countofrequestheaderbytessent
func (u_ URLSessionTaskTransactionMetrics) SetCountOfRequestHeaderBytesSent(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCountOfRequestHeaderBytesSent:"), value)
}


// The size of data delivered to your delegate or completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/countofresponsebodybytesafterdecoding
func (u_ URLSessionTaskTransactionMetrics) CountOfResponseBodyBytesAfterDecoding() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("countOfResponseBodyBytesAfterDecoding"))
	return rv
}


// The size of data delivered to your delegate or completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/countofresponsebodybytesafterdecoding
func (u_ URLSessionTaskTransactionMetrics) SetCountOfResponseBodyBytesAfterDecoding(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCountOfResponseBodyBytesAfterDecoding:"), value)
}


// The number of bytes transferred for the response body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/countofresponsebodybytesreceived
func (u_ URLSessionTaskTransactionMetrics) CountOfResponseBodyBytesReceived() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("countOfResponseBodyBytesReceived"))
	return rv
}


// The number of bytes transferred for the response body.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/countofresponsebodybytesreceived
func (u_ URLSessionTaskTransactionMetrics) SetCountOfResponseBodyBytesReceived(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCountOfResponseBodyBytesReceived:"), value)
}


// The number of bytes transferred for the response header.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/countofresponseheaderbytesreceived
func (u_ URLSessionTaskTransactionMetrics) CountOfResponseHeaderBytesReceived() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("countOfResponseHeaderBytesReceived"))
	return rv
}


// The number of bytes transferred for the response header.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/countofresponseheaderbytesreceived
func (u_ URLSessionTaskTransactionMetrics) SetCountOfResponseHeaderBytesReceived(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCountOfResponseHeaderBytesReceived:"), value)
}


// The time after the name lookup was completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/domainlookupenddate
func (u_ URLSessionTaskTransactionMetrics) DomainLookupEndDate() Date {
	rv := objc.Send[Date](u_.ID, objc.Sel("domainLookupEndDate"))
	return rv
}


// The time after the name lookup was completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/domainlookupenddate
func (u_ URLSessionTaskTransactionMetrics) SetDomainLookupEndDate(value IDate) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDomainLookupEndDate:"), value)
}


// The time immediately before the task started the name lookup for the resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/domainlookupstartdate
func (u_ URLSessionTaskTransactionMetrics) DomainLookupStartDate() Date {
	rv := objc.Send[Date](u_.ID, objc.Sel("domainLookupStartDate"))
	return rv
}


// The time immediately before the task started the name lookup for the resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/domainlookupstartdate
func (u_ URLSessionTaskTransactionMetrics) SetDomainLookupStartDate(value IDate) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDomainLookupStartDate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/domainresolutionprotocol
func (u_ URLSessionTaskTransactionMetrics) DomainResolutionProtocol() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("domainResolutionProtocol"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/domainresolutionprotocol
func (u_ URLSessionTaskTransactionMetrics) SetDomainResolutionProtocol(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDomainResolutionProtocol:"), value)
}


// The time when the task started fetching the resource, from the server or locally.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/fetchstartdate
func (u_ URLSessionTaskTransactionMetrics) FetchStartDate() Date {
	rv := objc.Send[Date](u_.ID, objc.Sel("fetchStartDate"))
	return rv
}


// The time when the task started fetching the resource, from the server or locally.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/fetchstartdate
func (u_ URLSessionTaskTransactionMetrics) SetFetchStartDate(value IDate) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setFetchStartDate:"), value)
}


// A Boolean value that indicates whether the connection operates over a cellular interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/iscellular
func (u_ URLSessionTaskTransactionMetrics) IsCellular() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isCellular"))
	return rv
}


// A Boolean value that indicates whether the connection operates over a cellular interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/iscellular
func (u_ URLSessionTaskTransactionMetrics) SetIsCellular(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsCellular:"), value)
}


// A Boolean value that indicates whether the connection operates over an interface marked as constrained.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/isconstrained
func (u_ URLSessionTaskTransactionMetrics) IsConstrained() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isConstrained"))
	return rv
}


// A Boolean value that indicates whether the connection operates over an interface marked as constrained.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/isconstrained
func (u_ URLSessionTaskTransactionMetrics) SetIsConstrained(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsConstrained:"), value)
}


// A Boolean value that indicates whether the connection operates over an expensive interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/isexpensive
func (u_ URLSessionTaskTransactionMetrics) IsExpensive() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isExpensive"))
	return rv
}


// A Boolean value that indicates whether the connection operates over an expensive interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/isexpensive
func (u_ URLSessionTaskTransactionMetrics) SetIsExpensive(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsExpensive:"), value)
}


// A Boolean value that indicates whether the connection uses a successfully negotiated multipath protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/ismultipath
func (u_ URLSessionTaskTransactionMetrics) IsMultipath() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isMultipath"))
	return rv
}


// A Boolean value that indicates whether the connection uses a successfully negotiated multipath protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/ismultipath
func (u_ URLSessionTaskTransactionMetrics) SetIsMultipath(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsMultipath:"), value)
}


// A Boolean value that indicastes whether the task used a proxy connection to fetch the resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/isproxyconnection
func (u_ URLSessionTaskTransactionMetrics) IsProxyConnection() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isProxyConnection"))
	return rv
}


// A Boolean value that indicastes whether the task used a proxy connection to fetch the resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/isproxyconnection
func (u_ URLSessionTaskTransactionMetrics) SetIsProxyConnection(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsProxyConnection:"), value)
}


// A Boolean value that indicates whether the task used a persistent connection to fetch the resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/isreusedconnection
func (u_ URLSessionTaskTransactionMetrics) IsReusedConnection() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isReusedConnection"))
	return rv
}


// A Boolean value that indicates whether the task used a persistent connection to fetch the resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/isreusedconnection
func (u_ URLSessionTaskTransactionMetrics) SetIsReusedConnection(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsReusedConnection:"), value)
}


// The IP address string of the local interface for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/localaddress
func (u_ URLSessionTaskTransactionMetrics) LocalAddress() string {
	rv := objc.Send[string](u_.ID, objc.Sel("localAddress"))
	return rv
}


// The IP address string of the local interface for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/localaddress
func (u_ URLSessionTaskTransactionMetrics) SetLocalAddress(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setLocalAddress:"), objc.String(value))
}


// The port number of the local interface for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/localport
func (u_ URLSessionTaskTransactionMetrics) LocalPort() int {
	rv := objc.Send[int](u_.ID, objc.Sel("localPort"))
	return rv
}


// The port number of the local interface for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/localport
func (u_ URLSessionTaskTransactionMetrics) SetLocalPort(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setLocalPort:"), value)
}


// The TLS cipher suite the task negotiated with the endpoint for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/negotiatedtlsciphersuite
func (u_ URLSessionTaskTransactionMetrics) NegotiatedTLSCipherSuite() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("negotiatedTLSCipherSuite"))
	return rv
}


// The TLS cipher suite the task negotiated with the endpoint for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/negotiatedtlsciphersuite
func (u_ URLSessionTaskTransactionMetrics) SetNegotiatedTLSCipherSuite(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNegotiatedTLSCipherSuite:"), value)
}


// The TLS protocol version the task negotiated with the endpoint for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/negotiatedtlsprotocolversion
func (u_ URLSessionTaskTransactionMetrics) NegotiatedTLSProtocolVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("negotiatedTLSProtocolVersion"))
	return rv
}


// The TLS protocol version the task negotiated with the endpoint for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/negotiatedtlsprotocolversion
func (u_ URLSessionTaskTransactionMetrics) SetNegotiatedTLSProtocolVersion(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNegotiatedTLSProtocolVersion:"), value)
}


// The network protocol used to fetch the resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/networkprotocolname
func (u_ URLSessionTaskTransactionMetrics) NetworkProtocolName() string {
	rv := objc.Send[string](u_.ID, objc.Sel("networkProtocolName"))
	return rv
}


// The network protocol used to fetch the resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/networkprotocolname
func (u_ URLSessionTaskTransactionMetrics) SetNetworkProtocolName(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNetworkProtocolName:"), objc.String(value))
}


// The IP address string of the remote interface for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/remoteaddress
func (u_ URLSessionTaskTransactionMetrics) RemoteAddress() string {
	rv := objc.Send[string](u_.ID, objc.Sel("remoteAddress"))
	return rv
}


// The IP address string of the remote interface for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/remoteaddress
func (u_ URLSessionTaskTransactionMetrics) SetRemoteAddress(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRemoteAddress:"), objc.String(value))
}


// The port number of the remote interface for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/remoteport
func (u_ URLSessionTaskTransactionMetrics) RemotePort() int {
	rv := objc.Send[int](u_.ID, objc.Sel("remotePort"))
	return rv
}


// The port number of the remote interface for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/remoteport
func (u_ URLSessionTaskTransactionMetrics) SetRemotePort(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRemotePort:"), value)
}


// The transaction request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/request
func (u_ URLSessionTaskTransactionMetrics) Request() URLRequest {
	rv := objc.Send[URLRequest](u_.ID, objc.Sel("request"))
	return rv
}


// The transaction request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/request
func (u_ URLSessionTaskTransactionMetrics) SetRequest(value IURLRequest) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRequest:"), value)
}


// The time immediately after the task finished requesting the resource, regardless of whether it was retrieved from the server or local resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/requestenddate
func (u_ URLSessionTaskTransactionMetrics) RequestEndDate() Date {
	rv := objc.Send[Date](u_.ID, objc.Sel("requestEndDate"))
	return rv
}


// The time immediately after the task finished requesting the resource, regardless of whether it was retrieved from the server or local resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/requestenddate
func (u_ URLSessionTaskTransactionMetrics) SetRequestEndDate(value IDate) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRequestEndDate:"), value)
}


// The time immediately before the task started requesting the resource, regardless of whether it is retrieved from the server or local resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/requeststartdate
func (u_ URLSessionTaskTransactionMetrics) RequestStartDate() Date {
	rv := objc.Send[Date](u_.ID, objc.Sel("requestStartDate"))
	return rv
}


// The time immediately before the task started requesting the resource, regardless of whether it is retrieved from the server or local resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/requeststartdate
func (u_ URLSessionTaskTransactionMetrics) SetRequestStartDate(value IDate) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRequestStartDate:"), value)
}


// A value that indicates whether the resource was loaded, pushed, or retrieved from the local cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/resourcefetchtype
func (u_ URLSessionTaskTransactionMetrics) ResourceFetchType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("resourceFetchType"))
	return rv
}


// A value that indicates whether the resource was loaded, pushed, or retrieved from the local cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/resourcefetchtype
func (u_ URLSessionTaskTransactionMetrics) SetResourceFetchType(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setResourceFetchType:"), value)
}


// The transaction response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/response
func (u_ URLSessionTaskTransactionMetrics) Response() URLResponse {
	rv := objc.Send[URLResponse](u_.ID, objc.Sel("response"))
	return rv
}


// The transaction response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/response
func (u_ URLSessionTaskTransactionMetrics) SetResponse(value IURLResponse) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setResponse:"), value)
}


// The time immediately after the task received the last byte of the resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/responseenddate
func (u_ URLSessionTaskTransactionMetrics) ResponseEndDate() Date {
	rv := objc.Send[Date](u_.ID, objc.Sel("responseEndDate"))
	return rv
}


// The time immediately after the task received the last byte of the resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/responseenddate
func (u_ URLSessionTaskTransactionMetrics) SetResponseEndDate(value IDate) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setResponseEndDate:"), value)
}


// The time immediately after the task received the first byte of the response from the server or from local resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/responsestartdate
func (u_ URLSessionTaskTransactionMetrics) ResponseStartDate() Date {
	rv := objc.Send[Date](u_.ID, objc.Sel("responseStartDate"))
	return rv
}


// The time immediately after the task received the first byte of the response from the server or from local resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/responsestartdate
func (u_ URLSessionTaskTransactionMetrics) SetResponseStartDate(value IDate) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setResponseStartDate:"), value)
}


// The time immediately after the security handshake completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/secureconnectionenddate
func (u_ URLSessionTaskTransactionMetrics) SecureConnectionEndDate() Date {
	rv := objc.Send[Date](u_.ID, objc.Sel("secureConnectionEndDate"))
	return rv
}


// The time immediately after the security handshake completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/secureconnectionenddate
func (u_ URLSessionTaskTransactionMetrics) SetSecureConnectionEndDate(value IDate) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSecureConnectionEndDate:"), value)
}


// The time immediately before the task started the TLS security handshake to secure the current connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/secureconnectionstartdate
func (u_ URLSessionTaskTransactionMetrics) SecureConnectionStartDate() Date {
	rv := objc.Send[Date](u_.ID, objc.Sel("secureConnectionStartDate"))
	return rv
}


// The time immediately before the task started the TLS security handshake to secure the current connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessiontasktransactionmetrics/secureconnectionstartdate
func (u_ URLSessionTaskTransactionMetrics) SetSecureConnectionStartDate(value IDate) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSecureConnectionStartDate:"), value)
}




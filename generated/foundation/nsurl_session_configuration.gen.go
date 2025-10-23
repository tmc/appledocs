// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLSessionConfiguration] class.
var (
	URLSessionConfigurationClass     _URLSessionConfigurationClass
	URLSessionConfigurationClassOnce sync.Once
)

func getURLSessionConfigurationClass() _URLSessionConfigurationClass {
	URLSessionConfigurationClassOnce.Do(func() {
		URLSessionConfigurationClass = _URLSessionConfigurationClass{objc.GetClass("NSURLSessionConfiguration")}
	})
	return URLSessionConfigurationClass
}

type _URLSessionConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [URLSessionConfiguration] class.
type IURLSessionConfiguration interface {
	objectivec.IObject
	// properties:
	NetworkServiceType() URLRequestNetworkServiceType
	SetNetworkServiceType(value URLRequestNetworkServiceType)
	SharedContainerIdentifier() string /* primitive/slice/pointer. */
	SetSharedContainerIdentifier(value string /* primitive/slice/pointer. */)
	Configuration() IURLSessionConfiguration
	SetConfiguration(value IURLSessionConfiguration)
	AllowsCellularAccess() bool /* primitive/slice/pointer. */
	SetAllowsCellularAccess(value bool /* primitive/slice/pointer. */)
	AllowsConstrainedNetworkAccess() bool /* primitive/slice/pointer. */
	SetAllowsConstrainedNetworkAccess(value bool /* primitive/slice/pointer. */)
	AllowsExpensiveNetworkAccess() bool /* primitive/slice/pointer. */
	SetAllowsExpensiveNetworkAccess(value bool /* primitive/slice/pointer. */)
	AllowsUltraConstrainedNetworkAccess() bool /* primitive/slice/pointer. */
	SetAllowsUltraConstrainedNetworkAccess(value bool /* primitive/slice/pointer. */)
	ConnectionProxyDictionary() unsafe.Pointer
	SetConnectionProxyDictionary(value unsafe.Pointer)
	EnablesEarlyData() bool /* primitive/slice/pointer. */
	SetEnablesEarlyData(value bool /* primitive/slice/pointer. */)
	HttpAdditionalHeaders() unsafe.Pointer
	SetHttpAdditionalHeaders(value unsafe.Pointer)
	HttpCookieAcceptPolicy() unsafe.Pointer
	SetHttpCookieAcceptPolicy(value unsafe.Pointer)
	HttpCookieStorage() IHTTPCookieStorage
	SetHttpCookieStorage(value IHTTPCookieStorage)
	HttpMaximumConnectionsPerHost() int /* primitive/slice/pointer. */
	SetHttpMaximumConnectionsPerHost(value int /* primitive/slice/pointer. */)
	HttpShouldSetCookies() bool /* primitive/slice/pointer. */
	SetHttpShouldSetCookies(value bool /* primitive/slice/pointer. */)
	HttpShouldUsePipelining() bool /* primitive/slice/pointer. */
	SetHttpShouldUsePipelining(value bool /* primitive/slice/pointer. */)
	Identifier() string /* primitive/slice/pointer. */
	SetIdentifier(value string /* primitive/slice/pointer. */)
	IsDiscretionary() bool /* primitive/slice/pointer. */
	SetIsDiscretionary(value bool /* primitive/slice/pointer. */)
	MultipathServiceType() unsafe.Pointer
	SetMultipathServiceType(value unsafe.Pointer)
	ProtocolClasses() unsafe.Pointer
	SetProtocolClasses(value unsafe.Pointer)
	ProxyConfigurations() unsafe.Pointer
	SetProxyConfigurations(value unsafe.Pointer)
	RequestCachePolicy() unsafe.Pointer
	SetRequestCachePolicy(value unsafe.Pointer)
	RequiresDNSSECValidation() bool /* primitive/slice/pointer. */
	SetRequiresDNSSECValidation(value bool /* primitive/slice/pointer. */)
	SessionSendsLaunchEvents() bool /* primitive/slice/pointer. */
	SetSessionSendsLaunchEvents(value bool /* primitive/slice/pointer. */)
	ShouldUseExtendedBackgroundIdleMode() bool /* primitive/slice/pointer. */
	SetShouldUseExtendedBackgroundIdleMode(value bool /* primitive/slice/pointer. */)
	TimeoutIntervalForRequest() objc.IObject /* cross-framework: TimeInterval */
	SetTimeoutIntervalForRequest(value objc.IObject /* cross-framework: TimeInterval */)
	TimeoutIntervalForResource() objc.IObject /* cross-framework: TimeInterval */
	SetTimeoutIntervalForResource(value objc.IObject /* cross-framework: TimeInterval */)
	TlsMaximumSupportedProtocol() unsafe.Pointer
	SetTlsMaximumSupportedProtocol(value unsafe.Pointer)
	TlsMaximumSupportedProtocolVersion() unsafe.Pointer
	SetTlsMaximumSupportedProtocolVersion(value unsafe.Pointer)
	TlsMinimumSupportedProtocol() unsafe.Pointer
	SetTlsMinimumSupportedProtocol(value unsafe.Pointer)
	TlsMinimumSupportedProtocolVersion() unsafe.Pointer
	SetTlsMinimumSupportedProtocolVersion(value unsafe.Pointer)
	UrlCache() IURLCache
	SetUrlCache(value IURLCache)
	UrlCredentialStorage() IURLCredentialStorage
	SetUrlCredentialStorage(value IURLCredentialStorage)
	UsesClassicLoadingMode() bool /* primitive/slice/pointer. */
	SetUsesClassicLoadingMode(value bool /* primitive/slice/pointer. */)
	WaitsForConnectivity() bool /* primitive/slice/pointer. */
	SetWaitsForConnectivity(value bool /* primitive/slice/pointer. */)
	// methods:
}

// A configuration object that defines behavior and policies for a URL session.
//
// An object defines the behavior and policies to use when uploading and downloading data using an object. When uploading or downloading data, creating a configuration object is always the first step you must take. You use this object to configure the timeout values, caching policies, connection requirements, and other types of information that you intend to use with your object. It is important to configure your object appropriately before using it to initialize a session object. Session objects make a copy of the configuration settings you provide and use those settings to configure the session. Once configured, the session object ignores any changes you make to the object. If you need to modify your transfer policies, you must update the session configuration object and use it to create a new object. For more information about using configuration objects to create sessions, see .


// A configuration object that defines behavior and policies for a URL session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration
type URLSessionConfiguration struct {
	objectivec.Object
}

// URLSessionConfigurationFrom constructs a [URLSessionConfiguration] from an unsafe.Pointer.
//
// A configuration object that defines behavior and policies for a URL session.
func URLSessionConfigurationFrom(ptr unsafe.Pointer) URLSessionConfiguration {
	return URLSessionConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLSessionConfigurationClass) Alloc() URLSessionConfiguration {
	rv := objc.Send[URLSessionConfiguration](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLSessionConfigurationClass) New() URLSessionConfiguration {
	rv := objc.Send[URLSessionConfiguration](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLSessionConfiguration) Init() URLSessionConfiguration {
	rv := objc.Send[URLSessionConfiguration](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLSessionConfiguration) Autorelease() URLSessionConfiguration {
	rv := objc.Send[URLSessionConfiguration](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSessionConfiguration creates a new URLSessionConfiguration instance.
func NewURLSessionConfiguration() URLSessionConfiguration {
	return getURLSessionConfigurationClass().New()
}



// The type of network service for all tasks within network sessions to enable Cellular Network Slicing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/networkServiceType
func (u_ URLSessionConfiguration) NetworkServiceType() URLRequestNetworkServiceType {
	rv := objc.Send[URLRequestNetworkServiceType](u_.ID, objc.Sel("networkServiceType"))
	return rv
}


// The type of network service for all tasks within network sessions to enable Cellular Network Slicing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/networkServiceType
func (u_ URLSessionConfiguration) SetNetworkServiceType(value URLRequestNetworkServiceType) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNetworkServiceType:"), value)
}


// The identifier for the shared container into which files in background URL sessions should be downloaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/sharedContainerIdentifier
func (u_ URLSessionConfiguration) SharedContainerIdentifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](u_.ID, objc.Sel("sharedContainerIdentifier"))
	return rv
}


// The identifier for the shared container into which files in background URL sessions should be downloaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/sharedContainerIdentifier
func (u_ URLSessionConfiguration) SetSharedContainerIdentifier(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSharedContainerIdentifier:"), objc.String(value))
}


// A copy of the configuration object for this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsession/configuration
func (u_ URLSessionConfiguration) Configuration() IURLSessionConfiguration {
	rv := objc.Send[URLSessionConfiguration](u_.ID, objc.Sel("configuration"))
	return rv
}


// A copy of the configuration object for this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsession/configuration
func (u_ URLSessionConfiguration) SetConfiguration(value IURLSessionConfiguration) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setConfiguration:"), value)
}


// A Boolean value that determines whether connections should be made over a cellular network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/allowscellularaccess
func (u_ URLSessionConfiguration) AllowsCellularAccess() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsCellularAccess"))
	return rv
}


// A Boolean value that determines whether connections should be made over a cellular network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/allowscellularaccess
func (u_ URLSessionConfiguration) SetAllowsCellularAccess(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsCellularAccess:"), value)
}


// A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/allowsconstrainednetworkaccess
func (u_ URLSessionConfiguration) AllowsConstrainedNetworkAccess() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsConstrainedNetworkAccess"))
	return rv
}


// A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/allowsconstrainednetworkaccess
func (u_ URLSessionConfiguration) SetAllowsConstrainedNetworkAccess(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsConstrainedNetworkAccess:"), value)
}


// A Boolean value that indicates whether connections may use a network interface that the system considers expensive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/allowsexpensivenetworkaccess
func (u_ URLSessionConfiguration) AllowsExpensiveNetworkAccess() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsExpensiveNetworkAccess"))
	return rv
}


// A Boolean value that indicates whether connections may use a network interface that the system considers expensive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/allowsexpensivenetworkaccess
func (u_ URLSessionConfiguration) SetAllowsExpensiveNetworkAccess(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsExpensiveNetworkAccess:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/allowsultraconstrainednetworkaccess
func (u_ URLSessionConfiguration) AllowsUltraConstrainedNetworkAccess() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsUltraConstrainedNetworkAccess"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/allowsultraconstrainednetworkaccess
func (u_ URLSessionConfiguration) SetAllowsUltraConstrainedNetworkAccess(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsUltraConstrainedNetworkAccess:"), value)
}


// A dictionary containing information about the proxy to use within this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/connectionproxydictionary
func (u_ URLSessionConfiguration) ConnectionProxyDictionary() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("connectionProxyDictionary"))
	return rv
}


// A dictionary containing information about the proxy to use within this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/connectionproxydictionary
func (u_ URLSessionConfiguration) SetConnectionProxyDictionary(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setConnectionProxyDictionary:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/enablesearlydata
func (u_ URLSessionConfiguration) EnablesEarlyData() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("enablesEarlyData"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/enablesearlydata
func (u_ URLSessionConfiguration) SetEnablesEarlyData(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setEnablesEarlyData:"), value)
}


// A dictionary of additional headers to send with requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/httpadditionalheaders
func (u_ URLSessionConfiguration) HttpAdditionalHeaders() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("httpAdditionalHeaders"))
	return rv
}


// A dictionary of additional headers to send with requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/httpadditionalheaders
func (u_ URLSessionConfiguration) SetHttpAdditionalHeaders(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHttpAdditionalHeaders:"), value)
}


// A policy constant that determines when cookies should be accepted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/httpcookieacceptpolicy
func (u_ URLSessionConfiguration) HttpCookieAcceptPolicy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("httpCookieAcceptPolicy"))
	return rv
}


// A policy constant that determines when cookies should be accepted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/httpcookieacceptpolicy
func (u_ URLSessionConfiguration) SetHttpCookieAcceptPolicy(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHttpCookieAcceptPolicy:"), value)
}


// The cookie store for storing cookies within this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/httpcookiestorage
func (u_ URLSessionConfiguration) HttpCookieStorage() IHTTPCookieStorage {
	rv := objc.Send[HTTPCookieStorage](u_.ID, objc.Sel("httpCookieStorage"))
	return rv
}


// The cookie store for storing cookies within this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/httpcookiestorage
func (u_ URLSessionConfiguration) SetHttpCookieStorage(value IHTTPCookieStorage) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHttpCookieStorage:"), value)
}


// The maximum number of simultaneous connections to make to a given host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/httpmaximumconnectionsperhost
func (u_ URLSessionConfiguration) HttpMaximumConnectionsPerHost() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](u_.ID, objc.Sel("httpMaximumConnectionsPerHost"))
	return rv
}


// The maximum number of simultaneous connections to make to a given host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/httpmaximumconnectionsperhost
func (u_ URLSessionConfiguration) SetHttpMaximumConnectionsPerHost(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHttpMaximumConnectionsPerHost:"), value)
}


// A Boolean value that determines whether requests should contain cookies from the cookie store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/httpshouldsetcookies
func (u_ URLSessionConfiguration) HttpShouldSetCookies() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("httpShouldSetCookies"))
	return rv
}


// A Boolean value that determines whether requests should contain cookies from the cookie store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/httpshouldsetcookies
func (u_ URLSessionConfiguration) SetHttpShouldSetCookies(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHttpShouldSetCookies:"), value)
}


// A Boolean value that determines whether the session should use HTTP pipelining.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/httpshouldusepipelining
func (u_ URLSessionConfiguration) HttpShouldUsePipelining() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("httpShouldUsePipelining"))
	return rv
}


// A Boolean value that determines whether the session should use HTTP pipelining.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/httpshouldusepipelining
func (u_ URLSessionConfiguration) SetHttpShouldUsePipelining(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHttpShouldUsePipelining:"), value)
}


// The background session identifier of the configuration object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/identifier
func (u_ URLSessionConfiguration) Identifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](u_.ID, objc.Sel("identifier"))
	return rv
}


// The background session identifier of the configuration object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/identifier
func (u_ URLSessionConfiguration) SetIdentifier(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}


// A Boolean value that determines whether background tasks can be scheduled at the discretion of the system for optimal performance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/isdiscretionary
func (u_ URLSessionConfiguration) IsDiscretionary() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("isDiscretionary"))
	return rv
}


// A Boolean value that determines whether background tasks can be scheduled at the discretion of the system for optimal performance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/isdiscretionary
func (u_ URLSessionConfiguration) SetIsDiscretionary(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsDiscretionary:"), value)
}


// A service type that specifies the Multipath TCP connection policy for transmitting data over Wi-Fi and cellular interfaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/multipathservicetype-swift.property
func (u_ URLSessionConfiguration) MultipathServiceType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("multipathServiceType"))
	return rv
}


// A service type that specifies the Multipath TCP connection policy for transmitting data over Wi-Fi and cellular interfaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/multipathservicetype-swift.property
func (u_ URLSessionConfiguration) SetMultipathServiceType(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setMultipathServiceType:"), value)
}


// An array of extra protocol subclasses that handle requests in a session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/protocolclasses
func (u_ URLSessionConfiguration) ProtocolClasses() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("protocolClasses"))
	return rv
}


// An array of extra protocol subclasses that handle requests in a session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/protocolclasses
func (u_ URLSessionConfiguration) SetProtocolClasses(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setProtocolClasses:"), value)
}


// An array of proxy configuration objects containing information about the proxies to use within this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/proxyconfigurations
func (u_ URLSessionConfiguration) ProxyConfigurations() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("proxyConfigurations"))
	return rv
}


// An array of proxy configuration objects containing information about the proxies to use within this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/proxyconfigurations
func (u_ URLSessionConfiguration) SetProxyConfigurations(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setProxyConfigurations:"), value)
}


// A predefined constant that determines when to return a response from the cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/requestcachepolicy
func (u_ URLSessionConfiguration) RequestCachePolicy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("requestCachePolicy"))
	return rv
}


// A predefined constant that determines when to return a response from the cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/requestcachepolicy
func (u_ URLSessionConfiguration) SetRequestCachePolicy(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRequestCachePolicy:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/requiresdnssecvalidation
func (u_ URLSessionConfiguration) RequiresDNSSECValidation() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("requiresDNSSECValidation"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/requiresdnssecvalidation
func (u_ URLSessionConfiguration) SetRequiresDNSSECValidation(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRequiresDNSSECValidation:"), value)
}


// A Boolean value that indicates whether the app should be resumed or launched in the background when transfers finish.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/sessionsendslaunchevents
func (u_ URLSessionConfiguration) SessionSendsLaunchEvents() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("sessionSendsLaunchEvents"))
	return rv
}


// A Boolean value that indicates whether the app should be resumed or launched in the background when transfers finish.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/sessionsendslaunchevents
func (u_ URLSessionConfiguration) SetSessionSendsLaunchEvents(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSessionSendsLaunchEvents:"), value)
}


// A Boolean value that indicates whether TCP connections should be kept open when the app moves to the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/shoulduseextendedbackgroundidlemode
func (u_ URLSessionConfiguration) ShouldUseExtendedBackgroundIdleMode() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("shouldUseExtendedBackgroundIdleMode"))
	return rv
}


// A Boolean value that indicates whether TCP connections should be kept open when the app moves to the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/shoulduseextendedbackgroundidlemode
func (u_ URLSessionConfiguration) SetShouldUseExtendedBackgroundIdleMode(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setShouldUseExtendedBackgroundIdleMode:"), value)
}


// The timeout interval to use when waiting for additional data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/timeoutintervalforrequest
func (u_ URLSessionConfiguration) TimeoutIntervalForRequest() objc.IObject /* cross-framework: TimeInterval */ {
	rv := objc.Send[TimeInterval](u_.ID, objc.Sel("timeoutIntervalForRequest"))
	return rv
}


// The timeout interval to use when waiting for additional data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/timeoutintervalforrequest
func (u_ URLSessionConfiguration) SetTimeoutIntervalForRequest(value objc.IObject /* cross-framework: TimeInterval */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTimeoutIntervalForRequest:"), value)
}


// The maximum amount of time that a resource request should be allowed to take.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/timeoutintervalforresource
func (u_ URLSessionConfiguration) TimeoutIntervalForResource() objc.IObject /* cross-framework: TimeInterval */ {
	rv := objc.Send[TimeInterval](u_.ID, objc.Sel("timeoutIntervalForResource"))
	return rv
}


// The maximum amount of time that a resource request should be allowed to take.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/timeoutintervalforresource
func (u_ URLSessionConfiguration) SetTimeoutIntervalForResource(value objc.IObject /* cross-framework: TimeInterval */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTimeoutIntervalForResource:"), value)
}


// The maximum TLS protocol version that the client should request when making connections in this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/tlsmaximumsupportedprotocol
func (u_ URLSessionConfiguration) TlsMaximumSupportedProtocol() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("tlsMaximumSupportedProtocol"))
	return rv
}


// The maximum TLS protocol version that the client should request when making connections in this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/tlsmaximumsupportedprotocol
func (u_ URLSessionConfiguration) SetTlsMaximumSupportedProtocol(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTlsMaximumSupportedProtocol:"), value)
}


// The maximum TLS protocol version that the client should request when making connections in this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/tlsmaximumsupportedprotocolversion
func (u_ URLSessionConfiguration) TlsMaximumSupportedProtocolVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("tlsMaximumSupportedProtocolVersion"))
	return rv
}


// The maximum TLS protocol version that the client should request when making connections in this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/tlsmaximumsupportedprotocolversion
func (u_ URLSessionConfiguration) SetTlsMaximumSupportedProtocolVersion(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTlsMaximumSupportedProtocolVersion:"), value)
}


// The minimum TLS protocol to accept during protocol negotiation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/tlsminimumsupportedprotocol
func (u_ URLSessionConfiguration) TlsMinimumSupportedProtocol() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("tlsMinimumSupportedProtocol"))
	return rv
}


// The minimum TLS protocol to accept during protocol negotiation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/tlsminimumsupportedprotocol
func (u_ URLSessionConfiguration) SetTlsMinimumSupportedProtocol(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTlsMinimumSupportedProtocol:"), value)
}


// The minimum TLS protocol version that the client should accept when making connections in this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/tlsminimumsupportedprotocolversion
func (u_ URLSessionConfiguration) TlsMinimumSupportedProtocolVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("tlsMinimumSupportedProtocolVersion"))
	return rv
}


// The minimum TLS protocol version that the client should accept when making connections in this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/tlsminimumsupportedprotocolversion
func (u_ URLSessionConfiguration) SetTlsMinimumSupportedProtocolVersion(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTlsMinimumSupportedProtocolVersion:"), value)
}


// The URL cache for providing cached responses to requests within the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/urlcache
func (u_ URLSessionConfiguration) UrlCache() IURLCache {
	rv := objc.Send[URLCache](u_.ID, objc.Sel("urlCache"))
	return rv
}


// The URL cache for providing cached responses to requests within the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/urlcache
func (u_ URLSessionConfiguration) SetUrlCache(value IURLCache) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUrlCache:"), value)
}


// A credential store that provides credentials for authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/urlcredentialstorage
func (u_ URLSessionConfiguration) UrlCredentialStorage() IURLCredentialStorage {
	rv := objc.Send[URLCredentialStorage](u_.ID, objc.Sel("urlCredentialStorage"))
	return rv
}


// A credential store that provides credentials for authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/urlcredentialstorage
func (u_ URLSessionConfiguration) SetUrlCredentialStorage(value IURLCredentialStorage) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUrlCredentialStorage:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/usesclassicloadingmode
func (u_ URLSessionConfiguration) UsesClassicLoadingMode() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("usesClassicLoadingMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/usesclassicloadingmode
func (u_ URLSessionConfiguration) SetUsesClassicLoadingMode(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUsesClassicLoadingMode:"), value)
}


// A Boolean value that indicates whether the session should wait for connectivity to become available, or fail immediately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/waitsforconnectivity
func (u_ URLSessionConfiguration) WaitsForConnectivity() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("waitsForConnectivity"))
	return rv
}


// A Boolean value that indicates whether the session should wait for connectivity to become available, or fail immediately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/waitsforconnectivity
func (u_ URLSessionConfiguration) SetWaitsForConnectivity(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setWaitsForConnectivity:"), value)
}




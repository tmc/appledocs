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
	ProxyConfigurations() []Object
	SetProxyConfigurations(value []Object)
	AllowsCellularAccess() bool
	SetAllowsCellularAccess(value bool)
	AllowsConstrainedNetworkAccess() bool
	SetAllowsConstrainedNetworkAccess(value bool)
	AllowsExpensiveNetworkAccess() bool
	SetAllowsExpensiveNetworkAccess(value bool)
	ConnectionProxyDictionary() objc.ID
	SetConnectionProxyDictionary(value objc.ID)
	HTTPAdditionalHeaders() objc.ID
	SetHTTPAdditionalHeaders(value objc.ID)
	HTTPCookieAcceptPolicy() unsafe.Pointer
	SetHTTPCookieAcceptPolicy(value unsafe.Pointer)
	HTTPCookieStorage() NSHTTPCookieStorage
	SetHTTPCookieStorage(value IHTTPCookieStorage)
	HTTPMaximumConnectionsPerHost() int
	SetHTTPMaximumConnectionsPerHost(value int)
	HTTPShouldSetCookies() bool
	SetHTTPShouldSetCookies(value bool)
	HTTPShouldUsePipelining() bool
	SetHTTPShouldUsePipelining(value bool)
	Identifier() string
	Discretionary() bool
	SetDiscretionary(value bool)
	MultipathServiceType() URLSessionMultipathServiceType
	SetMultipathServiceType(value URLSessionMultipathServiceType)
	NetworkServiceType() URLRequestNetworkServiceType
	SetNetworkServiceType(value URLRequestNetworkServiceType)
	SessionSendsLaunchEvents() bool
	SetSessionSendsLaunchEvents(value bool)
	SharedContainerIdentifier() string
	SetSharedContainerIdentifier(value string)
	ShouldUseExtendedBackgroundIdleMode() bool
	SetShouldUseExtendedBackgroundIdleMode(value bool)
	TimeoutIntervalForRequest() TimeInterval
	SetTimeoutIntervalForRequest(value ITimeInterval)
	TimeoutIntervalForResource() TimeInterval
	SetTimeoutIntervalForResource(value ITimeInterval)
	TLSMaximumSupportedProtocol() unsafe.Pointer
	SetTLSMaximumSupportedProtocol(value unsafe.Pointer)
	UsesClassicLoadingMode() bool
	SetUsesClassicLoadingMode(value bool)
	WaitsForConnectivity() bool
	SetWaitsForConnectivity(value bool)
	Configuration() NSURLSessionConfiguration
	SetConfiguration(value IURLSessionConfiguration)
	AllowsUltraConstrainedNetworkAccess() bool
	SetAllowsUltraConstrainedNetworkAccess(value bool)
	EnablesEarlyData() bool
	SetEnablesEarlyData(value bool)
	IsDiscretionary() bool
	SetIsDiscretionary(value bool)
	ProtocolClasses() unsafe.Pointer
	SetProtocolClasses(value unsafe.Pointer)
	RequestCachePolicy() unsafe.Pointer
	SetRequestCachePolicy(value unsafe.Pointer)
	RequiresDNSSECValidation() bool
	SetRequiresDNSSECValidation(value bool)
	TlsMaximumSupportedProtocolVersion() unsafe.Pointer
	SetTlsMaximumSupportedProtocolVersion(value unsafe.Pointer)
	TlsMinimumSupportedProtocol() unsafe.Pointer
	SetTlsMinimumSupportedProtocol(value unsafe.Pointer)
	TlsMinimumSupportedProtocolVersion() unsafe.Pointer
	SetTlsMinimumSupportedProtocolVersion(value unsafe.Pointer)
	UrlCache() NSURLCache
	SetUrlCache(value IURLCache)
	UrlCredentialStorage() NSURLCredentialStorage
	SetUrlCredentialStorage(value IURLCredentialStorage)
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




// Creates a session configuration object that allows HTTP and HTTPS uploads or downloads to be performed in the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/background(withIdentifier:)
func (uc _URLSessionConfigurationClass) BackgroundSessionConfigurationWithIdentifier(identifier string) URLSessionConfiguration {
	rv := objc.Send[URLSessionConfiguration](objc.ID(uc.class), objc.Sel("backgroundSessionConfigurationWithIdentifier:"), objc.String(identifier))
	return rv
}


// Returns a session configuration object that allows HTTP and HTTPS uploads or downloads to be performed in the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/backgroundSessionConfiguration(_:)
func (uc _URLSessionConfigurationClass) BackgroundSessionConfiguration(identifier string) URLSessionConfiguration {
	rv := objc.Send[URLSessionConfiguration](objc.ID(uc.class), objc.Sel("backgroundSessionConfiguration:"), objc.String(identifier))
	return rv
}


// A default session configuration object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/default
func (uc _URLSessionConfigurationClass) DefaultSessionConfiguration() URLSessionConfiguration {
	rv := objc.Send[NSURLSessionConfiguration](objc.ID(uc.class), objc.Sel("defaultSessionConfiguration"))
	return rv
}

// A session configuration that uses no persistent storage for caches, cookies, or credentials.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/ephemeral
func (uc _URLSessionConfigurationClass) EphemeralSessionConfiguration() URLSessionConfiguration {
	rv := objc.Send[NSURLSessionConfiguration](objc.ID(uc.class), objc.Sel("ephemeralSessionConfiguration"))
	return rv
}

// An array of proxy configuration objects containing information about the proxies to use within this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionConfiguration/proxyConfigurations
func (u_ URLSessionConfiguration) ProxyConfigurations() []Object {
	rv := objc.Send[[]Object](u_.ID, objc.Sel("proxyConfigurations"))
	return rv
}


// An array of proxy configuration objects containing information about the proxies to use within this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionConfiguration/proxyConfigurations
func (u_ URLSessionConfiguration) SetProxyConfigurations(value []Object) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](u_.ID, objc.Sel("setProxyConfigurations:"), nsArray)
}


// A Boolean value that determines whether connections should be made over a cellular network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/allowsCellularAccess
func (u_ URLSessionConfiguration) AllowsCellularAccess() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsCellularAccess"))
	return rv
}


// A Boolean value that determines whether connections should be made over a cellular network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/allowsCellularAccess
func (u_ URLSessionConfiguration) SetAllowsCellularAccess(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsCellularAccess:"), value)
}


// A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/allowsConstrainedNetworkAccess
func (u_ URLSessionConfiguration) AllowsConstrainedNetworkAccess() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsConstrainedNetworkAccess"))
	return rv
}


// A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/allowsConstrainedNetworkAccess
func (u_ URLSessionConfiguration) SetAllowsConstrainedNetworkAccess(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsConstrainedNetworkAccess:"), value)
}


// A Boolean value that indicates whether connections may use a network interface that the system considers expensive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/allowsExpensiveNetworkAccess
func (u_ URLSessionConfiguration) AllowsExpensiveNetworkAccess() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsExpensiveNetworkAccess"))
	return rv
}


// A Boolean value that indicates whether connections may use a network interface that the system considers expensive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/allowsExpensiveNetworkAccess
func (u_ URLSessionConfiguration) SetAllowsExpensiveNetworkAccess(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsExpensiveNetworkAccess:"), value)
}


// A dictionary containing information about the proxy to use within this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/connectionProxyDictionary
func (u_ URLSessionConfiguration) ConnectionProxyDictionary() objc.ID {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("connectionProxyDictionary"))
	return rv
}


// A dictionary containing information about the proxy to use within this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/connectionProxyDictionary
func (u_ URLSessionConfiguration) SetConnectionProxyDictionary(value objc.ID) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setConnectionProxyDictionary:"), value)
}


// A default session configuration object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/default
func (u_ URLSessionConfiguration) DefaultSessionConfiguration() NSURLSessionConfiguration {
	rv := objc.Send[NSURLSessionConfiguration](u_.ID, objc.Sel("defaultSessionConfiguration"))
	return rv
}


// A session configuration that uses no persistent storage for caches, cookies, or credentials.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/ephemeral
func (u_ URLSessionConfiguration) EphemeralSessionConfiguration() NSURLSessionConfiguration {
	rv := objc.Send[NSURLSessionConfiguration](u_.ID, objc.Sel("ephemeralSessionConfiguration"))
	return rv
}


// A dictionary of additional headers to send with requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpAdditionalHeaders
func (u_ URLSessionConfiguration) HTTPAdditionalHeaders() objc.ID {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("HTTPAdditionalHeaders"))
	return rv
}


// A dictionary of additional headers to send with requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpAdditionalHeaders
func (u_ URLSessionConfiguration) SetHTTPAdditionalHeaders(value objc.ID) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHTTPAdditionalHeaders:"), value)
}


// A policy constant that determines when cookies should be accepted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpCookieAcceptPolicy
func (u_ URLSessionConfiguration) HTTPCookieAcceptPolicy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("HTTPCookieAcceptPolicy"))
	return rv
}


// A policy constant that determines when cookies should be accepted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpCookieAcceptPolicy
func (u_ URLSessionConfiguration) SetHTTPCookieAcceptPolicy(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHTTPCookieAcceptPolicy:"), value)
}


// The cookie store for storing cookies within this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpCookieStorage
func (u_ URLSessionConfiguration) HTTPCookieStorage() NSHTTPCookieStorage {
	rv := objc.Send[NSHTTPCookieStorage](u_.ID, objc.Sel("HTTPCookieStorage"))
	return rv
}


// The cookie store for storing cookies within this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpCookieStorage
func (u_ URLSessionConfiguration) SetHTTPCookieStorage(value IHTTPCookieStorage) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHTTPCookieStorage:"), value)
}


// The maximum number of simultaneous connections to make to a given host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpMaximumConnectionsPerHost
func (u_ URLSessionConfiguration) HTTPMaximumConnectionsPerHost() int {
	rv := objc.Send[int](u_.ID, objc.Sel("HTTPMaximumConnectionsPerHost"))
	return rv
}


// The maximum number of simultaneous connections to make to a given host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpMaximumConnectionsPerHost
func (u_ URLSessionConfiguration) SetHTTPMaximumConnectionsPerHost(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHTTPMaximumConnectionsPerHost:"), value)
}


// A Boolean value that determines whether requests should contain cookies from the cookie store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpShouldSetCookies
func (u_ URLSessionConfiguration) HTTPShouldSetCookies() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("HTTPShouldSetCookies"))
	return rv
}


// A Boolean value that determines whether requests should contain cookies from the cookie store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpShouldSetCookies
func (u_ URLSessionConfiguration) SetHTTPShouldSetCookies(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHTTPShouldSetCookies:"), value)
}


// A Boolean value that determines whether the session should use HTTP pipelining.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpShouldUsePipelining
func (u_ URLSessionConfiguration) HTTPShouldUsePipelining() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("HTTPShouldUsePipelining"))
	return rv
}


// A Boolean value that determines whether the session should use HTTP pipelining.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpShouldUsePipelining
func (u_ URLSessionConfiguration) SetHTTPShouldUsePipelining(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHTTPShouldUsePipelining:"), value)
}


// The background session identifier of the configuration object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/identifier
func (u_ URLSessionConfiguration) Identifier() string {
	rv := objc.Send[string](u_.ID, objc.Sel("identifier"))
	return rv
}


// A Boolean value that determines whether background tasks can be scheduled at the discretion of the system for optimal performance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/isDiscretionary
func (u_ URLSessionConfiguration) Discretionary() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("discretionary"))
	return rv
}


// A Boolean value that determines whether background tasks can be scheduled at the discretion of the system for optimal performance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/isDiscretionary
func (u_ URLSessionConfiguration) SetDiscretionary(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDiscretionary:"), value)
}


// A service type that specifies the Multipath TCP connection policy for transmitting data over Wi-Fi and cellular interfaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/multipathServiceType-swift.property
func (u_ URLSessionConfiguration) MultipathServiceType() URLSessionMultipathServiceType {
	rv := objc.Send[URLSessionMultipathServiceType](u_.ID, objc.Sel("multipathServiceType"))
	return rv
}


// A service type that specifies the Multipath TCP connection policy for transmitting data over Wi-Fi and cellular interfaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/multipathServiceType-swift.property
func (u_ URLSessionConfiguration) SetMultipathServiceType(value URLSessionMultipathServiceType) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setMultipathServiceType:"), value)
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


// A Boolean value that indicates whether the app should be resumed or launched in the background when transfers finish.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/sessionSendsLaunchEvents
func (u_ URLSessionConfiguration) SessionSendsLaunchEvents() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("sessionSendsLaunchEvents"))
	return rv
}


// A Boolean value that indicates whether the app should be resumed or launched in the background when transfers finish.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/sessionSendsLaunchEvents
func (u_ URLSessionConfiguration) SetSessionSendsLaunchEvents(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSessionSendsLaunchEvents:"), value)
}


// The identifier for the shared container into which files in background URL sessions should be downloaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/sharedContainerIdentifier
func (u_ URLSessionConfiguration) SharedContainerIdentifier() string {
	rv := objc.Send[string](u_.ID, objc.Sel("sharedContainerIdentifier"))
	return rv
}


// The identifier for the shared container into which files in background URL sessions should be downloaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/sharedContainerIdentifier
func (u_ URLSessionConfiguration) SetSharedContainerIdentifier(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSharedContainerIdentifier:"), objc.String(value))
}


// A Boolean value that indicates whether TCP connections should be kept open when the app moves to the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/shouldUseExtendedBackgroundIdleMode
func (u_ URLSessionConfiguration) ShouldUseExtendedBackgroundIdleMode() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("shouldUseExtendedBackgroundIdleMode"))
	return rv
}


// A Boolean value that indicates whether TCP connections should be kept open when the app moves to the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/shouldUseExtendedBackgroundIdleMode
func (u_ URLSessionConfiguration) SetShouldUseExtendedBackgroundIdleMode(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setShouldUseExtendedBackgroundIdleMode:"), value)
}


// The timeout interval to use when waiting for additional data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/timeoutIntervalForRequest
func (u_ URLSessionConfiguration) TimeoutIntervalForRequest() TimeInterval {
	rv := objc.Send[TimeInterval](u_.ID, objc.Sel("timeoutIntervalForRequest"))
	return rv
}


// The timeout interval to use when waiting for additional data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/timeoutIntervalForRequest
func (u_ URLSessionConfiguration) SetTimeoutIntervalForRequest(value ITimeInterval) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTimeoutIntervalForRequest:"), value)
}


// The maximum amount of time that a resource request should be allowed to take.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/timeoutIntervalForResource
func (u_ URLSessionConfiguration) TimeoutIntervalForResource() TimeInterval {
	rv := objc.Send[TimeInterval](u_.ID, objc.Sel("timeoutIntervalForResource"))
	return rv
}


// The maximum amount of time that a resource request should be allowed to take.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/timeoutIntervalForResource
func (u_ URLSessionConfiguration) SetTimeoutIntervalForResource(value ITimeInterval) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTimeoutIntervalForResource:"), value)
}


// The maximum TLS protocol version that the client should request when making connections in this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/tlsMaximumSupportedProtocol
func (u_ URLSessionConfiguration) TLSMaximumSupportedProtocol() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("TLSMaximumSupportedProtocol"))
	return rv
}


// The maximum TLS protocol version that the client should request when making connections in this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/tlsMaximumSupportedProtocol
func (u_ URLSessionConfiguration) SetTLSMaximumSupportedProtocol(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTLSMaximumSupportedProtocol:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/usesClassicLoadingMode
func (u_ URLSessionConfiguration) UsesClassicLoadingMode() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("usesClassicLoadingMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/usesClassicLoadingMode
func (u_ URLSessionConfiguration) SetUsesClassicLoadingMode(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUsesClassicLoadingMode:"), value)
}


// A Boolean value that indicates whether the session should wait for connectivity to become available, or fail immediately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/waitsForConnectivity
func (u_ URLSessionConfiguration) WaitsForConnectivity() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("waitsForConnectivity"))
	return rv
}


// A Boolean value that indicates whether the session should wait for connectivity to become available, or fail immediately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/waitsForConnectivity
func (u_ URLSessionConfiguration) SetWaitsForConnectivity(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setWaitsForConnectivity:"), value)
}


// A copy of the configuration object for this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsession/configuration
func (u_ URLSessionConfiguration) Configuration() NSURLSessionConfiguration {
	rv := objc.Send[NSURLSessionConfiguration](u_.ID, objc.Sel("configuration"))
	return rv
}


// A copy of the configuration object for this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsession/configuration
func (u_ URLSessionConfiguration) SetConfiguration(value IURLSessionConfiguration) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setConfiguration:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/allowsultraconstrainednetworkaccess
func (u_ URLSessionConfiguration) AllowsUltraConstrainedNetworkAccess() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsUltraConstrainedNetworkAccess"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/allowsultraconstrainednetworkaccess
func (u_ URLSessionConfiguration) SetAllowsUltraConstrainedNetworkAccess(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsUltraConstrainedNetworkAccess:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/enablesearlydata
func (u_ URLSessionConfiguration) EnablesEarlyData() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("enablesEarlyData"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/enablesearlydata
func (u_ URLSessionConfiguration) SetEnablesEarlyData(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setEnablesEarlyData:"), value)
}


// A Boolean value that determines whether background tasks can be scheduled at the discretion of the system for optimal performance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/isdiscretionary
func (u_ URLSessionConfiguration) IsDiscretionary() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isDiscretionary"))
	return rv
}


// A Boolean value that determines whether background tasks can be scheduled at the discretion of the system for optimal performance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/isdiscretionary
func (u_ URLSessionConfiguration) SetIsDiscretionary(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsDiscretionary:"), value)
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
func (u_ URLSessionConfiguration) RequiresDNSSECValidation() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("requiresDNSSECValidation"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/requiresdnssecvalidation
func (u_ URLSessionConfiguration) SetRequiresDNSSECValidation(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRequiresDNSSECValidation:"), value)
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
func (u_ URLSessionConfiguration) UrlCache() NSURLCache {
	rv := objc.Send[NSURLCache](u_.ID, objc.Sel("urlCache"))
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
func (u_ URLSessionConfiguration) UrlCredentialStorage() NSURLCredentialStorage {
	rv := objc.Send[NSURLCredentialStorage](u_.ID, objc.Sel("urlCredentialStorage"))
	return rv
}


// A credential store that provides credentials for authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/urlcredentialstorage
func (u_ URLSessionConfiguration) SetUrlCredentialStorage(value IURLCredentialStorage) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUrlCredentialStorage:"), value)
}



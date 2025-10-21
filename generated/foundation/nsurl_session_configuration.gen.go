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
}

// A configuration object that defines behavior and policies for a URL session.
//
// An object defines the behavior and policies to use when uploading and downloading data using an object. When uploading or downloading data, creating a configuration object is always the first step you must take. You use this object to configure the timeout values, caching policies, connection requirements, and other types of information that you intend to use with your object. It is important to configure your object appropriately before using it to initialize a session object. Session objects make a copy of the configuration settings you provide and use those settings to configure the session. Once configured, the session object ignores any changes you make to the object. If you need to modify your transfer policies, you must update the session configuration object and use it to create a new object. For more information about using configuration objects to create sessions, see .
//
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
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/background(withIdentifier:)
func (uc _URLSessionConfigurationClass) BackgroundSessionConfigurationWithIdentifier(identifier string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("backgroundSessionConfigurationWithIdentifier:"), objc.String(identifier))
	return rv
}

// Returns a session configuration object that allows HTTP and HTTPS uploads or downloads to be performed in the background.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/backgroundSessionConfiguration(_:)
func (uc _URLSessionConfigurationClass) BackgroundSessionConfiguration(identifier string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("backgroundSessionConfiguration:"), objc.String(identifier))
	return rv
}

// A default session configuration object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/default
func (uc _URLSessionConfigurationClass) DefaultSessionConfiguration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("defaultSessionConfiguration"))
	return rv
}

// A session configuration that uses no persistent storage for caches, cookies, or credentials.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/ephemeral
func (uc _URLSessionConfigurationClass) EphemeralSessionConfiguration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("ephemeralSessionConfiguration"))
	return rv
}

// An array of proxy configuration objects containing information about the proxies to use within this session.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionConfiguration/proxyConfigurations
func (u_ URLSessionConfiguration) ProxyConfigurations() []NSObject {
	rv := objc.Send[[]NSObject](u_.ID, objc.Sel("proxyConfigurations"))
	return rv
}

// SetProxyConfigurations sets the value of the proxyConfigurations property.
// An array of proxy configuration objects containing information about the proxies to use within this session.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLSessionConfiguration/proxyConfigurations
func (u_ URLSessionConfiguration) SetProxyConfigurations(value []NSObject) {
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
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/allowsCellularAccess
func (u_ URLSessionConfiguration) AllowsCellularAccess() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsCellularAccess"))
	return rv
}

// SetAllowsCellularAccess sets the value of the allowsCellularAccess property.
// A Boolean value that determines whether connections should be made over a cellular network.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/allowsCellularAccess
func (u_ URLSessionConfiguration) SetAllowsCellularAccess(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsCellularAccess:"), value)
}

// A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/allowsConstrainedNetworkAccess
func (u_ URLSessionConfiguration) AllowsConstrainedNetworkAccess() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsConstrainedNetworkAccess"))
	return rv
}

// SetAllowsConstrainedNetworkAccess sets the value of the allowsConstrainedNetworkAccess property.
// A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/allowsConstrainedNetworkAccess
func (u_ URLSessionConfiguration) SetAllowsConstrainedNetworkAccess(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsConstrainedNetworkAccess:"), value)
}

// A Boolean value that indicates whether connections may use a network interface that the system considers expensive.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/allowsExpensiveNetworkAccess
func (u_ URLSessionConfiguration) AllowsExpensiveNetworkAccess() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsExpensiveNetworkAccess"))
	return rv
}

// SetAllowsExpensiveNetworkAccess sets the value of the allowsExpensiveNetworkAccess property.
// A Boolean value that indicates whether connections may use a network interface that the system considers expensive.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/allowsExpensiveNetworkAccess
func (u_ URLSessionConfiguration) SetAllowsExpensiveNetworkAccess(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsExpensiveNetworkAccess:"), value)
}

// A dictionary containing information about the proxy to use within this session.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/connectionProxyDictionary
func (u_ URLSessionConfiguration) ConnectionProxyDictionary() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("connectionProxyDictionary"))
	return rv
}

// SetConnectionProxyDictionary sets the value of the connectionProxyDictionary property.
// A dictionary containing information about the proxy to use within this session.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/connectionProxyDictionary
func (u_ URLSessionConfiguration) SetConnectionProxyDictionary(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setConnectionProxyDictionary:"), value)
}

// A default session configuration object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/default
func (u_ URLSessionConfiguration) DefaultSessionConfiguration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("defaultSessionConfiguration"))
	return rv
}

// A session configuration that uses no persistent storage for caches, cookies, or credentials.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/ephemeral
func (u_ URLSessionConfiguration) EphemeralSessionConfiguration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("ephemeralSessionConfiguration"))
	return rv
}

// A dictionary of additional headers to send with requests.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpAdditionalHeaders
func (u_ URLSessionConfiguration) HTTPAdditionalHeaders() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("HTTPAdditionalHeaders"))
	return rv
}

// SetHTTPAdditionalHeaders sets the value of the HTTPAdditionalHeaders property.
// A dictionary of additional headers to send with requests.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpAdditionalHeaders
func (u_ URLSessionConfiguration) SetHTTPAdditionalHeaders(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHTTPAdditionalHeaders:"), value)
}

// A policy constant that determines when cookies should be accepted.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpCookieAcceptPolicy
func (u_ URLSessionConfiguration) HTTPCookieAcceptPolicy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("HTTPCookieAcceptPolicy"))
	return rv
}

// SetHTTPCookieAcceptPolicy sets the value of the HTTPCookieAcceptPolicy property.
// A policy constant that determines when cookies should be accepted.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpCookieAcceptPolicy
func (u_ URLSessionConfiguration) SetHTTPCookieAcceptPolicy(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHTTPCookieAcceptPolicy:"), value)
}

// The cookie store for storing cookies within this session.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpCookieStorage
func (u_ URLSessionConfiguration) HTTPCookieStorage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("HTTPCookieStorage"))
	return rv
}

// SetHTTPCookieStorage sets the value of the HTTPCookieStorage property.
// The cookie store for storing cookies within this session.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpCookieStorage
func (u_ URLSessionConfiguration) SetHTTPCookieStorage(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHTTPCookieStorage:"), value)
}

// The maximum number of simultaneous connections to make to a given host.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpMaximumConnectionsPerHost
func (u_ URLSessionConfiguration) HTTPMaximumConnectionsPerHost() int {
	rv := objc.Send[int](u_.ID, objc.Sel("HTTPMaximumConnectionsPerHost"))
	return rv
}

// SetHTTPMaximumConnectionsPerHost sets the value of the HTTPMaximumConnectionsPerHost property.
// The maximum number of simultaneous connections to make to a given host.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpMaximumConnectionsPerHost
func (u_ URLSessionConfiguration) SetHTTPMaximumConnectionsPerHost(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHTTPMaximumConnectionsPerHost:"), value)
}

// A Boolean value that determines whether requests should contain cookies from the cookie store.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpShouldSetCookies
func (u_ URLSessionConfiguration) HTTPShouldSetCookies() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("HTTPShouldSetCookies"))
	return rv
}

// SetHTTPShouldSetCookies sets the value of the HTTPShouldSetCookies property.
// A Boolean value that determines whether requests should contain cookies from the cookie store.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpShouldSetCookies
func (u_ URLSessionConfiguration) SetHTTPShouldSetCookies(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHTTPShouldSetCookies:"), value)
}

// A Boolean value that determines whether the session should use HTTP pipelining.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpShouldUsePipelining
func (u_ URLSessionConfiguration) HTTPShouldUsePipelining() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("HTTPShouldUsePipelining"))
	return rv
}

// SetHTTPShouldUsePipelining sets the value of the HTTPShouldUsePipelining property.
// A Boolean value that determines whether the session should use HTTP pipelining.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpShouldUsePipelining
func (u_ URLSessionConfiguration) SetHTTPShouldUsePipelining(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHTTPShouldUsePipelining:"), value)
}

// The background session identifier of the configuration object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/identifier
func (u_ URLSessionConfiguration) Identifier() string {
	rv := objc.Send[string](u_.ID, objc.Sel("identifier"))
	return rv
}

// A Boolean value that determines whether background tasks can be scheduled at the discretion of the system for optimal performance.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/isDiscretionary
func (u_ URLSessionConfiguration) Discretionary() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("discretionary"))
	return rv
}

// SetDiscretionary sets the value of the discretionary property.
// A Boolean value that determines whether background tasks can be scheduled at the discretion of the system for optimal performance.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/isDiscretionary
func (u_ URLSessionConfiguration) SetDiscretionary(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDiscretionary:"), value)
}

// A service type that specifies the Multipath TCP connection policy for transmitting data over Wi-Fi and cellular interfaces.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/multipathServiceType-swift.property
func (u_ URLSessionConfiguration) MultipathServiceType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("multipathServiceType"))
	return rv
}

// SetMultipathServiceType sets the value of the multipathServiceType property.
// A service type that specifies the Multipath TCP connection policy for transmitting data over Wi-Fi and cellular interfaces.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/multipathServiceType-swift.property
func (u_ URLSessionConfiguration) SetMultipathServiceType(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setMultipathServiceType:"), value)
}

// The type of network service for all tasks within network sessions to enable Cellular Network Slicing.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/networkServiceType
func (u_ URLSessionConfiguration) NetworkServiceType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("networkServiceType"))
	return rv
}

// SetNetworkServiceType sets the value of the networkServiceType property.
// The type of network service for all tasks within network sessions to enable Cellular Network Slicing.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/networkServiceType
func (u_ URLSessionConfiguration) SetNetworkServiceType(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNetworkServiceType:"), value)
}

// A Boolean value that indicates whether the app should be resumed or launched in the background when transfers finish.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/sessionSendsLaunchEvents
func (u_ URLSessionConfiguration) SessionSendsLaunchEvents() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("sessionSendsLaunchEvents"))
	return rv
}

// SetSessionSendsLaunchEvents sets the value of the sessionSendsLaunchEvents property.
// A Boolean value that indicates whether the app should be resumed or launched in the background when transfers finish.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/sessionSendsLaunchEvents
func (u_ URLSessionConfiguration) SetSessionSendsLaunchEvents(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSessionSendsLaunchEvents:"), value)
}

// The identifier for the shared container into which files in background URL sessions should be downloaded.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/sharedContainerIdentifier
func (u_ URLSessionConfiguration) SharedContainerIdentifier() string {
	rv := objc.Send[string](u_.ID, objc.Sel("sharedContainerIdentifier"))
	return rv
}

// SetSharedContainerIdentifier sets the value of the sharedContainerIdentifier property.
// The identifier for the shared container into which files in background URL sessions should be downloaded.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/sharedContainerIdentifier
func (u_ URLSessionConfiguration) SetSharedContainerIdentifier(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSharedContainerIdentifier:"), objc.String(value))
}

// A Boolean value that indicates whether TCP connections should be kept open when the app moves to the background.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/shouldUseExtendedBackgroundIdleMode
func (u_ URLSessionConfiguration) ShouldUseExtendedBackgroundIdleMode() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("shouldUseExtendedBackgroundIdleMode"))
	return rv
}

// SetShouldUseExtendedBackgroundIdleMode sets the value of the shouldUseExtendedBackgroundIdleMode property.
// A Boolean value that indicates whether TCP connections should be kept open when the app moves to the background.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/shouldUseExtendedBackgroundIdleMode
func (u_ URLSessionConfiguration) SetShouldUseExtendedBackgroundIdleMode(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setShouldUseExtendedBackgroundIdleMode:"), value)
}

// The timeout interval to use when waiting for additional data.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/timeoutIntervalForRequest
func (u_ URLSessionConfiguration) TimeoutIntervalForRequest() TimeInterval {
	rv := objc.Send[TimeInterval](u_.ID, objc.Sel("timeoutIntervalForRequest"))
	return rv
}

// SetTimeoutIntervalForRequest sets the value of the timeoutIntervalForRequest property.
// The timeout interval to use when waiting for additional data.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/timeoutIntervalForRequest
func (u_ URLSessionConfiguration) SetTimeoutIntervalForRequest(value TimeInterval) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTimeoutIntervalForRequest:"), value)
}

// The maximum amount of time that a resource request should be allowed to take.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/timeoutIntervalForResource
func (u_ URLSessionConfiguration) TimeoutIntervalForResource() TimeInterval {
	rv := objc.Send[TimeInterval](u_.ID, objc.Sel("timeoutIntervalForResource"))
	return rv
}

// SetTimeoutIntervalForResource sets the value of the timeoutIntervalForResource property.
// The maximum amount of time that a resource request should be allowed to take.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/timeoutIntervalForResource
func (u_ URLSessionConfiguration) SetTimeoutIntervalForResource(value TimeInterval) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTimeoutIntervalForResource:"), value)
}

// The maximum TLS protocol version that the client should request when making connections in this session.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/tlsMaximumSupportedProtocol
func (u_ URLSessionConfiguration) TLSMaximumSupportedProtocol() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("TLSMaximumSupportedProtocol"))
	return rv
}

// SetTLSMaximumSupportedProtocol sets the value of the TLSMaximumSupportedProtocol property.
// The maximum TLS protocol version that the client should request when making connections in this session.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/tlsMaximumSupportedProtocol
func (u_ URLSessionConfiguration) SetTLSMaximumSupportedProtocol(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTLSMaximumSupportedProtocol:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/usesClassicLoadingMode
func (u_ URLSessionConfiguration) UsesClassicLoadingMode() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("usesClassicLoadingMode"))
	return rv
}

// SetUsesClassicLoadingMode sets the value of the usesClassicLoadingMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/usesClassicLoadingMode
func (u_ URLSessionConfiguration) SetUsesClassicLoadingMode(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUsesClassicLoadingMode:"), value)
}

// A Boolean value that indicates whether the session should wait for connectivity to become available, or fail immediately.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/waitsForConnectivity
func (u_ URLSessionConfiguration) WaitsForConnectivity() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("waitsForConnectivity"))
	return rv
}

// SetWaitsForConnectivity sets the value of the waitsForConnectivity property.
// A Boolean value that indicates whether the session should wait for connectivity to become available, or fail immediately.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/waitsForConnectivity
func (u_ URLSessionConfiguration) SetWaitsForConnectivity(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setWaitsForConnectivity:"), value)
}

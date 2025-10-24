// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSURLSessionConfiguration */


/* debug [class_header]: Header for NSURLSessionConfiguration */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for URLSessionConfiguration */
// An interface definition for the [URLSessionConfiguration] class.
type IURLSessionConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for URLSessionConfiguration */
	// properties:
	AllowsCellularAccess() bool
	SetAllowsCellularAccess(value bool)
	AllowsConstrainedNetworkAccess() bool
	SetAllowsConstrainedNetworkAccess(value bool)
	AllowsExpensiveNetworkAccess() bool
	SetAllowsExpensiveNetworkAccess(value bool)
	HTTPAdditionalHeaders() IDictionary
	SetHTTPAdditionalHeaders(value IDictionary)
	Identifier() IString
	Discretionary() bool
	SetDiscretionary(value bool)
	NetworkServiceType() URLRequestNetworkServiceType
	SetNetworkServiceType(value URLRequestNetworkServiceType)
	SessionSendsLaunchEvents() bool
	SetSessionSendsLaunchEvents(value bool)
	SharedContainerIdentifier() IString
	SetSharedContainerIdentifier(value IString)
	ShouldUseExtendedBackgroundIdleMode() bool
	SetShouldUseExtendedBackgroundIdleMode(value bool)
	TimeoutIntervalForRequest() float64
	SetTimeoutIntervalForRequest(value float64)
	TimeoutIntervalForResource() float64
	SetTimeoutIntervalForResource(value float64)
	UsesClassicLoadingMode() bool
	SetUsesClassicLoadingMode(value bool)
	WaitsForConnectivity() bool
	SetWaitsForConnectivity(value bool)
	Configuration() IURLSessionConfiguration
	SetConfiguration(value IURLSessionConfiguration)
	AllowsUltraConstrainedNetworkAccess() bool
	SetAllowsUltraConstrainedNetworkAccess(value bool)
	ConnectionProxyDictionary() objectivec.IObject
	SetConnectionProxyDictionary(value objectivec.IObject)
	EnablesEarlyData() bool
	SetEnablesEarlyData(value bool)
	HttpCookieAcceptPolicy() objectivec.IObject
	SetHttpCookieAcceptPolicy(value objectivec.IObject)
	HttpCookieStorage() IHTTPCookieStorage
	SetHttpCookieStorage(value IHTTPCookieStorage)
	HttpMaximumConnectionsPerHost() int
	SetHttpMaximumConnectionsPerHost(value int)
	HttpShouldSetCookies() bool
	SetHttpShouldSetCookies(value bool)
	HttpShouldUsePipelining() bool
	SetHttpShouldUsePipelining(value bool)
	IsDiscretionary() bool
	SetIsDiscretionary(value bool)
	ProtocolClasses() objc.Class
	SetProtocolClasses(value objc.Class)
	ProxyConfigurations() objectivec.IObject
	SetProxyConfigurations(value objectivec.IObject)
	RequestCachePolicy() objectivec.IObject
	SetRequestCachePolicy(value objectivec.IObject)
	RequiresDNSSECValidation() bool
	SetRequiresDNSSECValidation(value bool)
	TlsMaximumSupportedProtocol() objectivec.IObject
	SetTlsMaximumSupportedProtocol(value objectivec.IObject)
	TlsMaximumSupportedProtocolVersion() objectivec.IObject
	SetTlsMaximumSupportedProtocolVersion(value objectivec.IObject)
	TlsMinimumSupportedProtocol() objectivec.IObject
	SetTlsMinimumSupportedProtocol(value objectivec.IObject)
	TlsMinimumSupportedProtocolVersion() objectivec.IObject
	SetTlsMinimumSupportedProtocolVersion(value objectivec.IObject)
	UrlCache() IURLCache
	SetUrlCache(value IURLCache)
	UrlCredentialStorage() IURLCredentialStorage
	SetUrlCredentialStorage(value IURLCredentialStorage)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for URLSessionConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for URLSessionConfiguration */
// Alloc allocates a new instance without initialization.
func (uc _URLSessionConfigurationClass) Alloc() URLSessionConfiguration {
	rv := objc.Send[URLSessionConfiguration](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for URLSessionConfiguration */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for URLSessionConfiguration */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for URLSessionConfiguration */

// Creates a session configuration object that allows HTTP and HTTPS uploads or downloads to be performed in the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/background(withIdentifier:)
func (uc _URLSessionConfigurationClass) BackgroundSessionConfigurationWithIdentifier(identifier IString) IURLSessionConfiguration {
	rv := objc.Send[URLSessionConfiguration](objc.ID(uc.class), objc.Sel("backgroundSessionConfigurationWithIdentifier:"), identifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=BackgroundSessionConfigurationWithIdentifier) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for URLSessionConfiguration */

// A default session configuration object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/default
func (uc _URLSessionConfigurationClass) DefaultSessionConfiguration() URLSessionConfiguration {
	rv := objc.Send[URLSessionConfiguration](objc.ID(uc.class), objc.Sel("defaultSessionConfiguration"))
	return rv
}/* debug [class_properties_class/property]: defaultSessionConfiguration */

// A session configuration that uses no persistent storage for caches, cookies, or credentials.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/ephemeral
func (uc _URLSessionConfigurationClass) EphemeralSessionConfiguration() URLSessionConfiguration {
	rv := objc.Send[URLSessionConfiguration](objc.ID(uc.class), objc.Sel("ephemeralSessionConfiguration"))
	return rv
}/* debug [class_properties_class/property]: ephemeralSessionConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for URLSessionConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for URLSessionConfiguration */

// A Boolean value that determines whether connections should be made over a cellular network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/allowsCellularAccess
func (u_ URLSessionConfiguration) AllowsCellularAccess() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsCellularAccess"))
	return rv
}/* debug [instance_properties/getter]: allowsCellularAccess */


// A Boolean value that determines whether connections should be made over a cellular network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/allowsCellularAccess
func (u_ URLSessionConfiguration) SetAllowsCellularAccess(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsCellularAccess:"), value)
}/* debug [instance_properties/setter]: allowsCellularAccess */


// A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/allowsConstrainedNetworkAccess
func (u_ URLSessionConfiguration) AllowsConstrainedNetworkAccess() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsConstrainedNetworkAccess"))
	return rv
}/* debug [instance_properties/getter]: allowsConstrainedNetworkAccess */


// A Boolean value that indicates whether connections may use the network when the user has specified Low Data Mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/allowsConstrainedNetworkAccess
func (u_ URLSessionConfiguration) SetAllowsConstrainedNetworkAccess(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsConstrainedNetworkAccess:"), value)
}/* debug [instance_properties/setter]: allowsConstrainedNetworkAccess */


// A Boolean value that indicates whether connections may use a network interface that the system considers expensive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/allowsExpensiveNetworkAccess
func (u_ URLSessionConfiguration) AllowsExpensiveNetworkAccess() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsExpensiveNetworkAccess"))
	return rv
}/* debug [instance_properties/getter]: allowsExpensiveNetworkAccess */


// A Boolean value that indicates whether connections may use a network interface that the system considers expensive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/allowsExpensiveNetworkAccess
func (u_ URLSessionConfiguration) SetAllowsExpensiveNetworkAccess(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsExpensiveNetworkAccess:"), value)
}/* debug [instance_properties/setter]: allowsExpensiveNetworkAccess */


// A default session configuration object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/default
func (u_ URLSessionConfiguration) DefaultSessionConfiguration() IURLSessionConfiguration {
	rv := objc.Send[URLSessionConfiguration](u_.ID, objc.Sel("defaultSessionConfiguration"))
	return rv
}/* debug [instance_properties/getter]: defaultSessionConfiguration */


// A session configuration that uses no persistent storage for caches, cookies, or credentials.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/ephemeral
func (u_ URLSessionConfiguration) EphemeralSessionConfiguration() IURLSessionConfiguration {
	rv := objc.Send[URLSessionConfiguration](u_.ID, objc.Sel("ephemeralSessionConfiguration"))
	return rv
}/* debug [instance_properties/getter]: ephemeralSessionConfiguration */


// A dictionary of additional headers to send with requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpAdditionalHeaders
func (u_ URLSessionConfiguration) HTTPAdditionalHeaders() IDictionary {
	rv := objc.Send[Dictionary](u_.ID, objc.Sel("HTTPAdditionalHeaders"))
	return rv
}/* debug [instance_properties/getter]: HTTPAdditionalHeaders */


// A dictionary of additional headers to send with requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/httpAdditionalHeaders
func (u_ URLSessionConfiguration) SetHTTPAdditionalHeaders(value IDictionary) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHTTPAdditionalHeaders:"), value)
}/* debug [instance_properties/setter]: HTTPAdditionalHeaders */


// The background session identifier of the configuration object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/identifier
func (u_ URLSessionConfiguration) Identifier() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// A Boolean value that determines whether background tasks can be scheduled at the discretion of the system for optimal performance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/isDiscretionary
func (u_ URLSessionConfiguration) Discretionary() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("discretionary"))
	return rv
}/* debug [instance_properties/getter]: discretionary */


// A Boolean value that determines whether background tasks can be scheduled at the discretion of the system for optimal performance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/isDiscretionary
func (u_ URLSessionConfiguration) SetDiscretionary(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDiscretionary:"), value)
}/* debug [instance_properties/setter]: discretionary */


// The type of network service for all tasks within network sessions to enable Cellular Network Slicing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/networkServiceType
func (u_ URLSessionConfiguration) NetworkServiceType() URLRequestNetworkServiceType {
	rv := objc.Send[URLRequestNetworkServiceType](u_.ID, objc.Sel("networkServiceType"))
	return rv
}/* debug [instance_properties/getter]: networkServiceType */


// The type of network service for all tasks within network sessions to enable Cellular Network Slicing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/networkServiceType
func (u_ URLSessionConfiguration) SetNetworkServiceType(value URLRequestNetworkServiceType) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNetworkServiceType:"), value)
}/* debug [instance_properties/setter]: networkServiceType */


// A Boolean value that indicates whether the app should be resumed or launched in the background when transfers finish.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/sessionSendsLaunchEvents
func (u_ URLSessionConfiguration) SessionSendsLaunchEvents() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("sessionSendsLaunchEvents"))
	return rv
}/* debug [instance_properties/getter]: sessionSendsLaunchEvents */


// A Boolean value that indicates whether the app should be resumed or launched in the background when transfers finish.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/sessionSendsLaunchEvents
func (u_ URLSessionConfiguration) SetSessionSendsLaunchEvents(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSessionSendsLaunchEvents:"), value)
}/* debug [instance_properties/setter]: sessionSendsLaunchEvents */


// The identifier for the shared container into which files in background URL sessions should be downloaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/sharedContainerIdentifier
func (u_ URLSessionConfiguration) SharedContainerIdentifier() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("sharedContainerIdentifier"))
	return rv
}/* debug [instance_properties/getter]: sharedContainerIdentifier */


// The identifier for the shared container into which files in background URL sessions should be downloaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/sharedContainerIdentifier
func (u_ URLSessionConfiguration) SetSharedContainerIdentifier(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSharedContainerIdentifier:"), value)
}/* debug [instance_properties/setter]: sharedContainerIdentifier */


// A Boolean value that indicates whether TCP connections should be kept open when the app moves to the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/shouldUseExtendedBackgroundIdleMode
func (u_ URLSessionConfiguration) ShouldUseExtendedBackgroundIdleMode() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("shouldUseExtendedBackgroundIdleMode"))
	return rv
}/* debug [instance_properties/getter]: shouldUseExtendedBackgroundIdleMode */


// A Boolean value that indicates whether TCP connections should be kept open when the app moves to the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/shouldUseExtendedBackgroundIdleMode
func (u_ URLSessionConfiguration) SetShouldUseExtendedBackgroundIdleMode(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setShouldUseExtendedBackgroundIdleMode:"), value)
}/* debug [instance_properties/setter]: shouldUseExtendedBackgroundIdleMode */


// The timeout interval to use when waiting for additional data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/timeoutIntervalForRequest
func (u_ URLSessionConfiguration) TimeoutIntervalForRequest() float64 {
	rv := objc.Send[float64](u_.ID, objc.Sel("timeoutIntervalForRequest"))
	return rv
}/* debug [instance_properties/getter]: timeoutIntervalForRequest */


// The timeout interval to use when waiting for additional data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/timeoutIntervalForRequest
func (u_ URLSessionConfiguration) SetTimeoutIntervalForRequest(value float64) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTimeoutIntervalForRequest:"), value)
}/* debug [instance_properties/setter]: timeoutIntervalForRequest */


// The maximum amount of time that a resource request should be allowed to take.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/timeoutIntervalForResource
func (u_ URLSessionConfiguration) TimeoutIntervalForResource() float64 {
	rv := objc.Send[float64](u_.ID, objc.Sel("timeoutIntervalForResource"))
	return rv
}/* debug [instance_properties/getter]: timeoutIntervalForResource */


// The maximum amount of time that a resource request should be allowed to take.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/timeoutIntervalForResource
func (u_ URLSessionConfiguration) SetTimeoutIntervalForResource(value float64) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTimeoutIntervalForResource:"), value)
}/* debug [instance_properties/setter]: timeoutIntervalForResource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/usesClassicLoadingMode
func (u_ URLSessionConfiguration) UsesClassicLoadingMode() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("usesClassicLoadingMode"))
	return rv
}/* debug [instance_properties/getter]: usesClassicLoadingMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/usesClassicLoadingMode
func (u_ URLSessionConfiguration) SetUsesClassicLoadingMode(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUsesClassicLoadingMode:"), value)
}/* debug [instance_properties/setter]: usesClassicLoadingMode */


// A Boolean value that indicates whether the session should wait for connectivity to become available, or fail immediately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/waitsForConnectivity
func (u_ URLSessionConfiguration) WaitsForConnectivity() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("waitsForConnectivity"))
	return rv
}/* debug [instance_properties/getter]: waitsForConnectivity */


// A Boolean value that indicates whether the session should wait for connectivity to become available, or fail immediately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/waitsForConnectivity
func (u_ URLSessionConfiguration) SetWaitsForConnectivity(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setWaitsForConnectivity:"), value)
}/* debug [instance_properties/setter]: waitsForConnectivity */


// A copy of the configuration object for this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsession/configuration
func (u_ URLSessionConfiguration) Configuration() IURLSessionConfiguration {
	rv := objc.Send[URLSessionConfiguration](u_.ID, objc.Sel("configuration"))
	return rv
}/* debug [instance_properties/getter]: configuration */


// A copy of the configuration object for this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsession/configuration
func (u_ URLSessionConfiguration) SetConfiguration(value IURLSessionConfiguration) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setConfiguration:"), value)
}/* debug [instance_properties/setter]: configuration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/allowsultraconstrainednetworkaccess
func (u_ URLSessionConfiguration) AllowsUltraConstrainedNetworkAccess() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("allowsUltraConstrainedNetworkAccess"))
	return rv
}/* debug [instance_properties/getter]: allowsUltraConstrainedNetworkAccess */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/allowsultraconstrainednetworkaccess
func (u_ URLSessionConfiguration) SetAllowsUltraConstrainedNetworkAccess(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setAllowsUltraConstrainedNetworkAccess:"), value)
}/* debug [instance_properties/setter]: allowsUltraConstrainedNetworkAccess */


// A dictionary containing information about the proxy to use within this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/connectionproxydictionary
func (u_ URLSessionConfiguration) ConnectionProxyDictionary() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](u_.ID, objc.Sel("connectionProxyDictionary"))
	return rv
}/* debug [instance_properties/getter]: connectionProxyDictionary */


// A dictionary containing information about the proxy to use within this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/connectionproxydictionary
func (u_ URLSessionConfiguration) SetConnectionProxyDictionary(value objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setConnectionProxyDictionary:"), value)
}/* debug [instance_properties/setter]: connectionProxyDictionary */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/enablesearlydata
func (u_ URLSessionConfiguration) EnablesEarlyData() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("enablesEarlyData"))
	return rv
}/* debug [instance_properties/getter]: enablesEarlyData */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/enablesearlydata
func (u_ URLSessionConfiguration) SetEnablesEarlyData(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setEnablesEarlyData:"), value)
}/* debug [instance_properties/setter]: enablesEarlyData */


// A policy constant that determines when cookies should be accepted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/httpcookieacceptpolicy
func (u_ URLSessionConfiguration) HttpCookieAcceptPolicy() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](u_.ID, objc.Sel("httpCookieAcceptPolicy"))
	return rv
}/* debug [instance_properties/getter]: httpCookieAcceptPolicy */


// A policy constant that determines when cookies should be accepted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/httpcookieacceptpolicy
func (u_ URLSessionConfiguration) SetHttpCookieAcceptPolicy(value objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHttpCookieAcceptPolicy:"), value)
}/* debug [instance_properties/setter]: httpCookieAcceptPolicy */


// The cookie store for storing cookies within this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/httpcookiestorage
func (u_ URLSessionConfiguration) HttpCookieStorage() IHTTPCookieStorage {
	rv := objc.Send[HTTPCookieStorage](u_.ID, objc.Sel("httpCookieStorage"))
	return rv
}/* debug [instance_properties/getter]: httpCookieStorage */


// The cookie store for storing cookies within this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/httpcookiestorage
func (u_ URLSessionConfiguration) SetHttpCookieStorage(value IHTTPCookieStorage) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHttpCookieStorage:"), value)
}/* debug [instance_properties/setter]: httpCookieStorage */


// The maximum number of simultaneous connections to make to a given host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/httpmaximumconnectionsperhost
func (u_ URLSessionConfiguration) HttpMaximumConnectionsPerHost() int {
	rv := objc.Send[int](u_.ID, objc.Sel("httpMaximumConnectionsPerHost"))
	return rv
}/* debug [instance_properties/getter]: httpMaximumConnectionsPerHost */


// The maximum number of simultaneous connections to make to a given host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/httpmaximumconnectionsperhost
func (u_ URLSessionConfiguration) SetHttpMaximumConnectionsPerHost(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHttpMaximumConnectionsPerHost:"), value)
}/* debug [instance_properties/setter]: httpMaximumConnectionsPerHost */


// A Boolean value that determines whether requests should contain cookies from the cookie store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/httpshouldsetcookies
func (u_ URLSessionConfiguration) HttpShouldSetCookies() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("httpShouldSetCookies"))
	return rv
}/* debug [instance_properties/getter]: httpShouldSetCookies */


// A Boolean value that determines whether requests should contain cookies from the cookie store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/httpshouldsetcookies
func (u_ URLSessionConfiguration) SetHttpShouldSetCookies(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHttpShouldSetCookies:"), value)
}/* debug [instance_properties/setter]: httpShouldSetCookies */


// A Boolean value that determines whether the session should use HTTP pipelining.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/httpshouldusepipelining
func (u_ URLSessionConfiguration) HttpShouldUsePipelining() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("httpShouldUsePipelining"))
	return rv
}/* debug [instance_properties/getter]: httpShouldUsePipelining */


// A Boolean value that determines whether the session should use HTTP pipelining.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/httpshouldusepipelining
func (u_ URLSessionConfiguration) SetHttpShouldUsePipelining(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHttpShouldUsePipelining:"), value)
}/* debug [instance_properties/setter]: httpShouldUsePipelining */


// A Boolean value that determines whether background tasks can be scheduled at the discretion of the system for optimal performance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/isdiscretionary
func (u_ URLSessionConfiguration) IsDiscretionary() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isDiscretionary"))
	return rv
}/* debug [instance_properties/getter]: isDiscretionary */


// A Boolean value that determines whether background tasks can be scheduled at the discretion of the system for optimal performance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/isdiscretionary
func (u_ URLSessionConfiguration) SetIsDiscretionary(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsDiscretionary:"), value)
}/* debug [instance_properties/setter]: isDiscretionary */


// An array of extra protocol subclasses that handle requests in a session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/protocolclasses
func (u_ URLSessionConfiguration) ProtocolClasses() objc.Class {
	rv := objc.Send[objc.Class](u_.ID, objc.Sel("protocolClasses"))
	return rv
}/* debug [instance_properties/getter]: protocolClasses */


// An array of extra protocol subclasses that handle requests in a session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/protocolclasses
func (u_ URLSessionConfiguration) SetProtocolClasses(value objc.Class) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setProtocolClasses:"), value)
}/* debug [instance_properties/setter]: protocolClasses */


// An array of proxy configuration objects containing information about the proxies to use within this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/proxyconfigurations
func (u_ URLSessionConfiguration) ProxyConfigurations() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](u_.ID, objc.Sel("proxyConfigurations"))
	return rv
}/* debug [instance_properties/getter]: proxyConfigurations */


// An array of proxy configuration objects containing information about the proxies to use within this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/proxyconfigurations
func (u_ URLSessionConfiguration) SetProxyConfigurations(value objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setProxyConfigurations:"), value)
}/* debug [instance_properties/setter]: proxyConfigurations */


// A predefined constant that determines when to return a response from the cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/requestcachepolicy
func (u_ URLSessionConfiguration) RequestCachePolicy() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](u_.ID, objc.Sel("requestCachePolicy"))
	return rv
}/* debug [instance_properties/getter]: requestCachePolicy */


// A predefined constant that determines when to return a response from the cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/requestcachepolicy
func (u_ URLSessionConfiguration) SetRequestCachePolicy(value objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRequestCachePolicy:"), value)
}/* debug [instance_properties/setter]: requestCachePolicy */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/requiresdnssecvalidation
func (u_ URLSessionConfiguration) RequiresDNSSECValidation() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("requiresDNSSECValidation"))
	return rv
}/* debug [instance_properties/getter]: requiresDNSSECValidation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/requiresdnssecvalidation
func (u_ URLSessionConfiguration) SetRequiresDNSSECValidation(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRequiresDNSSECValidation:"), value)
}/* debug [instance_properties/setter]: requiresDNSSECValidation */


// The maximum TLS protocol version that the client should request when making connections in this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/tlsmaximumsupportedprotocol
func (u_ URLSessionConfiguration) TlsMaximumSupportedProtocol() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](u_.ID, objc.Sel("tlsMaximumSupportedProtocol"))
	return rv
}/* debug [instance_properties/getter]: tlsMaximumSupportedProtocol */


// The maximum TLS protocol version that the client should request when making connections in this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/tlsmaximumsupportedprotocol
func (u_ URLSessionConfiguration) SetTlsMaximumSupportedProtocol(value objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTlsMaximumSupportedProtocol:"), value)
}/* debug [instance_properties/setter]: tlsMaximumSupportedProtocol */


// The maximum TLS protocol version that the client should request when making connections in this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/tlsmaximumsupportedprotocolversion
func (u_ URLSessionConfiguration) TlsMaximumSupportedProtocolVersion() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](u_.ID, objc.Sel("tlsMaximumSupportedProtocolVersion"))
	return rv
}/* debug [instance_properties/getter]: tlsMaximumSupportedProtocolVersion */


// The maximum TLS protocol version that the client should request when making connections in this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/tlsmaximumsupportedprotocolversion
func (u_ URLSessionConfiguration) SetTlsMaximumSupportedProtocolVersion(value objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTlsMaximumSupportedProtocolVersion:"), value)
}/* debug [instance_properties/setter]: tlsMaximumSupportedProtocolVersion */


// The minimum TLS protocol to accept during protocol negotiation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/tlsminimumsupportedprotocol
func (u_ URLSessionConfiguration) TlsMinimumSupportedProtocol() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](u_.ID, objc.Sel("tlsMinimumSupportedProtocol"))
	return rv
}/* debug [instance_properties/getter]: tlsMinimumSupportedProtocol */


// The minimum TLS protocol to accept during protocol negotiation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/tlsminimumsupportedprotocol
func (u_ URLSessionConfiguration) SetTlsMinimumSupportedProtocol(value objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTlsMinimumSupportedProtocol:"), value)
}/* debug [instance_properties/setter]: tlsMinimumSupportedProtocol */


// The minimum TLS protocol version that the client should accept when making connections in this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/tlsminimumsupportedprotocolversion
func (u_ URLSessionConfiguration) TlsMinimumSupportedProtocolVersion() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](u_.ID, objc.Sel("tlsMinimumSupportedProtocolVersion"))
	return rv
}/* debug [instance_properties/getter]: tlsMinimumSupportedProtocolVersion */


// The minimum TLS protocol version that the client should accept when making connections in this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/tlsminimumsupportedprotocolversion
func (u_ URLSessionConfiguration) SetTlsMinimumSupportedProtocolVersion(value objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTlsMinimumSupportedProtocolVersion:"), value)
}/* debug [instance_properties/setter]: tlsMinimumSupportedProtocolVersion */


// The URL cache for providing cached responses to requests within the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/urlcache
func (u_ URLSessionConfiguration) UrlCache() IURLCache {
	rv := objc.Send[URLCache](u_.ID, objc.Sel("urlCache"))
	return rv
}/* debug [instance_properties/getter]: urlCache */


// The URL cache for providing cached responses to requests within the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/urlcache
func (u_ URLSessionConfiguration) SetUrlCache(value IURLCache) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUrlCache:"), value)
}/* debug [instance_properties/setter]: urlCache */


// A credential store that provides credentials for authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/urlcredentialstorage
func (u_ URLSessionConfiguration) UrlCredentialStorage() IURLCredentialStorage {
	rv := objc.Send[URLCredentialStorage](u_.ID, objc.Sel("urlCredentialStorage"))
	return rv
}/* debug [instance_properties/getter]: urlCredentialStorage */


// A credential store that provides credentials for authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlsessionconfiguration/urlcredentialstorage
func (u_ URLSessionConfiguration) SetUrlCredentialStorage(value IURLCredentialStorage) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUrlCredentialStorage:"), value)
}/* debug [instance_properties/setter]: urlCredentialStorage */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSURLSessionConfiguration */



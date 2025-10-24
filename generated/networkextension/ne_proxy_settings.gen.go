// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEProxySettings */


/* debug [class_header]: Header for NEProxySettings */
// The class instance for the [NEProxySettings] class.
var (
	NEProxySettingsClass     _NEProxySettingsClass
	NEProxySettingsClassOnce sync.Once
)

func getNEProxySettingsClass() _NEProxySettingsClass {
	NEProxySettingsClassOnce.Do(func() {
		NEProxySettingsClass = _NEProxySettingsClass{objc.GetClass("NEProxySettings")}
	})
	return NEProxySettingsClass
}

type _NEProxySettingsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEProxySettings */
// An interface definition for the [NEProxySettings] class.
type INEProxySettings interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEProxySettings */
	// properties:
	AutoProxyConfigurationEnabled() bool
	SetAutoProxyConfigurationEnabled(value bool)
	ExceptionList() []string
	SetExceptionList(value []string)
	ExcludeSimpleHostnames() bool
	SetExcludeSimpleHostnames(value bool)
	HTTPEnabled() bool
	SetHTTPEnabled(value bool)
	HTTPSEnabled() bool
	SetHTTPSEnabled(value bool)
	HTTPServer() INEProxyServer
	SetHTTPServer(value INEProxyServer)
	HTTPSServer() INEProxyServer
	SetHTTPSServer(value INEProxyServer)
	MatchDomains() []string
	SetMatchDomains(value []string)
	ProxyAutoConfigurationJavaScript() objc.IObject /* cross-framework: NSString */
	SetProxyAutoConfigurationJavaScript(value objc.IObject /* cross-framework: NSString */)
	ProxyAutoConfigurationURL() objc.IObject /* cross-framework: NSURL */
	SetProxyAutoConfigurationURL(value objc.IObject /* cross-framework: NSURL */)
	DnsSettings() INEDNSSettings
	SetDnsSettings(value INEDNSSettings)
	ProxySettings() INEProxySettings
	SetProxySettings(value INEProxySettings)
	TunnelRemoteAddress() objc.IObject /* cross-framework: NSString */
	SetTunnelRemoteAddress(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEProxySettings */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEProxySettings */
// Alloc allocates a new instance without initialization.
func (nc _NEProxySettingsClass) Alloc() NEProxySettings {
	rv := objc.Send[NEProxySettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEProxySettingsClass) New() NEProxySettings {
	rv := objc.Send[NEProxySettings](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEProxySettings) Init() NEProxySettings {
	rv := objc.Send[NEProxySettings](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEProxySettings) Autorelease() NEProxySettings {
	rv := objc.Send[NEProxySettings](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEProxySettings creates a new NEProxySettings instance.
func NewNEProxySettings() NEProxySettings {
	return getNEProxySettingsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEProxySettings */
// contains HTTP proxy settings.
//
// is used in the context of a VPN configuration to specify the proxy that should be used for network traffic when the VPN is active. Instances of this class are thread safe.


// contains HTTP proxy settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings
type NEProxySettings struct {
	objectivec.Object
}

// NEProxySettingsFrom constructs a [NEProxySettings] from an unsafe.Pointer.
//
// contains HTTP proxy settings.
func NEProxySettingsFrom(ptr unsafe.Pointer) NEProxySettings {
	return NEProxySettings{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEProxySettings *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEProxySettings */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEProxySettings */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEProxySettings */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEProxySettings */

// A Boolean indicating if proxy auto-configuration is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/autoProxyConfigurationEnabled
func (n_ NEProxySettings) AutoProxyConfigurationEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("autoProxyConfigurationEnabled"))
	return rv
}/* debug [instance_properties/getter]: autoProxyConfigurationEnabled */


// A Boolean indicating if proxy auto-configuration is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/autoProxyConfigurationEnabled
func (n_ NEProxySettings) SetAutoProxyConfigurationEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAutoProxyConfigurationEnabled:"), value)
}/* debug [instance_properties/setter]: autoProxyConfigurationEnabled */


// An array of domain name patterns. If the destination host name of an HTTP connection matches one of these patterns then the proxy settings will not be used for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/exceptionList
func (n_ NEProxySettings) ExceptionList() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("exceptionList"))
	return rv
}/* debug [instance_properties/getter]: exceptionList */


// An array of domain name patterns. If the destination host name of an HTTP connection matches one of these patterns then the proxy settings will not be used for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/exceptionList
func (n_ NEProxySettings) SetExceptionList(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](n_.ID, objc.Sel("setExceptionList:"), nsArray)
}/* debug [instance_properties/setter]: exceptionList */


// A Boolean indicating if HTTP requests using single-label host names should be excluded from using the proxy settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/excludeSimpleHostnames
func (n_ NEProxySettings) ExcludeSimpleHostnames() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("excludeSimpleHostnames"))
	return rv
}/* debug [instance_properties/getter]: excludeSimpleHostnames */


// A Boolean indicating if HTTP requests using single-label host names should be excluded from using the proxy settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/excludeSimpleHostnames
func (n_ NEProxySettings) SetExcludeSimpleHostnames(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setExcludeSimpleHostnames:"), value)
}/* debug [instance_properties/setter]: excludeSimpleHostnames */


// A Boolean indicating if a static HTTP proxy will be used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/httpEnabled
func (n_ NEProxySettings) HTTPEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("HTTPEnabled"))
	return rv
}/* debug [instance_properties/getter]: HTTPEnabled */


// A Boolean indicating if a static HTTP proxy will be used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/httpEnabled
func (n_ NEProxySettings) SetHTTPEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHTTPEnabled:"), value)
}/* debug [instance_properties/setter]: HTTPEnabled */


// A Boolean indicating if a static HTTPS proxy will be used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/httpsEnabled
func (n_ NEProxySettings) HTTPSEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("HTTPSEnabled"))
	return rv
}/* debug [instance_properties/getter]: HTTPSEnabled */


// A Boolean indicating if a static HTTPS proxy will be used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/httpsEnabled
func (n_ NEProxySettings) SetHTTPSEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHTTPSEnabled:"), value)
}/* debug [instance_properties/setter]: HTTPSEnabled */


// An object containing the static HTTP proxy server settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/httpServer
func (n_ NEProxySettings) HTTPServer() INEProxyServer {
	rv := objc.Send[NEProxyServer](n_.ID, objc.Sel("HTTPServer"))
	return rv
}/* debug [instance_properties/getter]: HTTPServer */


// An object containing the static HTTP proxy server settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/httpServer
func (n_ NEProxySettings) SetHTTPServer(value INEProxyServer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHTTPServer:"), value)
}/* debug [instance_properties/setter]: HTTPServer */


// An object containing the static HTTPS proxy server settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/httpsServer
func (n_ NEProxySettings) HTTPSServer() INEProxyServer {
	rv := objc.Send[NEProxyServer](n_.ID, objc.Sel("HTTPSServer"))
	return rv
}/* debug [instance_properties/getter]: HTTPSServer */


// An object containing the static HTTPS proxy server settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/httpsServer
func (n_ NEProxySettings) SetHTTPSServer(value INEProxyServer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHTTPSServer:"), value)
}/* debug [instance_properties/setter]: HTTPSServer */


// An array of domain strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/matchDomains
func (n_ NEProxySettings) MatchDomains() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("matchDomains"))
	return rv
}/* debug [instance_properties/getter]: matchDomains */


// An array of domain strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/matchDomains
func (n_ NEProxySettings) SetMatchDomains(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchDomains:"), nsArray)
}/* debug [instance_properties/setter]: matchDomains */


// A string containing the Proxy Auto Configuration (PAC) JavaScript source code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/proxyAutoConfigurationJavaScript
func (n_ NEProxySettings) ProxyAutoConfigurationJavaScript() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("proxyAutoConfigurationJavaScript"))
	return rv
}/* debug [instance_properties/getter]: proxyAutoConfigurationJavaScript */


// A string containing the Proxy Auto Configuration (PAC) JavaScript source code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/proxyAutoConfigurationJavaScript
func (n_ NEProxySettings) SetProxyAutoConfigurationJavaScript(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProxyAutoConfigurationJavaScript:"), value)
}/* debug [instance_properties/setter]: proxyAutoConfigurationJavaScript */


// A URL specifying the location from where the Proxy Auto Configuration (PAC) script should be downloaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/proxyAutoConfigurationURL
func (n_ NEProxySettings) ProxyAutoConfigurationURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](n_.ID, objc.Sel("proxyAutoConfigurationURL"))
	return rv
}/* debug [instance_properties/getter]: proxyAutoConfigurationURL */


// A URL specifying the location from where the Proxy Auto Configuration (PAC) script should be downloaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/proxyAutoConfigurationURL
func (n_ NEProxySettings) SetProxyAutoConfigurationURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProxyAutoConfigurationURL:"), value)
}/* debug [instance_properties/setter]: proxyAutoConfigurationURL */


// The tunnel DNS settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelnetworksettings/dnssettings
func (n_ NEProxySettings) DnsSettings() INEDNSSettings {
	rv := objc.Send[NEDNSSettings](n_.ID, objc.Sel("dnsSettings"))
	return rv
}/* debug [instance_properties/getter]: dnsSettings */


// The tunnel DNS settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelnetworksettings/dnssettings
func (n_ NEProxySettings) SetDnsSettings(value INEDNSSettings) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDnsSettings:"), value)
}/* debug [instance_properties/setter]: dnsSettings */


// The tunnel HTTP proxy settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelnetworksettings/proxysettings
func (n_ NEProxySettings) ProxySettings() INEProxySettings {
	rv := objc.Send[NEProxySettings](n_.ID, objc.Sel("proxySettings"))
	return rv
}/* debug [instance_properties/getter]: proxySettings */


// The tunnel HTTP proxy settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelnetworksettings/proxysettings
func (n_ NEProxySettings) SetProxySettings(value INEProxySettings) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProxySettings:"), value)
}/* debug [instance_properties/setter]: proxySettings */


// The IP address of the tunnel server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelnetworksettings/tunnelremoteaddress
func (n_ NEProxySettings) TunnelRemoteAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("tunnelRemoteAddress"))
	return rv
}/* debug [instance_properties/getter]: tunnelRemoteAddress */


// The IP address of the tunnel server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelnetworksettings/tunnelremoteaddress
func (n_ NEProxySettings) SetTunnelRemoteAddress(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTunnelRemoteAddress:"), value)
}/* debug [instance_properties/setter]: tunnelRemoteAddress */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEProxySettings */




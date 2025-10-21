// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [NEProxySettings] class.
type INEProxySettings interface {
	objectivec.IObject
}

// contains HTTP proxy settings.
//
// is used in the context of a VPN configuration to specify the proxy that should be used for network traffic when the VPN is active. Instances of this class are thread safe.
//
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

// Alloc allocates a new instance without initialization.
func (nc _NEProxySettingsClass) Alloc() NEProxySettings {
	rv := objc.Send[NEProxySettings](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// A Boolean indicating if proxy auto-configuration is enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/autoproxyconfigurationenabled
func (n_ NEProxySettings) AutoProxyConfigurationEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("autoProxyConfigurationEnabled"))
	return rv
}


// SetAutoProxyConfigurationEnabled sets the value of the autoProxyConfigurationEnabled property.
// A Boolean indicating if proxy auto-configuration is enabled.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/autoproxyconfigurationenabled
func (n_ NEProxySettings) SetAutoProxyConfigurationEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAutoProxyConfigurationEnabled:"), value)
}

// An array of domain name patterns. If the destination host name of an HTTP connection matches one of these patterns then the proxy settings will not be used for the connection.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/exceptionlist
func (n_ NEProxySettings) ExceptionList() string {
	rv := objc.Send[string](n_.ID, objc.Sel("exceptionList"))
	return rv
}


// SetExceptionList sets the value of the exceptionList property.
// An array of domain name patterns. If the destination host name of an HTTP connection matches one of these patterns then the proxy settings will not be used for the connection.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/exceptionlist
func (n_ NEProxySettings) SetExceptionList(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setExceptionList:"), objc.String(value))
}

// A Boolean indicating if HTTP requests using single-label host names should be excluded from using the proxy settings.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/excludesimplehostnames
func (n_ NEProxySettings) ExcludeSimpleHostnames() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("excludeSimpleHostnames"))
	return rv
}


// SetExcludeSimpleHostnames sets the value of the excludeSimpleHostnames property.
// A Boolean indicating if HTTP requests using single-label host names should be excluded from using the proxy settings.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/excludesimplehostnames
func (n_ NEProxySettings) SetExcludeSimpleHostnames(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setExcludeSimpleHostnames:"), value)
}

// A Boolean indicating if a static HTTP proxy will be used.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/httpenabled
func (n_ NEProxySettings) HttpEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("httpEnabled"))
	return rv
}


// SetHttpEnabled sets the value of the httpEnabled property.
// A Boolean indicating if a static HTTP proxy will be used.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/httpenabled
func (n_ NEProxySettings) SetHttpEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHttpEnabled:"), value)
}

// An
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/httpserver
func (n_ NEProxySettings) HttpServer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("httpServer"))
	return rv
}


// SetHttpServer sets the value of the httpServer property.
// An

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/httpserver
func (n_ NEProxySettings) SetHttpServer(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHttpServer:"), value)
}

// A Boolean indicating if a static HTTPS proxy will be used.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/httpsenabled
func (n_ NEProxySettings) HttpsEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("httpsEnabled"))
	return rv
}


// SetHttpsEnabled sets the value of the httpsEnabled property.
// A Boolean indicating if a static HTTPS proxy will be used.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/httpsenabled
func (n_ NEProxySettings) SetHttpsEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHttpsEnabled:"), value)
}

// An
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/httpsserver
func (n_ NEProxySettings) HttpsServer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("httpsServer"))
	return rv
}


// SetHttpsServer sets the value of the httpsServer property.
// An

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/httpsserver
func (n_ NEProxySettings) SetHttpsServer(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHttpsServer:"), value)
}

// An array of domain strings.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/matchdomains
func (n_ NEProxySettings) MatchDomains() string {
	rv := objc.Send[string](n_.ID, objc.Sel("matchDomains"))
	return rv
}


// SetMatchDomains sets the value of the matchDomains property.
// An array of domain strings.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/matchdomains
func (n_ NEProxySettings) SetMatchDomains(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchDomains:"), objc.String(value))
}

// A string containing the Proxy Auto Configuration (PAC) JavaScript source code.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/proxyautoconfigurationjavascript
func (n_ NEProxySettings) ProxyAutoConfigurationJavaScript() string {
	rv := objc.Send[string](n_.ID, objc.Sel("proxyAutoConfigurationJavaScript"))
	return rv
}


// SetProxyAutoConfigurationJavaScript sets the value of the proxyAutoConfigurationJavaScript property.
// A string containing the Proxy Auto Configuration (PAC) JavaScript source code.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/proxyautoconfigurationjavascript
func (n_ NEProxySettings) SetProxyAutoConfigurationJavaScript(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProxyAutoConfigurationJavaScript:"), objc.String(value))
}

// A URL specifying the location from where the Proxy Auto Configuration (PAC) script should be downloaded.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/proxyautoconfigurationurl
func (n_ NEProxySettings) ProxyAutoConfigurationURL() foundation.URL {
	rv := objc.Send[foundation.URL](n_.ID, objc.Sel("proxyAutoConfigurationURL"))
	return rv
}


// SetProxyAutoConfigurationURL sets the value of the proxyAutoConfigurationURL property.
// A URL specifying the location from where the Proxy Auto Configuration (PAC) script should be downloaded.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neproxysettings/proxyautoconfigurationurl
func (n_ NEProxySettings) SetProxyAutoConfigurationURL(value foundation.URL) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProxyAutoConfigurationURL:"), value)
}

// The tunnel DNS settings.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelnetworksettings/dnssettings
func (n_ NEProxySettings) DnsSettings() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("dnsSettings"))
	return rv
}


// SetDnsSettings sets the value of the dnsSettings property.
// The tunnel DNS settings.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelnetworksettings/dnssettings
func (n_ NEProxySettings) SetDnsSettings(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDnsSettings:"), value)
}

// The tunnel HTTP proxy settings.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelnetworksettings/proxysettings
func (n_ NEProxySettings) ProxySettings() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("proxySettings"))
	return rv
}


// SetProxySettings sets the value of the proxySettings property.
// The tunnel HTTP proxy settings.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelnetworksettings/proxysettings
func (n_ NEProxySettings) SetProxySettings(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProxySettings:"), value)
}

// The IP address of the tunnel server.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelnetworksettings/tunnelremoteaddress
func (n_ NEProxySettings) TunnelRemoteAddress() string {
	rv := objc.Send[string](n_.ID, objc.Sel("tunnelRemoteAddress"))
	return rv
}


// SetTunnelRemoteAddress sets the value of the tunnelRemoteAddress property.
// The IP address of the tunnel server.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelnetworksettings/tunnelremoteaddress
func (n_ NEProxySettings) SetTunnelRemoteAddress(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTunnelRemoteAddress:"), objc.String(value))
}




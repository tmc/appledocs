// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	ProxyAutoConfigurationJavaScript() foundation.foundation.INSString
	SetProxyAutoConfigurationJavaScript(value foundation.foundation.INSString)
	ProxyAutoConfigurationURL() foundation.foundation.INSURL
	SetProxyAutoConfigurationURL(value foundation.foundation.INSURL)
	DnsSettings() INEDNSSettings
	SetDnsSettings(value INEDNSSettings)
	ProxySettings() INEProxySettings
	SetProxySettings(value INEProxySettings)
	TunnelRemoteAddress() foundation.foundation.INSString
	SetTunnelRemoteAddress(value foundation.foundation.INSString)


	

	// methods:


}





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

























// A Boolean indicating if proxy auto-configuration is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/autoProxyConfigurationEnabled
func (n_ NEProxySettings) AutoProxyConfigurationEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("autoProxyConfigurationEnabled"))
	return rv
}


// A Boolean indicating if proxy auto-configuration is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/autoProxyConfigurationEnabled
func (n_ NEProxySettings) SetAutoProxyConfigurationEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAutoProxyConfigurationEnabled:"), value)
}


// An array of domain name patterns. If the destination host name of an HTTP connection matches one of these patterns then the proxy settings will not be used for the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/exceptionList
func (n_ NEProxySettings) ExceptionList() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("exceptionList"))
	return rv
}


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
}


// A Boolean indicating if HTTP requests using single-label host names should be excluded from using the proxy settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/excludeSimpleHostnames
func (n_ NEProxySettings) ExcludeSimpleHostnames() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("excludeSimpleHostnames"))
	return rv
}


// A Boolean indicating if HTTP requests using single-label host names should be excluded from using the proxy settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/excludeSimpleHostnames
func (n_ NEProxySettings) SetExcludeSimpleHostnames(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setExcludeSimpleHostnames:"), value)
}


// A Boolean indicating if a static HTTP proxy will be used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/httpEnabled
func (n_ NEProxySettings) HTTPEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("HTTPEnabled"))
	return rv
}


// A Boolean indicating if a static HTTP proxy will be used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/httpEnabled
func (n_ NEProxySettings) SetHTTPEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHTTPEnabled:"), value)
}


// A Boolean indicating if a static HTTPS proxy will be used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/httpsEnabled
func (n_ NEProxySettings) HTTPSEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("HTTPSEnabled"))
	return rv
}


// A Boolean indicating if a static HTTPS proxy will be used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/httpsEnabled
func (n_ NEProxySettings) SetHTTPSEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHTTPSEnabled:"), value)
}


// An object containing the static HTTP proxy server settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/httpServer
func (n_ NEProxySettings) HTTPServer() INEProxyServer {
	rv := objc.Send[NEProxyServer](n_.ID, objc.Sel("HTTPServer"))
	return rv
}


// An object containing the static HTTP proxy server settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/httpServer
func (n_ NEProxySettings) SetHTTPServer(value INEProxyServer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHTTPServer:"), value)
}


// An object containing the static HTTPS proxy server settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/httpsServer
func (n_ NEProxySettings) HTTPSServer() INEProxyServer {
	rv := objc.Send[NEProxyServer](n_.ID, objc.Sel("HTTPSServer"))
	return rv
}


// An object containing the static HTTPS proxy server settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/httpsServer
func (n_ NEProxySettings) SetHTTPSServer(value INEProxyServer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHTTPSServer:"), value)
}


// An array of domain strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/matchDomains
func (n_ NEProxySettings) MatchDomains() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("matchDomains"))
	return rv
}


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
}


// A string containing the Proxy Auto Configuration (PAC) JavaScript source code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/proxyAutoConfigurationJavaScript
func (n_ NEProxySettings) ProxyAutoConfigurationJavaScript() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("proxyAutoConfigurationJavaScript"))
	return rv
}


// A string containing the Proxy Auto Configuration (PAC) JavaScript source code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/proxyAutoConfigurationJavaScript
func (n_ NEProxySettings) SetProxyAutoConfigurationJavaScript(value foundation.foundation.INSString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProxyAutoConfigurationJavaScript:"), value)
}


// A URL specifying the location from where the Proxy Auto Configuration (PAC) script should be downloaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/proxyAutoConfigurationURL
func (n_ NEProxySettings) ProxyAutoConfigurationURL() foundation.foundation.INSURL {
	rv := objc.Send[foundation.NSURL](n_.ID, objc.Sel("proxyAutoConfigurationURL"))
	return rv
}


// A URL specifying the location from where the Proxy Auto Configuration (PAC) script should be downloaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEProxySettings/proxyAutoConfigurationURL
func (n_ NEProxySettings) SetProxyAutoConfigurationURL(value foundation.foundation.INSURL) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProxyAutoConfigurationURL:"), value)
}


// The tunnel DNS settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelnetworksettings/dnssettings
func (n_ NEProxySettings) DnsSettings() INEDNSSettings {
	rv := objc.Send[NEDNSSettings](n_.ID, objc.Sel("dnsSettings"))
	return rv
}


// The tunnel DNS settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelnetworksettings/dnssettings
func (n_ NEProxySettings) SetDnsSettings(value INEDNSSettings) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDnsSettings:"), value)
}


// The tunnel HTTP proxy settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelnetworksettings/proxysettings
func (n_ NEProxySettings) ProxySettings() INEProxySettings {
	rv := objc.Send[NEProxySettings](n_.ID, objc.Sel("proxySettings"))
	return rv
}


// The tunnel HTTP proxy settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelnetworksettings/proxysettings
func (n_ NEProxySettings) SetProxySettings(value INEProxySettings) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setProxySettings:"), value)
}


// The IP address of the tunnel server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelnetworksettings/tunnelremoteaddress
func (n_ NEProxySettings) TunnelRemoteAddress() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("tunnelRemoteAddress"))
	return rv
}


// The IP address of the tunnel server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/netunnelnetworksettings/tunnelremoteaddress
func (n_ NEProxySettings) SetTunnelRemoteAddress(value foundation.foundation.INSString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTunnelRemoteAddress:"), value)
}









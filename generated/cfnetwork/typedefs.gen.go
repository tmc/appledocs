// Code generated from Apple documentation for CFNetwork. DO NOT EDIT.

package cfnetwork

// Type aliases and typedefs
// CFHTTPAuthenticationRef - An opaque reference representing HTTP authentication information.
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPAuthentication
// CFHTTPAuthenticationRef has base type: struct _CFHTTPAuthentication *
type CFHTTPAuthenticationRef uintptr
// CFHTTPMessageRef - An opaque reference representing an HTTP message.
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPMessage
// CFHTTPMessageRef has base type: struct __CFHTTPMessage *
type CFHTTPMessageRef uintptr
// CFHostRef - An opaque reference representing an CFHost object.
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHost
// CFHostRef has base type: struct __CFHost *
type CFHostRef uintptr
// CFHostClientCallBack - Defines a pointer to the callback function that is called when an asynchronous resolution of a CFHost completes or an error occurs for an asynchronous CFHost resolution.
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHostClientCallBack
// CFHostClientCallBack has base type: void (*)(struct __CFHost *, enum CFHostInfoType, const CFStreamError *, void *)
type CFHostClientCallBack uintptr
// CFNetDiagnosticRef - An opaque reference representing a CFNetDiagnostic.
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetDiagnostic
// CFNetDiagnosticRef has base type: struct __CFNetDiagnostic *
type CFNetDiagnosticRef uintptr
// CFNetDiagnosticStatus - A CFIndex type that is used to return status values from   status and diagnostic functions. For a list of possible values, see  .
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetDiagnosticStatus
// CFNetDiagnosticStatus has base type: CFIndex
type CFNetDiagnosticStatus uintptr
// CFNetServiceRef - An opaque reference representing a CFNetService.
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetService
// CFNetServiceRef has base type: struct __CFNetService *
type CFNetServiceRef uintptr
// CFNetServiceBrowserRef - An opaque reference representing a CFNetServiceBrowser.
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceBrowser
// CFNetServiceBrowserRef has base type: struct __CFNetServiceBrowser *
type CFNetServiceBrowserRef uintptr
// CFNetServiceBrowserClientCallBack - Defines a pointer to the callback function for a CFNetServiceBrowser.
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceBrowserClientCallBack
// CFNetServiceBrowserClientCallBack has base type: void (*)(struct __CFNetServiceBrowser *, unsigned long, const void *, CFStreamError *, void *)
type CFNetServiceBrowserClientCallBack uintptr
// CFNetServiceClientCallBack - Defines a pointer to the callback function for a CFNetService.
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceClientCallBack
// CFNetServiceClientCallBack has base type: void (*)(struct __CFNetService *, CFStreamError *, void *)
type CFNetServiceClientCallBack uintptr
// CFNetServiceMonitorRef - An opaque reference for a service monitor.
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceMonitor
// CFNetServiceMonitorRef has base type: struct __CFNetServiceMonitor *
type CFNetServiceMonitorRef uintptr
// CFNetServiceMonitorClientCallBack - Defines a pointer to the callback function that is to be called when a monitored record type changes.
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceMonitorClientCallBack
// CFNetServiceMonitorClientCallBack has base type: void (*)(struct __CFNetServiceMonitor *, struct __CFNetService *, enum CFNetServiceMonitorType, const struct __CFData *, CFStreamError *, void *)
type CFNetServiceMonitorClientCallBack uintptr
// CFProxyAutoConfigurationResultCallback - Callback function called when a proxy autoconfiguration computation has completed.
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFProxyAutoConfigurationResultCallback
// CFProxyAutoConfigurationResultCallback has base type: void (*)(void *, const struct __CFArray *, struct __CFError *)
type CFProxyAutoConfigurationResultCallback uintptr


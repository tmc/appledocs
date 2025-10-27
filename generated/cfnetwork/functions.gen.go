// Code generated from Apple documentation for CFNetwork. DO NOT EDIT.

package cfnetwork


import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// CFNetwork Functions (97 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_CFFTPCreateParsedResourceListing func(AllocatorRef, unsafe.Pointer, Index, unsafe.Pointer) Index
	_CFHTTPAuthenticationAppliesToRequest func(HTTPAuthenticationRef, HTTPMessageRef) unsafe.Pointer
	_CFHTTPAuthenticationCopyDomains func(HTTPAuthenticationRef) ArrayRef
	_CFHTTPAuthenticationCopyMethod func(HTTPAuthenticationRef) StringRef
	_CFHTTPAuthenticationCopyRealm func(HTTPAuthenticationRef) StringRef
	_CFHTTPAuthenticationCreateFromResponse func(AllocatorRef, HTTPMessageRef) HTTPAuthenticationRef
	_CFHTTPAuthenticationGetTypeID func() TypeID
	_CFHTTPAuthenticationIsValid func(HTTPAuthenticationRef, unsafe.Pointer) unsafe.Pointer
	_CFHTTPAuthenticationRequiresAccountDomain func(HTTPAuthenticationRef) unsafe.Pointer
	_CFHTTPAuthenticationRequiresOrderedRequests func(HTTPAuthenticationRef) unsafe.Pointer
	_CFHTTPAuthenticationRequiresUserNameAndPassword func(HTTPAuthenticationRef) unsafe.Pointer
	_CFHTTPMessageAddAuthentication func(HTTPMessageRef, HTTPMessageRef, StringRef, StringRef, StringRef, unsafe.Pointer) unsafe.Pointer
	_CFHTTPMessageAppendBytes func(HTTPMessageRef, unsafe.Pointer, Index) unsafe.Pointer
	_CFHTTPMessageApplyCredentialDictionary func(HTTPMessageRef, HTTPAuthenticationRef, DictionaryRef, unsafe.Pointer) unsafe.Pointer
	_CFHTTPMessageApplyCredentials func(HTTPMessageRef, HTTPAuthenticationRef, StringRef, StringRef, unsafe.Pointer) unsafe.Pointer
	_CFHTTPMessageCopyAllHeaderFields func(HTTPMessageRef) DictionaryRef
	_CFHTTPMessageCopyBody func(HTTPMessageRef) DataRef
	_CFHTTPMessageCopyHeaderFieldValue func(HTTPMessageRef, StringRef) StringRef
	_CFHTTPMessageCopyRequestMethod func(HTTPMessageRef) StringRef
	_CFHTTPMessageCopyRequestURL func(HTTPMessageRef) URLRef
	_CFHTTPMessageCopyResponseStatusLine func(HTTPMessageRef) StringRef
	_CFHTTPMessageCopySerializedMessage func(HTTPMessageRef) DataRef
	_CFHTTPMessageCopyVersion func(HTTPMessageRef) StringRef
	_CFHTTPMessageCreateCopy func(AllocatorRef, HTTPMessageRef) HTTPMessageRef
	_CFHTTPMessageCreateEmpty func(AllocatorRef, unsafe.Pointer) HTTPMessageRef
	_CFHTTPMessageCreateRequest func(AllocatorRef, StringRef, URLRef, StringRef) HTTPMessageRef
	_CFHTTPMessageCreateResponse func(AllocatorRef, Index, StringRef, StringRef) HTTPMessageRef
	_CFHTTPMessageGetResponseStatusCode func(HTTPMessageRef) Index
	_CFHTTPMessageGetTypeID func() TypeID
	_CFHTTPMessageIsHeaderComplete func(HTTPMessageRef) unsafe.Pointer
	_CFHTTPMessageIsRequest func(HTTPMessageRef) unsafe.Pointer
	_CFHTTPMessageSetBody func(HTTPMessageRef, DataRef)
	_CFHTTPMessageSetHeaderFieldValue func(HTTPMessageRef, StringRef, StringRef)
	_CFHostCancelInfoResolution func(HostRef, HostInfoType)
	_CFHostCreateCopy func(AllocatorRef, HostRef) HostRef
	_CFHostCreateWithAddress func(AllocatorRef, DataRef) HostRef
	_CFHostCreateWithName func(AllocatorRef, StringRef) HostRef
	_CFHostGetAddressing func(HostRef, unsafe.Pointer) ArrayRef
	_CFHostGetNames func(HostRef, unsafe.Pointer) ArrayRef
	_CFHostGetReachability func(HostRef, unsafe.Pointer) DataRef
	_CFHostGetTypeID func() TypeID
	_CFHostScheduleWithRunLoop func(HostRef, RunLoopRef, StringRef)
	_CFHostSetClient func(HostRef, HostClientCallBack, unsafe.Pointer) unsafe.Pointer
	_CFHostStartInfoResolution func(HostRef, HostInfoType, unsafe.Pointer) unsafe.Pointer
	_CFHostUnscheduleFromRunLoop func(HostRef, RunLoopRef, StringRef)
	_CFNetDiagnosticCopyNetworkStatusPassively func(NetDiagnosticRef, unsafe.Pointer) NetDiagnosticStatus
	_CFNetDiagnosticCreateWithStreams func(AllocatorRef, ReadStreamRef, WriteStreamRef) NetDiagnosticRef
	_CFNetDiagnosticCreateWithURL func(AllocatorRef, URLRef) NetDiagnosticRef
	_CFNetDiagnosticDiagnoseProblemInteractively func(NetDiagnosticRef) NetDiagnosticStatus
	_CFNetDiagnosticSetName func(NetDiagnosticRef, StringRef)
	_CFNetServiceBrowserCreate func(AllocatorRef, NetServiceBrowserClientCallBack, unsafe.Pointer) NetServiceBrowserRef
	_CFNetServiceBrowserGetTypeID func() TypeID
	_CFNetServiceBrowserInvalidate func(NetServiceBrowserRef)
	_CFNetServiceBrowserScheduleWithRunLoop func(NetServiceBrowserRef, RunLoopRef, StringRef)
	_CFNetServiceBrowserSearchForDomains func(NetServiceBrowserRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFNetServiceBrowserSearchForServices func(NetServiceBrowserRef, StringRef, StringRef, unsafe.Pointer) unsafe.Pointer
	_CFNetServiceBrowserStopSearch func(NetServiceBrowserRef, unsafe.Pointer)
	_CFNetServiceBrowserUnscheduleFromRunLoop func(NetServiceBrowserRef, RunLoopRef, StringRef)
	_CFNetServiceCancel func(NetServiceRef)
	_CFNetServiceCreate func(AllocatorRef, StringRef, StringRef, StringRef, unsafe.Pointer) NetServiceRef
	_CFNetServiceCreateCopy func(AllocatorRef, NetServiceRef) NetServiceRef
	_CFNetServiceCreateDictionaryWithTXTData func(AllocatorRef, DataRef) DictionaryRef
	_CFNetServiceCreateTXTDataWithDictionary func(AllocatorRef, DictionaryRef) DataRef
	_CFNetServiceGetAddressing func(NetServiceRef) ArrayRef
	_CFNetServiceGetDomain func(NetServiceRef) StringRef
	_CFNetServiceGetName func(NetServiceRef) StringRef
	_CFNetServiceGetPortNumber func(NetServiceRef) unsafe.Pointer
	_CFNetServiceGetTXTData func(NetServiceRef) DataRef
	_CFNetServiceGetTargetHost func(NetServiceRef) StringRef
	_CFNetServiceGetType func(NetServiceRef) StringRef
	_CFNetServiceGetTypeID func() TypeID
	_CFNetServiceMonitorCreate func(AllocatorRef, NetServiceRef, NetServiceMonitorClientCallBack, unsafe.Pointer) NetServiceMonitorRef
	_CFNetServiceMonitorGetTypeID func() TypeID
	_CFNetServiceMonitorInvalidate func(NetServiceMonitorRef)
	_CFNetServiceMonitorScheduleWithRunLoop func(NetServiceMonitorRef, RunLoopRef, StringRef)
	_CFNetServiceMonitorStart func(NetServiceMonitorRef, NetServiceMonitorType, unsafe.Pointer) unsafe.Pointer
	_CFNetServiceMonitorStop func(NetServiceMonitorRef, unsafe.Pointer)
	_CFNetServiceMonitorUnscheduleFromRunLoop func(NetServiceMonitorRef, RunLoopRef, StringRef)
	_CFNetServiceRegister func(NetServiceRef, unsafe.Pointer) unsafe.Pointer
	_CFNetServiceRegisterWithOptions func(NetServiceRef, OptionFlags, unsafe.Pointer) unsafe.Pointer
	_CFNetServiceResolve func(NetServiceRef, unsafe.Pointer) unsafe.Pointer
	_CFNetServiceResolveWithTimeout func(NetServiceRef, TimeInterval, unsafe.Pointer) unsafe.Pointer
	_CFNetServiceScheduleWithRunLoop func(NetServiceRef, RunLoopRef, StringRef)
	_CFNetServiceSetClient func(NetServiceRef, NetServiceClientCallBack, unsafe.Pointer) unsafe.Pointer
	_CFNetServiceSetTXTData func(NetServiceRef, DataRef) unsafe.Pointer
	_CFNetServiceUnscheduleFromRunLoop func(NetServiceRef, RunLoopRef, StringRef)
	_CFNetworkCopyProxiesForAutoConfigurationScript func(StringRef, URLRef, unsafe.Pointer) ArrayRef
	_CFNetworkCopyProxiesForURL func(URLRef, DictionaryRef) ArrayRef
	_CFNetworkCopySystemProxySettings func() DictionaryRef
	_CFNetworkExecuteProxyAutoConfigurationScript func(StringRef, URLRef, ProxyAutoConfigurationResultCallback, unsafe.Pointer) RunLoopSourceRef
	_CFNetworkExecuteProxyAutoConfigurationURL func(URLRef, URLRef, ProxyAutoConfigurationResultCallback, unsafe.Pointer) RunLoopSourceRef
	_CFReadStreamCreateForHTTPRequest func(AllocatorRef, HTTPMessageRef) ReadStreamRef
	_CFReadStreamCreateForStreamedHTTPRequest func(AllocatorRef, HTTPMessageRef, ReadStreamRef) ReadStreamRef
	_CFReadStreamCreateWithFTPURL func(AllocatorRef, URLRef) ReadStreamRef
	_CFStreamCreatePairWithSocketToCFHost func(AllocatorRef, HostRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFStreamCreatePairWithSocketToNetService func(AllocatorRef, NetServiceRef, unsafe.Pointer, unsafe.Pointer)
	_CFWriteStreamCreateWithFTPURL func(AllocatorRef, URLRef) WriteStreamRef
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_CFFTPCreateParsedResourceListing, lib, "CFFTPCreateParsedResourceListing")
	tryRegister(&_CFHTTPAuthenticationAppliesToRequest, lib, "CFHTTPAuthenticationAppliesToRequest")
	tryRegister(&_CFHTTPAuthenticationCopyDomains, lib, "CFHTTPAuthenticationCopyDomains")
	tryRegister(&_CFHTTPAuthenticationCopyMethod, lib, "CFHTTPAuthenticationCopyMethod")
	tryRegister(&_CFHTTPAuthenticationCopyRealm, lib, "CFHTTPAuthenticationCopyRealm")
	tryRegister(&_CFHTTPAuthenticationCreateFromResponse, lib, "CFHTTPAuthenticationCreateFromResponse")
	tryRegister(&_CFHTTPAuthenticationGetTypeID, lib, "CFHTTPAuthenticationGetTypeID")
	tryRegister(&_CFHTTPAuthenticationIsValid, lib, "CFHTTPAuthenticationIsValid")
	tryRegister(&_CFHTTPAuthenticationRequiresAccountDomain, lib, "CFHTTPAuthenticationRequiresAccountDomain")
	tryRegister(&_CFHTTPAuthenticationRequiresOrderedRequests, lib, "CFHTTPAuthenticationRequiresOrderedRequests")
	tryRegister(&_CFHTTPAuthenticationRequiresUserNameAndPassword, lib, "CFHTTPAuthenticationRequiresUserNameAndPassword")
	tryRegister(&_CFHTTPMessageAddAuthentication, lib, "CFHTTPMessageAddAuthentication")
	tryRegister(&_CFHTTPMessageAppendBytes, lib, "CFHTTPMessageAppendBytes")
	tryRegister(&_CFHTTPMessageApplyCredentialDictionary, lib, "CFHTTPMessageApplyCredentialDictionary")
	tryRegister(&_CFHTTPMessageApplyCredentials, lib, "CFHTTPMessageApplyCredentials")
	tryRegister(&_CFHTTPMessageCopyAllHeaderFields, lib, "CFHTTPMessageCopyAllHeaderFields")
	tryRegister(&_CFHTTPMessageCopyBody, lib, "CFHTTPMessageCopyBody")
	tryRegister(&_CFHTTPMessageCopyHeaderFieldValue, lib, "CFHTTPMessageCopyHeaderFieldValue")
	tryRegister(&_CFHTTPMessageCopyRequestMethod, lib, "CFHTTPMessageCopyRequestMethod")
	tryRegister(&_CFHTTPMessageCopyRequestURL, lib, "CFHTTPMessageCopyRequestURL")
	tryRegister(&_CFHTTPMessageCopyResponseStatusLine, lib, "CFHTTPMessageCopyResponseStatusLine")
	tryRegister(&_CFHTTPMessageCopySerializedMessage, lib, "CFHTTPMessageCopySerializedMessage")
	tryRegister(&_CFHTTPMessageCopyVersion, lib, "CFHTTPMessageCopyVersion")
	tryRegister(&_CFHTTPMessageCreateCopy, lib, "CFHTTPMessageCreateCopy")
	tryRegister(&_CFHTTPMessageCreateEmpty, lib, "CFHTTPMessageCreateEmpty")
	tryRegister(&_CFHTTPMessageCreateRequest, lib, "CFHTTPMessageCreateRequest")
	tryRegister(&_CFHTTPMessageCreateResponse, lib, "CFHTTPMessageCreateResponse")
	tryRegister(&_CFHTTPMessageGetResponseStatusCode, lib, "CFHTTPMessageGetResponseStatusCode")
	tryRegister(&_CFHTTPMessageGetTypeID, lib, "CFHTTPMessageGetTypeID")
	tryRegister(&_CFHTTPMessageIsHeaderComplete, lib, "CFHTTPMessageIsHeaderComplete")
	tryRegister(&_CFHTTPMessageIsRequest, lib, "CFHTTPMessageIsRequest")
	tryRegister(&_CFHTTPMessageSetBody, lib, "CFHTTPMessageSetBody")
	tryRegister(&_CFHTTPMessageSetHeaderFieldValue, lib, "CFHTTPMessageSetHeaderFieldValue")
	tryRegister(&_CFHostCancelInfoResolution, lib, "CFHostCancelInfoResolution")
	tryRegister(&_CFHostCreateCopy, lib, "CFHostCreateCopy")
	tryRegister(&_CFHostCreateWithAddress, lib, "CFHostCreateWithAddress")
	tryRegister(&_CFHostCreateWithName, lib, "CFHostCreateWithName")
	tryRegister(&_CFHostGetAddressing, lib, "CFHostGetAddressing")
	tryRegister(&_CFHostGetNames, lib, "CFHostGetNames")
	tryRegister(&_CFHostGetReachability, lib, "CFHostGetReachability")
	tryRegister(&_CFHostGetTypeID, lib, "CFHostGetTypeID")
	tryRegister(&_CFHostScheduleWithRunLoop, lib, "CFHostScheduleWithRunLoop")
	tryRegister(&_CFHostSetClient, lib, "CFHostSetClient")
	tryRegister(&_CFHostStartInfoResolution, lib, "CFHostStartInfoResolution")
	tryRegister(&_CFHostUnscheduleFromRunLoop, lib, "CFHostUnscheduleFromRunLoop")
	tryRegister(&_CFNetDiagnosticCopyNetworkStatusPassively, lib, "CFNetDiagnosticCopyNetworkStatusPassively")
	tryRegister(&_CFNetDiagnosticCreateWithStreams, lib, "CFNetDiagnosticCreateWithStreams")
	tryRegister(&_CFNetDiagnosticCreateWithURL, lib, "CFNetDiagnosticCreateWithURL")
	tryRegister(&_CFNetDiagnosticDiagnoseProblemInteractively, lib, "CFNetDiagnosticDiagnoseProblemInteractively")
	tryRegister(&_CFNetDiagnosticSetName, lib, "CFNetDiagnosticSetName")
	tryRegister(&_CFNetServiceBrowserCreate, lib, "CFNetServiceBrowserCreate")
	tryRegister(&_CFNetServiceBrowserGetTypeID, lib, "CFNetServiceBrowserGetTypeID")
	tryRegister(&_CFNetServiceBrowserInvalidate, lib, "CFNetServiceBrowserInvalidate")
	tryRegister(&_CFNetServiceBrowserScheduleWithRunLoop, lib, "CFNetServiceBrowserScheduleWithRunLoop")
	tryRegister(&_CFNetServiceBrowserSearchForDomains, lib, "CFNetServiceBrowserSearchForDomains")
	tryRegister(&_CFNetServiceBrowserSearchForServices, lib, "CFNetServiceBrowserSearchForServices")
	tryRegister(&_CFNetServiceBrowserStopSearch, lib, "CFNetServiceBrowserStopSearch")
	tryRegister(&_CFNetServiceBrowserUnscheduleFromRunLoop, lib, "CFNetServiceBrowserUnscheduleFromRunLoop")
	tryRegister(&_CFNetServiceCancel, lib, "CFNetServiceCancel")
	tryRegister(&_CFNetServiceCreate, lib, "CFNetServiceCreate")
	tryRegister(&_CFNetServiceCreateCopy, lib, "CFNetServiceCreateCopy")
	tryRegister(&_CFNetServiceCreateDictionaryWithTXTData, lib, "CFNetServiceCreateDictionaryWithTXTData")
	tryRegister(&_CFNetServiceCreateTXTDataWithDictionary, lib, "CFNetServiceCreateTXTDataWithDictionary")
	tryRegister(&_CFNetServiceGetAddressing, lib, "CFNetServiceGetAddressing")
	tryRegister(&_CFNetServiceGetDomain, lib, "CFNetServiceGetDomain")
	tryRegister(&_CFNetServiceGetName, lib, "CFNetServiceGetName")
	tryRegister(&_CFNetServiceGetPortNumber, lib, "CFNetServiceGetPortNumber")
	tryRegister(&_CFNetServiceGetTXTData, lib, "CFNetServiceGetTXTData")
	tryRegister(&_CFNetServiceGetTargetHost, lib, "CFNetServiceGetTargetHost")
	tryRegister(&_CFNetServiceGetType, lib, "CFNetServiceGetType")
	tryRegister(&_CFNetServiceGetTypeID, lib, "CFNetServiceGetTypeID")
	tryRegister(&_CFNetServiceMonitorCreate, lib, "CFNetServiceMonitorCreate")
	tryRegister(&_CFNetServiceMonitorGetTypeID, lib, "CFNetServiceMonitorGetTypeID")
	tryRegister(&_CFNetServiceMonitorInvalidate, lib, "CFNetServiceMonitorInvalidate")
	tryRegister(&_CFNetServiceMonitorScheduleWithRunLoop, lib, "CFNetServiceMonitorScheduleWithRunLoop")
	tryRegister(&_CFNetServiceMonitorStart, lib, "CFNetServiceMonitorStart")
	tryRegister(&_CFNetServiceMonitorStop, lib, "CFNetServiceMonitorStop")
	tryRegister(&_CFNetServiceMonitorUnscheduleFromRunLoop, lib, "CFNetServiceMonitorUnscheduleFromRunLoop")
	tryRegister(&_CFNetServiceRegister, lib, "CFNetServiceRegister")
	tryRegister(&_CFNetServiceRegisterWithOptions, lib, "CFNetServiceRegisterWithOptions")
	tryRegister(&_CFNetServiceResolve, lib, "CFNetServiceResolve")
	tryRegister(&_CFNetServiceResolveWithTimeout, lib, "CFNetServiceResolveWithTimeout")
	tryRegister(&_CFNetServiceScheduleWithRunLoop, lib, "CFNetServiceScheduleWithRunLoop")
	tryRegister(&_CFNetServiceSetClient, lib, "CFNetServiceSetClient")
	tryRegister(&_CFNetServiceSetTXTData, lib, "CFNetServiceSetTXTData")
	tryRegister(&_CFNetServiceUnscheduleFromRunLoop, lib, "CFNetServiceUnscheduleFromRunLoop")
	tryRegister(&_CFNetworkCopyProxiesForAutoConfigurationScript, lib, "CFNetworkCopyProxiesForAutoConfigurationScript")
	tryRegister(&_CFNetworkCopyProxiesForURL, lib, "CFNetworkCopyProxiesForURL")
	tryRegister(&_CFNetworkCopySystemProxySettings, lib, "CFNetworkCopySystemProxySettings")
	tryRegister(&_CFNetworkExecuteProxyAutoConfigurationScript, lib, "CFNetworkExecuteProxyAutoConfigurationScript")
	tryRegister(&_CFNetworkExecuteProxyAutoConfigurationURL, lib, "CFNetworkExecuteProxyAutoConfigurationURL")
	tryRegister(&_CFReadStreamCreateForHTTPRequest, lib, "CFReadStreamCreateForHTTPRequest")
	tryRegister(&_CFReadStreamCreateForStreamedHTTPRequest, lib, "CFReadStreamCreateForStreamedHTTPRequest")
	tryRegister(&_CFReadStreamCreateWithFTPURL, lib, "CFReadStreamCreateWithFTPURL")
	tryRegister(&_CFStreamCreatePairWithSocketToCFHost, lib, "CFStreamCreatePairWithSocketToCFHost")
	tryRegister(&_CFStreamCreatePairWithSocketToNetService, lib, "CFStreamCreatePairWithSocketToNetService")
	tryRegister(&_CFWriteStreamCreateWithFTPURL, lib, "CFWriteStreamCreateWithFTPURL")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}



// Parses an FTP listing to a dictionary.
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.3.
// Parses an FTP listing to a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFFTPCreateParsedResourceListing(_:_:_:_:)
func CFFTPCreateParsedResourceListing(alloc AllocatorRef, buffer unsafe.Pointer, bufferLength Index, parsed unsafe.Pointer) Index {
	return _CFFTPCreateParsedResourceListing(alloc, buffer, bufferLength, parsed)
}

// Returns a Boolean value that indicates whether a CFHTTPAuthentication object is associated with a CFHTTPMessage object.
//
// Added in macOS 10.2.
// Returns a Boolean value that indicates whether a CFHTTPAuthentication object is associated with a CFHTTPMessage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPAuthenticationAppliesToRequest(_:_:)
func CFHTTPAuthenticationAppliesToRequest(auth HTTPAuthenticationRef, request HTTPMessageRef) unsafe.Pointer {
	return _CFHTTPAuthenticationAppliesToRequest(auth, request)
}

// Returns an array of domain URLs to which a given CFHTTPAuthentication object can be applied.
//
// Added in macOS 10.2.
// Returns an array of domain URLs to which a given CFHTTPAuthentication object can be applied.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPAuthenticationCopyDomains(_:)
func CFHTTPAuthenticationCopyDomains(auth HTTPAuthenticationRef) ArrayRef {
	return _CFHTTPAuthenticationCopyDomains(auth)
}

// Gets the strongest authentication method that will be used when a CFHTTPAuthentication object is applied to a request.
//
// Added in macOS 10.2.
// Gets the strongest authentication method that will be used when a CFHTTPAuthentication object is applied to a request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPAuthenticationCopyMethod(_:)
func CFHTTPAuthenticationCopyMethod(auth HTTPAuthenticationRef) StringRef {
	return _CFHTTPAuthenticationCopyMethod(auth)
}

// Gets an authentication information’s namespace.
//
// Added in macOS 10.2.
// Gets an authentication information’s namespace.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPAuthenticationCopyRealm(_:)
func CFHTTPAuthenticationCopyRealm(auth HTTPAuthenticationRef) StringRef {
	return _CFHTTPAuthenticationCopyRealm(auth)
}

// Uses an authentication failure response to create a CFHTTPAuthentication object.
//
// Added in macOS 10.2.
// Uses an authentication failure response to create a CFHTTPAuthentication object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPAuthenticationCreateFromResponse(_:_:)
func CFHTTPAuthenticationCreateFromResponse(alloc AllocatorRef, response HTTPMessageRef) HTTPAuthenticationRef {
	return _CFHTTPAuthenticationCreateFromResponse(alloc, response)
}

// Gets the Core Foundation type identifier for the CFHTTPAuthentication opaque type.
//
// Added in macOS 10.2.
// Gets the Core Foundation type identifier for the CFHTTPAuthentication opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPAuthenticationGetTypeID()
func CFHTTPAuthenticationGetTypeID() TypeID {
	return _CFHTTPAuthenticationGetTypeID()
}

// Returns a Boolean value that indicates whether a CFHTTPAuthentication object is valid.
//
// Added in macOS 10.2.
// Returns a Boolean value that indicates whether a CFHTTPAuthentication object is valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPAuthenticationIsValid(_:_:)
func CFHTTPAuthenticationIsValid(auth HTTPAuthenticationRef, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFHTTPAuthenticationIsValid(auth, error_)
}

// Returns a Boolean value that indicates whether a CFHTTPAuthentication object uses an authentication method that requires an account domain.
//
// Added in macOS 10.4.
// Returns a Boolean value that indicates whether a CFHTTPAuthentication object uses an authentication method that requires an account domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPAuthenticationRequiresAccountDomain(_:)
func CFHTTPAuthenticationRequiresAccountDomain(auth HTTPAuthenticationRef) unsafe.Pointer {
	return _CFHTTPAuthenticationRequiresAccountDomain(auth)
}

// Returns a Boolean value that indicates whether authentication requests should be made one at a time.
//
// Added in macOS 10.2.
// Returns a Boolean value that indicates whether authentication requests should be made one at a time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPAuthenticationRequiresOrderedRequests(_:)
func CFHTTPAuthenticationRequiresOrderedRequests(auth HTTPAuthenticationRef) unsafe.Pointer {
	return _CFHTTPAuthenticationRequiresOrderedRequests(auth)
}

// Returns a Boolean value that indicates whether a CFHTTPAuthentication object uses an authentication method that requires a username and a password.
//
// Added in macOS 10.3.
// Returns a Boolean value that indicates whether a CFHTTPAuthentication object uses an authentication method that requires a username and a password.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPAuthenticationRequiresUserNameAndPassword(_:)
func CFHTTPAuthenticationRequiresUserNameAndPassword(auth HTTPAuthenticationRef) unsafe.Pointer {
	return _CFHTTPAuthenticationRequiresUserNameAndPassword(auth)
}

// Adds authentication information to a request.
//
// Added in macOS 10.1.
// Adds authentication information to a request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPMessageAddAuthentication(_:_:_:_:_:_:)
func CFHTTPMessageAddAuthentication(request HTTPMessageRef, authenticationFailureResponse HTTPMessageRef, username StringRef, password StringRef, authenticationScheme StringRef, forProxy unsafe.Pointer) unsafe.Pointer {
	return _CFHTTPMessageAddAuthentication(request, authenticationFailureResponse, username, password, authenticationScheme, forProxy)
}

// Appends data to a object.
//
// Added in macOS 10.1.
// Appends data to a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPMessageAppendBytes(_:_:_:)
func CFHTTPMessageAppendBytes(message HTTPMessageRef, newBytes unsafe.Pointer, numBytes Index) unsafe.Pointer {
	return _CFHTTPMessageAppendBytes(message, newBytes, numBytes)
}

// Use a dictionary containing authentication credentials to perform the authentication method specified by a object.
//
// Added in macOS 10.4.
// Use a dictionary containing authentication credentials to perform the authentication method specified by a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPMessageApplyCredentialDictionary(_:_:_:_:)
func CFHTTPMessageApplyCredentialDictionary(request HTTPMessageRef, auth HTTPAuthenticationRef, dict DictionaryRef, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFHTTPMessageApplyCredentialDictionary(request, auth, dict, error_)
}

// Performs the authentication method specified by a object.
//
// Added in macOS 10.2.
// Performs the authentication method specified by a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPMessageApplyCredentials(_:_:_:_:_:)
func CFHTTPMessageApplyCredentials(request HTTPMessageRef, auth HTTPAuthenticationRef, username StringRef, password StringRef, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFHTTPMessageApplyCredentials(request, auth, username, password, error_)
}

// Gets all header fields from a object.
//
// Added in macOS 10.1.
// Gets all header fields from a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPMessageCopyAllHeaderFields(_:)
func CFHTTPMessageCopyAllHeaderFields(message HTTPMessageRef) DictionaryRef {
	return _CFHTTPMessageCopyAllHeaderFields(message)
}

// Gets the body from a object.
//
// Added in macOS 10.1.
// Gets the body from a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPMessageCopyBody(_:)
func CFHTTPMessageCopyBody(message HTTPMessageRef) DataRef {
	return _CFHTTPMessageCopyBody(message)
}

// Gets the value of a header field from a object.
//
// Added in macOS 10.1.
// Gets the value of a header field from a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPMessageCopyHeaderFieldValue(_:_:)
func CFHTTPMessageCopyHeaderFieldValue(message HTTPMessageRef, headerField StringRef) StringRef {
	return _CFHTTPMessageCopyHeaderFieldValue(message, headerField)
}

// Gets the request method from a object.
//
// Added in macOS 10.1.
// Gets the request method from a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPMessageCopyRequestMethod(_:)
func CFHTTPMessageCopyRequestMethod(request HTTPMessageRef) StringRef {
	return _CFHTTPMessageCopyRequestMethod(request)
}

// Gets the URL from a object.
//
// Added in macOS 10.1.
// Gets the URL from a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPMessageCopyRequestURL(_:)
func CFHTTPMessageCopyRequestURL(request HTTPMessageRef) URLRef {
	return _CFHTTPMessageCopyRequestURL(request)
}

// Gets the status line from a object.
//
// Added in macOS 10.1.
// Gets the status line from a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPMessageCopyResponseStatusLine(_:)
func CFHTTPMessageCopyResponseStatusLine(response HTTPMessageRef) StringRef {
	return _CFHTTPMessageCopyResponseStatusLine(response)
}

// Serializes a CFHTTPMessage object.
//
// Added in macOS 10.1.
// Serializes a CFHTTPMessage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPMessageCopySerializedMessage(_:)
func CFHTTPMessageCopySerializedMessage(message HTTPMessageRef) DataRef {
	return _CFHTTPMessageCopySerializedMessage(message)
}

// Gets the HTTP version from a object.
//
// Added in macOS 10.1.
// Gets the HTTP version from a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPMessageCopyVersion(_:)
func CFHTTPMessageCopyVersion(message HTTPMessageRef) StringRef {
	return _CFHTTPMessageCopyVersion(message)
}

// Gets a copy of a CFHTTPMessage object.
//
// Added in macOS 10.1.
// Gets a copy of a CFHTTPMessage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPMessageCreateCopy(_:_:)
func CFHTTPMessageCreateCopy(alloc AllocatorRef, message HTTPMessageRef) HTTPMessageRef {
	return _CFHTTPMessageCreateCopy(alloc, message)
}

// Creates and returns a new, empty object.
//
// Added in macOS 10.1.
// Creates and returns a new, empty object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPMessageCreateEmpty(_:_:)
func CFHTTPMessageCreateEmpty(alloc AllocatorRef, isRequest unsafe.Pointer) HTTPMessageRef {
	return _CFHTTPMessageCreateEmpty(alloc, isRequest)
}

// Creates and returns a object for an HTTP request.
//
// Added in macOS 10.1.
// Creates and returns a object for an HTTP request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPMessageCreateRequest(_:_:_:_:)
func CFHTTPMessageCreateRequest(alloc AllocatorRef, requestMethod StringRef, url URLRef, httpVersion StringRef) HTTPMessageRef {
	return _CFHTTPMessageCreateRequest(alloc, requestMethod, url, httpVersion)
}

// Creates and returns a object for an HTTP response.
//
// Added in macOS 10.1.
// Creates and returns a object for an HTTP response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPMessageCreateResponse(_:_:_:_:)
func CFHTTPMessageCreateResponse(alloc AllocatorRef, statusCode Index, statusDescription StringRef, httpVersion StringRef) HTTPMessageRef {
	return _CFHTTPMessageCreateResponse(alloc, statusCode, statusDescription, httpVersion)
}

// Gets the status code from a object representing an HTTP response.
//
// Added in macOS 10.1.
// Gets the status code from a object representing an HTTP response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPMessageGetResponseStatusCode(_:)
func CFHTTPMessageGetResponseStatusCode(response HTTPMessageRef) Index {
	return _CFHTTPMessageGetResponseStatusCode(response)
}

// Returns the Core Foundation type identifier for the opaque type.
//
// Added in macOS 10.1.
// Returns the Core Foundation type identifier for the opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPMessageGetTypeID()
func CFHTTPMessageGetTypeID() TypeID {
	return _CFHTTPMessageGetTypeID()
}

// Determines whether a message header is complete.
//
// Added in macOS 10.1.
// Determines whether a message header is complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPMessageIsHeaderComplete(_:)
func CFHTTPMessageIsHeaderComplete(message HTTPMessageRef) unsafe.Pointer {
	return _CFHTTPMessageIsHeaderComplete(message)
}

// Returns a Boolean indicating whether the HTTP message is a request or a response.
//
// Added in macOS 10.1.
// Returns a Boolean indicating whether the HTTP message is a request or a response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPMessageIsRequest(_:)
func CFHTTPMessageIsRequest(message HTTPMessageRef) unsafe.Pointer {
	return _CFHTTPMessageIsRequest(message)
}

// Sets the body of a object.
//
// Added in macOS 10.1.
// Sets the body of a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPMessageSetBody(_:_:)
func CFHTTPMessageSetBody(message HTTPMessageRef, bodyData DataRef) {
	_CFHTTPMessageSetBody(message, bodyData)
}

// Sets the value of a header field in an HTTP message.
//
// Added in macOS 10.1.
// Sets the value of a header field in an HTTP message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHTTPMessageSetHeaderFieldValue(_:_:_:)
func CFHTTPMessageSetHeaderFieldValue(message HTTPMessageRef, headerField StringRef, value StringRef) {
	_CFHTTPMessageSetHeaderFieldValue(message, headerField, value)
}

// Cancels the resolution of a host.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.3.
// Cancels the resolution of a host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHostCancelInfoResolution(_:_:)
func CFHostCancelInfoResolution(theHost HostRef, info HostInfoType) {
	_CFHostCancelInfoResolution(theHost, info)
}

// Creates a new host object by copying.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.3.
// Creates a new host object by copying.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHostCreateCopy(_:_:)
func CFHostCreateCopy(alloc AllocatorRef, host HostRef) HostRef {
	return _CFHostCreateCopy(alloc, host)
}

// Uses an address to create an instance of a host object.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.3.
// Uses an address to create an instance of a host object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHostCreateWithAddress(_:_:)
func CFHostCreateWithAddress(allocator AllocatorRef, addr DataRef) HostRef {
	return _CFHostCreateWithAddress(allocator, addr)
}

// Uses a name to create an instance of a host object.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.3.
// Uses a name to create an instance of a host object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHostCreateWithName(_:_:)
func CFHostCreateWithName(allocator AllocatorRef, hostname StringRef) HostRef {
	return _CFHostCreateWithName(allocator, hostname)
}

// Gets the addresses from a host.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.3.
// Gets the addresses from a host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHostGetAddressing(_:_:)
func CFHostGetAddressing(theHost HostRef, hasBeenResolved unsafe.Pointer) ArrayRef {
	return _CFHostGetAddressing(theHost, hasBeenResolved)
}

// Gets the names from a CFHost.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.3.
// Gets the names from a CFHost.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHostGetNames(_:_:)
func CFHostGetNames(theHost HostRef, hasBeenResolved unsafe.Pointer) ArrayRef {
	return _CFHostGetNames(theHost, hasBeenResolved)
}

// Gets reachability information from a host.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.3.
// Gets reachability information from a host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHostGetReachability(_:_:)
func CFHostGetReachability(theHost HostRef, hasBeenResolved unsafe.Pointer) DataRef {
	return _CFHostGetReachability(theHost, hasBeenResolved)
}

// Gets the Core Foundation type identifier for the CFHost opaque type.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.3.
// Gets the Core Foundation type identifier for the CFHost opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHostGetTypeID()
func CFHostGetTypeID() TypeID {
	return _CFHostGetTypeID()
}

// Schedules a CFHost on a run loop.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.3.
// Schedules a CFHost on a run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHostScheduleWithRunLoop(_:_:_:)
func CFHostScheduleWithRunLoop(theHost HostRef, runLoop RunLoopRef, runLoopMode StringRef) {
	_CFHostScheduleWithRunLoop(theHost, runLoop, runLoopMode)
}

// Associates a client context and a callback function with a CFHost object or disassociates a client context and callback function that were previously set.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.3.
// Associates a client context and a callback function with a CFHost object or disassociates a client context and callback function that were previously set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHostSetClient(_:_:_:)
func CFHostSetClient(theHost HostRef, clientCB HostClientCallBack, clientContext unsafe.Pointer) unsafe.Pointer {
	return _CFHostSetClient(theHost, clientCB, clientContext)
}

// Starts resolution for a host object.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.3.
// Starts resolution for a host object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHostStartInfoResolution(_:_:_:)
func CFHostStartInfoResolution(theHost HostRef, info HostInfoType, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFHostStartInfoResolution(theHost, info, error_)
}

// Unschedules a CFHost from a run loop.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.3.
// Unschedules a CFHost from a run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHostUnscheduleFromRunLoop(_:_:_:)
func CFHostUnscheduleFromRunLoop(theHost HostRef, runLoop RunLoopRef, runLoopMode StringRef) {
	_CFHostUnscheduleFromRunLoop(theHost, runLoop, runLoopMode)
}

// Gets a network status value.
//
// Deprecated: This function was deprecated in macOS 10.13.
//
// Added in macOS 10.4.
// Gets a network status value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetDiagnosticCopyNetworkStatusPassively(_:_:)
func CFNetDiagnosticCopyNetworkStatusPassively(details NetDiagnosticRef, description unsafe.Pointer) NetDiagnosticStatus {
	return _CFNetDiagnosticCopyNetworkStatusPassively(details, description)
}

// Creates a network diagnostic object from a pair of CFStreams.
//
// Deprecated: This function was deprecated in macOS 10.13.
//
// Added in macOS 10.4.
// Creates a network diagnostic object from a pair of CFStreams.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetDiagnosticCreateWithStreams(_:_:_:)
func CFNetDiagnosticCreateWithStreams(alloc AllocatorRef, readStream ReadStreamRef, writeStream WriteStreamRef) NetDiagnosticRef {
	return _CFNetDiagnosticCreateWithStreams(alloc, readStream, writeStream)
}

// Creates a CFNetDiagnosticRef from a CFURLRef.
//
// Deprecated: This function was deprecated in macOS 10.13.
//
// Added in macOS 10.4.
// Creates a CFNetDiagnosticRef from a CFURLRef.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetDiagnosticCreateWithURL(_:_:)
func CFNetDiagnosticCreateWithURL(alloc AllocatorRef, url URLRef) NetDiagnosticRef {
	return _CFNetDiagnosticCreateWithURL(alloc, url)
}

// Opens a Network Diagnostics window.
//
// Deprecated: This function was deprecated in macOS 10.13.
//
// Added in macOS 10.4.
// Opens a Network Diagnostics window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetDiagnosticDiagnoseProblemInteractively(_:)
func CFNetDiagnosticDiagnoseProblemInteractively(details NetDiagnosticRef) NetDiagnosticStatus {
	return _CFNetDiagnosticDiagnoseProblemInteractively(details)
}

// Overrides the displayed application name.
//
// Deprecated: This function was deprecated in macOS 10.13.
//
// Added in macOS 10.4.
// Overrides the displayed application name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetDiagnosticSetName(_:_:)
func CFNetDiagnosticSetName(details NetDiagnosticRef, name StringRef) {
	_CFNetDiagnosticSetName(details, name)
}

// Creates an instance of a Network Service browser object.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.2.
// Creates an instance of a Network Service browser object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceBrowserCreate(_:_:_:)
func CFNetServiceBrowserCreate(alloc AllocatorRef, clientCB NetServiceBrowserClientCallBack, clientContext unsafe.Pointer) NetServiceBrowserRef {
	return _CFNetServiceBrowserCreate(alloc, clientCB, clientContext)
}

// Gets the Core Foundation type identifier for the Network Service browser object.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.2.
// Gets the Core Foundation type identifier for the Network Service browser object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceBrowserGetTypeID()
func CFNetServiceBrowserGetTypeID() TypeID {
	return _CFNetServiceBrowserGetTypeID()
}

// Invalidates an instance of a Network Service browser object.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.2.
// Invalidates an instance of a Network Service browser object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceBrowserInvalidate(_:)
func CFNetServiceBrowserInvalidate(browser NetServiceBrowserRef) {
	_CFNetServiceBrowserInvalidate(browser)
}

// Schedules a CFNetServiceBrowser on a run loop.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.2.
// Schedules a CFNetServiceBrowser on a run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceBrowserScheduleWithRunLoop(_:_:_:)
func CFNetServiceBrowserScheduleWithRunLoop(browser NetServiceBrowserRef, runLoop RunLoopRef, runLoopMode StringRef) {
	_CFNetServiceBrowserScheduleWithRunLoop(browser, runLoop, runLoopMode)
}

// Searches for domains.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.2.
// Searches for domains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceBrowserSearchForDomains(_:_:_:)
func CFNetServiceBrowserSearchForDomains(browser NetServiceBrowserRef, registrationDomains unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFNetServiceBrowserSearchForDomains(browser, registrationDomains, error_)
}

// Searches a domain for services of a specified type.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.2.
// Searches a domain for services of a specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceBrowserSearchForServices(_:_:_:_:)
func CFNetServiceBrowserSearchForServices(browser NetServiceBrowserRef, domain StringRef, serviceType StringRef, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFNetServiceBrowserSearchForServices(browser, domain, serviceType, error_)
}

// Stops a search for domains or services.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.2.
// Stops a search for domains or services.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceBrowserStopSearch(_:_:)
func CFNetServiceBrowserStopSearch(browser NetServiceBrowserRef, error_ unsafe.Pointer) {
	_CFNetServiceBrowserStopSearch(browser, error_)
}

// Unschedules a CFNetServiceBrowser from a run loop and mode.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.2.
// Unschedules a CFNetServiceBrowser from a run loop and mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceBrowserUnscheduleFromRunLoop(_:_:_:)
func CFNetServiceBrowserUnscheduleFromRunLoop(browser NetServiceBrowserRef, runLoop RunLoopRef, runLoopMode StringRef) {
	_CFNetServiceBrowserUnscheduleFromRunLoop(browser, runLoop, runLoopMode)
}

// Cancels a service registration or a service resolution.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.2.
// Cancels a service registration or a service resolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceCancel(_:)
func CFNetServiceCancel(theService NetServiceRef) {
	_CFNetServiceCancel(theService)
}

// Creates an instance of a Network Service object.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.2.
// Creates an instance of a Network Service object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceCreate(_:_:_:_:_:)
func CFNetServiceCreate(alloc AllocatorRef, domain StringRef, serviceType StringRef, name StringRef, port unsafe.Pointer) NetServiceRef {
	return _CFNetServiceCreate(alloc, domain, serviceType, name, port)
}

// Creates a copy of a CFNetService object.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.3.
// Creates a copy of a CFNetService object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceCreateCopy(_:_:)
func CFNetServiceCreateCopy(alloc AllocatorRef, service NetServiceRef) NetServiceRef {
	return _CFNetServiceCreateCopy(alloc, service)
}

// Uses TXT record data to create a dictionary.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Uses TXT record data to create a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceCreateDictionaryWithTXTData(_:_:)
func CFNetServiceCreateDictionaryWithTXTData(alloc AllocatorRef, txtRecord DataRef) DictionaryRef {
	return _CFNetServiceCreateDictionaryWithTXTData(alloc, txtRecord)
}

// Flattens a set of key/value pairs into a CFDataRef suitable for passing to .
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Flattens a set of key/value pairs into a CFDataRef suitable for passing to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceCreateTXTDataWithDictionary(_:_:)
func CFNetServiceCreateTXTDataWithDictionary(alloc AllocatorRef, keyValuePairs DictionaryRef) DataRef {
	return _CFNetServiceCreateTXTDataWithDictionary(alloc, keyValuePairs)
}

// Gets the IP addressing from a CFNetService.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.2.
// Gets the IP addressing from a CFNetService.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceGetAddressing(_:)
func CFNetServiceGetAddressing(theService NetServiceRef) ArrayRef {
	return _CFNetServiceGetAddressing(theService)
}

// Gets the domain from a CFNetService.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.2.
// Gets the domain from a CFNetService.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceGetDomain(_:)
func CFNetServiceGetDomain(theService NetServiceRef) StringRef {
	return _CFNetServiceGetDomain(theService)
}

// Gets the name from a CFNetService.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.2.
// Gets the name from a CFNetService.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceGetName(_:)
func CFNetServiceGetName(theService NetServiceRef) StringRef {
	return _CFNetServiceGetName(theService)
}

// This function gets the port number from a CFNetService.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.5.
// This function gets the port number from a CFNetService.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceGetPortNumber(_:)
func CFNetServiceGetPortNumber(theService NetServiceRef) unsafe.Pointer {
	return _CFNetServiceGetPortNumber(theService)
}

// Queries a network service for the contents of its TXT records.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Queries a network service for the contents of its TXT records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceGetTXTData(_:)
func CFNetServiceGetTXTData(theService NetServiceRef) DataRef {
	return _CFNetServiceGetTXTData(theService)
}

// Queries a CFNetService for its target hosts.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Queries a CFNetService for its target hosts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceGetTargetHost(_:)
func CFNetServiceGetTargetHost(theService NetServiceRef) StringRef {
	return _CFNetServiceGetTargetHost(theService)
}

// Gets the type from a CFNetService.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.2.
// Gets the type from a CFNetService.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceGetType(_:)
func CFNetServiceGetType(theService NetServiceRef) StringRef {
	return _CFNetServiceGetType(theService)
}

// Gets the Core Foundation type identifier for the Network Service object.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.2.
// Gets the Core Foundation type identifier for the Network Service object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceGetTypeID()
func CFNetServiceGetTypeID() TypeID {
	return _CFNetServiceGetTypeID()
}

// Creates an instance of a NetServiceMonitor object that watches for record changes.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Creates an instance of a NetServiceMonitor object that watches for record changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceMonitorCreate(_:_:_:_:)
func CFNetServiceMonitorCreate(alloc AllocatorRef, theService NetServiceRef, clientCB NetServiceMonitorClientCallBack, clientContext unsafe.Pointer) NetServiceMonitorRef {
	return _CFNetServiceMonitorCreate(alloc, theService, clientCB, clientContext)
}

// Gets the Core Foundation type identifier for all CFNetServiceMonitor instances.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Gets the Core Foundation type identifier for all CFNetServiceMonitor instances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceMonitorGetTypeID()
func CFNetServiceMonitorGetTypeID() TypeID {
	return _CFNetServiceMonitorGetTypeID()
}

// Invalidates an instance of a Network Service monitor object.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Invalidates an instance of a Network Service monitor object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceMonitorInvalidate(_:)
func CFNetServiceMonitorInvalidate(monitor NetServiceMonitorRef) {
	_CFNetServiceMonitorInvalidate(monitor)
}

// Schedules a CFNetServiceMonitor on a run loop.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Schedules a CFNetServiceMonitor on a run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceMonitorScheduleWithRunLoop(_:_:_:)
func CFNetServiceMonitorScheduleWithRunLoop(monitor NetServiceMonitorRef, runLoop RunLoopRef, runLoopMode StringRef) {
	_CFNetServiceMonitorScheduleWithRunLoop(monitor, runLoop, runLoopMode)
}

// Starts monitoring.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Starts monitoring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceMonitorStart(_:_:_:)
func CFNetServiceMonitorStart(monitor NetServiceMonitorRef, recordType NetServiceMonitorType, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFNetServiceMonitorStart(monitor, recordType, error_)
}

// Stops a CFNetServiceMonitor.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Stops a CFNetServiceMonitor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceMonitorStop(_:_:)
func CFNetServiceMonitorStop(monitor NetServiceMonitorRef, error_ unsafe.Pointer) {
	_CFNetServiceMonitorStop(monitor, error_)
}

// Unschedules a CFNetServiceMonitor from a run loop.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Unschedules a CFNetServiceMonitor from a run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceMonitorUnscheduleFromRunLoop(_:_:_:)
func CFNetServiceMonitorUnscheduleFromRunLoop(monitor NetServiceMonitorRef, runLoop RunLoopRef, runLoopMode StringRef) {
	_CFNetServiceMonitorUnscheduleFromRunLoop(monitor, runLoop, runLoopMode)
}

// Makes a CFNetService available on the network.
//
// Deprecated: This function was deprecated in macOS 10.4.
//
// Added in macOS 10.2.
// Makes a CFNetService available on the network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceRegister
func CFNetServiceRegister(theService NetServiceRef, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFNetServiceRegister(theService, error_)
}

// Makes a CFNetService available on the network.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Makes a CFNetService available on the network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceRegisterWithOptions(_:_:_:)
func CFNetServiceRegisterWithOptions(theService NetServiceRef, options OptionFlags, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFNetServiceRegisterWithOptions(theService, options, error_)
}

// This function updates the specified CFNetService with the IP address or addresses associated with the service. Call to get the addresses.
//
// Deprecated: This function was deprecated in macOS 10.4.
//
// Added in macOS 10.2.
// This function updates the specified CFNetService with the IP address or addresses associated with the service. Call to get the addresses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceResolve
func CFNetServiceResolve(theService NetServiceRef, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFNetServiceResolve(theService, error_)
}

// Gets the IP address or addresses for a CFNetService.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Gets the IP address or addresses for a CFNetService.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceResolveWithTimeout(_:_:_:)
func CFNetServiceResolveWithTimeout(theService NetServiceRef, timeout TimeInterval, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFNetServiceResolveWithTimeout(theService, timeout, error_)
}

// Schedules a CFNetService on a run loop.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.2.
// Schedules a CFNetService on a run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceScheduleWithRunLoop(_:_:_:)
func CFNetServiceScheduleWithRunLoop(theService NetServiceRef, runLoop RunLoopRef, runLoopMode StringRef) {
	_CFNetServiceScheduleWithRunLoop(theService, runLoop, runLoopMode)
}

// Associates a callback function with a CFNetService or disassociates a callback function from a CFNetService.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.2.
// Associates a callback function with a CFNetService or disassociates a callback function from a CFNetService.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceSetClient(_:_:_:)
func CFNetServiceSetClient(theService NetServiceRef, clientCB NetServiceClientCallBack, clientContext unsafe.Pointer) unsafe.Pointer {
	return _CFNetServiceSetClient(theService, clientCB, clientContext)
}

// Sets the TXT record for a CFNetService.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.4.
// Sets the TXT record for a CFNetService.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceSetTXTData(_:_:)
func CFNetServiceSetTXTData(theService NetServiceRef, txtRecord DataRef) unsafe.Pointer {
	return _CFNetServiceSetTXTData(theService, txtRecord)
}

// Unschedules a CFNetService from a run loop.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.2.
// Unschedules a CFNetService from a run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceUnscheduleFromRunLoop(_:_:_:)
func CFNetServiceUnscheduleFromRunLoop(theService NetServiceRef, runLoop RunLoopRef, runLoopMode StringRef) {
	_CFNetServiceUnscheduleFromRunLoop(theService, runLoop, runLoopMode)
}

// Executes a proxy autoconfiguration script to determine the best proxy to use to retrieve a specified URL.
//
// Added in macOS 10.5.
// Executes a proxy autoconfiguration script to determine the best proxy to use to retrieve a specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkCopyProxiesForAutoConfigurationScript(_:_:_:)
func CFNetworkCopyProxiesForAutoConfigurationScript(proxyAutoConfigurationScript StringRef, targetURL URLRef, error_ unsafe.Pointer) ArrayRef {
	return _CFNetworkCopyProxiesForAutoConfigurationScript(proxyAutoConfigurationScript, targetURL, error_)
}

// Returns the list of proxies that should be used to download a given URL.
//
// Added in macOS 10.5.
// Returns the list of proxies that should be used to download a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkCopyProxiesForURL(_:_:)
func CFNetworkCopyProxiesForURL(url URLRef, proxySettings DictionaryRef) ArrayRef {
	return _CFNetworkCopyProxiesForURL(url, proxySettings)
}

// Returns a CFDictionary containing the current systemwide internet proxy settings.
//
// Added in macOS 10.6.
// Returns a CFDictionary containing the current systemwide internet proxy settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkCopySystemProxySettings()
func CFNetworkCopySystemProxySettings() DictionaryRef {
	return _CFNetworkCopySystemProxySettings()
}

// Downloads a proxy autoconfiguration script and executes it.
//
// Added in macOS 10.5.
// Downloads a proxy autoconfiguration script and executes it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkExecuteProxyAutoConfigurationScript(_:_:_:_:)
func CFNetworkExecuteProxyAutoConfigurationScript(proxyAutoConfigurationScript StringRef, targetURL URLRef, cb ProxyAutoConfigurationResultCallback, clientContext unsafe.Pointer) RunLoopSourceRef {
	return _CFNetworkExecuteProxyAutoConfigurationScript(proxyAutoConfigurationScript, targetURL, cb, clientContext)
}

// Downloads a proxy autoconfiguration script and executes it.
//
// Added in macOS 10.5.
// Downloads a proxy autoconfiguration script and executes it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkExecuteProxyAutoConfigurationURL(_:_:_:_:)
func CFNetworkExecuteProxyAutoConfigurationURL(proxyAutoConfigURL URLRef, targetURL URLRef, cb ProxyAutoConfigurationResultCallback, clientContext unsafe.Pointer) RunLoopSourceRef {
	return _CFNetworkExecuteProxyAutoConfigurationURL(proxyAutoConfigURL, targetURL, cb, clientContext)
}

// Creates a read stream for a CFHTTP request message.
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.2.
// Creates a read stream for a CFHTTP request message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFReadStreamCreateForHTTPRequest(_:_:)
func CFReadStreamCreateForHTTPRequest(alloc AllocatorRef, request HTTPMessageRef) ReadStreamRef {
	return _CFReadStreamCreateForHTTPRequest(alloc, request)
}

// Creates a read stream for a CFHTTP request message object whose body is too long to keep in memory.
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.2.
// Creates a read stream for a CFHTTP request message object whose body is too long to keep in memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFReadStreamCreateForStreamedHTTPRequest(_:_:_:)
func CFReadStreamCreateForStreamedHTTPRequest(alloc AllocatorRef, requestHeaders HTTPMessageRef, requestBody ReadStreamRef) ReadStreamRef {
	return _CFReadStreamCreateForStreamedHTTPRequest(alloc, requestHeaders, requestBody)
}

// Creates an FTP read stream.
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.3.
// Creates an FTP read stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFReadStreamCreateWithFTPURL(_:_:)
func CFReadStreamCreateWithFTPURL(alloc AllocatorRef, ftpURL URLRef) ReadStreamRef {
	return _CFReadStreamCreateWithFTPURL(alloc, ftpURL)
}

// Creates readable and writable streams connected to a given object.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.3.
// Creates readable and writable streams connected to a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFStreamCreatePairWithSocketToCFHost(_:_:_:_:_:)
func CFStreamCreatePairWithSocketToCFHost(alloc AllocatorRef, host HostRef, port unsafe.Pointer, readStream unsafe.Pointer, writeStream unsafe.Pointer) {
	_CFStreamCreatePairWithSocketToCFHost(alloc, host, port, readStream, writeStream)
}

// Creates a pair of streams for a CFNetService.
//
// Deprecated: This function was deprecated in macOS 15.0.
//
// Added in macOS 10.3.
// Creates a pair of streams for a CFNetService.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFStreamCreatePairWithSocketToNetService(_:_:_:_:)
func CFStreamCreatePairWithSocketToNetService(alloc AllocatorRef, service NetServiceRef, readStream unsafe.Pointer, writeStream unsafe.Pointer) {
	_CFStreamCreatePairWithSocketToNetService(alloc, service, readStream, writeStream)
}

// Creates an FTP write stream.
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.3.
// Creates an FTP write stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFWriteStreamCreateWithFTPURL(_:_:)
func CFWriteStreamCreateWithFTPURL(alloc AllocatorRef, ftpURL URLRef) WriteStreamRef {
	return _CFWriteStreamCreateWithFTPURL(alloc, ftpURL)
}





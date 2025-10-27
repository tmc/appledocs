// Code generated from Apple documentation for CFNetwork. DO NOT EDIT.

package cfnetwork


// Enum types and constants

// CFHostInfoType - Values indicating the type of data that is to be resolved or the type of data that was resolved.
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHostInfoType
type CFHostInfoType uint

const (
	// kCFHostAddresses - Specifies that addresses are to be resolved or that addresses were resolved.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHostInfoType/addresses
	kCFHostAddresses CFHostInfoType = 0
	// kCFHostNames - Specifies that names are to be resolved or that names were resolved.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHostInfoType/names
	kCFHostNames CFHostInfoType = 0
	// kCFHostReachability - Specifies that reachability information is to be resolved or that reachability information was resolved.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHostInfoType/reachability
	kCFHostReachability CFHostInfoType = 0
)


// CFNetDiagnosticStatusValues - Constants for diagnostic status values.
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetDiagnosticStatusValues
type CFNetDiagnosticStatusValues uint

const (
	// kCFNetDiagnosticConnectionDown - The connection does not appear to be working.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetDiagnosticStatusValues/connectionDown
	kCFNetDiagnosticConnectionDown CFNetDiagnosticStatusValues = 0
	// kCFNetDiagnosticConnectionIndeterminate - The status of the connection is not known.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetDiagnosticStatusValues/connectionIndeterminate
	kCFNetDiagnosticConnectionIndeterminate CFNetDiagnosticStatusValues = 0
	// kCFNetDiagnosticConnectionUp - The connection appears to be working.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetDiagnosticStatusValues/connectionUp
	kCFNetDiagnosticConnectionUp CFNetDiagnosticStatusValues = 0
	// kCFNetDiagnosticErr - An error occurred that prevented the call from completing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetDiagnosticStatusValues/err
	kCFNetDiagnosticErr CFNetDiagnosticStatusValues = 0
	// kCFNetDiagnosticNoErr - No error occurred but there is no status.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetDiagnosticStatusValues/noErr
	kCFNetDiagnosticNoErr CFNetDiagnosticStatusValues = 0
)


// CFNetServiceBrowserFlags - Flags that the system passes to net service browser callbacks.
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceBrowserFlags
type CFNetServiceBrowserFlags uint

const (
	// kCFNetServiceFlagIsDefault - Specifies whether the resulting domain is the default registration or browse domain.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceBrowserFlags/isDefault
	kCFNetServiceFlagIsDefault CFNetServiceBrowserFlags = 0
	// kCFNetServiceFlagIsDomain - Specifies whether the result pertains to a search for domains or services.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceBrowserFlags/isDomain
	kCFNetServiceFlagIsDomain CFNetServiceBrowserFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceBrowserFlags/isRegistrationDomain
	kCFNetServiceFlagIsRegistrationDomain CFNetServiceBrowserFlags = 0
	// kCFNetServiceFlagMoreComing - A hint that the system will call the client’s callback function again soon.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceBrowserFlags/moreComing
	kCFNetServiceFlagMoreComing CFNetServiceBrowserFlags = 0
	// kCFNetServiceFlagRemove - Specifies whether the client should remove the result instead of adding it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceBrowserFlags/remove
	kCFNetServiceFlagRemove CFNetServiceBrowserFlags = 0
)


// CFNetServiceMonitorType - Record type specifier used to tell a service monitor the type of record changes to watch for.
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceMonitorType
type CFNetServiceMonitorType uint

const (
	// kCFNetServiceMonitorTXT - Watch for TXT record changes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceMonitorType/TXT
	kCFNetServiceMonitorTXT CFNetServiceMonitorType = 0
)


// CFNetServiceRegisterFlags - Options to use when registering a service on the network.
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceRegisterFlags
type CFNetServiceRegisterFlags uint

const (
	// kCFNetServiceFlagNoAutoRename - Causes registrations to fail if a name conflict occurs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceRegisterFlags/noAutoRename
	kCFNetServiceFlagNoAutoRename CFNetServiceRegisterFlags = 0
)


// CFNetServicesError - Error codes that may be returned by CFNetServices functions or passed to CFNetServices callback functions.
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServicesError
type CFNetServicesError uint

const (
	// kCFNetServicesErrorBadArgument - A required argument was not provided.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServicesError/badArgument
	kCFNetServicesErrorBadArgument CFNetServicesError = 0
	// kCFNetServicesErrorCancel - The search or service was canceled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServicesError/cancel
	kCFNetServicesErrorCancel CFNetServicesError = 0
	// kCFNetServicesErrorCollision - An attempt was made to use a name that is already in use.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServicesError/collision
	kCFNetServicesErrorCollision CFNetServicesError = 0
	// kCFNetServicesErrorInProgress - A search is already in progress.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServicesError/inProgress
	kCFNetServicesErrorInProgress CFNetServicesError = 0
	// kCFNetServicesErrorInvalid - Invalid data was passed to a CFNetServices function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServicesError/invalid
	kCFNetServicesErrorInvalid CFNetServicesError = 0
	// kCFNetServicesErrorMissingRequiredConfiguration - A required configuration for local network access is missing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServicesError/missingRequiredConfiguration
	kCFNetServicesErrorMissingRequiredConfiguration CFNetServicesError = 0
	// kCFNetServicesErrorNotFound - Not used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServicesError/notFound
	kCFNetServicesErrorNotFound CFNetServicesError = 0
	// kCFNetServicesErrorTimeout - Resolution failed because the timeout was reached.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServicesError/timeout
	kCFNetServicesErrorTimeout CFNetServicesError = 0
	// kCFNetServicesErrorUnknown - An unknown CFNetService error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServicesError/unknown
	kCFNetServicesErrorUnknown CFNetServicesError = 0
)


// CFNetworkErrors - This enumeration contains error codes returned under the error domain 
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors
type CFNetworkErrors uint

const (
	// kCFErrorHTTPAuthenticationTypeUnsupported - The client and server couldn’t agree on a supported authentication type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfErrorHTTPAuthenticationTypeUnsupported
	kCFErrorHTTPAuthenticationTypeUnsupported CFNetworkErrors = 0
	// kCFErrorHTTPBadCredentials - The server rejected the credentials provided for an authenticated connection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfErrorHTTPBadCredentials
	kCFErrorHTTPBadCredentials CFNetworkErrors = 0
	// kCFErrorHTTPBadProxyCredentials - The proxy rejected the authentication credentials provided for logging in.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfErrorHTTPBadProxyCredentials
	kCFErrorHTTPBadProxyCredentials CFNetworkErrors = 0
	// kCFErrorHTTPBadURL - The requested URL couldn’t be retrieved.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfErrorHTTPBadURL
	kCFErrorHTTPBadURL CFNetworkErrors = 0
	// kCFErrorHTTPConnectionLost - The connection to the server was dropped.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfErrorHTTPConnectionLost
	kCFErrorHTTPConnectionLost CFNetworkErrors = 0
	// kCFErrorHTTPParseFailure - The HTTP server response couldn’t be parsed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfErrorHTTPParseFailure
	kCFErrorHTTPParseFailure CFNetworkErrors = 0
	// kCFErrorHTTPProxyConnectionFailure - A connection to the HTTPS proxy couldn’t be established.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfErrorHTTPProxyConnectionFailure
	kCFErrorHTTPProxyConnectionFailure CFNetworkErrors = 0
	// kCFErrorHTTPRedirectionLoopDetected - Too many HTTP redirects occurred before reaching a page that didn’t redirect the client to another page.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfErrorHTTPRedirectionLoopDetected
	kCFErrorHTTPRedirectionLoopDetected CFNetworkErrors = 0
	// kCFErrorHTTPSProxyConnectionFailure - A connection couldn’t be established to the HTTPS proxy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfErrorHTTPSProxyConnectionFailure
	kCFErrorHTTPSProxyConnectionFailure CFNetworkErrors = 0
	// kCFErrorPACFileAuth - The authentication credentials provided by the proxy autoconfiguration file were rejected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfErrorPACFileAuth
	kCFErrorPACFileAuth CFNetworkErrors = 0
	// kCFErrorPACFileError - An error occurred with the proxy autoconfiguration file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfErrorPACFileError
	kCFErrorPACFileError CFNetworkErrors = 0
	// kCFFTPErrorUnexpectedStatusCode - The server returned an unexpected status code.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfftpErrorUnexpectedStatusCode
	kCFFTPErrorUnexpectedStatusCode CFNetworkErrors = 0
	// kCFHostErrorHostNotFound - The specified host wasn’t found.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfHostErrorHostNotFound
	kCFHostErrorHostNotFound CFNetworkErrors = 0
	// kCFHostErrorUnknown - An unknown error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfHostErrorUnknown
	kCFHostErrorUnknown CFNetworkErrors = 0
	// kCFHTTPCookieCannotParseCookieFile - The cookie file can’t be parsed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfhttpCookieCannotParseCookieFile
	kCFHTTPCookieCannotParseCookieFile CFNetworkErrors = 0
	// kCFNetServiceErrorBadArgument - A required argument either wasn’t provided or wasn’t valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfNetServiceErrorBadArgument
	kCFNetServiceErrorBadArgument CFNetworkErrors = 0
	// kCFNetServiceErrorCancel - The search or service was canceled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfNetServiceErrorCancel
	kCFNetServiceErrorCancel CFNetworkErrors = 0
	// kCFNetServiceErrorCollision - An attempt was made to use a name that’s already in use.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfNetServiceErrorCollision
	kCFNetServiceErrorCollision CFNetworkErrors = 0
	// kCFNetServiceErrorDNSServiceFailure - The DNS service discovery returned an error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfNetServiceErrorDNSServiceFailure
	kCFNetServiceErrorDNSServiceFailure CFNetworkErrors = 0
	// kCFNetServiceErrorInProgress - A new search couldn’t be started because a search is already in progress.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfNetServiceErrorInProgress
	kCFNetServiceErrorInProgress CFNetworkErrors = 0
	// kCFNetServiceErrorInvalid - Invalid data was passed to a CFNetServices function.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfNetServiceErrorInvalid
	kCFNetServiceErrorInvalid CFNetworkErrors = 0
	// kCFNetServiceErrorNotFound - This error isn’t used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfNetServiceErrorNotFound
	kCFNetServiceErrorNotFound CFNetworkErrors = 0
	// kCFNetServiceErrorTimeout - Resolution failed because the timeout was reached.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfNetServiceErrorTimeout
	kCFNetServiceErrorTimeout CFNetworkErrors = 0
	// kCFNetServiceErrorUnknown - An error of unknown type has occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfNetServiceErrorUnknown
	kCFNetServiceErrorUnknown CFNetworkErrors = 0
	// kCFSOCKS4ErrorIdConflict - The server rejected the request because the client program and the   daemon reported different user IDs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfsocks4ErrorIdConflict
	kCFSOCKS4ErrorIdConflict CFNetworkErrors = 0
	// kCFSOCKS4ErrorIdentdFailed - The server couldn’t connect to the   daemon on the client and rejected the request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfsocks4ErrorIdentdFailed
	kCFSOCKS4ErrorIdentdFailed CFNetworkErrors = 0
	// kCFSOCKS4ErrorRequestFailed - The server rejected the request, or the request failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfsocks4ErrorRequestFailed
	kCFSOCKS4ErrorRequestFailed CFNetworkErrors = 0
	// kCFSOCKS4ErrorUnknownStatusCode - The server returned an unknown status code.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfsocks4ErrorUnknownStatusCode
	kCFSOCKS4ErrorUnknownStatusCode CFNetworkErrors = 0
	// kCFSOCKS5ErrorBadCredentials - The SOCKS server refused the client connection because of bad login credentials.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfsocks5ErrorBadCredentials
	kCFSOCKS5ErrorBadCredentials CFNetworkErrors = 0
	// kCFSOCKS5ErrorBadResponseAddr - The address type returned isn’t supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfsocks5ErrorBadResponseAddr
	kCFSOCKS5ErrorBadResponseAddr CFNetworkErrors = 0
	// kCFSOCKS5ErrorBadState - The stream isn’t in a state that allows the requested operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfsocks5ErrorBadState
	kCFSOCKS5ErrorBadState CFNetworkErrors = 0
	// kCFSOCKS5ErrorNoAcceptableMethod - The client and server couldn’t find a mutually agreeable authentication method.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfsocks5ErrorNoAcceptableMethod
	kCFSOCKS5ErrorNoAcceptableMethod CFNetworkErrors = 0
	// kCFSOCKS5ErrorUnsupportedNegotiationMethod - The requested method isn’t supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfsocks5ErrorUnsupportedNegotiationMethod
	kCFSOCKS5ErrorUnsupportedNegotiationMethod CFNetworkErrors = 0
	// kCFSOCKSErrorUnknownClientVersion - The SOCKS server rejected access because it doesn’t support connections with the requested SOCKS version.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfsocksErrorUnknownClientVersion
	kCFSOCKSErrorUnknownClientVersion CFNetworkErrors = 0
	// kCFSOCKSErrorUnsupportedServerVersion - The SOCKS server doesn’t support the requested version.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfsocksErrorUnsupportedServerVersion
	kCFSOCKSErrorUnsupportedServerVersion CFNetworkErrors = 0
	// kCFStreamErrorHTTPSProxyFailureUnexpectedResponseToCONNECTMethod - The HTTPS proxy returned an unexpected status code, such as a   redirect.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfStreamErrorHTTPSProxyFailureUnexpectedResponseToCONNECTMethod
	kCFStreamErrorHTTPSProxyFailureUnexpectedResponseToCONNECTMethod CFNetworkErrors = 0
	// kCFURLErrorAppTransportSecurityRequiresSecureConnection - The connection failed because the App Transport Security configuration requires a secure connection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorAppTransportSecurityRequiresSecureConnection
	kCFURLErrorAppTransportSecurityRequiresSecureConnection CFNetworkErrors = 0
	// kCFURLErrorBackgroundSessionInUseByAnotherProcess - The background session failed because it was in use by another process.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorBackgroundSessionInUseByAnotherProcess
	kCFURLErrorBackgroundSessionInUseByAnotherProcess CFNetworkErrors = 0
	// kCFURLErrorBackgroundSessionWasDisconnected - The background session failed because it was disconnected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorBackgroundSessionWasDisconnected
	kCFURLErrorBackgroundSessionWasDisconnected CFNetworkErrors = 0
	// kCFURLErrorBadServerResponse - The connection received an invalid server response.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorBadServerResponse
	kCFURLErrorBadServerResponse CFNetworkErrors = 0
	// kCFURLErrorBadURL - The connection failed due to a malformed URL.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorBadURL
	kCFURLErrorBadURL CFNetworkErrors = 0
	// kCFURLErrorCallIsActive - The connection failed because a call is active.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorCallIsActive
	kCFURLErrorCallIsActive CFNetworkErrors = 0
	// kCFURLErrorCancelled - The connection was cancelled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorCancelled
	kCFURLErrorCancelled CFNetworkErrors = 0
	// kCFURLErrorCannotCloseFile - The file can’t be closed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorCannotCloseFile
	kCFURLErrorCannotCloseFile CFNetworkErrors = 0
	// kCFURLErrorCannotConnectToHost - The connection failed because a connection can’t be made to the host.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorCannotConnectToHost
	kCFURLErrorCannotConnectToHost CFNetworkErrors = 0
	// kCFURLErrorCannotCreateFile - The file can’t be created.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorCannotCreateFile
	kCFURLErrorCannotCreateFile CFNetworkErrors = 0
	// kCFURLErrorCannotDecodeContentData - The connection can’t decode data encoded with an unknown content encoding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorCannotDecodeContentData
	kCFURLErrorCannotDecodeContentData CFNetworkErrors = 0
	// kCFURLErrorCannotDecodeRawData - The connection can’t decode data encoded with a known content encoding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorCannotDecodeRawData
	kCFURLErrorCannotDecodeRawData CFNetworkErrors = 0
	// kCFURLErrorCannotFindHost - The connection failed because the host couldn’t be found.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorCannotFindHost
	kCFURLErrorCannotFindHost CFNetworkErrors = 0
	// kCFURLErrorCannotLoadFromNetwork - The connection failed because it’s being required to return a cached resource, but one isn’t available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorCannotLoadFromNetwork
	kCFURLErrorCannotLoadFromNetwork CFNetworkErrors = 0
	// kCFURLErrorCannotMoveFile - The file can’t be moved.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorCannotMoveFile
	kCFURLErrorCannotMoveFile CFNetworkErrors = 0
	// kCFURLErrorCannotOpenFile - The file can’t be opened.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorCannotOpenFile
	kCFURLErrorCannotOpenFile CFNetworkErrors = 0
	// kCFURLErrorCannotParseResponse - The connection can’t parse the server’s response.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorCannotParseResponse
	kCFURLErrorCannotParseResponse CFNetworkErrors = 0
	// kCFURLErrorCannotRemoveFile - The file can’t be removed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorCannotRemoveFile
	kCFURLErrorCannotRemoveFile CFNetworkErrors = 0
	// kCFURLErrorCannotWriteToFile - The file can’t be written.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorCannotWriteToFile
	kCFURLErrorCannotWriteToFile CFNetworkErrors = 0
	// kCFURLErrorClientCertificateRejected - The secure connection failed because the client’s certificate was rejected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorClientCertificateRejected
	kCFURLErrorClientCertificateRejected CFNetworkErrors = 0
	// kCFURLErrorClientCertificateRequired - The secure connection failed because the server requires a client certificate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorClientCertificateRequired
	kCFURLErrorClientCertificateRequired CFNetworkErrors = 0
	// kCFURLErrorDataLengthExceedsMaximum - The file operation failed because the file is too large.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorDataLengthExceedsMaximum
	kCFURLErrorDataLengthExceedsMaximum CFNetworkErrors = 0
	// kCFURLErrorDataNotAllowed - The connection failed because data use isn’t currently allowed on the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorDataNotAllowed
	kCFURLErrorDataNotAllowed CFNetworkErrors = 0
	// kCFURLErrorDNSLookupFailed - The connection failed because the DNS lookup failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorDNSLookupFailed
	kCFURLErrorDNSLookupFailed CFNetworkErrors = 0
	// kCFURLErrorDownloadDecodingFailedMidStream - The download failed because decoding of the downloaded data failed midstream.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorDownloadDecodingFailedMidStream
	kCFURLErrorDownloadDecodingFailedMidStream CFNetworkErrors = 0
	// kCFURLErrorDownloadDecodingFailedToComplete - The download failed because decoding of the downloaded data failed to complete.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorDownloadDecodingFailedToComplete
	kCFURLErrorDownloadDecodingFailedToComplete CFNetworkErrors = 0
	// kCFURLErrorFileDoesNotExist - The file operation failed because the file doesn’t exist.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorFileDoesNotExist
	kCFURLErrorFileDoesNotExist CFNetworkErrors = 0
	// kCFURLErrorFileIsDirectory - The file operation failed because the file is a directory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorFileIsDirectory
	kCFURLErrorFileIsDirectory CFNetworkErrors = 0
	// kCFURLErrorFileOutsideSafeArea - The file is outside of the safe area.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorFileOutsideSafeArea
	kCFURLErrorFileOutsideSafeArea CFNetworkErrors = 0
	// kCFURLErrorHTTPTooManyRedirects - The HTTP connection failed due to too many redirects.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorHTTPTooManyRedirects
	kCFURLErrorHTTPTooManyRedirects CFNetworkErrors = 0
	// kCFURLErrorInternationalRoamingOff - The connection failed because international roaming is disabled on the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorInternationalRoamingOff
	kCFURLErrorInternationalRoamingOff CFNetworkErrors = 0
	// kCFURLErrorNetworkConnectionLost - The connection failed because the network connection was lost.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorNetworkConnectionLost
	kCFURLErrorNetworkConnectionLost CFNetworkErrors = 0
	// kCFURLErrorNoPermissionsToReadFile - The file operation failed because it doesn’t have permission to read the file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorNoPermissionsToReadFile
	kCFURLErrorNoPermissionsToReadFile CFNetworkErrors = 0
	// kCFURLErrorNotConnectedToInternet - The connection failed because the device isn’t connected to the internet.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorNotConnectedToInternet
	kCFURLErrorNotConnectedToInternet CFNetworkErrors = 0
	// kCFURLErrorRedirectToNonExistentLocation - The connection was redirected to a nonexistent location.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorRedirectToNonExistentLocation
	kCFURLErrorRedirectToNonExistentLocation CFNetworkErrors = 0
	// kCFURLErrorRequestBodyStreamExhausted - The connection failed because the request’s body stream was exhausted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorRequestBodyStreamExhausted
	kCFURLErrorRequestBodyStreamExhausted CFNetworkErrors = 0
	// kCFURLErrorResourceUnavailable - The connection’s resource is unavailable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorResourceUnavailable
	kCFURLErrorResourceUnavailable CFNetworkErrors = 0
	// kCFURLErrorSecureConnectionFailed - The secure connection failed for an unknown reason.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorSecureConnectionFailed
	kCFURLErrorSecureConnectionFailed CFNetworkErrors = 0
	// kCFURLErrorServerCertificateHasBadDate - The secure connection failed because the server’s certificate has an invalid date.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorServerCertificateHasBadDate
	kCFURLErrorServerCertificateHasBadDate CFNetworkErrors = 0
	// kCFURLErrorServerCertificateHasUnknownRoot - The secure connection failed because the server’s certificate has an unknown root.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorServerCertificateHasUnknownRoot
	kCFURLErrorServerCertificateHasUnknownRoot CFNetworkErrors = 0
	// kCFURLErrorServerCertificateNotYetValid - The secure connection failed because the server’s certificate isn’t valid yet.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorServerCertificateNotYetValid
	kCFURLErrorServerCertificateNotYetValid CFNetworkErrors = 0
	// kCFURLErrorServerCertificateUntrusted - The secure connection failed because the server’s certificate isn’t trusted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorServerCertificateUntrusted
	kCFURLErrorServerCertificateUntrusted CFNetworkErrors = 0
	// kCFURLErrorTimedOut - The connection timed out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorTimedOut
	kCFURLErrorTimedOut CFNetworkErrors = 0
	// kCFURLErrorUnknown - An unknown error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorUnknown
	kCFURLErrorUnknown CFNetworkErrors = 0
	// kCFURLErrorUnsupportedURL - The connection failed due to an unsupported URL scheme.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorUnsupportedURL
	kCFURLErrorUnsupportedURL CFNetworkErrors = 0
	// kCFURLErrorUserAuthenticationRequired - The connection failed because it requires authentication.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorUserAuthenticationRequired
	kCFURLErrorUserAuthenticationRequired CFNetworkErrors = 0
	// kCFURLErrorUserCancelledAuthentication - The connection failed because the user cancelled required authentication.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorUserCancelledAuthentication
	kCFURLErrorUserCancelledAuthentication CFNetworkErrors = 0
	// kCFURLErrorZeroByteResource - The resource retrieved by the connection is zero bytes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorZeroByteResource
	kCFURLErrorZeroByteResource CFNetworkErrors = 0
)


// CFStreamErrorHTTP - Error codes that a read stream for an HTTP request may return.
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFStreamErrorHTTP
type CFStreamErrorHTTP uint

const (
	// kCFStreamErrorHTTPBadURL - The URL is not properly formatted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFStreamErrorHTTP/badURL
	kCFStreamErrorHTTPBadURL CFStreamErrorHTTP = 0
	// kCFStreamErrorHTTPParseFailure - A parsing error occurred while an incoming message was being deserialized and appended to a message object. The headers of the incoming message may be formatted improperly.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFStreamErrorHTTP/parseFailure
	kCFStreamErrorHTTPParseFailure CFStreamErrorHTTP = 0
	// kCFStreamErrorHTTPRedirectionLoop - A redirection loop has been detected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFStreamErrorHTTP/redirectionLoop
	kCFStreamErrorHTTPRedirectionLoop CFStreamErrorHTTP = 0
)


// CFStreamErrorHTTPAuthentication - Authentication error codes that may be returned when trying to apply authentication to a request.
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFStreamErrorHTTPAuthentication
type CFStreamErrorHTTPAuthentication uint

const (
	// kCFStreamErrorHTTPAuthenticationBadPassword - Password is in a format that is not suitable for the request. Currently, passwords are decoded using  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFStreamErrorHTTPAuthentication/badPassword
	kCFStreamErrorHTTPAuthenticationBadPassword CFStreamErrorHTTPAuthentication = 0
	// kCFStreamErrorHTTPAuthenticationBadUserName - User name is in a format that is not suitable for the request. Currently, user names are decoded using  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFStreamErrorHTTPAuthentication/badUserName
	kCFStreamErrorHTTPAuthenticationBadUserName CFStreamErrorHTTPAuthentication = 0
	// kCFStreamErrorHTTPAuthenticationTypeUnsupported - Specified authentication type is not supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFStreamErrorHTTPAuthentication/typeUnsupported
	kCFStreamErrorHTTPAuthenticationTypeUnsupported CFStreamErrorHTTPAuthentication = 0
)



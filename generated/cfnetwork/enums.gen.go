// Code generated from Apple documentation for CFNetwork. DO NOT EDIT.

package cfnetwork

// Enum types and constants
// CFHostInfoType - Values indicating the type of data that is to be resolved or the type of data that was resolved.
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHostInfoType
type HostInfoType uint

const (
// kCFHostAddresses - Specifies that addresses are to be resolved or that addresses were resolved.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHostInfoType/addresses
kCFHostAddresses HostInfoType = 0
// kCFHostNames - Specifies that names are to be resolved or that names were resolved.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHostInfoType/names
kCFHostNames HostInfoType = 0
// kCFHostReachability - Specifies that reachability information is to be resolved or that reachability information was resolved.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFHostInfoType/reachability
kCFHostReachability HostInfoType = 0
)

// CFNetDiagnosticStatusValues - Constants for diagnostic status values.
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetDiagnosticStatusValues
type NetDiagnosticStatusValues uint

const (
// kCFNetDiagnosticConnectionDown - The connection does not appear to be working.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetDiagnosticStatusValues/connectionDown
kCFNetDiagnosticConnectionDown NetDiagnosticStatusValues = 0
// kCFNetDiagnosticConnectionIndeterminate - The status of the connection is not known.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetDiagnosticStatusValues/connectionIndeterminate
kCFNetDiagnosticConnectionIndeterminate NetDiagnosticStatusValues = 0
// kCFNetDiagnosticConnectionUp - The connection appears to be working.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetDiagnosticStatusValues/connectionUp
kCFNetDiagnosticConnectionUp NetDiagnosticStatusValues = 0
// kCFNetDiagnosticErr - An error occurred that prevented the call from completing.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetDiagnosticStatusValues/err
kCFNetDiagnosticErr NetDiagnosticStatusValues = 0
// kCFNetDiagnosticNoErr - No error occurred but there is no status.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetDiagnosticStatusValues/noErr
kCFNetDiagnosticNoErr NetDiagnosticStatusValues = 0
)

// CFNetServiceBrowserFlags - Flags that the system passes to net service browser callbacks.
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceBrowserFlags
type NetServiceBrowserFlags uint

const (
// kCFNetServiceFlagIsDefault - Specifies whether the resulting domain is the default registration or browse domain.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceBrowserFlags/isDefault
kCFNetServiceFlagIsDefault NetServiceBrowserFlags = 0
// kCFNetServiceFlagIsDomain - Specifies whether the result pertains to a search for domains or services.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceBrowserFlags/isDomain
kCFNetServiceFlagIsDomain NetServiceBrowserFlags = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceBrowserFlags/isRegistrationDomain
kCFNetServiceFlagIsRegistrationDomain NetServiceBrowserFlags = 0
// kCFNetServiceFlagMoreComing - A hint that the system will call the client’s callback function again soon.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceBrowserFlags/moreComing
kCFNetServiceFlagMoreComing NetServiceBrowserFlags = 0
// kCFNetServiceFlagRemove - Specifies whether the client should remove the result instead of adding it.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceBrowserFlags/remove
kCFNetServiceFlagRemove NetServiceBrowserFlags = 0
)

// CFNetServiceMonitorType - Record type specifier used to tell a service monitor the type of record changes to watch for.
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceMonitorType
type NetServiceMonitorType uint

const (
// kCFNetServiceMonitorTXT - Watch for TXT record changes.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceMonitorType/TXT
kCFNetServiceMonitorTXT NetServiceMonitorType = 0
)

// CFNetServiceRegisterFlags - Options to use when registering a service on the network.
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceRegisterFlags
type NetServiceRegisterFlags uint

const (
// kCFNetServiceFlagNoAutoRename - Causes registrations to fail if a name conflict occurs.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServiceRegisterFlags/noAutoRename
kCFNetServiceFlagNoAutoRename NetServiceRegisterFlags = 0
)

// CFNetServicesError - Error codes that may be returned by CFNetServices functions or passed to CFNetServices callback functions.
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServicesError
type NetServicesError uint

const (
// kCFNetServicesErrorBadArgument - A required argument was not provided.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServicesError/badArgument
kCFNetServicesErrorBadArgument NetServicesError = 0
// kCFNetServicesErrorCancel - The search or service was canceled.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServicesError/cancel
kCFNetServicesErrorCancel NetServicesError = 0
// kCFNetServicesErrorCollision - An attempt was made to use a name that is already in use.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServicesError/collision
kCFNetServicesErrorCollision NetServicesError = 0
// kCFNetServicesErrorInProgress - A search is already in progress.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServicesError/inProgress
kCFNetServicesErrorInProgress NetServicesError = 0
// kCFNetServicesErrorInvalid - Invalid data was passed to a CFNetServices function.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServicesError/invalid
kCFNetServicesErrorInvalid NetServicesError = 0
// kCFNetServicesErrorMissingRequiredConfiguration - A required configuration for local network access is missing.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServicesError/missingRequiredConfiguration
kCFNetServicesErrorMissingRequiredConfiguration NetServicesError = 0
// kCFNetServicesErrorNotFound - Not used.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServicesError/notFound
kCFNetServicesErrorNotFound NetServicesError = 0
// kCFNetServicesErrorTimeout - Resolution failed because the timeout was reached.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServicesError/timeout
kCFNetServicesErrorTimeout NetServicesError = 0
// kCFNetServicesErrorUnknown - An unknown CFNetService error occurred.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetServicesError/unknown
kCFNetServicesErrorUnknown NetServicesError = 0
)

// CFNetworkErrors - This enumeration contains error codes returned under the error domain 
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors
type NetworkErrors uint

const (
// kCFErrorHTTPAuthenticationTypeUnsupported - The client and server couldn’t agree on a supported authentication type.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfErrorHTTPAuthenticationTypeUnsupported
kCFErrorHTTPAuthenticationTypeUnsupported NetworkErrors = 0
// kCFErrorHTTPBadCredentials - The server rejected the credentials provided for an authenticated connection.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfErrorHTTPBadCredentials
kCFErrorHTTPBadCredentials NetworkErrors = 0
// kCFErrorHTTPBadProxyCredentials - The proxy rejected the authentication credentials provided for logging in.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfErrorHTTPBadProxyCredentials
kCFErrorHTTPBadProxyCredentials NetworkErrors = 0
// kCFErrorHTTPBadURL - The requested URL couldn’t be retrieved.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfErrorHTTPBadURL
kCFErrorHTTPBadURL NetworkErrors = 0
// kCFErrorHTTPConnectionLost - The connection to the server was dropped.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfErrorHTTPConnectionLost
kCFErrorHTTPConnectionLost NetworkErrors = 0
// kCFErrorHTTPParseFailure - The HTTP server response couldn’t be parsed.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfErrorHTTPParseFailure
kCFErrorHTTPParseFailure NetworkErrors = 0
// kCFErrorHTTPProxyConnectionFailure - A connection to the HTTPS proxy couldn’t be established.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfErrorHTTPProxyConnectionFailure
kCFErrorHTTPProxyConnectionFailure NetworkErrors = 0
// kCFErrorHTTPRedirectionLoopDetected - Too many HTTP redirects occurred before reaching a page that didn’t redirect the client to another page.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfErrorHTTPRedirectionLoopDetected
kCFErrorHTTPRedirectionLoopDetected NetworkErrors = 0
// kCFErrorHTTPSProxyConnectionFailure - A connection couldn’t be established to the HTTPS proxy.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfErrorHTTPSProxyConnectionFailure
kCFErrorHTTPSProxyConnectionFailure NetworkErrors = 0
// kCFErrorPACFileAuth - The authentication credentials provided by the proxy autoconfiguration file were rejected.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfErrorPACFileAuth
kCFErrorPACFileAuth NetworkErrors = 0
// kCFErrorPACFileError - An error occurred with the proxy autoconfiguration file.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfErrorPACFileError
kCFErrorPACFileError NetworkErrors = 0
// kCFHostErrorHostNotFound - The specified host wasn’t found.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfHostErrorHostNotFound
kCFHostErrorHostNotFound NetworkErrors = 0
// kCFHostErrorUnknown - An unknown error.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfHostErrorUnknown
kCFHostErrorUnknown NetworkErrors = 0
// kCFNetServiceErrorBadArgument - A required argument either wasn’t provided or wasn’t valid.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfNetServiceErrorBadArgument
kCFNetServiceErrorBadArgument NetworkErrors = 0
// kCFNetServiceErrorCancel - The search or service was canceled.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfNetServiceErrorCancel
kCFNetServiceErrorCancel NetworkErrors = 0
// kCFNetServiceErrorCollision - An attempt was made to use a name that’s already in use.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfNetServiceErrorCollision
kCFNetServiceErrorCollision NetworkErrors = 0
// kCFNetServiceErrorDNSServiceFailure - The DNS service discovery returned an error.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfNetServiceErrorDNSServiceFailure
kCFNetServiceErrorDNSServiceFailure NetworkErrors = 0
// kCFNetServiceErrorInProgress - A new search couldn’t be started because a search is already in progress.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfNetServiceErrorInProgress
kCFNetServiceErrorInProgress NetworkErrors = 0
// kCFNetServiceErrorInvalid - Invalid data was passed to a CFNetServices function.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfNetServiceErrorInvalid
kCFNetServiceErrorInvalid NetworkErrors = 0
// kCFNetServiceErrorNotFound - This error isn’t used.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfNetServiceErrorNotFound
kCFNetServiceErrorNotFound NetworkErrors = 0
// kCFNetServiceErrorTimeout - Resolution failed because the timeout was reached.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfNetServiceErrorTimeout
kCFNetServiceErrorTimeout NetworkErrors = 0
// kCFNetServiceErrorUnknown - An error of unknown type has occurred.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfNetServiceErrorUnknown
kCFNetServiceErrorUnknown NetworkErrors = 0
// kCFStreamErrorHTTPSProxyFailureUnexpectedResponseToCONNECTMethod - The HTTPS proxy returned an unexpected status code, such as a   redirect.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfStreamErrorHTTPSProxyFailureUnexpectedResponseToCONNECTMethod
kCFStreamErrorHTTPSProxyFailureUnexpectedResponseToCONNECTMethod NetworkErrors = 0
// kCFFTPErrorUnexpectedStatusCode - The server returned an unexpected status code.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfftpErrorUnexpectedStatusCode
kCFFTPErrorUnexpectedStatusCode NetworkErrors = 0
// kCFHTTPCookieCannotParseCookieFile - The cookie file can’t be parsed.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfhttpCookieCannotParseCookieFile
kCFHTTPCookieCannotParseCookieFile NetworkErrors = 0
// kCFSOCKS4ErrorIdConflict - The server rejected the request because the client program and the   daemon reported different user IDs.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfsocks4ErrorIdConflict
kCFSOCKS4ErrorIdConflict NetworkErrors = 0
// kCFSOCKS4ErrorIdentdFailed - The server couldn’t connect to the   daemon on the client and rejected the request.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfsocks4ErrorIdentdFailed
kCFSOCKS4ErrorIdentdFailed NetworkErrors = 0
// kCFSOCKS4ErrorRequestFailed - The server rejected the request, or the request failed.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfsocks4ErrorRequestFailed
kCFSOCKS4ErrorRequestFailed NetworkErrors = 0
// kCFSOCKS4ErrorUnknownStatusCode - The server returned an unknown status code.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfsocks4ErrorUnknownStatusCode
kCFSOCKS4ErrorUnknownStatusCode NetworkErrors = 0
// kCFSOCKS5ErrorBadCredentials - The SOCKS server refused the client connection because of bad login credentials.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfsocks5ErrorBadCredentials
kCFSOCKS5ErrorBadCredentials NetworkErrors = 0
// kCFSOCKS5ErrorBadResponseAddr - The address type returned isn’t supported.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfsocks5ErrorBadResponseAddr
kCFSOCKS5ErrorBadResponseAddr NetworkErrors = 0
// kCFSOCKS5ErrorBadState - The stream isn’t in a state that allows the requested operation.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfsocks5ErrorBadState
kCFSOCKS5ErrorBadState NetworkErrors = 0
// kCFSOCKS5ErrorNoAcceptableMethod - The client and server couldn’t find a mutually agreeable authentication method.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfsocks5ErrorNoAcceptableMethod
kCFSOCKS5ErrorNoAcceptableMethod NetworkErrors = 0
// kCFSOCKS5ErrorUnsupportedNegotiationMethod - The requested method isn’t supported.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfsocks5ErrorUnsupportedNegotiationMethod
kCFSOCKS5ErrorUnsupportedNegotiationMethod NetworkErrors = 0
// kCFSOCKSErrorUnknownClientVersion - The SOCKS server rejected access because it doesn’t support connections with the requested SOCKS version.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfsocksErrorUnknownClientVersion
kCFSOCKSErrorUnknownClientVersion NetworkErrors = 0
// kCFSOCKSErrorUnsupportedServerVersion - The SOCKS server doesn’t support the requested version.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfsocksErrorUnsupportedServerVersion
kCFSOCKSErrorUnsupportedServerVersion NetworkErrors = 0
// kCFURLErrorAppTransportSecurityRequiresSecureConnection - The connection failed because the App Transport Security configuration requires a secure connection.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorAppTransportSecurityRequiresSecureConnection
kCFURLErrorAppTransportSecurityRequiresSecureConnection NetworkErrors = 0
// kCFURLErrorBackgroundSessionInUseByAnotherProcess - The background session failed because it was in use by another process.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorBackgroundSessionInUseByAnotherProcess
kCFURLErrorBackgroundSessionInUseByAnotherProcess NetworkErrors = 0
// kCFURLErrorBackgroundSessionWasDisconnected - The background session failed because it was disconnected.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorBackgroundSessionWasDisconnected
kCFURLErrorBackgroundSessionWasDisconnected NetworkErrors = 0
// kCFURLErrorBadServerResponse - The connection received an invalid server response.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorBadServerResponse
kCFURLErrorBadServerResponse NetworkErrors = 0
// kCFURLErrorBadURL - The connection failed due to a malformed URL.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorBadURL
kCFURLErrorBadURL NetworkErrors = 0
// kCFURLErrorCallIsActive - The connection failed because a call is active.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorCallIsActive
kCFURLErrorCallIsActive NetworkErrors = 0
// kCFURLErrorCancelled - The connection was cancelled.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorCancelled
kCFURLErrorCancelled NetworkErrors = 0
// kCFURLErrorCannotCloseFile - The file can’t be closed.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorCannotCloseFile
kCFURLErrorCannotCloseFile NetworkErrors = 0
// kCFURLErrorCannotConnectToHost - The connection failed because a connection can’t be made to the host.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorCannotConnectToHost
kCFURLErrorCannotConnectToHost NetworkErrors = 0
// kCFURLErrorCannotCreateFile - The file can’t be created.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorCannotCreateFile
kCFURLErrorCannotCreateFile NetworkErrors = 0
// kCFURLErrorCannotDecodeContentData - The connection can’t decode data encoded with an unknown content encoding.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorCannotDecodeContentData
kCFURLErrorCannotDecodeContentData NetworkErrors = 0
// kCFURLErrorCannotDecodeRawData - The connection can’t decode data encoded with a known content encoding.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorCannotDecodeRawData
kCFURLErrorCannotDecodeRawData NetworkErrors = 0
// kCFURLErrorCannotFindHost - The connection failed because the host couldn’t be found.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorCannotFindHost
kCFURLErrorCannotFindHost NetworkErrors = 0
// kCFURLErrorCannotLoadFromNetwork - The connection failed because it’s being required to return a cached resource, but one isn’t available.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorCannotLoadFromNetwork
kCFURLErrorCannotLoadFromNetwork NetworkErrors = 0
// kCFURLErrorCannotMoveFile - The file can’t be moved.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorCannotMoveFile
kCFURLErrorCannotMoveFile NetworkErrors = 0
// kCFURLErrorCannotOpenFile - The file can’t be opened.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorCannotOpenFile
kCFURLErrorCannotOpenFile NetworkErrors = 0
// kCFURLErrorCannotParseResponse - The connection can’t parse the server’s response.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorCannotParseResponse
kCFURLErrorCannotParseResponse NetworkErrors = 0
// kCFURLErrorCannotRemoveFile - The file can’t be removed.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorCannotRemoveFile
kCFURLErrorCannotRemoveFile NetworkErrors = 0
// kCFURLErrorCannotWriteToFile - The file can’t be written.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorCannotWriteToFile
kCFURLErrorCannotWriteToFile NetworkErrors = 0
// kCFURLErrorClientCertificateRejected - The secure connection failed because the client’s certificate was rejected.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorClientCertificateRejected
kCFURLErrorClientCertificateRejected NetworkErrors = 0
// kCFURLErrorClientCertificateRequired - The secure connection failed because the server requires a client certificate.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorClientCertificateRequired
kCFURLErrorClientCertificateRequired NetworkErrors = 0
// kCFURLErrorDNSLookupFailed - The connection failed because the DNS lookup failed.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorDNSLookupFailed
kCFURLErrorDNSLookupFailed NetworkErrors = 0
// kCFURLErrorDataLengthExceedsMaximum - The file operation failed because the file is too large.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorDataLengthExceedsMaximum
kCFURLErrorDataLengthExceedsMaximum NetworkErrors = 0
// kCFURLErrorDataNotAllowed - The connection failed because data use isn’t currently allowed on the device.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorDataNotAllowed
kCFURLErrorDataNotAllowed NetworkErrors = 0
// kCFURLErrorDownloadDecodingFailedMidStream - The download failed because decoding of the downloaded data failed midstream.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorDownloadDecodingFailedMidStream
kCFURLErrorDownloadDecodingFailedMidStream NetworkErrors = 0
// kCFURLErrorDownloadDecodingFailedToComplete - The download failed because decoding of the downloaded data failed to complete.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorDownloadDecodingFailedToComplete
kCFURLErrorDownloadDecodingFailedToComplete NetworkErrors = 0
// kCFURLErrorFileDoesNotExist - The file operation failed because the file doesn’t exist.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorFileDoesNotExist
kCFURLErrorFileDoesNotExist NetworkErrors = 0
// kCFURLErrorFileIsDirectory - The file operation failed because the file is a directory.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorFileIsDirectory
kCFURLErrorFileIsDirectory NetworkErrors = 0
// kCFURLErrorFileOutsideSafeArea - The file is outside of the safe area.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorFileOutsideSafeArea
kCFURLErrorFileOutsideSafeArea NetworkErrors = 0
// kCFURLErrorHTTPTooManyRedirects - The HTTP connection failed due to too many redirects.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorHTTPTooManyRedirects
kCFURLErrorHTTPTooManyRedirects NetworkErrors = 0
// kCFURLErrorInternationalRoamingOff - The connection failed because international roaming is disabled on the device.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorInternationalRoamingOff
kCFURLErrorInternationalRoamingOff NetworkErrors = 0
// kCFURLErrorNetworkConnectionLost - The connection failed because the network connection was lost.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorNetworkConnectionLost
kCFURLErrorNetworkConnectionLost NetworkErrors = 0
// kCFURLErrorNoPermissionsToReadFile - The file operation failed because it doesn’t have permission to read the file.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorNoPermissionsToReadFile
kCFURLErrorNoPermissionsToReadFile NetworkErrors = 0
// kCFURLErrorNotConnectedToInternet - The connection failed because the device isn’t connected to the internet.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorNotConnectedToInternet
kCFURLErrorNotConnectedToInternet NetworkErrors = 0
// kCFURLErrorRedirectToNonExistentLocation - The connection was redirected to a nonexistent location.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorRedirectToNonExistentLocation
kCFURLErrorRedirectToNonExistentLocation NetworkErrors = 0
// kCFURLErrorRequestBodyStreamExhausted - The connection failed because the request’s body stream was exhausted.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorRequestBodyStreamExhausted
kCFURLErrorRequestBodyStreamExhausted NetworkErrors = 0
// kCFURLErrorResourceUnavailable - The connection’s resource is unavailable.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorResourceUnavailable
kCFURLErrorResourceUnavailable NetworkErrors = 0
// kCFURLErrorSecureConnectionFailed - The secure connection failed for an unknown reason.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorSecureConnectionFailed
kCFURLErrorSecureConnectionFailed NetworkErrors = 0
// kCFURLErrorServerCertificateHasBadDate - The secure connection failed because the server’s certificate has an invalid date.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorServerCertificateHasBadDate
kCFURLErrorServerCertificateHasBadDate NetworkErrors = 0
// kCFURLErrorServerCertificateHasUnknownRoot - The secure connection failed because the server’s certificate has an unknown root.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorServerCertificateHasUnknownRoot
kCFURLErrorServerCertificateHasUnknownRoot NetworkErrors = 0
// kCFURLErrorServerCertificateNotYetValid - The secure connection failed because the server’s certificate isn’t valid yet.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorServerCertificateNotYetValid
kCFURLErrorServerCertificateNotYetValid NetworkErrors = 0
// kCFURLErrorServerCertificateUntrusted - The secure connection failed because the server’s certificate isn’t trusted.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorServerCertificateUntrusted
kCFURLErrorServerCertificateUntrusted NetworkErrors = 0
// kCFURLErrorTimedOut - The connection timed out.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorTimedOut
kCFURLErrorTimedOut NetworkErrors = 0
// kCFURLErrorUnknown - An unknown error occurred.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorUnknown
kCFURLErrorUnknown NetworkErrors = 0
// kCFURLErrorUnsupportedURL - The connection failed due to an unsupported URL scheme.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorUnsupportedURL
kCFURLErrorUnsupportedURL NetworkErrors = 0
// kCFURLErrorUserAuthenticationRequired - The connection failed because it requires authentication.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorUserAuthenticationRequired
kCFURLErrorUserAuthenticationRequired NetworkErrors = 0
// kCFURLErrorUserCancelledAuthentication - The connection failed because the user cancelled required authentication.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorUserCancelledAuthentication
kCFURLErrorUserCancelledAuthentication NetworkErrors = 0
// kCFURLErrorZeroByteResource - The resource retrieved by the connection is zero bytes.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFNetworkErrors/cfurlErrorZeroByteResource
kCFURLErrorZeroByteResource NetworkErrors = 0
)

// CFStreamErrorHTTP - Error codes that a read stream for an HTTP request may return.
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFStreamErrorHTTP
type StreamErrorHTTP uint

const (
// kCFStreamErrorHTTPBadURL - The URL is not properly formatted.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFStreamErrorHTTP/badURL
kCFStreamErrorHTTPBadURL StreamErrorHTTP = 0
// kCFStreamErrorHTTPParseFailure - A parsing error occurred while an incoming message was being deserialized and appended to a message object. The headers of the incoming message may be formatted improperly.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFStreamErrorHTTP/parseFailure
kCFStreamErrorHTTPParseFailure StreamErrorHTTP = 0
// kCFStreamErrorHTTPRedirectionLoop - A redirection loop has been detected.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFStreamErrorHTTP/redirectionLoop
kCFStreamErrorHTTPRedirectionLoop StreamErrorHTTP = 0
)

// CFStreamErrorHTTPAuthentication - Authentication error codes that may be returned when trying to apply authentication to a request.
//
// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFStreamErrorHTTPAuthentication
type StreamErrorHTTPAuthentication uint

const (
// kCFStreamErrorHTTPAuthenticationBadPassword - Password is in a format that is not suitable for the request. Currently, passwords are decoded using  .
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFStreamErrorHTTPAuthentication/badPassword
kCFStreamErrorHTTPAuthenticationBadPassword StreamErrorHTTPAuthentication = 0
// kCFStreamErrorHTTPAuthenticationBadUserName - User name is in a format that is not suitable for the request. Currently, user names are decoded using  .
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFStreamErrorHTTPAuthentication/badUserName
kCFStreamErrorHTTPAuthenticationBadUserName StreamErrorHTTPAuthentication = 0
// kCFStreamErrorHTTPAuthenticationTypeUnsupported - Specified authentication type is not supported.
//
	// [Full Topic]: https://developer.apple.com/documentation/CFNetwork/CFStreamErrorHTTPAuthentication/typeUnsupported
kCFStreamErrorHTTPAuthenticationTypeUnsupported StreamErrorHTTPAuthentication = 0
)



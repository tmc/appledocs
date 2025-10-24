// Code generated from Apple documentation for MultipeerConnectivity. DO NOT EDIT.

package multipeerconnectivity

/* debug [enums.gen.go]: Generating 4 enums for MultipeerConnectivity */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum MCErrorCode (7 cases) */
// MCErrorCode - Error codes found in 
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCError/Code
type MCErrorCode uint

const (
	// MCErrorCancelled - The operation was cancelled by the user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCError/Code/cancelled
	MCErrorCancelled MCErrorCode = 0
	// MCErrorInvalidParameter - Your app passed an invalid value as a parameter.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCError/Code/invalidParameter
	MCErrorInvalidParameter MCErrorCode = 0
	// MCErrorNotConnected - Your app attempted to send data to a peer that is not connected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCError/Code/notConnected
	MCErrorNotConnected MCErrorCode = 0
	// MCErrorTimedOut - The connection attempt timed out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCError/Code/timedOut
	MCErrorTimedOut MCErrorCode = 0
	// MCErrorUnavailable - Multipeer connectivity is currently unavailable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCError/Code/unavailable
	MCErrorUnavailable MCErrorCode = 0
	// MCErrorUnknown - An unknown error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCError/Code/unknown
	MCErrorUnknown MCErrorCode = 0
	// MCErrorUnsupported - The operation is unsupported. For example, this error is returned if you call   with a URL that is neither a local file nor a web URL.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCError/Code/unsupported
	MCErrorUnsupported MCErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum MCEncryptionPreference (3 cases) */
// MCEncryptionPreference - Indicates whether a session should use encryption when communicating with nearby peers.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCEncryptionPreference
type MCEncryptionPreference uint

const (
	// MCEncryptionNone - The session should not be encrypted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCEncryptionPreference/none
	MCEncryptionNone MCEncryptionPreference = 0
	// MCEncryptionOptional - The session prefers to use encryption, but accepts unencrypted connections. A connection uses encryption when all the peers choose either   or  . If some peers choose  , then the session will not be encrypted. For this reason, if some peers running your app can be configured without encryption, you should always assume that the session is unencrypted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCEncryptionPreference/optional
	MCEncryptionOptional MCEncryptionPreference = 0
	// MCEncryptionRequired - The session requires encryption.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCEncryptionPreference/required
	MCEncryptionRequired MCEncryptionPreference = 0
)

/* debug [enums.gen.go]: Processing enum MCSessionSendDataMode (2 cases) */
// MCSessionSendDataMode - Indicates whether delivery of data should be guaranteed.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSessionSendDataMode
type MCSessionSendDataMode uint

const (
	// MCSessionSendDataReliable - The framework should guarantee delivery of each message, enqueueing and retransmitting data as needed, and ensuring in-order delivery.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSessionSendDataMode/reliable
	MCSessionSendDataReliable MCSessionSendDataMode = 0
	// MCSessionSendDataUnreliable - Messages to peers should be sent immediately without socket-level queueing. If a message cannot be sent immediately, it should be dropped. The order of messages is not guaranteed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSessionSendDataMode/unreliable
	MCSessionSendDataUnreliable MCSessionSendDataMode = 0
)

/* debug [enums.gen.go]: Processing enum MCSessionState (3 cases) */
// MCSessionState - Indicates the current state of a given peer within a session.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSessionState
type MCSessionState uint

const (
	// MCSessionStateConnected - The peer is connected to this session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSessionState/connected
	MCSessionStateConnected MCSessionState = 0
	// MCSessionStateConnecting - A connection to the peer is currently being established.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSessionState/connecting
	MCSessionStateConnecting MCSessionState = 0
	// MCSessionStateNotConnected - The peer is not (or is no longer) in this session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCSessionState/notConnected
	MCSessionStateNotConnected MCSessionState = 0
)



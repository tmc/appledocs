// Code generated from Apple documentation for SystemConfiguration. DO NOT EDIT.

package systemconfiguration


// Enum types and constants

// SCNetworkConnectionPPPStatus - The PPP-specific status of the network connection.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionPPPStatus
type SCNetworkConnectionPPPStatus uint

const (
	// kSCNetworkConnectionPPPAuthenticating - PPP is authenticating to the server (PAP, CHAP, MS-CHAP, or EAP protocols).
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionPPPStatus/authenticating
	kSCNetworkConnectionPPPAuthenticating SCNetworkConnectionPPPStatus = 0
	// kSCNetworkConnectionPPPConnected - PPP is now fully connected for at least one networking layer. Additional networking protocol might still be negotiating.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionPPPStatus/connected
	kSCNetworkConnectionPPPConnected SCNetworkConnectionPPPStatus = 0
	// kSCNetworkConnectionPPPConnectingLink - PPP is connecting the lower connection layer (for example, the modem is dialing out).
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionPPPStatus/connectingLink
	kSCNetworkConnectionPPPConnectingLink SCNetworkConnectionPPPStatus = 0
	// kSCNetworkConnectionPPPDialOnTraffic - PPP is waiting for networking traffic to automatically establish the connection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionPPPStatus/dialOnTraffic
	kSCNetworkConnectionPPPDialOnTraffic SCNetworkConnectionPPPStatus = 0
	// kSCNetworkConnectionPPPDisconnected - PPP is disconnected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionPPPStatus/disconnected
	kSCNetworkConnectionPPPDisconnected SCNetworkConnectionPPPStatus = 0
	// kSCNetworkConnectionPPPDisconnectingLink - PPP is disconnecting the lower level (for example, the modem is hanging up).
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionPPPStatus/disconnectingLink
	kSCNetworkConnectionPPPDisconnectingLink SCNetworkConnectionPPPStatus = 0
	// kSCNetworkConnectionPPPHoldingLinkOff - PPP is disconnected and maintaining the link temporarily off.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionPPPStatus/holdingLinkOff
	kSCNetworkConnectionPPPHoldingLinkOff SCNetworkConnectionPPPStatus = 0
	// kSCNetworkConnectionPPPInitializing - PPP is initializing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionPPPStatus/initializing
	kSCNetworkConnectionPPPInitializing SCNetworkConnectionPPPStatus = 0
	// kSCNetworkConnectionPPPNegotiatingLink - The PPP lower layer is connected and PPP is negotiating the link layer (LCP protocol).
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionPPPStatus/negotiatingLink
	kSCNetworkConnectionPPPNegotiatingLink SCNetworkConnectionPPPStatus = 0
	// kSCNetworkConnectionPPPNegotiatingNetwork - PPP is now authenticated and negotiating the networking layer (IPCP or IPv6CP protocols).
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionPPPStatus/negotiatingNetwork
	kSCNetworkConnectionPPPNegotiatingNetwork SCNetworkConnectionPPPStatus = 0
	// kSCNetworkConnectionPPPSuspended - PPP is suspended as a result of the suspend command (for example, when a V.92 Modem is On Hold).
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionPPPStatus/suspended
	kSCNetworkConnectionPPPSuspended SCNetworkConnectionPPPStatus = 0
	// kSCNetworkConnectionPPPTerminating - PPP networking and link protocols are terminating.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionPPPStatus/terminating
	kSCNetworkConnectionPPPTerminating SCNetworkConnectionPPPStatus = 0
	// kSCNetworkConnectionPPPWaitingForCallBack - PPP is waiting for the server to call back.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionPPPStatus/waitingForCallBack
	kSCNetworkConnectionPPPWaitingForCallBack SCNetworkConnectionPPPStatus = 0
	// kSCNetworkConnectionPPPWaitingForRedial - PPP has found a busy server and is waiting for redial.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionPPPStatus/waitingForRedial
	kSCNetworkConnectionPPPWaitingForRedial SCNetworkConnectionPPPStatus = 0
)


// SCNetworkConnectionStatus - The current status of the network connection.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionStatus
type SCNetworkConnectionStatus uint

const (
	// kSCNetworkConnectionConnected - The network connection is connected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionStatus/connected
	kSCNetworkConnectionConnected SCNetworkConnectionStatus = 0
	// kSCNetworkConnectionConnecting - The network connection is connecting.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionStatus/connecting
	kSCNetworkConnectionConnecting SCNetworkConnectionStatus = 0
	// kSCNetworkConnectionDisconnected - The network connection is disconnected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionStatus/disconnected
	kSCNetworkConnectionDisconnected SCNetworkConnectionStatus = 0
	// kSCNetworkConnectionDisconnecting - The network connection is disconnecting.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionStatus/disconnecting
	kSCNetworkConnectionDisconnecting SCNetworkConnectionStatus = 0
	// kSCNetworkConnectionInvalid - The network connection refers to an invalid service.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionStatus/invalid
	kSCNetworkConnectionInvalid SCNetworkConnectionStatus = 0
)


// SCNetworkReachabilityFlags - Flags that indicate the reachability of a network node name or address, including whether a connection is required, and whether some user intervention might be required when establishing a connection.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityFlags
type SCNetworkReachabilityFlags uint

const (
	// kSCNetworkReachabilityFlagsConnectionAutomatic - The specified node name or address can be reached using the current network configuration, but a connection must first be established. Any traffic directed to the specified name or address will initiate the connection. This flag is a synonym for  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityFlags/connectionAutomatic
	kSCNetworkReachabilityFlagsConnectionAutomatic SCNetworkReachabilityFlags = 0
	// kSCNetworkReachabilityFlagsConnectionOnDemand - The specified node name or address can be reached using the current network configuration, but a connection must first be established.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityFlags/connectionOnDemand
	kSCNetworkReachabilityFlagsConnectionOnDemand SCNetworkReachabilityFlags = 0
	// kSCNetworkReachabilityFlagsConnectionOnTraffic - The specified node name or address can be reached using the current network configuration, but a connection must first be established. Any traffic directed to the specified name or address will initiate the connection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityFlags/connectionOnTraffic
	kSCNetworkReachabilityFlagsConnectionOnTraffic SCNetworkReachabilityFlags = 0
	// kSCNetworkReachabilityFlagsConnectionRequired - The specified node name or address can be reached using the current network configuration, but a connection must first be established. If this flag is set, the   flag,   flag, or   flag is also typically set to indicate the type of connection required. If the user must manually make the connection, the   flag is also set.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityFlags/connectionRequired
	kSCNetworkReachabilityFlagsConnectionRequired SCNetworkReachabilityFlags = 0
	// kSCNetworkReachabilityFlagsInterventionRequired - The specified node name or address can be reached using the current network configuration, but a connection must first be established.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityFlags/interventionRequired
	kSCNetworkReachabilityFlagsInterventionRequired SCNetworkReachabilityFlags = 0
	// kSCNetworkReachabilityFlagsIsDirect - Network traffic to the specified node name or address will not go through a gateway, but is routed directly to one of the interfaces in the system.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityFlags/isDirect
	kSCNetworkReachabilityFlagsIsDirect SCNetworkReachabilityFlags = 0
	// kSCNetworkReachabilityFlagsIsLocalAddress - The specified node name or address is one that is associated with a network interface on the current system.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityFlags/isLocalAddress
	kSCNetworkReachabilityFlagsIsLocalAddress SCNetworkReachabilityFlags = 0
	// kSCNetworkReachabilityFlagsIsWWAN - The specified node name or address can be reached via a cellular connection, such as EDGE or GPRS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityFlags/isWWAN
	kSCNetworkReachabilityFlagsIsWWAN SCNetworkReachabilityFlags = 0
	// kSCNetworkReachabilityFlagsReachable - The specified node name or address can be reached using the current network configuration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityFlags/reachable
	kSCNetworkReachabilityFlagsReachable SCNetworkReachabilityFlags = 0
	// kSCNetworkReachabilityFlagsTransientConnection - The specified node name or address can be reached via a transient connection, such as PPP.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityFlags/transientConnection
	kSCNetworkReachabilityFlagsTransientConnection SCNetworkReachabilityFlags = 0
)


// SCPreferencesNotification - The type of notification (used with the 
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesNotification
type SCPreferencesNotification uint

const (
	// kSCPreferencesNotificationApply - Indicates when a request has been made to apply the currently saved preferences to the active system configuration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesNotification/apply
	kSCPreferencesNotificationApply SCPreferencesNotification = 0
	// kSCPreferencesNotificationCommit - Indicates when new preferences have been saved.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesNotification/commit
	kSCPreferencesNotificationCommit SCPreferencesNotification = 0
)



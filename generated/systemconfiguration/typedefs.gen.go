// Code generated from Apple documentation for SystemConfiguration. DO NOT EDIT.

package systemconfiguration

// Type aliases and typedefs
// AuthorizationRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/AuthorizationRef
// AuthorizationRef has base type: const struct AuthorizationOpaqueRef *
type AuthorizationRef uintptr
// SCBondInterfaceRef - The reference to an object that represents an Ethernet bond interface.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterface
// SCBondInterfaceRef has base type: SCNetworkInterfaceRef
type SCBondInterfaceRef uintptr
// SCBondStatusRef - The reference to an object that represents the status of an Ethernet bond interface.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondStatus
// SCBondStatusRef has base type: const struct __SCBondStatus *
type SCBondStatusRef uintptr
// SCDynamicStoreRef - The handle to an open dynamic store session with the system configuration daemon.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStore
// SCDynamicStoreRef has base type: const struct __SCDynamicStore *
type SCDynamicStoreRef uintptr
// SCDynamicStoreCallBack - Callback used when notification of changes made to the dynamic store is delivered.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCallBack
// SCDynamicStoreCallBack has base type: void (*)(const struct __SCDynamicStore *, const struct __CFArray *, void *)
type SCDynamicStoreCallBack uintptr
// SCNetworkConnectionRef - The handle to manage a connection-oriented service.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnection
// SCNetworkConnectionRef has base type: const struct __SCNetworkConnection *
type SCNetworkConnectionRef uintptr
// SCNetworkConnectionCallBack - The type of callback function used when a status event is delivered.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionCallBack
// SCNetworkConnectionCallBack has base type: void (*)(const struct __SCNetworkConnection *, enum SCNetworkConnectionStatus, void *)
type SCNetworkConnectionCallBack uintptr
// SCNetworkConnectionFlags - Flags that indicate whether the specified network node name or address is reachable, whether a connection is required, and whether some user intervention may be required when establishing a connection.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionFlags
// SCNetworkConnectionFlags has base type: uint32_t
type SCNetworkConnectionFlags uintptr
// SCNetworkInterfaceRef - The reference to an object that represents a network interface.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterface
// SCNetworkInterfaceRef has base type: const struct __SCNetworkInterface *
type SCNetworkInterfaceRef uintptr
// SCNetworkProtocolRef - The reference to an object that represents a network protocol.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkProtocol
// SCNetworkProtocolRef has base type: const struct __SCNetworkProtocol *
type SCNetworkProtocolRef uintptr
// SCNetworkReachabilityRef - The handle to a network address or name.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachability
// SCNetworkReachabilityRef has base type: const struct __SCNetworkReachability *
type SCNetworkReachabilityRef uintptr
// SCNetworkReachabilityCallBack - Type of callback function used when the reachability of a network address or name changes.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityCallBack
// SCNetworkReachabilityCallBack has base type: void (*)(const struct __SCNetworkReachability *, enum SCNetworkReachabilityFlags, void *)
type SCNetworkReachabilityCallBack uintptr
// SCNetworkServiceRef - The reference to an object that represents a network service.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkService
// SCNetworkServiceRef has base type: const struct __SCNetworkService *
type SCNetworkServiceRef uintptr
// SCNetworkSetRef - The reference to an object that represents a network set.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSet
// SCNetworkSetRef has base type: const struct __SCNetworkSet *
type SCNetworkSetRef uintptr
// SCPreferencesRef - The handle to an open preferences session for accessing system configuration preferences.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferences
// SCPreferencesRef has base type: const struct __SCPreferences *
type SCPreferencesRef uintptr
// SCPreferencesCallBack - Type of the callback function used when the preferences have been updated or applied.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesCallBack
// SCPreferencesCallBack has base type: void (*)(const struct __SCPreferences *, enum SCPreferencesNotification, void *)
type SCPreferencesCallBack uintptr
// SCVLANInterfaceRef - The reference to an object that represents a virtual LAN (VLAN) interface.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterface
// SCVLANInterfaceRef has base type: SCNetworkInterfaceRef
type SCVLANInterfaceRef uintptr


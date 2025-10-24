// Code generated from Apple documentation for SystemConfiguration. DO NOT EDIT.

package systemconfiguration

import (
	"unsafe"
)

// Type aliases and typedefs
// AuthorizationRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/AuthorizationRef
// AuthorizationRef has base type: const struct AuthorizationOpaqueRef *
type AuthorizationRef uintptr

// BondInterfaceRef - The reference to an object that represents an Ethernet bond interface.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterface
// SCBondInterfaceRef has base type: SCNetworkInterfaceRef
type BondInterfaceRef uintptr

// BondStatusRef - The reference to an object that represents the status of an Ethernet bond interface.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondStatus
// SCBondStatusRef has base type: const struct __SCBondStatus *
type BondStatusRef uintptr

// DynamicStoreRef - The handle to an open dynamic store session with the system configuration daemon.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStore
// SCDynamicStoreRef has base type: const struct __SCDynamicStore *
type DynamicStoreRef uintptr

// DynamicStoreCallBack - Callback used when notification of changes made to the dynamic store is delivered.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCallBack
// SCDynamicStoreCallBack is a callback function
// C type: void (*)(const struct __SCDynamicStore *, const struct __CFArray *, void *)
type DynamicStoreCallBack = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)

// NetworkConnectionRef - The handle to manage a connection-oriented service.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnection
// SCNetworkConnectionRef has base type: const struct __SCNetworkConnection *
type NetworkConnectionRef uintptr

// NetworkConnectionCallBack - The type of callback function used when a status event is delivered.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionCallBack
// SCNetworkConnectionCallBack is a callback function
// C type: void (*)(const struct __SCNetworkConnection *, enum SCNetworkConnectionStatus, void *)
type NetworkConnectionCallBack = func(unsafe.Pointer, SCNetworkConnectionStatus, unsafe.Pointer)

// NetworkConnectionFlags - Flags that indicate whether the specified network node name or address is reachable, whether a connection is required, and whether some user intervention may be required when establishing a connection.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionFlags
// SCNetworkConnectionFlags has base type: uint32_t
type NetworkConnectionFlags uintptr

// NetworkInterfaceRef - The reference to an object that represents a network interface.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterface
// SCNetworkInterfaceRef has base type: const struct __SCNetworkInterface *
type NetworkInterfaceRef uintptr

// NetworkProtocolRef - The reference to an object that represents a network protocol.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkProtocol
// SCNetworkProtocolRef has base type: const struct __SCNetworkProtocol *
type NetworkProtocolRef uintptr

// NetworkReachabilityRef - The handle to a network address or name.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachability
// SCNetworkReachabilityRef has base type: const struct __SCNetworkReachability *
type NetworkReachabilityRef uintptr

// NetworkReachabilityCallBack - Type of callback function used when the reachability of a network address or name changes.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityCallBack
// SCNetworkReachabilityCallBack is a callback function
// C type: void (*)(const struct __SCNetworkReachability *, enum SCNetworkReachabilityFlags, void *)
type NetworkReachabilityCallBack = func(unsafe.Pointer, SCNetworkReachabilityFlags, unsafe.Pointer)

// NetworkServiceRef - The reference to an object that represents a network service.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkService
// SCNetworkServiceRef has base type: const struct __SCNetworkService *
type NetworkServiceRef uintptr

// NetworkSetRef - The reference to an object that represents a network set.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSet
// SCNetworkSetRef has base type: const struct __SCNetworkSet *
type NetworkSetRef uintptr

// PreferencesRef - The handle to an open preferences session for accessing system configuration preferences.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferences
// SCPreferencesRef has base type: const struct __SCPreferences *
type PreferencesRef uintptr

// PreferencesCallBack - Type of the callback function used when the preferences have been updated or applied.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesCallBack
// SCPreferencesCallBack is a callback function
// C type: void (*)(const struct __SCPreferences *, enum SCPreferencesNotification, void *)
type PreferencesCallBack = func(unsafe.Pointer, SCPreferencesNotification, unsafe.Pointer)

// VLANInterfaceRef - The reference to an object that represents a virtual LAN (VLAN) interface.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterface
// SCVLANInterfaceRef has base type: SCNetworkInterfaceRef
type VLANInterfaceRef uintptr

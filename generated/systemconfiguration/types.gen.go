// Code generated from Apple documentation for SystemConfiguration. DO NOT EDIT.

package systemconfiguration

import (
	"unsafe"
)

// C struct types
// SCDynamicStoreContext - Structure containing user-specified data and callbacks for a dynamic store session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreContext
type SCDynamicStoreContext struct {
	CopyDescription unsafe.Pointer // The callback used to provide a description of the   field.
	Info            unsafe.Pointer // A C pointer to a user-specified block of data.
	Release         unsafe.Pointer // The callback used to remove a retain previously added for the   field. If this parameter is not a pointer to a function of the correct prototype, the behavior is undefined. The value of this parameter can be  .
	Retain          unsafe.Pointer // The callback used to add a retain for the   field. If this parameter is not a pointer to a function of the correct prototype, the behavior is undefined. The value of this parameter can be  .
	Version         Index          // The version number of the structure type being passed in as a parameter to the   creation function (such as  ). This structure is version  .
} /* debug [types.gen.go/struct]: SCDynamicStoreContext */

// SCNetworkConnectionContext - A structure containing user-specified data and callbacks for a network connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionContext
type SCNetworkConnectionContext struct {
	CopyDescription unsafe.Pointer // The callback used to provide a description of the   field.
	Info            unsafe.Pointer // A C pointer to a user-specified block of data.
	Release         unsafe.Pointer // The calllback used to remove a retain previously added for the info field. If this parameter is not a pointer to a function of the correct prototype, the behavior is undefined. The value may be  .
	Retain          unsafe.Pointer // The callback used to add a retain for the info field. If this parameter is not a pointer to a function of the correct prototype, the behavior is undefined. The value may be  .
	Version         Index          // The version number of the structure type being passed in as a parameter to the   function. This structure is version  .
} /* debug [types.gen.go/struct]: SCNetworkConnectionContext */

// SCNetworkReachabilityContext - Structure containing user-specified data and callbacks used with
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityContext
type SCNetworkReachabilityContext struct {
	CopyDescription unsafe.Pointer // The callback used to provide a description of the   field.
	Info            unsafe.Pointer // A C pointer to a user-specified block of data.
	Release         unsafe.Pointer // The callback used to remove a retain previously added for the info field. If this parameter is not a pointer to a function of the correct prototype, the behavior is undefined. The value can be  .
	Retain          unsafe.Pointer // The callback used to add a retain for the info field. If this parameter is not a pointer to a function of the correct prototype, the behavior is undefined. The value can be  .
	Version         Index          // The version number of the structure type being passed in as a parameter to an   creation function. This structure is version  .
} /* debug [types.gen.go/struct]: SCNetworkReachabilityContext */

// SCPreferencesContext - A structure containing user-specified data and callbacks for accessing system configuration preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesContext
type SCPreferencesContext struct {
	CopyDescription unsafe.Pointer // The callback used to provide a description of the   field.
	Info            unsafe.Pointer // A C pointer to a user-specified block of data.
	Release         unsafe.Pointer // The calllback used to remove a retain previously added for the   field. If this parameter is not a pointer to a function of the correct prototype, the behavior is undefined. The value may be  .
	Retain          unsafe.Pointer // The callback used to add a retain for the   field. If this parameter is not a pointer to a function of the correct prototype, the behavior is undefined. The value may be  .
	Version         Index          // The version number of the structure type being passed in as a parameter to  . This structure is version  .
} /* debug [types.gen.go/struct]: SCPreferencesContext */

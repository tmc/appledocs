// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// CoreWLAN Functions (18 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_CWKeychainCopyEAPIdentity func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CWKeychainCopyEAPIdentityList func(unsafe.Pointer) unsafe.Pointer
	_CWKeychainCopyEAPUsernameAndPassword func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CWKeychainCopyPassword func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CWKeychainCopyWiFiEAPIdentity func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CWKeychainDeleteEAPUsernameAndPassword func(unsafe.Pointer) unsafe.Pointer
	_CWKeychainDeletePassword func(unsafe.Pointer) unsafe.Pointer
	_CWKeychainDeleteWiFiEAPUsernameAndPassword func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CWKeychainDeleteWiFiPassword func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CWKeychainFindWiFiEAPUsernameAndPassword func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CWKeychainFindWiFiPassword func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CWKeychainSetEAPIdentity func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CWKeychainSetEAPUsernameAndPassword func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CWKeychainSetPassword func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CWKeychainSetWiFiEAPIdentity func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CWKeychainSetWiFiEAPUsernameAndPassword func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CWKeychainSetWiFiPassword func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CWMergeNetworks func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_CWKeychainCopyEAPIdentity, lib, "CWKeychainCopyEAPIdentity")
	tryRegister(&_CWKeychainCopyEAPIdentityList, lib, "CWKeychainCopyEAPIdentityList")
	tryRegister(&_CWKeychainCopyEAPUsernameAndPassword, lib, "CWKeychainCopyEAPUsernameAndPassword")
	tryRegister(&_CWKeychainCopyPassword, lib, "CWKeychainCopyPassword")
	tryRegister(&_CWKeychainCopyWiFiEAPIdentity, lib, "CWKeychainCopyWiFiEAPIdentity")
	tryRegister(&_CWKeychainDeleteEAPUsernameAndPassword, lib, "CWKeychainDeleteEAPUsernameAndPassword")
	tryRegister(&_CWKeychainDeletePassword, lib, "CWKeychainDeletePassword")
	tryRegister(&_CWKeychainDeleteWiFiEAPUsernameAndPassword, lib, "CWKeychainDeleteWiFiEAPUsernameAndPassword")
	tryRegister(&_CWKeychainDeleteWiFiPassword, lib, "CWKeychainDeleteWiFiPassword")
	tryRegister(&_CWKeychainFindWiFiEAPUsernameAndPassword, lib, "CWKeychainFindWiFiEAPUsernameAndPassword")
	tryRegister(&_CWKeychainFindWiFiPassword, lib, "CWKeychainFindWiFiPassword")
	tryRegister(&_CWKeychainSetEAPIdentity, lib, "CWKeychainSetEAPIdentity")
	tryRegister(&_CWKeychainSetEAPUsernameAndPassword, lib, "CWKeychainSetEAPUsernameAndPassword")
	tryRegister(&_CWKeychainSetPassword, lib, "CWKeychainSetPassword")
	tryRegister(&_CWKeychainSetWiFiEAPIdentity, lib, "CWKeychainSetWiFiEAPIdentity")
	tryRegister(&_CWKeychainSetWiFiEAPUsernameAndPassword, lib, "CWKeychainSetWiFiEAPUsernameAndPassword")
	tryRegister(&_CWKeychainSetWiFiPassword, lib, "CWKeychainSetWiFiPassword")
	tryRegister(&_CWMergeNetworks, lib, "CWMergeNetworks")
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



// Finds and returns the identity stored for corresponding network with the specified SSID.
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.7.

// Finds and returns the identity stored for corresponding network with the specified SSID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWKeychainCopyEAPIdentity
func CWKeychainCopyEAPIdentity(ssidData unsafe.Pointer, identity unsafe.Pointer) unsafe.Pointer {
	return _CWKeychainCopyEAPIdentity(ssidData, identity)
	}


// Finds and returns the available identities stored in the keychain.
//
// Added in macOS 10.7.

// Finds and returns the available identities stored in the keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWKeychainCopyEAPIdentityList(_:)
func CWKeychainCopyEAPIdentityList(list unsafe.Pointer) unsafe.Pointer {
	return _CWKeychainCopyEAPIdentityList(list)
	}


// Finds and returns the username and password stored for corresponding network with the specified SSID.
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.7.

// Finds and returns the username and password stored for corresponding network with the specified SSID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWKeychainCopyEAPUsernameAndPassword
func CWKeychainCopyEAPUsernameAndPassword(ssidData unsafe.Pointer, username unsafe.Pointer, password unsafe.Pointer) unsafe.Pointer {
	return _CWKeychainCopyEAPUsernameAndPassword(ssidData, username, password)
	}


// Finds and returns the keychain password stored for the corresponding network with the specified SSID.
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.7.

// Finds and returns the keychain password stored for the corresponding network with the specified SSID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWKeychainCopyPassword
func CWKeychainCopyPassword(ssidData unsafe.Pointer, password unsafe.Pointer) unsafe.Pointer {
	return _CWKeychainCopyPassword(ssidData, password)
	}


// Finds and returns the identity stored for the SSID and keychain domain you specify.
//
// Added in macOS 10.9.

// Finds and returns the identity stored for the SSID and keychain domain you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWKeychainCopyWiFiEAPIdentity(_:_:_:)
func CWKeychainCopyWiFiEAPIdentity(domain unsafe.Pointer, ssid unsafe.Pointer, identity unsafe.Pointer) unsafe.Pointer {
	return _CWKeychainCopyWiFiEAPIdentity(domain, ssid, identity)
	}


// Deletes the keychain item containing the 802.1X username and password for the specified SSID.
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.7.

// Deletes the keychain item containing the 802.1X username and password for the specified SSID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWKeychainDeleteEAPUsernameAndPassword
func CWKeychainDeleteEAPUsernameAndPassword(ssidData unsafe.Pointer) unsafe.Pointer {
	return _CWKeychainDeleteEAPUsernameAndPassword(ssidData)
	}


// Deletes the network password for the specified SSID from the default keychain.
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.7.

// Deletes the network password for the specified SSID from the default keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWKeychainDeletePassword
func CWKeychainDeletePassword(ssidData unsafe.Pointer) unsafe.Pointer {
	return _CWKeychainDeletePassword(ssidData)
	}


// Deletes the 802.1X username and password for the SSID and keychain domain you specify.
//
// Added in macOS 10.9.

// Deletes the 802.1X username and password for the SSID and keychain domain you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWKeychainDeleteWiFiEAPUsernameAndPassword(_:_:)
func CWKeychainDeleteWiFiEAPUsernameAndPassword(domain unsafe.Pointer, ssid unsafe.Pointer) unsafe.Pointer {
	return _CWKeychainDeleteWiFiEAPUsernameAndPassword(domain, ssid)
	}


// Deletes the password for the SSID and keychain domain you specify.
//
// Added in macOS 10.9.

// Deletes the password for the SSID and keychain domain you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWKeychainDeleteWiFiPassword(_:_:)
func CWKeychainDeleteWiFiPassword(domain unsafe.Pointer, ssid unsafe.Pointer) unsafe.Pointer {
	return _CWKeychainDeleteWiFiPassword(domain, ssid)
	}


// Finds and returns the 802.1X username and password stored for the SSID and keychain domain you specify.
//
// Added in macOS 10.9.

// Finds and returns the 802.1X username and password stored for the SSID and keychain domain you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWKeychainFindWiFiEAPUsernameAndPassword(_:_:_:_:)
func CWKeychainFindWiFiEAPUsernameAndPassword(domain unsafe.Pointer, ssid unsafe.Pointer, username unsafe.Pointer, password unsafe.Pointer) unsafe.Pointer {
	return _CWKeychainFindWiFiEAPUsernameAndPassword(domain, ssid, username, password)
	}


// Finds and returns, by reference, the password for the SSID and keychain domain you specify.
//
// Added in macOS 10.9.

// Finds and returns, by reference, the password for the SSID and keychain domain you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWKeychainFindWiFiPassword(_:_:_:)
func CWKeychainFindWiFiPassword(domain unsafe.Pointer, ssid unsafe.Pointer, password unsafe.Pointer) unsafe.Pointer {
	return _CWKeychainFindWiFiPassword(domain, ssid, password)
	}


// Associates an exisiting identity item to the specified SSID.
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.7.

// Associates an exisiting identity item to the specified SSID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWKeychainSetEAPIdentity
func CWKeychainSetEAPIdentity(ssidData unsafe.Pointer, identity unsafe.Pointer) unsafe.Pointer {
	return _CWKeychainSetEAPIdentity(ssidData, identity)
	}


// Sets the keychain item containing the 802.1X username and password for the specified SSID.
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.7.

// Sets the keychain item containing the 802.1X username and password for the specified SSID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWKeychainSetEAPUsernameAndPassword
func CWKeychainSetEAPUsernameAndPassword(ssidData unsafe.Pointer, username unsafe.Pointer, password unsafe.Pointer) unsafe.Pointer {
	return _CWKeychainSetEAPUsernameAndPassword(ssidData, username, password)
	}


// Sets the network keychain password for the specified SSID.
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.7.

// Sets the network keychain password for the specified SSID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWKeychainSetPassword
func CWKeychainSetPassword(ssidData unsafe.Pointer, password unsafe.Pointer) unsafe.Pointer {
	return _CWKeychainSetPassword(ssidData, password)
	}


// Associates an identity to the SSID and keychain domain you specify.
//
// Added in macOS 10.9.

// Associates an identity to the SSID and keychain domain you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWKeychainSetWiFiEAPIdentity(_:_:_:)
func CWKeychainSetWiFiEAPIdentity(domain unsafe.Pointer, ssid unsafe.Pointer, identity unsafe.Pointer) unsafe.Pointer {
	return _CWKeychainSetWiFiEAPIdentity(domain, ssid, identity)
	}


// Sets the 802.1X username and password for the SSID and keychain domain you specify.
//
// Added in macOS 10.9.

// Sets the 802.1X username and password for the SSID and keychain domain you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWKeychainSetWiFiEAPUsernameAndPassword(_:_:_:_:)
func CWKeychainSetWiFiEAPUsernameAndPassword(domain unsafe.Pointer, ssid unsafe.Pointer, username unsafe.Pointer, password unsafe.Pointer) unsafe.Pointer {
	return _CWKeychainSetWiFiEAPUsernameAndPassword(domain, ssid, username, password)
	}


// Sets the Wi-Fi network keychain password for the SSID and keychain domain you specify.
//
// Added in macOS 10.9.

// Sets the Wi-Fi network keychain password for the SSID and keychain domain you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWKeychainSetWiFiPassword(_:_:_:)
func CWKeychainSetWiFiPassword(domain unsafe.Pointer, ssid unsafe.Pointer, password unsafe.Pointer) unsafe.Pointer {
	return _CWKeychainSetWiFiPassword(domain, ssid, password)
	}


// Merges the specified set of CWNetwork objects.
//
// Added in macOS 10.7.

// Merges the specified set of CWNetwork objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWMergeNetworks(_:)
func CWMergeNetworks(networks unsafe.Pointer) unsafe.Pointer {
	return _CWMergeNetworks(networks)
	}





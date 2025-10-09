// Security Framework Keychain Example using purego bindings
//
// This example demonstrates using the macOS Security framework for keychain
// operations using only purego bindings without cgo.
//
// Features demonstrated:
// - Loading the Security framework
// - Keychain item creation
// - Keychain item queries
// - Password storage and retrieval
// - Proper error handling with Security framework error codes
package main

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/ebitengine/purego"
)

// Security framework types
type (
	OSStatus           int32
	CFTypeRef          unsafe.Pointer
	CFStringRef        CFTypeRef
	CFDictionaryRef    CFTypeRef
	CFDataRef          CFTypeRef
	CFAllocatorRef     unsafe.Pointer
	CFStringEncoding   uint32
	CFIndex            int64
	SecKeychainItemRef CFTypeRef
)

// Security framework constants
const (
	// OSStatus codes
	errSecSuccess          OSStatus = 0
	errSecItemNotFound     OSStatus = -25300
	errSecDuplicateItem    OSStatus = -25299
	errSecAuthFailed       OSStatus = -25293

	// CFString encoding
	kCFStringEncodingUTF8 CFStringEncoding = 0x08000100
)

// Security framework function bindings
var (
	securityLib uintptr
	cfLib       uintptr

	// Security functions
	SecItemAdd          func(attributes CFDictionaryRef, result *CFTypeRef) OSStatus
	SecItemCopyMatching func(query CFDictionaryRef, result *CFTypeRef) OSStatus
	SecItemDelete       func(query CFDictionaryRef) OSStatus

	// CoreFoundation functions for building queries
	CFStringCreateWithCString func(alloc CFAllocatorRef, cStr *byte, encoding CFStringEncoding) CFStringRef
	CFDataCreate              func(allocator CFAllocatorRef, bytes *byte, length CFIndex) CFDataRef
	CFDictionaryCreate        func(allocator CFAllocatorRef, keys *unsafe.Pointer, values *unsafe.Pointer, numValues CFIndex, keyCallbacks unsafe.Pointer, valueCallbacks unsafe.Pointer) CFDictionaryRef
	CFRelease                 func(cf CFTypeRef)
	CFDataGetBytePtr          func(theData CFDataRef) *byte
	CFDataGetLength           func(theData CFDataRef) CFIndex

	// Keychain attribute constants (loaded as symbols)
	kSecClass                CFStringRef
	kSecClassGenericPassword CFStringRef
	kSecAttrAccount          CFStringRef
	kSecAttrService          CFStringRef
	kSecValueData            CFStringRef
	kSecReturnData           CFStringRef
	kSecMatchLimit           CFStringRef
	kSecMatchLimitOne        CFStringRef
	kCFBooleanTrue           CFTypeRef
)

func init() {
	runtime.LockOSThread()

	var err error

	// Load Security framework
	securityLib, err = purego.Dlopen("/System/Library/Frameworks/Security.framework/Security", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(fmt.Sprintf("Failed to load Security framework: %v", err))
	}

	// Load CoreFoundation framework
	cfLib, err = purego.Dlopen("/System/Library/Frameworks/CoreFoundation.framework/CoreFoundation", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(fmt.Sprintf("Failed to load CoreFoundation framework: %v", err))
	}

	// Register Security functions
	purego.RegisterLibFunc(&SecItemAdd, securityLib, "SecItemAdd")
	purego.RegisterLibFunc(&SecItemCopyMatching, securityLib, "SecItemCopyMatching")
	purego.RegisterLibFunc(&SecItemDelete, securityLib, "SecItemDelete")

	// Register CoreFoundation functions
	purego.RegisterLibFunc(&CFStringCreateWithCString, cfLib, "CFStringCreateWithCString")
	purego.RegisterLibFunc(&CFDataCreate, cfLib, "CFDataCreate")
	purego.RegisterLibFunc(&CFDictionaryCreate, cfLib, "CFDictionaryCreate")
	purego.RegisterLibFunc(&CFRelease, cfLib, "CFRelease")
	purego.RegisterLibFunc(&CFDataGetBytePtr, cfLib, "CFDataGetBytePtr")
	purego.RegisterLibFunc(&CFDataGetLength, cfLib, "CFDataGetLength")

	// Load keychain attribute constants
	// These are exported as CFStringRef pointers, so we need to dereference them
	loadSymbol := func(name string) CFStringRef {
		sym, err := purego.Dlsym(securityLib, name)
		if err != nil {
			panic(fmt.Sprintf("Failed to load symbol %s: %v", name, err))
		}
		// Dereference the pointer to get the actual CFStringRef
		return CFStringRef(*(*unsafe.Pointer)(unsafe.Pointer(sym)))
	}

	kSecClass = loadSymbol("kSecClass")
	kSecClassGenericPassword = loadSymbol("kSecClassGenericPassword")
	kSecAttrAccount = loadSymbol("kSecAttrAccount")
	kSecAttrService = loadSymbol("kSecAttrService")
	kSecValueData = loadSymbol("kSecValueData")
	kSecReturnData = loadSymbol("kSecReturnData")
	kSecMatchLimit = loadSymbol("kSecMatchLimit")
	kSecMatchLimitOne = loadSymbol("kSecMatchLimitOne")

	// Load kCFBooleanTrue from CoreFoundation
	trueSym, err := purego.Dlsym(cfLib, "kCFBooleanTrue")
	if err != nil {
		panic(fmt.Sprintf("Failed to load kCFBooleanTrue: %v", err))
	}
	kCFBooleanTrue = CFTypeRef(*(*unsafe.Pointer)(unsafe.Pointer(trueSym)))
}

func cString(s string) *byte {
	b := append([]byte(s), 0)
	return &b[0]
}

func createCFString(s string) CFStringRef {
	return CFStringCreateWithCString(nil, cString(s), kCFStringEncodingUTF8)
}

func createCFData(data []byte) CFDataRef {
	if len(data) == 0 {
		return CFDataCreate(nil, nil, 0)
	}
	return CFDataCreate(nil, &data[0], CFIndex(len(data)))
}

func createCFDictionary(keys, values []CFTypeRef) CFDictionaryRef {
	if len(keys) != len(values) {
		panic("keys and values must have same length")
	}
	if len(keys) == 0 {
		return nil
	}

	keyPtrs := make([]unsafe.Pointer, len(keys))
	valPtrs := make([]unsafe.Pointer, len(values))
	for i := range keys {
		keyPtrs[i] = unsafe.Pointer(keys[i])
		valPtrs[i] = unsafe.Pointer(values[i])
	}

	return CFDictionaryCreate(
		nil,
		&keyPtrs[0],
		&valPtrs[0],
		CFIndex(len(keys)),
		nil, // Use default key callbacks
		nil, // Use default value callbacks
	)
}

func cfDataToBytes(data CFDataRef) []byte {
	if data == nil {
		return nil
	}
	length := CFDataGetLength(data)
	if length == 0 {
		return []byte{}
	}
	ptr := CFDataGetBytePtr(data)
	return unsafe.Slice(ptr, length)
}

func statusString(status OSStatus) string {
	switch status {
	case errSecSuccess:
		return "Success"
	case errSecItemNotFound:
		return "Item not found"
	case errSecDuplicateItem:
		return "Duplicate item"
	case errSecAuthFailed:
		return "Authentication failed"
	default:
		return fmt.Sprintf("Error %d", status)
	}
}

func main() {
	fmt.Println("=== macOS Security Framework Keychain Example ===\n")

	service := "com.example.testapp"
	account := "testuser"
	password := "MySecretPassword123!"

	// Example 1: Add a password to the keychain
	fmt.Println("1. Adding password to keychain:")
	fmt.Printf("   Service: %s\n", service)
	fmt.Printf("   Account: %s\n", account)

	serviceStr := createCFString(service)
	accountStr := createCFString(account)
	passwordData := createCFData([]byte(password))

	defer CFRelease(CFTypeRef(serviceStr))
	defer CFRelease(CFTypeRef(accountStr))
	defer CFRelease(CFTypeRef(passwordData))

	// Build attributes dictionary
	keys := []CFTypeRef{
		CFTypeRef(kSecClass),
		CFTypeRef(kSecAttrService),
		CFTypeRef(kSecAttrAccount),
		CFTypeRef(kSecValueData),
	}
	values := []CFTypeRef{
		CFTypeRef(kSecClassGenericPassword),
		CFTypeRef(serviceStr),
		CFTypeRef(accountStr),
		CFTypeRef(passwordData),
	}

	attributes := createCFDictionary(keys, values)
	defer CFRelease(CFTypeRef(attributes))

	// First, delete any existing item
	queryKeys := []CFTypeRef{
		CFTypeRef(kSecClass),
		CFTypeRef(kSecAttrService),
		CFTypeRef(kSecAttrAccount),
	}
	queryValues := []CFTypeRef{
		CFTypeRef(kSecClassGenericPassword),
		CFTypeRef(serviceStr),
		CFTypeRef(accountStr),
	}
	query := createCFDictionary(queryKeys, queryValues)
	SecItemDelete(query)
	CFRelease(CFTypeRef(query))

	// Add the item
	status := SecItemAdd(attributes, nil)
	if status == errSecSuccess {
		fmt.Println("   ✅ Password added successfully")
	} else {
		fmt.Printf("   ❌ Failed to add password: %s\n", statusString(status))
		return
	}
	fmt.Println()

	// Example 2: Retrieve the password from the keychain
	fmt.Println("2. Retrieving password from keychain:")

	// Build query dictionary with return data flag
	retrieveKeys := []CFTypeRef{
		CFTypeRef(kSecClass),
		CFTypeRef(kSecAttrService),
		CFTypeRef(kSecAttrAccount),
		CFTypeRef(kSecReturnData),
		CFTypeRef(kSecMatchLimit),
	}
	retrieveValues := []CFTypeRef{
		CFTypeRef(kSecClassGenericPassword),
		CFTypeRef(serviceStr),
		CFTypeRef(accountStr),
		kCFBooleanTrue,
		CFTypeRef(kSecMatchLimitOne),
	}

	retrieveQuery := createCFDictionary(retrieveKeys, retrieveValues)
	defer CFRelease(CFTypeRef(retrieveQuery))

	var result CFTypeRef
	status = SecItemCopyMatching(retrieveQuery, &result)

	if status == errSecSuccess {
		defer CFRelease(result)
		retrievedPassword := cfDataToBytes(CFDataRef(result))
		fmt.Printf("   ✅ Password retrieved: %s\n", string(retrievedPassword))

		if string(retrievedPassword) == password {
			fmt.Println("   ✅ Retrieved password matches original!")
		}
	} else {
		fmt.Printf("   ❌ Failed to retrieve password: %s\n", statusString(status))
	}
	fmt.Println()

	// Example 3: Delete the password from the keychain
	fmt.Println("3. Cleaning up - deleting password from keychain:")

	deleteQuery := createCFDictionary(queryKeys, queryValues)
	defer CFRelease(CFTypeRef(deleteQuery))

	status = SecItemDelete(deleteQuery)
	if status == errSecSuccess {
		fmt.Println("   ✅ Password deleted successfully")
	} else {
		fmt.Printf("   ❌ Failed to delete password: %s\n", statusString(status))
	}
	fmt.Println()

	fmt.Println("✅ Security framework example completed successfully!")
	fmt.Println("   Operations demonstrated:")
	fmt.Println("   - Adding passwords to keychain")
	fmt.Println("   - Retrieving passwords from keychain")
	fmt.Println("   - Deleting keychain items")
	fmt.Println("   - Proper CoreFoundation memory management")
}

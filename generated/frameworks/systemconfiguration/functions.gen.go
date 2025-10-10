// Code generated from Apple documentation for SystemConfiguration. DO NOT EDIT.

package systemconfiguration

// SystemConfiguration Functions
//
// This file contains function declarations discovered from Apple's documentation.
// To use these functions, you need to:
//   1. Map C types to Go types
//   2. Create function variables
//   3. Register them with purego.RegisterLibFunc
//
// Example:
//   var CGContextSetRGBFillColor func(c CGContextRef, red, green, blue, alpha CGFloat)
//   purego.RegisterLibFunc(&CGContextSetRGBFillColor, lib, "CGContextSetRGBFillColor")

// Discovered functions (15 total):

// CNCopyCurrentNetworkInfo(CFStringRef interfaceName);) CFDictionaryRef
//
// Availability:
//   - iOS 4.1+ (Deprecated in 14.0)
//   - iPadOS 4.1+ (Deprecated in 14.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//
// Deprecated: This function is deprecated.

// CNCopyCurrentNetworkInfo(interfaceName CFStringRef, );) CFDictionaryRef
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 4.1+ (Deprecated in 14.0)
//   - iPadOS 4.1+ (Deprecated in 14.0)
//
// Deprecated: This function is deprecated.

// DHCPClientPreferencesSetApplicationOptions(applicationID CFStringRef, options ,  const  UInt8 *, count ,  CFIndex, );) Boolean
//
// Availability:
//   - macOS 10.1+


// DHCPInfoGetOptionData(info CFDictionaryRef, code ,  UInt8, );) CFDataRef
//
// Availability:
//   - macOS 10.1+

// SCDynamicStoreCopyLocation(store _, :  SCDynamicStore?) ->  CFString?) func
//
// Availability:
//   - macOS 10.1+

// SCDynamicStoreGetTypeID() func
//
// Availability:
//   - macOS 10.1+


// SCDynamicStoreKeyCreateComputerName(allocator _, :  CFAllocator?) ->  CFString) func
//
// Availability:
//   - macOS 10.1+

// SCNetworkCheckReachabilityByName(nodename const  char *, flags ,  SCNetworkConnectionFlags *, );) Boolean
//
// Availability:
//   - macOS 10.1+ (Deprecated in 10.6)
//
// Deprecated: This function is deprecated.

// SCNetworkConnectionUnscheduleFromRunLoop(connection _, runLoop :  SCNetworkConnection,  _, runLoopMode :  CFRunLoop,  _, :  CFString) ->  Bool) func
//
// Availability:
//   - macOS 10.3+


// SCNetworkReachabilityCreateWithAddressPair(allocator _, localAddress :  CFAllocator?,  _, remoteAddress :  UnsafePointer< sockaddr>?,  _, :  UnsafePointer< sockaddr>?) ->  SCNetworkReachability?) func
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 17.4)
//   - iOS 2.0+ (Deprecated in 17.4)
//   - iPadOS 2.0+ (Deprecated in 17.4)
//   - macOS 10.3+ (Deprecated in 14.4)
//   - visionOS 1.0+ (Deprecated in 1.1)
//
// Deprecated: This function is deprecated.

// SCNetworkReachabilityGetFlags(target _, flags :  SCNetworkReachability,  _, :  UnsafeMutablePointer< SCNetworkReachabilityFlags>) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 17.4)
//   - iOS 2.0+ (Deprecated in 17.4)
//   - iPadOS 2.0+ (Deprecated in 17.4)
//   - macOS 10.3+ (Deprecated in 14.4)
//   - visionOS 1.0+ (Deprecated in 1.1)
//
// Deprecated: This function is deprecated.

// SCPreferencesCommitChanges(prefs _, :  SCPreferences) ->  Bool) func
//
// Availability:
//   - macOS 10.1+


// SCPreferencesCreate(allocator _, name :  CFAllocator?,  _, prefsID :  CFString,  _, :  CFString?) ->  SCPreferences?) func
//
// Availability:
//   - macOS 10.1+

// SCPreferencesPathGetValue(prefs _, path :  SCPreferences,  _, :  CFString) ->  CFDictionary?) func
//
// Availability:
//   - macOS 10.1+

// SCPreferencesSetComputerName(prefs _, name :  SCPreferences,  _, nameEncoding :  CFString?,  _, :  CFStringEncoding) ->  Bool) func
//
// Availability:
//   - macOS 10.1+



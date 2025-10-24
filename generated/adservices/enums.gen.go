// Code generated from Apple documentation for AdServices. DO NOT EDIT.

package adservices

/* debug [enums.gen.go]: Generating 1 enums for AdServices */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum AAAttributionErrorCode (3 cases) */
// AAAttributionErrorCode - The error code that the parent class issues.
//
// [Full Topic]: https://developer.apple.com/documentation/AdServices/AAAttributionError/Code
type AAAttributionErrorCode uint

const (
	// AAAttributionErrorCodeInternalError - The server is unable to provide a token because of an internal error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AdServices/AAAttributionError/Code/internalError
	AAAttributionErrorCodeInternalError AAAttributionErrorCode = 0
	// AAAttributionErrorCodeNetworkError - The server is unable to provide a token because the internet isn’t available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AdServices/AAAttributionError/Code/networkError
	AAAttributionErrorCodeNetworkError AAAttributionErrorCode = 0
	// AAAttributionErrorCodePlatformNotSupported - The server is unable to provide a token because of an unsupported operating system.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AdServices/AAAttributionError/Code/platformNotSupported
	AAAttributionErrorCodePlatformNotSupported AAAttributionErrorCode = 0
)



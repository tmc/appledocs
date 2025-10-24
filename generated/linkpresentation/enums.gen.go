// Code generated from Apple documentation for LinkPresentation. DO NOT EDIT.

package linkpresentation

/* debug [enums.gen.go]: Generating 1 enums for LinkPresentation */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum LPErrorCode (5 cases) */
// LPErrorCode - Possible error values that can be returned from LinkPresentation APIs.
//
// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPError/Code
type LPErrorCode uint

const (
	// LPErrorMetadataFetchCancelled - An error indicating that the metadata fetch was canceled by the client.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPError/Code/metadataFetchCancelled
	LPErrorMetadataFetchCancelled LPErrorCode = 0
	// LPErrorMetadataFetchFailed - An error indicating that a metadata fetch failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPError/Code/metadataFetchFailed
	LPErrorMetadataFetchFailed LPErrorCode = 0
	// LPErrorMetadataFetchNotAllowed - An error indicating that the metadata fetch was not allowed due to system policies.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPError/Code/metadataFetchNotAllowed
	LPErrorMetadataFetchNotAllowed LPErrorCode = 0
	// LPErrorMetadataFetchTimedOut - An error indicating that the metadata fetch took longer than allowed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPError/Code/metadataFetchTimedOut
	LPErrorMetadataFetchTimedOut LPErrorCode = 0
	// LPErrorUnknown - An unknown error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LinkPresentation/LPError/Code/unknown
	LPErrorUnknown LPErrorCode = 0
)



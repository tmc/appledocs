// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

// PURLHandleClient is the NSURLHandleClient protocol interface.
//
// The interface implemented by URL handle clients.
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.4)
//
// See: doc://com.apple.foundation/documentation/Foundation/NSURLHandleClient
type PURLHandleClient interface {
	// Required methods
	URLHandleResourceDataDidBecomeAvailable(sender IURLHandle, newBytes IData)
	URLHandleResourceDidFailLoadingWithReason(sender IURLHandle, reason IString)
	URLHandleResourceDidBeginLoading(sender IURLHandle)
	URLHandleResourceDidCancelLoading(sender IURLHandle)
	URLHandleResourceDidFinishLoading(sender IURLHandle)
}

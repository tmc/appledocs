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
	URLHandleResourceDataDidBecomeAvailable(sender IURLHandle, newBytes IData)/* debug [protocol_interface/required_method]: URLHandleResourceDataDidBecomeAvailable */
	URLHandleResourceDidFailLoadingWithReason(sender IURLHandle, reason IString)/* debug [protocol_interface/required_method]: URLHandleResourceDidFailLoadingWithReason */
	URLHandleResourceDidBeginLoading(sender IURLHandle)/* debug [protocol_interface/required_method]: URLHandleResourceDidBeginLoading */
	URLHandleResourceDidCancelLoading(sender IURLHandle)/* debug [protocol_interface/required_method]: URLHandleResourceDidCancelLoading */
	URLHandleResourceDidFinishLoading(sender IURLHandle)/* debug [protocol_interface/required_method]: URLHandleResourceDidFinishLoading */
}

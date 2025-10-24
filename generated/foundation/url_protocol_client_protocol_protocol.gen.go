// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

// PURLProtocolClient is the NSURLProtocolClient protocol interface.
//
// The interface used by   subclasses to communicate with the URL Loading System.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// See: doc://com.apple.foundation/documentation/Foundation/URLProtocolClient
type PURLProtocolClient interface {
	// Required methods
	URLProtocolCachedResponseIsValid(protocol_ IURLProtocol, cachedResponse ICachedURLResponse)
	URLProtocolDidCancelAuthenticationChallenge(protocol_ IURLProtocol, challenge IURLAuthenticationChallenge)
	URLProtocolDidFailWithError(protocol_ IURLProtocol, error_ IError)
	URLProtocolDidLoadData(protocol_ IURLProtocol, data IData)
	URLProtocolDidReceiveAuthenticationChallenge(protocol_ IURLProtocol, challenge IURLAuthenticationChallenge)
	URLProtocolDidReceiveResponseCacheStoragePolicy(protocol_ IURLProtocol, response IURLResponse, policy URLCacheStoragePolicy /* not a class type */)
	URLProtocolWasRedirectedToRequestRedirectResponse(protocol_ IURLProtocol, request IURLRequest, redirectResponse IURLResponse)
	URLProtocolDidFinishLoading(protocol_ IURLProtocol)
}

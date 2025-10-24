// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

// PHTTPCookieStoreObserver is the WKHTTPCookieStoreObserver protocol interface.
//
// The methods to adopt in an object that monitors changes to a webpage’s cookies.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.13+
//   - visionOS 1.0+
//
// See: doc://com.apple.webkit/documentation/WebKit/WKHTTPCookieStoreObserver
type PHTTPCookieStoreObserver interface {
	// Optional methods
	CookiesDidChangeInCookieStore(cookieStore IWKHTTPCookieStore)
	HasCookiesDidChangeInCookieStore() bool
}

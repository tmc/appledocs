// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"unsafe"
)

// PURLSchemeHandler is the WKURLSchemeHandler protocol interface.
//
// A protocol for loading resources with URL schemes that WebKit doesn’t handle.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.13+
//   - visionOS 1.0+
//
// See: doc://com.apple.webkit/documentation/WebKit/WKURLSchemeHandler
type PURLSchemeHandler interface {
	// Required methods
	WebViewStartURLSchemeTask(webView IWKWebView, urlSchemeTask unsafe.Pointer)/* debug [protocol_interface/required_method]: WebViewStartURLSchemeTask */
	WebViewStopURLSchemeTask(webView IWKWebView, urlSchemeTask unsafe.Pointer)/* debug [protocol_interface/required_method]: WebViewStopURLSchemeTask */
}

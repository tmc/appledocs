// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewURL

// ExampleNewURLAbsoluteURLWithDataRepresentationRelativeToURL demonstrates how to create a URL instance using NewURLAbsoluteURLWithDataRepresentationRelativeToURL.
func ExampleNewURLAbsoluteURLWithDataRepresentationRelativeToURL() {
	_ = foundation.NewURLAbsoluteURLWithDataRepresentationRelativeToURL(
		foundation.NSData{}, // data NSData
		foundation.URL{}, // baseURL URL
	)
	// Output:
}
// ExampleNewURLByResolvingAliasFileAtURLOptionsError demonstrates how to create a URL instance using NewURLByResolvingAliasFileAtURLOptionsError.
// Returns a new URL made by resolving the alias file at  .
func ExampleNewURLByResolvingAliasFileAtURLOptionsError() {
	_ = foundation.NewURLByResolvingAliasFileAtURLOptionsError(
		foundation.URL{}, // url URL
		foundation.URLBookmarkResolutionOptions{}, // options URLBookmarkResolutionOptions
		foundation.NSError{}, // error NSError
	)
	// Output:
}
// ExampleNewURLWithDataRepresentationRelativeToURL demonstrates how to create a URL instance using NewURLWithDataRepresentationRelativeToURL.
func ExampleNewURLWithDataRepresentationRelativeToURL() {
	_ = foundation.NewURLWithDataRepresentationRelativeToURL(
		foundation.NSData{}, // data NSData
		foundation.URL{}, // baseURL URL
	)
	// Output:
}

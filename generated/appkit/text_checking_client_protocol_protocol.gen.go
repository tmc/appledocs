// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

// PTextCheckingClient is the NSTextCheckingClient protocol interface.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSTextCheckingClient
type PTextCheckingClient interface {
	// Required methods
	ReplaceCharactersInRangeWithAnnotatedString(range_ corefoundation.Range, annotatedString foundation.AttributedString)
}

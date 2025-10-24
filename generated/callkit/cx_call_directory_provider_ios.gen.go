//go:build darwin && ios

// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CXCallDirectoryProvider


// Tells the extension to prepare for a host app’s request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryProvider/beginRequest(with:)
func (c_ CXCallDirectoryProvider) BeginRequestWithExtensionContext(context ICXCallDirectoryExtensionContext) {
	objc.Send[objc.ID](c_.ID, objc.Sel("beginRequestWithExtensionContext:"), context)
}

// iOS-only properties






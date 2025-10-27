//go:build darwin && ios

// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for NEFilterFlow


// iOS-only properties

// A string containing the identifier of the source app of the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterFlow/sourceAppIdentifier
func (n_ NEFilterFlow) SourceAppIdentifier() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("sourceAppIdentifier"))
	return rv
}

// A byte string that uniquely identifies the binary for each build of the app that is the source of the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterFlow/sourceAppUniqueIdentifier
func (n_ NEFilterFlow) SourceAppUniqueIdentifier() foundation.foundation.INSData {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("sourceAppUniqueIdentifier"))
	return rv
}

// The short version string of the app that is the source of the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterFlow/sourceAppVersion
func (n_ NEFilterFlow) SourceAppVersion() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("sourceAppVersion"))
	return rv
}






//go:build darwin && ios

// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for ExtensionContext


// iOS-only properties

// The minimum size for a Siri hosted view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/hostedViewMinimumAllowedSize
func (e_ ExtensionContext) HostedViewMinimumAllowedSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](e_.ID, objc.Sel("hostedViewMinimumAllowedSize"))
	return rv
}






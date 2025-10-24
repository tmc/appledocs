//go:build darwin && ios

// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for ExtensionContext


// iOS-only properties

// The minimum size for a Siri hosted view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext/hostedViewMinimumAllowedSize
func (e_ ExtensionContext) HostedViewMinimumAllowedSize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[objc.ID](e_.ID, objc.Sel("hostedViewMinimumAllowedSize"))
	return rv
}






//go:build darwin && ios

// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for Layer


// iOS-only properties

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/wantsDynamicContentScaling
func (l_ Layer) WantsDynamicContentScaling() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("wantsDynamicContentScaling"))
	return rv
}
func (l_ Layer) SetWantsDynamicContentScaling(value bool) {
	l_.ID.Send(objc.RegisterName("setWantsDynamicContentScaling:"), value)
}





//go:build darwin && ios

// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for EAGLLayer


// iOS-only properties

// A Boolean value that determines whether the layer presents its content using a Core Animation transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEAGLLayer/presentsWithTransaction
func (e_ EAGLLayer) PresentsWithTransaction() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("presentsWithTransaction"))
	return rv
}
func (e_ EAGLLayer) SetPresentsWithTransaction(value bool) {
	e_.ID.Send(objc.RegisterName("setPresentsWithTransaction:"), value)
}






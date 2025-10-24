//go:build darwin && ios

// Code generated from Apple documentation for VisionKit. DO NOT EDIT.

package visionkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for DocumentCameraScan


// iOS-only properties

// The title of the scanned document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/VisionKit/VNDocumentCameraScan/title
func (d_ DocumentCameraScan) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("title"))
	return rv
}






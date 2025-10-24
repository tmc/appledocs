//go:build darwin && ios

// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for PreviewSceneOptions


// iOS-only properties

// The index of the item to preview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewSceneActivationConfiguration/Options/initialPreviewIndex
func (p_ PreviewSceneOptions) InitialPreviewIndex() int {
	rv := objc.Send[int](p_.ID, objc.Sel("initialPreviewIndex"))
	return rv
}
func (p_ PreviewSceneOptions) SetInitialPreviewIndex(value int) {
	p_.ID.Send(objc.RegisterName("setInitialPreviewIndex:"), value)
}






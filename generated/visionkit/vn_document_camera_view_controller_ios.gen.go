//go:build darwin && ios

// Code generated from Apple documentation for VisionKit. DO NOT EDIT.

package visionkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for DocumentCameraViewController

// iOS-only properties

// The delegate to be notified when the user saves or cancels the document scanner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/VisionKit/VNDocumentCameraViewController/delegate
func (d_ DocumentCameraViewController) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("delegate"))
	return rv
}
func (d_ DocumentCameraViewController) SetDelegate(value unsafe.Pointer) {
	d_.ID.Send(objc.RegisterName("setDelegate:"), value)
}

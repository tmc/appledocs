//go:build darwin && ios

// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for ToolPickerLassoItem


// iOS-only properties

// A lasso tool for selecting parts of a drawing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerLassoItem/lassoTool-1urgb
func (t_ ToolPickerLassoItem) LassoTool() IPKLassoTool {
	rv := objc.Send[LassoTool](t_.ID, objc.Sel("lassoTool"))
	return rv
}





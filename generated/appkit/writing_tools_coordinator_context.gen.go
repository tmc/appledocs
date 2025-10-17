// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [WritingToolsCoordinatorContext] class.
var writingToolsCoordinatorContextClass = _WritingToolsCoordinatorContextClass{objc.GetClass("NSWritingToolsCoordinatorContext")}

type _WritingToolsCoordinatorContextClass struct {
	class objc.Class
}

// A data object that you use to share your custom view’s text with Writing Tools. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/Context

type WritingToolsCoordinatorContext struct {
	objectivec.Object
}

// WritingToolsCoordinatorContextFrom constructs a [WritingToolsCoordinatorContext] from an unsafe.Pointer.
//
// A data object that you use to share your custom view’s text with Writing Tools.
func WritingToolsCoordinatorContextFrom(ptr unsafe.Pointer) WritingToolsCoordinatorContext {
	return WritingToolsCoordinatorContext{objectivec.Object{objc.ID(ptr)}}
}




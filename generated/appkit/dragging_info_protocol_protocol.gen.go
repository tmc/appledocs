// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// PDraggingInfo is the NSDraggingInfo protocol interface.
//
// A set of methods that supply information about a dragging session.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSDraggingInfo
type PDraggingInfo interface {
	// Required methods
	EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock(enumOpts DraggingItemEnumerationOptions, view IView, classArray []objc.Class, searchOptions foundation.IDictionary, block unsafe.Pointer)
	NamesOfPromisedFilesDroppedAtDestination(dropDestination objc.IObject /* cross-framework: NSURL */) []string
	ResetSpringLoading()
	SlideDraggedImageTo(screenPoint objc.IObject /* cross-framework: Point */)
}

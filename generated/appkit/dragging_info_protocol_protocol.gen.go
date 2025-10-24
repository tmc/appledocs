// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/vision"
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
	EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock(enumOpts DraggingItemEnumerationOptions, view IView, classArray []objc.Class, searchOptions foundation.IDictionary, block unsafe.Pointer)/* debug [protocol_interface/required_method]: EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock */
	NamesOfPromisedFilesDroppedAtDestination(dropDestination objc.IObject /* cross-framework: NSURL */) []string/* debug [protocol_interface/required_method]: NamesOfPromisedFilesDroppedAtDestination */
	ResetSpringLoading()/* debug [protocol_interface/required_method]: ResetSpringLoading */
	SlideDraggedImageTo(screenPoint vision.Point)/* debug [protocol_interface/required_method]: SlideDraggedImageTo */
}

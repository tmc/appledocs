// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"
)

// PViewToolTipOwner is the NSViewToolTipOwner protocol interface.
//
// A set of methods for dynamically associating a tool tip with a view.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSViewToolTipOwner
type PViewToolTipOwner interface {
	// Required methods
	ViewStringForToolTipPointUserData(view IView, tag objc.IObject /* cross-framework: ToolTipTag */, point objc.IObject /* cross-framework: Point */, data unsafe.Pointer) foundation.String
}

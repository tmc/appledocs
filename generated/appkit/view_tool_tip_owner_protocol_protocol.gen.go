// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
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
	ViewStringForToolTipPointUserData(view IView, tag ToolTipTag, point corefoundation.CGPoint, data objectivec.IObject) foundation.String
}

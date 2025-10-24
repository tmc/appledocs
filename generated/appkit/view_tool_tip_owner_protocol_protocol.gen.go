// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"

	"github.com/tmc/appledocs/generated/vision"
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
	ViewStringForToolTipPointUserData(view IView, tag ToolTipTag /* typedef */, point vision.Point, data objectivec.IObject) foundation.String/* debug [protocol_interface/required_method]: ViewStringForToolTipPointUserData */
}

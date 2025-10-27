// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PAction is the CAAction protocol interface.
//
// An interface that allows instances to respond to actions triggered by a Core Animation layer change.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.quartzcore/documentation/QuartzCore/CAAction
type PAction interface {
	// Required methods
	RunActionForKeyObjectArguments(event foundation.foundation.INSString, anObject objectivec.IObject, dict foundation.foundation.INSDictionary)
}

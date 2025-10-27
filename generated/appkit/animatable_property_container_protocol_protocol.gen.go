// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PAnimatablePropertyContainer is the NSAnimatablePropertyContainer protocol interface.
//
// A set of methods that defines a way to add animation to an existing class with a minimum of API impact.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.13+
//   - iPadOS 10.13+
//   - macOS 10.5+
//
// See: doc://com.apple.appkit/documentation/AppKit/NSAnimatablePropertyContainer
type PAnimatablePropertyContainer interface {
	// Required methods
	AnimationForKey(key AnimatablePropertyKey) objc.ID
	Animator() objectivec.IObject
}

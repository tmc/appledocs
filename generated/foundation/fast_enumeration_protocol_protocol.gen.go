// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PFastEnumeration is the NSFastEnumeration protocol interface.
//
// A protocol that objects adopt to support fast enumeration.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// See: doc://com.apple.foundation/documentation/Foundation/NSFastEnumeration
type PFastEnumeration interface {
	// Required methods
	CountByEnumeratingWithStateObjectsCount(state objc.IObject /* cross-framework: FastEnumerationState */, buffer []objc.ID, len_ uint) uint
}

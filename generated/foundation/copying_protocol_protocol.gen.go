// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PCopying is the NSCopying protocol interface.
//
// A protocol that objects adopt to provide functional copies of themselves.
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
// See: doc://com.apple.foundation/documentation/Foundation/NSCopying
type PCopying interface {
	// Required methods
	CopyWithZone(zone Zone /* not a class type */) objc.ID
}

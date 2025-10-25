// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objectivec"
)

// PCoding is the NSCoding protocol interface.
//
// A protocol that enables an object to be encoded and decoded for archiving and distribution.
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
// See: doc://com.apple.foundation/documentation/Foundation/NSCoding
type PCoding interface {
	// Required methods
	EncodeWithCoder(coder ICoder)/* debug [protocol_interface/required_method]: EncodeWithCoder */
	InitWithCoder(coder ICoder) objectivec.IObject/* debug [protocol_interface/required_method]: InitWithCoder */
}

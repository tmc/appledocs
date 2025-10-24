// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PDynamicLibrary is the MTLDynamicLibrary protocol interface.
//
// A dynamically linkable representation of compiled shader code for a specific Metal device object.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLDynamicLibrary
type PDynamicLibrary interface {
	// Required methods
	SerializeToURLError(url objc.IObject /* cross-framework: NSURL */, error_ objectivec.IObject) bool/* debug [protocol_interface/required_method]: SerializeToURLError */
}

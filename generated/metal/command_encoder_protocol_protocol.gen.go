// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PCommandEncoder is the MTLCommandEncoder protocol interface.
//
// An encoder that writes GPU commands into a command buffer.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.11+
//   - tvOS +
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLCommandEncoder
type PCommandEncoder interface {
	// Required methods
	BarrierAfterQueueStagesBeforeStages(afterQueueStages Stages, beforeStages Stages)/* debug [protocol_interface/required_method]: BarrierAfterQueueStagesBeforeStages */
	EndEncoding()/* debug [protocol_interface/required_method]: EndEncoding */
	InsertDebugSignpost(string_ objc.IObject /* cross-framework: NSString */)/* debug [protocol_interface/required_method]: InsertDebugSignpost */
	PopDebugGroup()/* debug [protocol_interface/required_method]: PopDebugGroup */
	PushDebugGroup(string_ objc.IObject /* cross-framework: NSString */)/* debug [protocol_interface/required_method]: PushDebugGroup */
}

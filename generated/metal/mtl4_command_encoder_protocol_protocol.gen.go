// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PMTL4CommandEncoder is the MTL4CommandEncoder protocol interface.
//
// An encoder that writes GPU commands into a command buffer.
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTL4CommandEncoder
type PMTL4CommandEncoder interface {
	// Required methods
	BarrierAfterEncoderStagesBeforeEncoderStagesVisibilityOptions(afterEncoderStages Stages, beforeEncoderStages Stages, visibilityOptions MTL4VisibilityOptions)
	BarrierAfterQueueStagesBeforeStagesVisibilityOptions(afterQueueStages Stages, beforeStages Stages, visibilityOptions MTL4VisibilityOptions)
	BarrierAfterStagesBeforeQueueStagesVisibilityOptions(afterStages Stages, beforeQueueStages Stages, visibilityOptions MTL4VisibilityOptions)
	EndEncoding()
	InsertDebugSignpost(string_ foundation.foundation.INSString)
	PopDebugGroup()
	PushDebugGroup(string_ foundation.foundation.INSString)
	UpdateFenceAfterEncoderStages(fence unsafe.Pointer, afterEncoderStages Stages)
	WaitForFenceBeforeEncoderStages(fence unsafe.Pointer, beforeEncoderStages Stages)
}

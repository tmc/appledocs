// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PMTL4PipelineDataSetSerializer is the MTL4PipelineDataSetSerializer protocol interface.
//
// A fast-addition container for collecting data during pipeline state creation.
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTL4PipelineDataSetSerializer
type PMTL4PipelineDataSetSerializer interface {
	// Required methods
	SerializeAsArchiveAndFlushToURLError(url foundation.foundation.INSURL, error_ foundation.foundation.INSError) bool
	SerializeAsPipelinesScriptWithError(error_ foundation.foundation.INSError) foundation.Data
}

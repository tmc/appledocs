// Code generated from Apple documentation for ModelIO. DO NOT EDIT.

package modelio

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PMDLMeshBuffer is the MDLMeshBuffer protocol interface.
//
// The general interface for managing storage of vertex and index data used in loading, processing, and rendering meshes.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.modelio/documentation/ModelIO/MDLMeshBuffer
type PMDLMeshBuffer interface {
	// Required methods
	FillDataOffset(data objc.IObject /* cross-framework: NSData */, offset uint)/* debug [protocol_interface/required_method]: FillDataOffset */
	Map() unsafe.Pointer/* debug [protocol_interface/required_method]: Map */
}

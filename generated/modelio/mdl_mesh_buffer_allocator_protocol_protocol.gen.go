// Code generated from Apple documentation for ModelIO. DO NOT EDIT.

package modelio

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PMDLMeshBufferAllocator is the MDLMeshBufferAllocator protocol interface.
//
// The general interface for managing allocation of data buffers to be used in loading, processing, and rendering meshes.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.modelio/documentation/ModelIO/MDLMeshBufferAllocator
type PMDLMeshBufferAllocator interface {
	// Required methods
	NewBufferType(length uint, type_ MDLMeshBufferType) unsafe.Pointer/* debug [protocol_interface/required_method]: NewBufferType */
	NewBufferFromZoneDataType(zone unsafe.Pointer, data objc.IObject /* cross-framework: NSData */, type_ MDLMeshBufferType) unsafe.Pointer/* debug [protocol_interface/required_method]: NewBufferFromZoneDataType */
	NewBufferFromZoneLengthType(zone unsafe.Pointer, length uint, type_ MDLMeshBufferType) unsafe.Pointer/* debug [protocol_interface/required_method]: NewBufferFromZoneLengthType */
	NewBufferWithDataType(data objc.IObject /* cross-framework: NSData */, type_ MDLMeshBufferType) unsafe.Pointer/* debug [protocol_interface/required_method]: NewBufferWithDataType */
	NewZone(capacity uint) unsafe.Pointer/* debug [protocol_interface/required_method]: NewZone */
	NewZoneForBuffersWithSizeAndType(sizes []foundation.Number, types []foundation.Number) unsafe.Pointer/* debug [protocol_interface/required_method]: NewZoneForBuffersWithSizeAndType */
}

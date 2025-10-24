// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"
)

// PRasterizationRateMap is the MTLRasterizationRateMap protocol interface.
//
// A compiled read-only instance that determines how to apply variable rasterization rates when rendering.
//
// Availability:
//   - Mac Catalyst 13.4+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15.4+
//   - tvOS 16.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLRasterizationRateMap
type PRasterizationRateMap interface {
	// Required methods
	CopyParameterDataToBufferOffset(buffer unsafe.Pointer, offset uint)/* debug [protocol_interface/required_method]: CopyParameterDataToBufferOffset */
	MapScreenToPhysicalCoordinatesForLayer(screenCoordinates Coordinate2D /* typedef */, layerIndex uint) Coordinate2D/* debug [protocol_interface/required_method]: MapScreenToPhysicalCoordinatesForLayer */
	PhysicalSizeForLayer(layerIndex uint) MTLSize/* debug [protocol_interface/required_method]: PhysicalSizeForLayer */
	MapPhysicalToScreenCoordinatesForLayer(physicalCoordinates Coordinate2D /* typedef */, layerIndex uint) Coordinate2D/* debug [protocol_interface/required_method]: MapPhysicalToScreenCoordinatesForLayer */
}

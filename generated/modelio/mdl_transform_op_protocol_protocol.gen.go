// Code generated from Apple documentation for ModelIO. DO NOT EDIT.

package modelio

import (
	"unsafe"
)

// PMDLTransformOp is the MDLTransformOp protocol interface.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.13+
//   - tvOS 11.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.modelio/documentation/ModelIO/MDLTransformOp
type PMDLTransformOp interface {
	// Required methods
	Double4x4AtTime(time float64) unsafe.Pointer/* debug [protocol_interface/required_method]: Double4x4AtTime */
	Float4x4AtTime(time float64) unsafe.Pointer/* debug [protocol_interface/required_method]: Float4x4AtTime */
	IsInverseOp() bool/* debug [protocol_interface/required_method]: IsInverseOp */
}

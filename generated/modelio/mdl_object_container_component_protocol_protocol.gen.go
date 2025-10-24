// Code generated from Apple documentation for ModelIO. DO NOT EDIT.

package modelio

import (
	"unsafe"
)

// PMDLObjectContainerComponent is the MDLObjectContainerComponent protocol interface.
//
// The general interface for classes that can act as containers in an object hierarchy.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.modelio/documentation/ModelIO/MDLObjectContainerComponent
type PMDLObjectContainerComponent interface {
	// Required methods
	AddObject(object unsafe.Pointer)/* debug [protocol_interface/required_method]: AddObject */
	RemoveObject(object unsafe.Pointer)/* debug [protocol_interface/required_method]: RemoveObject */
	ObjectAtIndexedSubscript(index uint) unsafe.Pointer/* debug [protocol_interface/required_method]: ObjectAtIndexedSubscript */
}

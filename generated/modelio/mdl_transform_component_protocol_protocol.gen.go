// Code generated from Apple documentation for ModelIO. DO NOT EDIT.

package modelio

import (
	"unsafe"
)

// PMDLTransformComponent is the MDLTransformComponent protocol interface.
//
// The general interface for classes that manage local coordinate space transforms for 3D objects
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.modelio/documentation/ModelIO/MDLTransformComponent
type PMDLTransformComponent interface {
	// Optional methods
	LocalTransformAtTime(time float64) unsafe.Pointer
	HasLocalTransformAtTime() bool
	SetLocalTransform(transform unsafe.Pointer)
	HasSetLocalTransform() bool
	SetLocalTransformForTime(transform unsafe.Pointer, time float64)
	HasSetLocalTransformForTime() bool
}

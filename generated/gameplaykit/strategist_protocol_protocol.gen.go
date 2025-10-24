// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"unsafe"
)

// PStrategist is the GKStrategist protocol interface.
//
// A general interface for objects that provide artificial intelligence for use in turn-based (and similar) games.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.gameplaykit/documentation/GameplayKit/GKStrategist
type PStrategist interface {
	// Required methods
	BestMoveForActivePlayer() unsafe.Pointer/* debug [protocol_interface/required_method]: BestMoveForActivePlayer */
}

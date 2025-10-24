// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"unsafe"
)

// PPlayerViewControllerAnimationCoordinator is the AVPlayerViewControllerAnimationCoordinator protocol interface.
//
// A protocol that defines the methods to implement to synchronize animations with playback controls’ visibility animation.
//
// Availability:
//   - tvOS 11.0+
//
// See: doc://com.apple.avkit/documentation/AVKit/AVPlayerViewControllerAnimationCoordinator
type PPlayerViewControllerAnimationCoordinator interface {
	// Required methods
	AddCoordinatedAnimationsCompletion(animations unsafe.Pointer, completion unsafe.Pointer)/* debug [protocol_interface/required_method]: AddCoordinatedAnimationsCompletion */
}

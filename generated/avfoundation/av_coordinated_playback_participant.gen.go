// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVCoordinatedPlaybackParticipant] class.
var aVCoordinatedPlaybackParticipantClass = _AVCoordinatedPlaybackParticipantClass{objc.GetClass("AVCoordinatedPlaybackParticipant")}

type _AVCoordinatedPlaybackParticipantClass struct {
	class objc.Class
}

// An object that represents a participant in a coordinated playback session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCoordinatedPlaybackParticipant

type AVCoordinatedPlaybackParticipant struct {
	objectivec.Object
}

// AVCoordinatedPlaybackParticipantFrom constructs a [AVCoordinatedPlaybackParticipant] from an unsafe.Pointer.
//
// An object that represents a participant in a coordinated playback session.
func AVCoordinatedPlaybackParticipantFrom(ptr unsafe.Pointer) AVCoordinatedPlaybackParticipant {
	return AVCoordinatedPlaybackParticipant{objectivec.Object{objc.ID(ptr)}}
}




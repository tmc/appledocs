// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVCoordinatedPlaybackParticipant] class.
var (
	aVCoordinatedPlaybackParticipantClass     _AVCoordinatedPlaybackParticipantClass
	aVCoordinatedPlaybackParticipantClassOnce sync.Once
)

func getAVCoordinatedPlaybackParticipantClass() _AVCoordinatedPlaybackParticipantClass {
	aVCoordinatedPlaybackParticipantClassOnce.Do(func() {
		aVCoordinatedPlaybackParticipantClass = _AVCoordinatedPlaybackParticipantClass{objc.GetClass("AVCoordinatedPlaybackParticipant")}
	})
	return aVCoordinatedPlaybackParticipantClass
}

type _AVCoordinatedPlaybackParticipantClass struct {
	class objc.Class
}

// An interface definition for the [AVCoordinatedPlaybackParticipant] class.
type IAVCoordinatedPlaybackParticipant interface {
	objectivec.IObject
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

// Alloc allocates a new instance without initialization.
func (ac _AVCoordinatedPlaybackParticipantClass) Alloc() AVCoordinatedPlaybackParticipant {
	rv := objc.Send[AVCoordinatedPlaybackParticipant](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVCoordinatedPlaybackParticipantClass) New() AVCoordinatedPlaybackParticipant {
	rv := objc.Send[AVCoordinatedPlaybackParticipant](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVCoordinatedPlaybackParticipant) Init() AVCoordinatedPlaybackParticipant {
	rv := objc.Send[AVCoordinatedPlaybackParticipant](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVCoordinatedPlaybackParticipant) Autorelease() AVCoordinatedPlaybackParticipant {
	rv := objc.Send[AVCoordinatedPlaybackParticipant](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCoordinatedPlaybackParticipant creates a new AVCoordinatedPlaybackParticipant instance.
func NewAVCoordinatedPlaybackParticipant() AVCoordinatedPlaybackParticipant {
	return getAVCoordinatedPlaybackParticipantClass().New()
}





//go:build darwin && ios

// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for MusicPlayerPlayParametersQueueDescriptor


// Sets the time the item with the associated play parameters is to stop playing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerPlayParametersQueueDescriptor/setEndTime(_:forItemWith:)
func (m_ MusicPlayerPlayParametersQueueDescriptor) SetEndTimeForItemWithPlayParameters(endTime float64, playParameters IMPMusicPlayerPlayParameters) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndTime:forItemWithPlayParameters:"), endTime, playParameters)
}

// Sets the time the item with the associated play parameters is to start playing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerPlayParametersQueueDescriptor/setStartTime(_:forItemWith:)
func (m_ MusicPlayerPlayParametersQueueDescriptor) SetStartTimeForItemWithPlayParameters(startTime float64, playParameters IMPMusicPlayerPlayParameters) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTime:forItemWithPlayParameters:"), startTime, playParameters)
}

// iOS-only properties

// An array containing the play parameters returned from querying MusicKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerPlayParametersQueueDescriptor/playParametersQueue
func (m_ MusicPlayerPlayParametersQueueDescriptor) PlayParametersQueue() []MusicPlayerPlayParameters {
	rv := objc.Send[[]MusicPlayerPlayParameters](m_.ID, objc.Sel("playParametersQueue"))
	return rv
}
func (m_ MusicPlayerPlayParametersQueueDescriptor) SetPlayParametersQueue(value []MusicPlayerPlayParameters) {
	m_.ID.Send(objc.RegisterName("setPlayParametersQueue:"), value)
}

// The item identified by the play parameters to play first.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerPlayParametersQueueDescriptor/startItemPlayParameters
func (m_ MusicPlayerPlayParametersQueueDescriptor) StartItemPlayParameters() IMPMusicPlayerPlayParameters {
	rv := objc.Send[MusicPlayerPlayParameters](m_.ID, objc.Sel("startItemPlayParameters"))
	return rv
}
func (m_ MusicPlayerPlayParametersQueueDescriptor) SetStartItemPlayParameters(value IMPMusicPlayerPlayParameters) {
	m_.ID.Send(objc.RegisterName("setStartItemPlayParameters:"), value)
}





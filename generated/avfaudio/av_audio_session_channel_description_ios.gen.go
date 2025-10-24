//go:build darwin && ios

// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for AudioSessionChannelDescription


// iOS-only properties

// A description of the physical location of this channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionChannelDescription/channelLabel
func (a_ AudioSessionChannelDescription) ChannelLabel() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("channelLabel"))
	return rv
}

// The descriptive name for the channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionChannelDescription/channelName
func (a_ AudioSessionChannelDescription) ChannelName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("channelName"))
	return rv
}

// The index of this channel in its owning port’s array of channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionChannelDescription/channelNumber
func (a_ AudioSessionChannelDescription) ChannelNumber() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("channelNumber"))
	return rv
}

// The unique identifier (UID) for this channel’s owning port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionChannelDescription/owningPortUID
func (a_ AudioSessionChannelDescription) OwningPortUID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("owningPortUID"))
	return rv
}






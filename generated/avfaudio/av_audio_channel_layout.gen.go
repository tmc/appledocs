// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AudioChannelLayout] class.
var (
	AudioChannelLayoutClass     _AudioChannelLayoutClass
	AudioChannelLayoutClassOnce sync.Once
)

func getAudioChannelLayoutClass() _AudioChannelLayoutClass {
	AudioChannelLayoutClassOnce.Do(func() {
		AudioChannelLayoutClass = _AudioChannelLayoutClass{objc.GetClass("AVAudioChannelLayout")}
	})
	return AudioChannelLayoutClass
}

type _AudioChannelLayoutClass struct {
	class objc.Class
}

// An interface definition for the [AudioChannelLayout] class.
type IAudioChannelLayout interface {
	objectivec.IObject
	ChannelCount() AudioChannelCount
	SetChannelCount(value IAudioChannelCount)
	Layout() AudioChannelLayout
	SetLayout(value IAudioChannelLayout)
	LayoutTag() unsafe.Pointer
	SetLayoutTag(value unsafe.Pointer)
	AVChannelLayoutKey() string
}

// An object that describes the roles of a set of audio channels.
//
// The class is a thin wrapper for Core Audio’s .


// An object that describes the roles of a set of audio channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelLayout
type AudioChannelLayout struct {
	objectivec.Object
}

// AudioChannelLayoutFrom constructs a [AudioChannelLayout] from an unsafe.Pointer.
//
// An object that describes the roles of a set of audio channels.
func AudioChannelLayoutFrom(ptr unsafe.Pointer) AudioChannelLayout {
	return AudioChannelLayout{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioChannelLayoutClass) Alloc() AudioChannelLayout {
	rv := objc.Send[AudioChannelLayout](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioChannelLayoutClass) New() AudioChannelLayout {
	rv := objc.Send[AudioChannelLayout](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioChannelLayout) Init() AudioChannelLayout {
	rv := objc.Send[AudioChannelLayout](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioChannelLayout) Autorelease() AudioChannelLayout {
	rv := objc.Send[AudioChannelLayout](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioChannelLayout creates a new AudioChannelLayout instance.
func NewAudioChannelLayout() AudioChannelLayout {
	return getAudioChannelLayoutClass().New()
}



// The number of channels of audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiochannellayout/channelcount
func (a_ AudioChannelLayout) ChannelCount() AudioChannelCount {
	rv := objc.Send[AudioChannelCount](a_.ID, objc.Sel("channelCount"))
	return rv
}


// The number of channels of audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiochannellayout/channelcount
func (a_ AudioChannelLayout) SetChannelCount(value IAudioChannelCount) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setChannelCount:"), value)
}


// The underlying audio channel layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiochannellayout/layout
func (a_ AudioChannelLayout) Layout() AudioChannelLayout {
	rv := objc.Send[AudioChannelLayout](a_.ID, objc.Sel("layout"))
	return rv
}


// The underlying audio channel layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiochannellayout/layout
func (a_ AudioChannelLayout) SetLayout(value IAudioChannelLayout) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLayout:"), value)
}


// The audio channel’s underlying layout tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiochannellayout/layouttag
func (a_ AudioChannelLayout) LayoutTag() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("layoutTag"))
	return rv
}


// The audio channel’s underlying layout tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiochannellayout/layouttag
func (a_ AudioChannelLayout) SetLayoutTag(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLayoutTag:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avchannellayoutkey
func (a_ AudioChannelLayout) AVChannelLayoutKey() string {
	rv := objc.Send[string](a_.ID, objc.Sel("AVChannelLayoutKey"))
	return rv
}




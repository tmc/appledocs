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
	

	// properties:
	ChannelCount() AudioChannelCount /* typedef */
	Layout() IAudioChannelLayout
	LayoutTag() objectivec.IObject
	AVChannelLayoutKey() objc.IObject /* cross-framework: NSString */


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AudioChannelLayoutClass) Alloc() AudioChannelLayout {
	rv := objc.Send[AudioChannelLayout](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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






// Creates an audio channel layout object from an existing one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelLayout/init(layout:)
func NewAudioChannelLayoutWithLayout(layout IAudioChannelLayout) AudioChannelLayout {
	instance := getAudioChannelLayoutClass().Alloc()
	rv := objc.Send[AudioChannelLayout](instance.ID, objc.Sel("initWithLayout:"), layout)
	rv.Autorelease()
	return rv
}


// Creates an audio channel layout object from a layout tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelLayout/init(layoutTag:)
func NewAudioChannelLayoutWithLayoutTag(layoutTag objectivec.IObject) AudioChannelLayout {
	instance := getAudioChannelLayoutClass().Alloc()
	rv := objc.Send[AudioChannelLayout](instance.ID, objc.Sel("initWithLayoutTag:"), layoutTag)
	rv.Autorelease()
	return rv
}







// Creates an audio channel layout object from an existing one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelLayout/layoutWithLayout:
func (ac _AudioChannelLayoutClass) LayoutWithLayout(layout IAudioChannelLayout) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("layoutWithLayout:"), layout)
	return rv
}


// Creates an audio channel layout object from an audio channel layout tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelLayout/layoutWithLayoutTag:
func (ac _AudioChannelLayoutClass) LayoutWithLayoutTag(layoutTag objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("layoutWithLayoutTag:"), layoutTag)
	return rv
}












// Indicates whether another audio channel layout is exactly equal to the current layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelLayout/isEqual(_:)
func (a_ AudioChannelLayout) IsEqual(object objc.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isEqual:"), object)
	return rv
}







// The number of channels of audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelLayout/channelCount
func (a_ AudioChannelLayout) ChannelCount() AudioChannelCount /* typedef */ {
	rv := objc.Send[uint32](a_.ID, objc.Sel("channelCount"))
	return rv
}


// The underlying audio channel layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelLayout/layout
func (a_ AudioChannelLayout) Layout() IAudioChannelLayout {
	rv := objc.Send[AudioChannelLayout](a_.ID, objc.Sel("layout"))
	return rv
}


// The audio channel’s underlying layout tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelLayout/layoutTag
func (a_ AudioChannelLayout) LayoutTag() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("layoutTag"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avchannellayoutkey
func (a_ AudioChannelLayout) AVChannelLayoutKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVChannelLayoutKey"))
	return rv
}








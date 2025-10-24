// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioChannelLayout */


/* debug [class_header]: Header for AVAudioChannelLayout */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioChannelLayout */
// An interface definition for the [AudioChannelLayout] class.
type IAudioChannelLayout interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AudioChannelLayout */
	// properties:
	ChannelCount() AudioChannelCount /* typedef */
	Layout() IAudioChannelLayout
	LayoutTag() objectivec.IObject
	AVChannelLayoutKey() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioChannelLayout */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioChannelLayout */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioChannelLayout */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioChannelLayout */

// Creates an audio channel layout object from an existing one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelLayout/init(layout:)
func NewAudioChannelLayoutWithLayout(layout IAudioChannelLayout) AudioChannelLayout {
	instance := getAudioChannelLayoutClass().Alloc()
	rv := objc.Send[AudioChannelLayout](instance.ID, objc.Sel("initWithLayout:"), layout)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioChannelLayoutWithLayout */


// Creates an audio channel layout object from a layout tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelLayout/init(layoutTag:)
func NewAudioChannelLayoutWithLayoutTag(layoutTag objectivec.IObject) AudioChannelLayout {
	instance := getAudioChannelLayoutClass().Alloc()
	rv := objc.Send[AudioChannelLayout](instance.ID, objc.Sel("initWithLayoutTag:"), layoutTag)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioChannelLayoutWithLayoutTag */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioChannelLayout */

// Creates an audio channel layout object from an existing one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelLayout/layoutWithLayout:
func (ac _AudioChannelLayoutClass) LayoutWithLayout(layout IAudioChannelLayout) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("layoutWithLayout:"), layout)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayoutWithLayout) */


// Creates an audio channel layout object from an audio channel layout tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelLayout/layoutWithLayoutTag:
func (ac _AudioChannelLayoutClass) LayoutWithLayoutTag(layoutTag objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("layoutWithLayoutTag:"), layoutTag)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayoutWithLayoutTag) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioChannelLayout */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioChannelLayout */

// Indicates whether another audio channel layout is exactly equal to the current layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelLayout/isEqual(_:)
func (a_ AudioChannelLayout) IsEqual(object objc.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isEqual:"), object)
	return rv
}/* debug [instance_methods/method]: IsEqual */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioChannelLayout */

// The number of channels of audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelLayout/channelCount
func (a_ AudioChannelLayout) ChannelCount() AudioChannelCount /* typedef */ {
	rv := objc.Send[uint32](a_.ID, objc.Sel("channelCount"))
	return rv
}/* debug [instance_properties/getter]: channelCount */


// The underlying audio channel layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelLayout/layout
func (a_ AudioChannelLayout) Layout() IAudioChannelLayout {
	rv := objc.Send[AudioChannelLayout](a_.ID, objc.Sel("layout"))
	return rv
}/* debug [instance_properties/getter]: layout */


// The audio channel’s underlying layout tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelLayout/layoutTag
func (a_ AudioChannelLayout) LayoutTag() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("layoutTag"))
	return rv
}/* debug [instance_properties/getter]: layoutTag */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avchannellayoutkey
func (a_ AudioChannelLayout) AVChannelLayoutKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("AVChannelLayoutKey"))
	return rv
}/* debug [instance_properties/getter]: AVChannelLayoutKey */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioChannelLayout */



// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfaudio"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [InputPickerInteraction] class.
var (
	InputPickerInteractionClass     _InputPickerInteractionClass
	InputPickerInteractionClassOnce sync.Once
)

func getInputPickerInteractionClass() _InputPickerInteractionClass {
	InputPickerInteractionClassOnce.Do(func() {
		InputPickerInteractionClass = _InputPickerInteractionClass{objc.GetClass("AVInputPickerInteraction")}
	})
	return InputPickerInteractionClass
}

type _InputPickerInteractionClass struct {
	class objc.Class
}

// An interface definition for the [InputPickerInteraction] class.
type IInputPickerInteraction interface {
	objectivec.IObject
	Dismiss()
	Present()
	AudioSession() avfaudio.AudioSession
	SetAudioSession(value avfaudio.IAudioSession)
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	Presented() bool
	IsPresented() bool
	SetIsPresented(value bool)
}

// Use to present an input picker.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVInputPickerInteraction
type InputPickerInteraction struct {
	objectivec.Object
}

// InputPickerInteractionFrom constructs a [InputPickerInteraction] from an unsafe.Pointer.
//
// Use to present an input picker.
func InputPickerInteractionFrom(ptr unsafe.Pointer) InputPickerInteraction {
	return InputPickerInteraction{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _InputPickerInteractionClass) Alloc() InputPickerInteraction {
	rv := objc.Send[InputPickerInteraction](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _InputPickerInteractionClass) New() InputPickerInteraction {
	rv := objc.Send[InputPickerInteraction](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ InputPickerInteraction) Init() InputPickerInteraction {
	rv := objc.Send[InputPickerInteraction](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ InputPickerInteraction) Autorelease() InputPickerInteraction {
	rv := objc.Send[InputPickerInteraction](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInputPickerInteraction creates a new InputPickerInteraction instance.
func NewInputPickerInteraction() InputPickerInteraction {
	return getInputPickerInteractionClass().New()
}




// Creates a new instance of AVInputPickerInteraction using a specific .
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVInputPickerInteraction/init(audioSession:)
func NewInputPickerInteractionWithAudioSession(audioSession avfaudio.IAudioSession) InputPickerInteraction {
	instance := getInputPickerInteractionClass().Alloc()
	rv := objc.Send[InputPickerInteraction](instance.ID, objc.Sel("initWithAudioSession:"), audioSession)
	rv.Autorelease()
	return rv
}


// Dismisses the input picker.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVInputPickerInteraction/dismiss()
func (i_ InputPickerInteraction) Dismiss() {
	objc.Send[objc.ID](i_.ID, objc.Sel("dismiss"))
}

// Presents the input picker.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVInputPickerInteraction/present()
func (i_ InputPickerInteraction) Present() {
	objc.Send[objc.ID](i_.ID, objc.Sel("present"))
}

// The audio session for the picker.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVInputPickerInteraction/audioSession
func (i_ InputPickerInteraction) AudioSession() avfaudio.AudioSession {
	rv := objc.Send[avfaudio.AudioSession](i_.ID, objc.Sel("audioSession"))
	return rv
}


// SetAudioSession sets the value of the audioSession property.
// The audio session for the picker.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVInputPickerInteraction/audioSession
func (i_ InputPickerInteraction) SetAudioSession(value avfaudio.IAudioSession) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAudioSession:"), value)
}

// The input picker view’s delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVInputPickerInteraction/delegate-swift.property
func (i_ InputPickerInteraction) Delegate() objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The input picker view’s delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVInputPickerInteraction/delegate-swift.property
func (i_ InputPickerInteraction) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that indicates whether the picker is currently visible.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVInputPickerInteraction/isPresented
func (i_ InputPickerInteraction) Presented() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("presented"))
	return rv
}

// A Boolean value that indicates whether the picker is currently visible.
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avinputpickerinteraction/ispresented
func (i_ InputPickerInteraction) IsPresented() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isPresented"))
	return rv
}


// SetIsPresented sets the value of the isPresented property.
// A Boolean value that indicates whether the picker is currently visible.

//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avinputpickerinteraction/ispresented
func (i_ InputPickerInteraction) SetIsPresented(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsPresented:"), value)
}



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
	// properties:
	IsPresented() bool
	SetIsPresented(value bool)
	// methods:
}

// Use to present an input picker.


// Use to present an input picker.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVInputPickerInteraction/init(audioSession:)
func NewInputPickerInteractionWithAudioSession(audioSession objc.IObject /* cross-framework: AudioSession */) InputPickerInteraction {
	instance := getInputPickerInteractionClass().Alloc()
	rv := objc.Send[InputPickerInteraction](instance.ID, objc.Sel("initWithAudioSession:"), audioSession)
	rv.Autorelease()
	return rv
}



// A Boolean value that indicates whether the picker is currently visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avinputpickerinteraction/ispresented
func (i_ InputPickerInteraction) IsPresented() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isPresented"))
	return rv
}


// A Boolean value that indicates whether the picker is currently visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avinputpickerinteraction/ispresented
func (i_ InputPickerInteraction) SetIsPresented(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsPresented:"), value)
}



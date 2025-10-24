// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfaudio"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVInputPickerInteraction */


/* debug [class_header]: Header for AVInputPickerInteraction */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for InputPickerInteraction */
// An interface definition for the [InputPickerInteraction] class.
type IInputPickerInteraction interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for InputPickerInteraction */
	// properties:
	IsPresented() bool
	SetIsPresented(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for InputPickerInteraction */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for InputPickerInteraction */
// Alloc allocates a new instance without initialization.
func (ic _InputPickerInteractionClass) Alloc() InputPickerInteraction {
	rv := objc.Send[InputPickerInteraction](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for InputPickerInteraction */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for InputPickerInteraction */

// Creates a new instance of AVInputPickerInteraction using a specific .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVInputPickerInteraction/init(audioSession:)
func NewInputPickerInteractionWithAudioSession(audioSession avfaudio.AudioSession) InputPickerInteraction {
	instance := getInputPickerInteractionClass().Alloc()
	rv := objc.Send[InputPickerInteraction](instance.ID, objc.Sel("initWithAudioSession:"), audioSession)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewInputPickerInteractionWithAudioSession */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for InputPickerInteraction */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for InputPickerInteraction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for InputPickerInteraction */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for InputPickerInteraction */

// A Boolean value that indicates whether the picker is currently visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avinputpickerinteraction/ispresented
func (i_ InputPickerInteraction) IsPresented() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isPresented"))
	return rv
}/* debug [instance_properties/getter]: isPresented */


// A Boolean value that indicates whether the picker is currently visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avinputpickerinteraction/ispresented
func (i_ InputPickerInteraction) SetIsPresented(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsPresented:"), value)
}/* debug [instance_properties/setter]: isPresented */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVInputPickerInteraction */



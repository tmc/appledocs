// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfaudio"
)

/* debug [class.gen.go]: Generating class PHASEPushStreamNode */


/* debug [class_header]: Header for PHASEPushStreamNode */
// The class instance for the [PHASEPushStreamNode] class.
var (
	PHASEPushStreamNodeClass     _PHASEPushStreamNodeClass
	PHASEPushStreamNodeClassOnce sync.Once
)

func getPHASEPushStreamNodeClass() _PHASEPushStreamNodeClass {
	PHASEPushStreamNodeClassOnce.Do(func() {
		PHASEPushStreamNodeClass = _PHASEPushStreamNodeClass{objc.GetClass("PHASEPushStreamNode")}
	})
	return PHASEPushStreamNodeClass
}

type _PHASEPushStreamNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEPushStreamNode */
// An interface definition for the [PHASEPushStreamNode] class.
type IPHASEPushStreamNode interface {
	IPHASEStreamNode
	
/* debug [class_interface_properties]: Properties for PHASEPushStreamNode */
	// properties:
	Format() avfaudio.AudioFormat
	GainMetaParameter() IPHASENumberMetaParameter
	Mixer() IPHASEMixer
	RateMetaParameter() IPHASENumberMetaParameter
	PushStreamNodes() IPHASEPushStreamNode
	SetPushStreamNodes(value IPHASEPushStreamNode)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEPushStreamNode */
	// methods:
	ScheduleBuffer(buffer avfaudio.AudioPCMBuffer)
	ScheduleBufferCompletionCallbackTypeCompletionHandler(buffer avfaudio.AudioPCMBuffer, completionCallbackType PHASEPushStreamCompletionCallbackCondition, completionHandler unsafe.Pointer)
	ScheduleBufferAtTimeOptions(buffer avfaudio.AudioPCMBuffer, when avfaudio.AudioTime, options PHASEPushStreamBufferOptions)
	ScheduleBufferAtTimeOptionsCompletionCallbackTypeCompletionHandler(buffer avfaudio.AudioPCMBuffer, when avfaudio.AudioTime, options PHASEPushStreamBufferOptions, completionCallbackType PHASEPushStreamCompletionCallbackCondition, completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEPushStreamNode */
// Alloc allocates a new instance without initialization.
func (pc _PHASEPushStreamNodeClass) Alloc() PHASEPushStreamNode {
	rv := objc.Send[PHASEPushStreamNode](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEPushStreamNodeClass) New() PHASEPushStreamNode {
	rv := objc.Send[PHASEPushStreamNode](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEPushStreamNode) Init() PHASEPushStreamNode {
	rv := objc.Send[PHASEPushStreamNode](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEPushStreamNode) Autorelease() PHASEPushStreamNode {
	rv := objc.Send[PHASEPushStreamNode](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEPushStreamNode creates a new PHASEPushStreamNode instance.
func NewPHASEPushStreamNode() PHASEPushStreamNode {
	return getPHASEPushStreamNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEPushStreamNode */
// An audio stream you manage to provide a sound buffer data.
//
// A sound event’s dictionary populates with an instance of this class when PHASE invokes a in your event node tree. Your app provides the audio data that the sound event plays by calling one or more of this class’s buffer-scheduling functions, for example, .


// An audio stream you manage to provide a sound buffer data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamNode
type PHASEPushStreamNode struct {
	PHASEStreamNode
}

// PHASEPushStreamNodeFrom constructs a [PHASEPushStreamNode] from an unsafe.Pointer.
//
// An audio stream you manage to provide a sound buffer data.
func PHASEPushStreamNodeFrom(ptr unsafe.Pointer) PHASEPushStreamNode {
	return PHASEPushStreamNode{
		PHASEStreamNode: PHASEStreamNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEPushStreamNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEPushStreamNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEPushStreamNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEPushStreamNode */

// Schedules audio data for playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamNode/scheduleBuffer(buffer:)
func (p_ PHASEPushStreamNode) ScheduleBuffer(buffer avfaudio.AudioPCMBuffer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("scheduleBuffer:"), buffer)
}/* debug [instance_methods/method]: ScheduleBuffer */


// Schedules audio data playback with a completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamNode/scheduleBuffer(buffer:completionCallbackType:completionHandler:)
func (p_ PHASEPushStreamNode) ScheduleBufferCompletionCallbackTypeCompletionHandler(buffer avfaudio.AudioPCMBuffer, completionCallbackType PHASEPushStreamCompletionCallbackCondition, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("scheduleBuffer:completionCallbackType:completionHandler:"), buffer, completionCallbackType, completionHandler)
}/* debug [instance_methods/method]: ScheduleBufferCompletionCallbackTypeCompletionHandler */


// Schedules audio data playback at a specific time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamNode/scheduleBuffer(buffer:time:options:)
func (p_ PHASEPushStreamNode) ScheduleBufferAtTimeOptions(buffer avfaudio.AudioPCMBuffer, when avfaudio.AudioTime, options PHASEPushStreamBufferOptions) {
	objc.Send[objc.ID](p_.ID, objc.Sel("scheduleBuffer:atTime:options:"), buffer, when, options)
}/* debug [instance_methods/method]: ScheduleBufferAtTimeOptions */


// Schedules audio data playback at a specific time with a completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamNode/scheduleBuffer(buffer:time:options:completionCallbackType:completionHandler:)
func (p_ PHASEPushStreamNode) ScheduleBufferAtTimeOptionsCompletionCallbackTypeCompletionHandler(buffer avfaudio.AudioPCMBuffer, when avfaudio.AudioTime, options PHASEPushStreamBufferOptions, completionCallbackType PHASEPushStreamCompletionCallbackCondition, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("scheduleBuffer:atTime:options:completionCallbackType:completionHandler:"), buffer, when, options, completionCallbackType, completionHandler)
}/* debug [instance_methods/method]: ScheduleBufferAtTimeOptionsCompletionCallbackTypeCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEPushStreamNode */

// The format of the audio stream data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamNode/format
func (p_ PHASEPushStreamNode) Format() avfaudio.AudioFormat {
	rv := objc.Send[avfaudio.AudioFormat](p_.ID, objc.Sel("format"))
	return rv
}/* debug [instance_properties/getter]: format */


// A meta parameter for dynamic loudness control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamNode/gainMetaParameter
func (p_ PHASEPushStreamNode) GainMetaParameter() IPHASENumberMetaParameter {
	rv := objc.Send[PHASENumberMetaParameter](p_.ID, objc.Sel("gainMetaParameter"))
	return rv
}/* debug [instance_properties/getter]: gainMetaParameter */


// The audio stream’s output pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamNode/mixer
func (p_ PHASEPushStreamNode) Mixer() IPHASEMixer {
	rv := objc.Send[PHASEMixer](p_.ID, objc.Sel("mixer"))
	return rv
}/* debug [instance_properties/getter]: mixer */


// A meta parameter for dynamic rate control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamNode/rateMetaParameter
func (p_ PHASEPushStreamNode) RateMetaParameter() IPHASENumberMetaParameter {
	rv := objc.Send[PHASENumberMetaParameter](p_.ID, objc.Sel("rateMetaParameter"))
	return rv
}/* debug [instance_properties/getter]: rateMetaParameter */


// A collection of audio streams for playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/pushstreamnodes
func (p_ PHASEPushStreamNode) PushStreamNodes() IPHASEPushStreamNode {
	rv := objc.Send[PHASEPushStreamNode](p_.ID, objc.Sel("pushStreamNodes"))
	return rv
}/* debug [instance_properties/getter]: pushStreamNodes */


// A collection of audio streams for playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/pushstreamnodes
func (p_ PHASEPushStreamNode) SetPushStreamNodes(value IPHASEPushStreamNode) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPushStreamNodes:"), value)
}/* debug [instance_properties/setter]: pushStreamNodes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEPushStreamNode */



